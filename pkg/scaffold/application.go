package scaffold

import (
	"os"
	"path/filepath"

	"github.com/pluralsh/plural/pkg/output"
	"github.com/pluralsh/plural/pkg/utils"
	"github.com/pluralsh/plural/pkg/utils/git"
	"github.com/pluralsh/plural/pkg/utils/pathing"
	"sigs.k8s.io/yaml"
)

type Applications struct {
	Root string
}

func BuildApplications(root string) *Applications {
	return &Applications{Root: root}
}

func NewApplications() (*Applications, error) {
	root, err := git.Root()
	if err != nil {
		return nil, err
	}

	return BuildApplications(root), nil
}

func (apps *Applications) HelmValues(app string) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	valuesFile := pathing.SanitizeFilepath(filepath.Join(apps.Root, app, "helm", app, "values.yaml"))
	vals := make(map[string]interface{})
	valsContent, err := os.ReadFile(valuesFile)
	if err != nil {
		return res, err
	}
	if err := yaml.Unmarshal(valsContent, &vals); err != nil {
		return res, err
	}

	defaultValuesFile := pathing.SanitizeFilepath(filepath.Join(apps.Root, app, "helm", app, "default-values.yaml"))
	defaultVals := make(map[string]interface{})
	if utils.Exists(defaultValuesFile) {
		defaultValsContent, err := os.ReadFile(defaultValuesFile)
		if err != nil {
			return res, err
		}
		if err := yaml.Unmarshal(defaultValsContent, &defaultVals); err != nil {
			return res, err
		}
	}

	for k, v := range defaultVals {
		res[k] = v
	}
	for k, v := range vals {
		res[k] = v
	}
	return res, err
}

func (apps *Applications) TerraformValues(app string) (map[string]interface{}, error) {
	out, err := output.Read(pathing.SanitizeFilepath(filepath.Join(apps.Root, app, "output.yaml")))
	return out.Terraform, err
}
