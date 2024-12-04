package helm

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofrs/flock"
	"github.com/pluralsh/plural-cli/pkg/utils"
	"gopkg.in/yaml.v2"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/helmpath"
	"helm.sh/helm/v3/pkg/repo"
)

var settings = cli.New()

func debug(format string, v ...interface{}) {
	if utils.EnableDebug {
		format = fmt.Sprintf("INFO: %s\n", format)
		err := log.Output(2, fmt.Sprintf(format, v...))
		if err != nil {
			log.Panic(err)
		}
	}
}

func GetActionConfig(namespace string) (*action.Configuration, error) {
	actionConfig := new(action.Configuration)
	settings := cli.New()
	if os.Getenv("KUBECONFIG") != "" {
		settings.KubeConfig = os.Getenv("KUBECONFIG")
	}

	settings.SetNamespace(namespace)
	settings.Debug = false
	if err := actionConfig.Init(settings.RESTClientGetter(), namespace, "", debug); err != nil {
		return nil, err
	}
	return actionConfig, nil
}

func Template(conf *action.Configuration, name, namespace, path string, isUpgrade, validate bool, values map[string]interface{}) ([]byte, error) {
	// load chart from the path
	chart, err := loader.Load(path)
	if err != nil {
		return nil, err
	}

	client := action.NewInstall(conf)
	client.DryRun = true
	client.ReleaseName = name
	client.Replace = true // Skip the name check
	client.ClientOnly = !validate
	client.IsUpgrade = isUpgrade
	client.Namespace = namespace
	client.IncludeCRDs = false
	rel, err := client.Run(chart, values)
	if err != nil {
		return nil, err
	}
	var manifests bytes.Buffer
	fmt.Fprintln(&manifests, strings.TrimSpace(rel.Manifest))
	return manifests.Bytes(), nil
}

func AddRepo(repoName, repoUrl string) error {
	repoFile := getEnvVar("HELM_REPOSITORY_CONFIG", helmpath.ConfigPath("repositories.yaml"))
	err := os.MkdirAll(filepath.Dir(repoFile), os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	// Acquire a file lock for process synchronization.
	repoFileExt := filepath.Ext(repoFile)
	var lockPath string
	if len(repoFileExt) > 0 && len(repoFileExt) < len(repoFile) {
		lockPath = strings.TrimSuffix(repoFile, repoFileExt) + ".lock"
	} else {
		lockPath = repoFile + ".lock"
	}
	fileLock := flock.New(lockPath)
	lockCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	locked, err := fileLock.TryLockContext(lockCtx, time.Second)
	if err == nil && locked {
		defer func(fileLock *flock.Flock) {
			_ = fileLock.Unlock()
		}(fileLock)
	}
	if err != nil {
		return err
	}

	b, err := os.ReadFile(repoFile)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	var f repo.File
	if err := yaml.Unmarshal(b, &f); err != nil {
		return err
	}

	c := repo.Entry{
		Name:                  repoName,
		URL:                   repoUrl,
		InsecureSkipTLSverify: true,
	}

	// If the repo exists do one of two things:
	// 1. If the configuration for the name is the same continue without error.
	// 2. When the config is different require --force-update.

	// always force updates for now
	// if f.Has(repoName) {
	// 	return nil
	// }

	r, err := repo.NewChartRepository(&c, getter.All(settings))
	if err != nil {
		return err
	}

	if _, err := r.DownloadIndexFile(); err != nil {
		return fmt.Errorf("looks like %q is not a valid chart repository or cannot be reached", repoUrl)
	}

	f.Update(&c)
	return f.WriteFile(repoFile, 0644)
}

func Uninstall(name, namespace string) error {
	if available, err := IsReleaseAvailable(name, namespace); !available {
		return err
	}

	actionConfig, err := GetActionConfig(namespace)
	if err != nil {
		return err
	}
	client := action.NewUninstall(actionConfig)

	_, err = client.Run(name)
	return err
}

func IsReleaseAvailable(name, namespace string) (bool, error) {
	actionConfig, err := GetActionConfig(namespace)
	if err != nil {
		return false, err
	}
	client := action.NewList(actionConfig)
	resp, err := client.Run()
	if err != nil {
		return false, err
	}
	for _, rel := range resp {
		if rel.Name == name {
			return true, nil
		}
	}
	return false, nil
}

func getEnvVar(name, defaultValue string) string {
	if v, ok := os.LookupEnv(name); ok {
		return v
	}

	return defaultValue
}
