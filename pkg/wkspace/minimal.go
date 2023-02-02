package wkspace

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	helmdiff "github.com/databus23/helm-diff/v3/diff"
	diffmanifest "github.com/databus23/helm-diff/v3/manifest"
	"github.com/imdario/mergo"
	"github.com/pluralsh/plural/pkg/config"
	"github.com/pluralsh/plural/pkg/diff"
	"github.com/pluralsh/plural/pkg/helm"
	"github.com/pluralsh/plural/pkg/manifest"
	"github.com/pluralsh/plural/pkg/output"
	"github.com/pluralsh/plural/pkg/provider"
	"github.com/pluralsh/plural/pkg/utils"
	"github.com/pluralsh/plural/pkg/utils/git"
	"github.com/pluralsh/plural/pkg/utils/pathing"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/storage/driver"
	"sigs.k8s.io/yaml"
)

const (
	valuesYaml           = "values.yaml"
	defaultValuesYaml    = "default-values.yaml"
	helm2TestSuccessHook = "test-success"
	helm3TestHook        = "test"
)

type MinimalWorkspace struct {
	Name     string
	Provider provider.Provider
	Config   *config.Config
	Manifest *manifest.ProjectManifest
}

func Minimal(name string) (*MinimalWorkspace, error) {
	root, err := git.Root()
	if err != nil {
		return nil, err
	}

	prov, err := provider.GetProvider()
	if err != nil {
		return nil, err
	}

	project, _ := manifest.ReadProject(pathing.SanitizeFilepath(filepath.Join(root, "workspace.yaml")))
	conf := config.Read()
	return &MinimalWorkspace{Name: name, Provider: prov, Config: &conf, Manifest: project}, nil
}

func FormatValues(w io.Writer, vals string, output *output.Output) (err error) {
	tmpl, err := template.New("gotpl").Parse(vals)
	if err != nil {
		return
	}
	err = tmpl.Execute(w, map[string]interface{}{"Import": *output})
	return
}

func (m *MinimalWorkspace) BounceHelm(extraArgs ...string) error {
	path, err := filepath.Abs(pathing.SanitizeFilepath(filepath.Join("helm", m.Name)))
	if err != nil {
		return err
	}
	defaultVals, err := getValues(m.Name)
	if err != nil {
		return err
	}
	namespace := m.Config.Namespace(m.Name)
	actionConfig, err := helm.GetActionConfig(namespace)
	if err != nil {
		return err
	}
	utils.Warn("helm upgrade --install --skip-crds --namespace %s %s %s %s\n", namespace, m.Name, path, strings.Join(extraArgs, " "))
	chart, err := loader.Load(path)
	if err != nil {
		return err
	}
	// If a release does not exist, install it.
	histClient := action.NewHistory(actionConfig)
	histClient.Max = 1
	if _, err := histClient.Run(m.Name); errors.Is(err, driver.ErrReleaseNotFound) {
		instClient := action.NewInstall(actionConfig)
		instClient.Namespace = namespace
		instClient.ReleaseName = m.Name
		instClient.SkipCRDs = true

		if req := chart.Metadata.Dependencies; req != nil {
			if err := action.CheckDependencies(chart, req); err != nil {
				return err
			}
		}
		_, err = instClient.Run(chart, defaultVals)
		return err
	}
	client := action.NewUpgrade(actionConfig)
	client.Namespace = namespace
	client.SkipCRDs = true
	client.Timeout = time.Minute * 10
	_, err = client.Run(m.Name, chart, defaultVals)
	return err
}

func getValues(name string) (map[string]interface{}, error) {
	values := make(map[string]interface{})
	defaultVals := make(map[string]interface{})

	path, err := filepath.Abs(pathing.SanitizeFilepath(filepath.Join("helm", name)))
	if err != nil {
		return nil, err
	}

	backup, err := templateVals(name, path)
	if err == nil {
		defer func(oldpath, newpath string) {
			_ = os.Rename(oldpath, newpath)
		}(backup, pathing.SanitizeFilepath(filepath.Join(path, defaultValuesYaml)))
	}

	defaultValuesPath := pathing.SanitizeFilepath(filepath.Join(path, defaultValuesYaml))
	valuesPath := pathing.SanitizeFilepath(filepath.Join(path, valuesYaml))

	valsContent, err := os.ReadFile(valuesPath)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(valsContent, &values); err != nil {
		return nil, err
	}
	if utils.Exists(defaultValuesPath) {
		defaultValsContent, err := os.ReadFile(defaultValuesPath)
		if err != nil {
			return nil, err
		}
		if err := yaml.Unmarshal(defaultValsContent, &defaultVals); err != nil {
			return nil, err
		}
	}

	err = mergo.Merge(&defaultVals, values, mergo.WithOverride)
	if err != nil {
		return nil, err
	}
	return defaultVals, nil
}

