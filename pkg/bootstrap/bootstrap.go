package bootstrap

import (
	"os"
	"path/filepath"

	"github.com/pluralsh/plural/pkg/manifest"
	"github.com/pluralsh/plural/pkg/utils"
	"github.com/pluralsh/plural/pkg/utils/git"
	"github.com/pluralsh/plural/pkg/utils/pathing"
)

// getBootstrapSteps returns list of steps to run during cluster bootstrap.
func getBootstrapSteps(runPlural ActionFunc) ([]*Step, error) {
	projectManifest, err := manifest.FetchProject()
	if err != nil {
		return nil, err
	}

	gitRootDir, err := git.Root()
	if err != nil {
		return nil, err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	bootstrapPath := pathing.SanitizeFilepath(filepath.Join(gitRootDir, "bootstrap"))
	kubeconfigPath := pathing.SanitizeFilepath(filepath.Join(homeDir, ".kube", "config"))
	flags := getBootstrapFlags(projectManifest.Provider)

	return []*Step{
		{
			Name:       "Create local bootstrap cluster",
			Args:       []string{"plural", "bootstrap", "cluster", "create", "bootstrap", "--skip-if-exists"},
			Execute:    runPlural,
			TargetPath: bootstrapPath,
		},
		{
			Name:       "Bootstrap CRDs in local cluster",
			Args:       []string{"plural", "--bootstrap", "wkspace", "crds", "bootstrap"},
			Execute:    runPlural,
			TargetPath: bootstrapPath,
		},
		{
			Name:       "Install Cluster API operators in local cluster",
			Args:       append([]string{"plural", "--bootstrap", "wkspace", "helm", "bootstrap", "--skip", "cluster-api-cluster"}, flags...),
			Execute:    runPlural,
			TargetPath: bootstrapPath,
		},
		{
			Name:       "Deploy cluster",
			Args:       append([]string{"plural", "--bootstrap", "wkspace", "helm", "bootstrap"}, flags...),
			Execute:    runPlural,
			TargetPath: bootstrapPath,
		},
		{
			Name:       "Wait for cluster",
			Args:       []string{"plural", "--bootstrap", "clusters", "wait", "bootstrap", projectManifest.Cluster},
			Execute:    runPlural,
			TargetPath: bootstrapPath,
		},
		{
			Name:       "Wait for machine pools",
			Args:       []string{"plural", "--bootstrap", "clusters", "mpwait", "bootstrap", projectManifest.Cluster},
			Execute:    runPlural,
			TargetPath: bootstrapPath,
		},
		{
			Name:       "Initialize kubeconfig for target cluster",
			Args:       []string{"plural", "wkspace", "kube-init"},
			Execute:    runPlural,
			TargetPath: bootstrapPath,
		},
		{
			Name:       "Create bootstrap namespace in target cluster",
			Args:       []string{"plural", "bootstrap", "namespace", "create", "bootstrap"},
			Execute:    runPlural,
			TargetPath: bootstrapPath,
		},
		{
			Name:       "Bootstrap CRDs in target cluster",
			Args:       []string{"plural", "wkspace", "crds", "bootstrap"},
			Execute:    runPlural,
			TargetPath: bootstrapPath,
		},
		{
			Name:       "Install Cluster API operators in target cluster",
			Args:       append([]string{"plural", "wkspace", "helm", "bootstrap", "--skip", "cluster-api-cluster"}, flags...),
			Execute:    runPlural,
			TargetPath: bootstrapPath,
		},
		{
			Name:       "Move resources from local to target cluster",
			Args:       []string{"plural", "bootstrap", "cluster", "move", "--kubeconfig-context", "kind-bootstrap", "--to-kubeconfig", kubeconfigPath},
			Execute:    runPlural,
			TargetPath: bootstrapPath,
		},
		{
			Name:       "Destroy local cluster",
			Args:       []string{"plural", "--bootstrap", "bootstrap", "cluster", "delete", "bootstrap"},
			Execute:    runPlural,
			TargetPath: bootstrapPath,
		},
	}, nil
}

// BootstrapCluster bootstraps cluster with Cluster API.
func BootstrapCluster(runPlural ActionFunc) error {
	utils.Highlight("Bootstrapping cluster with Cluster API...\n")

	steps, err := getBootstrapSteps(runPlural)
	if err != nil {
		return err
	}

	for i, step := range steps {
		utils.Highlight("[%d/%d] %s \n", i+1, len(steps), step.Name)

		err := os.Chdir(step.TargetPath)
		if err != nil {
			return err
		}

		err = step.Execute(step.Args)
		if err != nil {
			return err
		}
	}

	utils.Success("Cluster bootstrapped successfully!\n")
	return nil
}
