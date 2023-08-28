package bootstrap

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pluralsh/plural/pkg/manifest"
	"github.com/pluralsh/plural/pkg/provider"
	"github.com/pluralsh/plural/pkg/utils"
	"github.com/pluralsh/plural/pkg/utils/git"
	"github.com/pluralsh/plural/pkg/utils/pathing"
)

// getEnvVar gets value of environment variable, if it is not set then default value is returned instead.
func getEnvVar(name, defaultValue string) string {
	if v, ok := os.LookupEnv(name); ok {
		return v
	}

	return defaultValue
}

// getBootstrapFlags returns list of provider-specific flags used during cluster bootstrap and destroy.
func getBootstrapFlags(provider string) []string {
	switch provider {
	case aws:
		return []string{
			"--set", "cluster-api-provider-aws.cluster-api-provider-aws.bootstrapMode=true",
			"--set", "bootstrap.aws-ebs-csi-driver.enabled=false",
			"--set", "bootstrap.aws-load-balancer-controller.enabled=false",
			"--set", "bootstrap.cluster-autoscaler.enabled=false",
			"--set", "bootstrap.metrics-server.enabled=false",
			"--set", "bootstrap.snapshot-controller.enabled=false",
			"--set", "bootstrap.snapshot-validation-webhook.enabled=false",
			"--set", "bootstrap.tigera-operator.enabled=false",
		}
	case "google":
		return []string{
			"--set", "bootstrap.cert-manager.serviceAccount.create=true",
			"--set", "cluster-api-provider-gcp.cluster-api-provider-gcp.bootstrapMode=true",
		}
	case "azure":
		return []string{
			"--set", "cluster-api-cluster.cluster.azure.clusterIdentity.bootstrapMode=true",
		}
	default:
		return []string{}
	}
}

// getKubeconfigPath returns path to kubeconfig in user home directory.
func getKubeconfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return pathing.SanitizeFilepath(filepath.Join(homeDir, ".kube", "config")), nil
}

// GetBootstrapPath returns bootstrap repository path.
func GetBootstrapPath() (string, error) {
	gitRootPath, err := git.Root()
	if err != nil {
		return "", err
	}

	return pathing.SanitizeFilepath(filepath.Join(gitRootPath, "bootstrap")), nil
}

// GetStepPath returns path from which step will be executed.
func GetStepPath(step *Step, defaultPath string) string {
	if step != nil && step.TargetPath != "" {
		return step.TargetPath
	}

	return defaultPath
}

// ExecuteSteps of a bootstrap, migration or destroy process.
func ExecuteSteps(steps []*Step) error {
	defaultPath, err := GetBootstrapPath()
	if err != nil {
		return err
	}

	for i, step := range steps {
		utils.Highlight("[%d/%d] %s \n", i+1, len(steps), step.Name)

		if step.Skip != nil && step.Skip() {
			continue
		}

		path := GetStepPath(step, defaultPath)
		err := os.Chdir(path)
		if err != nil {
			return err
		}

		err = step.Execute(step.Args)
		if err != nil {
			return err
		}
	}

	return nil
}

func RunWithTempCredentials(function ActionFunc) error {
	man, err := manifest.FetchProject()
	if err != nil {
		return err
	}

	var flags []string

	switch man.Provider {
	case provider.AZURE:
		acs, err := GetAzureCredentialsService()
		if err != nil {
			return err
		}

		clientId, clientSecret, err := acs.Setup(man.Cluster)
		if err != nil {
			return err
		}

		pathPrefix := "cluster-api-cluster.cluster.azure.clusterIdentity.bootstrapCredentials"
		flags = []string{
			"--set", fmt.Sprintf("%s.%s=%s", pathPrefix, "clientID", clientId),
			"--set", fmt.Sprintf("%s.%s=%s", pathPrefix, "clientSecret", clientSecret),
		}

		defer func(acs *AzureCredentialsService) {
			err := acs.Cleanup()
			if err != nil {
				utils.Error("%s", err)
			}
		}(acs)
	}

	return function(flags)
}