func (m *MinimalWorkspace) TemplateHelm() error {
	path, err := filepath.Abs(pathing.SanitizeFilepath(filepath.Join("helm", m.Name)))
	if err != nil {
		return err
	}
	namespace := m.Config.Namespace(m.Name)
	manifest, err := getTemplate(m.Name, namespace, false, false)
	if err != nil {
		return err
	}
	utils.Warn("helm template --skip-crds --namespace %s %s %s\n", namespace, m.Name, path)
	fmt.Printf("%s", manifest)
	return nil
}

func (m *MinimalWorkspace) DiffHelm() error {
	path, err := filepath.Abs(m.Name)
	if err != nil {
		return err
	}
	namespace := m.Config.Namespace(m.Name)
	utils.Warn("helm diff upgrade --install --show-secrets --reset-values  %s %s\n", m.Name, path)
	releaseManifest, err := getRelease(m.Name, namespace)
	if err != nil {
		return err
	}
	installManifest, err := getTemplate(m.Name, namespace, true, true)
	if err != nil {
		return err
	}

	currentSpecs := diffmanifest.Parse(string(releaseManifest), namespace, false, helm3TestHook, helm2TestSuccessHook)
	newSpecs := diffmanifest.Parse(string(installManifest), namespace, false, helm3TestHook, helm2TestSuccessHook)
	helmdiff.Manifests(currentSpecs, newSpecs, &helmdiff.Options{
		OutputFormat:    "diff",
		OutputContext:   -1,
		StripTrailingCR: false,
		ShowSecrets:     true,
		SuppressedKinds: []string{},
		FindRenames:     0,
	}, os.Stdout)
	return nil
}

func (m *MinimalWorkspace) DiffTerraform() error {
	return m.runDiff("terraform", "plan")
}

func (m *MinimalWorkspace) runDiff(command string, args ...string) error {
	diffFolder, err := m.constructDiffFolder()
	if err != nil {
		return err
	}
	outfile, err := os.Create(pathing.SanitizeFilepath(pathing.SanitizeFilepath(filepath.Join(diffFolder, command))))
	if err != nil {
		return err
	}
	defer func(outfile *os.File) {
		_ = outfile.Close()
	}(outfile)

	cmd := exec.Command(command, args...)
	cmd.Stdout = &diff.TeeWriter{File: outfile}
	cmd.Stderr = os.Stdout
	return cmd.Run()
}

func (m *MinimalWorkspace) constructDiffFolder() (string, error) {
	root, err := git.Root()
	if err != nil {
		return "", err
	}

	diffFolder, _ := filepath.Abs(pathing.SanitizeFilepath(filepath.Join(root, "diffs", m.Name)))
	if err := os.MkdirAll(diffFolder, os.ModePerm); err != nil {
		return diffFolder, err
	}

	return diffFolder, err
}

func getRelease(release, namespace string) ([]byte, error) {
	actionConfig, err := helm.GetActionConfig(namespace)
	if err != nil {
		return nil, err
	}
	client := action.NewGet(actionConfig)
	rel, err := client.Run(release)
	if err != nil {
		return nil, err
	}
	return []byte(rel.Manifest), nil
}

func getTemplate(release, namespace string, isUpgrade, validate bool) ([]byte, error) {
	path, err := filepath.Abs(pathing.SanitizeFilepath(filepath.Join("helm", release)))
	if err != nil {
		return nil, err
	}
	defaultVals, err := getValues(release)
	if err != nil {
		return nil, err
	}

	actionConfig, err := helm.GetActionConfig(namespace)
	if err != nil {
		return nil, err
	}

	return helm.Template(actionConfig, release, namespace, path, isUpgrade, validate, defaultVals)
}

func templateVals(app, path string) (backup string, err error) {
	root, _ := utils.ProjectRoot()
	valsFile := pathing.SanitizeFilepath(filepath.Join(path, defaultValuesYaml))
	vals, err := utils.ReadFile(valsFile)
	if err != nil {
		return
	}

	out, err := output.Read(pathing.SanitizeFilepath(filepath.Join(root, app, "output.yaml")))
	if err != nil {
		out = output.New()
	}

	backup = fmt.Sprintf("%s.bak", valsFile)
	err = os.Rename(valsFile, backup)
	if err != nil {
		return
	}

	f, err := os.Create(valsFile)
	if err != nil {
		return
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	err = FormatValues(f, vals, out)
	return
}
