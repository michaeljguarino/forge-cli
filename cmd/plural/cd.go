package plural

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	gqlclient "github.com/pluralsh/console-client-go"
	"github.com/pluralsh/plural/pkg/cd"
	"github.com/pluralsh/plural/pkg/config"
	"github.com/pluralsh/plural/pkg/console"
	"github.com/pluralsh/plural/pkg/utils"
	"github.com/pluralsh/polly/containers"
	"github.com/samber/lo"
	"github.com/urfave/cli"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func init() {
	consoleToken = ""
	consoleURL = ""
}

const (
	operatorNamespace = "plrl-deploy-operator"
)

var consoleToken string
var consoleURL string

func (p *Plural) cdCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "providers",
			Subcommands: p.cdProvidersCommands(),
			Usage:       "manage CD providers",
		},
		{
			Name:        "credentials",
			Subcommands: p.cdCredentialsCommands(),
			Usage:       "manage Provider credentials",
		},
		{
			Name:        "clusters",
			Subcommands: p.cdClusterCommands(),
			Usage:       "manage CD clusters",
		},
		{
			Name:        "services",
			Subcommands: p.cdServiceCommands(),
			Usage:       "manage CD services",
		},
		{
			Name:        "repositories",
			Subcommands: p.cdRepositoriesCommands(),
			Usage:       "manage CD repositories",
		},
		{
			Name:        "pipelines",
			Subcommands: p.pipelineCommands(),
			Usage:       "manage CD pipelines",
		},
		{
			Name:   "install",
			Action: p.handleInstallDeploymentsOperator,
			Usage:  "install deployments operator",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "url", Usage: "console url", Required: true},
				cli.StringFlag{Name: "token", Usage: "console token", Required: true},
			},
		},
		{
			Name:   "control-plane",
			Action: p.handleInstallControlPlane,
			Usage:  "sets up the plural console in an existing k8s cluster",
		},
		{
			Name:   "uninstall",
			Action: p.handleUninstallOperator,
			Usage:  "uninstalls the deployment operator from the current cluster",
		},
		{
			Name:   "login",
			Action: handleCdLogin,
			Usage:  "logs into your plural console",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "url", Usage: "console url", Required: true},
				cli.StringFlag{Name: "token", Usage: "console access token"},
			},
		},
	}
}

func (p *Plural) cdProvidersCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Action: latestVersion(p.handleListProviders),
			Usage:  "list providers",
		},
	}
}

func (p *Plural) cdCredentialsCommands() []cli.Command {
	return []cli.Command{
		{
			Name:      "create",
			ArgsUsage: "PROVIDER_NAME",
			Action:    latestVersion(requireArgs(p.handleCreateProviderCredentials, []string{"PROVIDER_NAME"})),
			Usage:     "create provider credentials",
		},
		{
			Name:      "delete",
			ArgsUsage: "ID",
			Action:    latestVersion(requireArgs(p.handleDeleteProviderCredentials, []string{"ID"})),
			Usage:     "delete provider credentials",
		},
	}
}

func (p *Plural) cdRepositoriesCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Action: latestVersion(p.handleListCDRepositories),
			Usage:  "list repositories",
		},
		{
			Name:   "create",
			Action: latestVersion(p.handleCreateCDRepository),
			Flags: []cli.Flag{
				cli.StringFlag{Name: "url", Usage: "git repo url", Required: true},
				cli.StringFlag{Name: "private-key", Usage: "git repo private key"},
				cli.StringFlag{Name: "passphrase", Usage: "git repo passphrase"},
				cli.StringFlag{Name: "username", Usage: "git repo username"},
				cli.StringFlag{Name: "password", Usage: "git repo password"},
			},
			Usage: "create repository",
		},
		{
			Name:      "update",
			ArgsUsage: "REPO_ID",
			Action:    latestVersion(requireArgs(p.handleUpdateCDRepository, []string{"REPO_ID"})),
			Flags: []cli.Flag{
				cli.StringFlag{Name: "url", Usage: "git repo url", Required: true},
				cli.StringFlag{Name: "private-key", Usage: "git repo private key"},
				cli.StringFlag{Name: "passphrase", Usage: "git repo passphrase"},
				cli.StringFlag{Name: "username", Usage: "git repo username"},
				cli.StringFlag{Name: "password", Usage: "git repo password"},
			},
			Usage: "update repository",
		},
	}
}

func (p *Plural) cdServiceCommands() []cli.Command {
	return []cli.Command{
		{
			Name:      "list",
			ArgsUsage: "CLUSTER_ID",
			Action:    latestVersion(requireArgs(p.handleListClusterServices, []string{"CLUSTER_ID"})),
			Usage:     "list cluster services",
		},
		{
			Name:      "create",
			ArgsUsage: "CLUSTER_ID",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "name", Usage: "service name", Required: true},
				cli.StringFlag{Name: "namespace", Usage: "service namespace. If not specified the 'default' will be used"},
				cli.StringFlag{Name: "version", Usage: "service version. If not specified the '0.0.1' will be used"},
				cli.StringFlag{Name: "repo-id", Usage: "repository ID", Required: true},
				cli.StringFlag{Name: "git-ref", Usage: "git ref, can be branch, tag or commit sha", Required: true},
				cli.StringFlag{Name: "git-folder", Usage: "folder within the source tree where manifests are located", Required: true},
				cli.StringSliceFlag{
					Name:  "conf",
					Usage: "config name value",
				},
				cli.StringFlag{Name: "config-file", Usage: "path for configuration file"},
			},
			Action: latestVersion(requireArgs(p.handleCreateClusterService, []string{"CLUSTER_ID"})),
			Usage:  "create cluster service",
		},
		{
			Name:      "update",
			ArgsUsage: "SERVICE_ID",
			Action:    latestVersion(requireArgs(p.handleUpdateClusterService, []string{"SERVICE_ID"})),
			Usage:     "update cluster service",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "version", Usage: "service version"},
				cli.StringFlag{Name: "git-ref", Usage: "git ref, can be branch, tag or commit sha"},
				cli.StringFlag{Name: "git-folder", Usage: "folder within the source tree where manifests are located"},
				cli.StringSliceFlag{
					Name:  "conf",
					Usage: "config name value",
				},
			},
		},
		{
			Name:      "describe",
			ArgsUsage: "SERVICE_ID",
			Action:    latestVersion(requireArgs(p.handleDescribeClusterService, []string{"SERVICE_ID"})),
			Flags: []cli.Flag{
				cli.StringFlag{Name: "o", Usage: "output format"},
			},
			Usage: "describe cluster service",
		},
		{
			Name:      "delete",
			ArgsUsage: "SERVICE_ID",
			Action:    latestVersion(requireArgs(p.handleDeleteClusterService, []string{"SERVICE_ID"})),
			Usage:     "delete cluster service",
		},
	}
}

func (p *Plural) cdClusterCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "list",
			Action: latestVersion(p.handleListClusters),
			Usage:  "list clusters",
		},
		{
			Name:      "describe",
			Action:    latestVersion(requireArgs(p.handleDescribeCluster, []string{"CLUSTER_ID"})),
			Usage:     "describe cluster",
			ArgsUsage: "CLUSTER_ID",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "o", Usage: "output format"},
			},
		},
		{
			Name:      "update",
			Action:    latestVersion(requireArgs(p.handleUpdateCluster, []string{"CLUSTER_ID"})),
			Usage:     "update cluster",
			ArgsUsage: "CLUSTER_ID",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "handle", Usage: "unique human readable name used to identify this cluster"},
				cli.StringFlag{Name: "kubeconf-path", Usage: "path to kubeconfig"},
				cli.StringFlag{Name: "kubeconf-context", Usage: "the kubeconfig context you want to use. If not specified, the current one will be used"},
			},
		},
		{
			Name:      "delete",
			Action:    latestVersion(requireArgs(p.handleDeleteCluster, []string{"CLUSTER_ID"})),
			Usage:     "deregisters a cluster in plural cd, and drains all services (unless --soft is specified)",
			ArgsUsage: "CLUSTER_ID",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "soft", Usage: "deletes a cluster in our system but doesn't drain resources, leaving them untouched"},
			},
		},
		{
			Name:      "get-credentials",
			Aliases:   []string{"kubeconfig"},
			Action:    latestVersion(requireArgs(p.handleGetClusterCredentials, []string{"CLUSTER_ID"})),
			Usage:     "updates kubeconfig file with appropriate credentials to point to specified cluster",
			ArgsUsage: "CLUSTER_ID",
		},
		{
			Name:      "create",
			Action:    latestVersion(requireArgs(p.handleCreateCluster, []string{"CLUSTER_NAME"})),
			Usage:     "create cluster",
			ArgsUsage: "CLUSTER_NAME",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "handle", Usage: "unique human readable name used to identify this cluster"},
				cli.StringFlag{Name: "version", Usage: "kubernetes cluster version", Required: true},
			},
		},
		{
			Name:   "bootstrap",
			Action: latestVersion(p.handleClusterBootstrap),
			Usage:  "creates a new BYOK cluster and installs the agent onto it using the current kubeconfig",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "name", Usage: "The name you'll give the cluster", Required: true},
				cli.StringFlag{Name: "handle", Usage: "optional handle for the cluster"},
				cli.StringSliceFlag{
					Name:  "tag",
					Usage: "a cluster tag to add, useful for targeting with global services",
				},
			},
		},
	}
}

func (p *Plural) pipelineCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "create",
			Action: latestVersion(requireArgs(p.handleCreatePipeline, []string{})),
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file",
					Usage: "the file this pipeline is defined in, use - for stdin",
				},
			},
		},
	}
}

func (p *Plural) handleCreateCDRepository(c *cli.Context) error {
	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}
	url := c.String("url")
	repo, err := p.ConsoleClient.CreateRepository(url, getFlag(c.String("privateKey")),
		getFlag(c.String("passphrase")), getFlag(c.String("username")), getFlag(c.String("password")))
	if err != nil {
		return err
	}

	headers := []string{"ID", "URL"}
	return utils.PrintTable([]gqlclient.GitRepositoryFragment{*repo.CreateGitRepository}, headers, func(r gqlclient.GitRepositoryFragment) ([]string, error) {
		return []string{r.ID, r.URL}, nil
	})
}

func (p *Plural) handleUpdateCDRepository(c *cli.Context) error {
	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}
	repoId := c.Args().Get(0)

	attr := gqlclient.GitAttributes{
		URL: c.String("url"),
	}

	if c.String("private-key") != "" {
		attr.PrivateKey = lo.ToPtr(c.String("private-key"))
	}

	if c.String("passphrase") != "" {
		attr.Passphrase = lo.ToPtr(c.String("passphrase"))
	}

	if c.String("password") != "" {
		attr.Password = lo.ToPtr(c.String("password"))
	}

	if c.String("username") != "" {
		attr.Username = lo.ToPtr(c.String("username"))
	}

	repo, err := p.ConsoleClient.UpdateRepository(repoId, attr)
	if err != nil {
		return err
	}

	headers := []string{"ID", "URL"}
	return utils.PrintTable([]gqlclient.GitRepositoryFragment{*repo.UpdateGitRepository}, headers, func(r gqlclient.GitRepositoryFragment) ([]string, error) {
		return []string{r.ID, r.URL}, nil
	})
}

func (p *Plural) handleListCDRepositories(_ *cli.Context) error {
	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}
	repos, err := p.ConsoleClient.ListRepositories()
	if err != nil {
		return err
	}
	if repos == nil {
		return fmt.Errorf("returned objects list [ListRepositories] is nil")
	}
	headers := []string{"ID", "URL", "Status", "Error"}
	return utils.PrintTable(repos.GitRepositories.Edges, headers, func(r *gqlclient.GitRepositoryEdgeFragment) ([]string, error) {
		health := "UNKNOWN"
		if r.Node.Health != nil {
			health = string(*r.Node.Health)
		}
		return []string{r.Node.ID, r.Node.URL, health, lo.FromPtr(r.Node.Error)}, nil
	})

}

func getIdAndName(input string) (id, name *string) {
	if strings.HasPrefix(input, "@") {
		h := strings.Trim(input, "@")
		name = &h
	} else {
		id = &input
	}
	return
}

func (p *Plural) handleListClusterServices(c *cli.Context) error {
	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}
	sd, err := p.ConsoleClient.ListClusterServices(getIdAndName(c.Args().Get(0)))
	if err != nil {
		return err
	}
	if sd == nil {
		return fmt.Errorf("returned objects list [ListClusterServices] is nil")
	}
	headers := []string{"Id", "Name", "Namespace", "Git Ref", "Git Folder", "Repo"}
	return utils.PrintTable(sd, headers, func(sd *gqlclient.ServiceDeploymentEdgeFragment) ([]string, error) {
		return []string{sd.Node.ID, sd.Node.Name, sd.Node.Namespace, sd.Node.Git.Ref, sd.Node.Git.Folder, sd.Node.Repository.URL}, nil
	})
}

type ServiceDeploymentAttributesConfiguration struct {
	Configuration []*gqlclient.ConfigAttributes
}

func (p *Plural) handleCreateClusterService(c *cli.Context) error {
	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}
	v, err := validateFlag(c, "version", "0.0.1")
	if err != nil {
		return err
	}
	name := c.String("name")
	namespace, err := validateFlag(c, "namespace", "default")
	if err != nil {
		return err
	}
	repoId := c.String("repo-id")
	gitRef := c.String("git-ref")
	gitFolder := c.String("git-folder")
	attributes := gqlclient.ServiceDeploymentAttributes{
		Name:         name,
		Namespace:    namespace,
		Version:      &v,
		RepositoryID: repoId,
		Git: gqlclient.GitRefAttributes{
			Ref:    gitRef,
			Folder: gitFolder,
		},
		Configuration: []*gqlclient.ConfigAttributes{},
	}

	if c.String("config-file") != "" {
		configFile, err := utils.ReadFile(c.String("config-file"))
		if err != nil {
			return err
		}
		sdc := ServiceDeploymentAttributesConfiguration{}
		if err := yaml.Unmarshal([]byte(configFile), &sdc); err != nil {
			return err
		}
		attributes.Configuration = append(attributes.Configuration, sdc.Configuration...)
	}
	var confArgs []string
	if c.IsSet("conf") {
		confArgs = append(confArgs, c.StringSlice("conf")...)
	}
	for _, conf := range confArgs {
		configurationPair := strings.Split(conf, "=")
		if len(configurationPair) == 2 {
			attributes.Configuration = append(attributes.Configuration, &gqlclient.ConfigAttributes{
				Name:  configurationPair[0],
				Value: &configurationPair[1],
			})
		}
	}

	clusterId, clusterName := getIdAndName(c.Args().Get(0))
	sd, err := p.ConsoleClient.CreateClusterService(clusterId, clusterName, attributes)
	if err != nil {
		return err
	}
	if sd == nil {
		return fmt.Errorf("the returned object is empty, check if all fields are set")
	}

	headers := []string{"Id", "Name", "Namespace", "Git Ref", "Git Folder", "Repo"}
	return utils.PrintTable([]*gqlclient.ServiceDeploymentFragment{sd}, headers, func(sd *gqlclient.ServiceDeploymentFragment) ([]string, error) {
		return []string{sd.ID, sd.Name, sd.Namespace, sd.Git.Ref, sd.Git.Folder, sd.Repository.URL}, nil
	})
}

func getServiceIdClusterNameServiceName(input string) (serviceId, clusterName, serviceName *string, err error) {
	if strings.HasPrefix(input, "@") {
		i := strings.Trim(input, "@")
		split := strings.Split(i, "/")
		if len(split) != 2 {
			err = fmt.Errorf("expected format @clusterName/serviceName")
			return
		}
		clusterName = &split[0]
		serviceName = &split[1]
	} else {
		serviceId = &input
	}
	return
}

func (p *Plural) handleDescribeClusterService(c *cli.Context) error {
	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}

	serviceId, clusterName, serviceName, err := getServiceIdClusterNameServiceName(c.Args().Get(0))
	if err != nil {
		return err
	}
	existing, err := p.ConsoleClient.GetClusterService(serviceId, serviceName, clusterName)
	if err != nil {
		return err
	}
	if existing == nil {
		return fmt.Errorf("existing service deployment is empty")
	}
	output := c.String("o")
	if output == "json" {
		utils.NewJsonPrinter(existing).PrettyPrint()
		return nil
	} else if output == "yaml" {
		utils.NewYAMLPrinter(existing).PrettyPrint()
		return nil
	}
	desc, err := console.DescribeService(existing)
	if err != nil {
		return err
	}
	fmt.Print(desc)

	return nil
}

func (p *Plural) handleUpdateClusterService(c *cli.Context) error {
	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}
	serviceId, clusterName, serviceName, err := getServiceIdClusterNameServiceName(c.Args().Get(0))
	if err != nil {
		return err
	}

	existing, err := p.ConsoleClient.GetClusterService(serviceId, serviceName, clusterName)
	if err != nil {
		return err
	}
	if existing == nil {
		return fmt.Errorf("existing service deployment is empty")
	}
	existingConfigurations := map[string]string{}
	attributes := gqlclient.ServiceUpdateAttributes{
		Version: &existing.Version,
		Git: &gqlclient.GitRefAttributes{
			Ref:    existing.Git.Ref,
			Folder: existing.Git.Folder,
		},
		Configuration: []*gqlclient.ConfigAttributes{},
	}

	for _, conf := range existing.Configuration {
		existingConfigurations[conf.Name] = conf.Value
	}

	v := c.String("version")
	if v != "" {
		attributes.Version = &v
	}
	if c.String("git-ref") != "" {
		attributes.Git.Ref = c.String("git-ref")
	}
	if c.String("git-folder") != "" {
		attributes.Git.Folder = c.String("git-folder")
	}
	var confArgs []string
	if c.IsSet("conf") {
		confArgs = append(confArgs, c.StringSlice("conf")...)
	}

	updateConfigurations := map[string]string{}
	for _, conf := range confArgs {
		configurationPair := strings.Split(conf, "=")
		if len(configurationPair) == 2 {
			updateConfigurations[configurationPair[0]] = configurationPair[1]
		}
	}
	for k, v := range updateConfigurations {
		existingConfigurations[k] = v
	}
	for key, value := range existingConfigurations {
		attributes.Configuration = append(attributes.Configuration, &gqlclient.ConfigAttributes{
			Name:  key,
			Value: lo.ToPtr(value),
		})
	}

	sd, err := p.ConsoleClient.UpdateClusterService(serviceId, serviceName, clusterName, attributes)
	if err != nil {
		return err
	}
	if sd == nil {
		return fmt.Errorf("returned object is nil")
	}

	headers := []string{"Id", "Name", "Namespace", "Git Ref", "Git Folder", "Repo"}
	return utils.PrintTable([]*gqlclient.ServiceDeploymentFragment{sd}, headers, func(sd *gqlclient.ServiceDeploymentFragment) ([]string, error) {
		return []string{sd.ID, sd.Name, sd.Namespace, sd.Git.Ref, sd.Git.Folder, sd.Repository.URL}, nil
	})
}

func (p *Plural) handleListClusters(_ *cli.Context) error {
	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}

	clusters, err := p.ConsoleClient.ListClusters()
	if err != nil {
		return err
	}
	if clusters == nil {
		return fmt.Errorf("returned objects list [ListClusters] is nil")
	}
	headers := []string{"Id", "Name", "Handle", "Version", "Provider"}
	return utils.PrintTable(clusters.Clusters.Edges, headers, func(cl *gqlclient.ClusterEdgeFragment) ([]string, error) {
		provider := ""
		if cl.Node.Provider != nil {
			provider = cl.Node.Provider.Name
		}
		handle := ""
		if cl.Node.Handle != nil {
			handle = *cl.Node.Handle
		}
		return []string{cl.Node.ID, cl.Node.Name, handle, *cl.Node.Version, provider}, nil
	})
}

func (p *Plural) handleInstallDeploymentsOperator(c *cli.Context) error {
	return p.doInstallOperator(c.String("url"), c.String("token"))
}

func (p *Plural) handleUninstallOperator(c *cli.Context) error {
	err := p.InitKube()
	if err != nil {
		return err
	}
	return console.UninstallAgent(operatorNamespace)
}

func (p *Plural) doInstallOperator(url, token string) error {
	err := p.InitKube()
	if err != nil {
		return err
	}
	err = p.Kube.CreateNamespace(operatorNamespace)
	if !apierrors.IsAlreadyExists(err) {
		return err
	}
	err = console.InstallAgent(url, token, operatorNamespace)
	if err == nil {
		utils.Success("deployment operator installed successfully\n")
	}
	return err
}

func (p *Plural) handleDeleteClusterService(c *cli.Context) error {
	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}
	serviceId := c.Args().Get(0)
	existing, err := p.ConsoleClient.DeleteClusterService(serviceId)
	if err != nil {
		return fmt.Errorf("could not delete service: %s", err)
	}

	utils.Success("Service %s/%s has been deleted successfully", existing.DeleteServiceDeployment.ID, existing.DeleteServiceDeployment.Name)
	return nil
}

func (p *Plural) handleDescribeCluster(c *cli.Context) error {
	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}
	existing, err := p.ConsoleClient.GetCluster(getIdAndName(c.Args().Get(0)))
	if err != nil {
		return err
	}
	if existing == nil {
		return fmt.Errorf("existing cluster is empty")
	}
	output := c.String("o")
	if output == "json" {
		utils.NewJsonPrinter(existing).PrettyPrint()
	} else if output == "yaml" {
		utils.NewYAMLPrinter(existing).PrettyPrint()
	}
	desc, err := console.DescribeCluster(existing)
	if err != nil {
		return err
	}
	fmt.Print(desc)
	return nil
}

func (p *Plural) handleCreatePipeline(c *cli.Context) error {
	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}

	var bytes []byte
	var err error
	file := c.String("file")
	if file == "-" {
		bytes, err = io.ReadAll(os.Stdin)
	} else {
		bytes, err = os.ReadFile(file)
	}

	if err != nil {
		return err
	}

	name, attrs, err := console.ConstructPipelineInput(bytes)
	if err != nil {
		return err
	}

	pipe, err := p.ConsoleClient.SavePipeline(name, *attrs)
	if err != nil {
		return err
	}

	utils.Success("Pipeline %s created successfully", pipe.Name)
	return nil
}

func (p *Plural) handleUpdateCluster(c *cli.Context) error {
	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}
	existing, err := p.ConsoleClient.GetCluster(getIdAndName(c.Args().Get(0)))
	if err != nil {
		return err
	}
	if existing == nil {
		return fmt.Errorf("this cluster does not exist")
	}
	updateAttr := gqlclient.ClusterUpdateAttributes{
		Version: existing.Version,
		Handle:  existing.Handle,
	}
	newHandle := c.String("handle")
	if newHandle != "" {
		updateAttr.Handle = &newHandle
	}
	kubeconfigPath := c.String("kubeconf-path")
	if kubeconfigPath != "" {
		kubeconfig, err := getKubeconfig(kubeconfigPath, c.String("kubeconf-context"))
		if err != nil {
			return err
		}

		updateAttr.Kubeconfig = &gqlclient.KubeconfigAttributes{
			Raw: &kubeconfig,
		}
	}

	result, err := p.ConsoleClient.UpdateCluster(existing.ID, updateAttr)
	if err != nil {
		return err
	}
	headers := []string{"Id", "Name", "Handle", "Version", "Provider"}
	return utils.PrintTable([]gqlclient.ClusterFragment{*result.UpdateCluster}, headers, func(cl gqlclient.ClusterFragment) ([]string, error) {
		provider := ""
		if cl.Provider != nil {
			provider = cl.Provider.Name
		}
		handle := ""
		if cl.Handle != nil {
			handle = *cl.Handle
		}
		return []string{cl.ID, cl.Name, handle, *cl.Version, provider}, nil
	})
}

func (p *Plural) handleDeleteCluster(c *cli.Context) error {
	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}

	existing, err := p.ConsoleClient.GetCluster(getIdAndName(c.Args().Get(0)))
	if err != nil {
		return err
	}
	if existing == nil {
		return fmt.Errorf("this cluster does not exist")
	}

	return p.ConsoleClient.DeleteCluster(existing.ID)
}

func handleCdLogin(c *cli.Context) (err error) {
	url := c.String("url")
	token := c.String("token")
	if token == "" {
		token, err = utils.ReadPwd("Enter your console access token")
		if err != nil {
			return
		}
	}
	conf := console.Config{Url: url, Token: token}
	return conf.Save()
}

func (p *Plural) handleGetClusterCredentials(c *cli.Context) error {
	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}

	cluster, err := p.ConsoleClient.GetCluster(getIdAndName(c.Args().Get(0)))
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster is nil")
	}

	return cd.SaveClusterKubeconfig(cluster, consoleToken)
}

func getKubeconfig(path, context string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		home, _ := os.UserHomeDir()
		path = filepath.Join(home, path[2:])
	}
	if !utils.Exists(path) {
		return "", fmt.Errorf("the specified path does not exist")
	}

	config, err := clientcmd.LoadFromFile(path)
	if err != nil {
		return "", err
	}

	if context != "" {
		if config.Contexts[context] == nil {
			return "", fmt.Errorf("the given context doesn't exist")
		}
		config.CurrentContext = context
	}
	newConfig := *clientcmdapi.NewConfig()
	newConfig.CurrentContext = config.CurrentContext
	newConfig.Contexts[config.CurrentContext] = config.Contexts[config.CurrentContext]
	newConfig.Clusters[config.CurrentContext] = config.Clusters[config.CurrentContext]
	newConfig.AuthInfos[config.CurrentContext] = config.AuthInfos[config.CurrentContext]
	newConfig.Extensions[config.CurrentContext] = config.Extensions[config.CurrentContext]
	newConfig.Preferences = config.Preferences
	result, err := clientcmd.Write(newConfig)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func getFlag(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func validateFlag(ctx *cli.Context, name string, defaultVal string) (string, error) {
	res := ctx.String(name)
	if res == "" {
		if defaultVal == "" {
			return "", fmt.Errorf("expected --%s flag", name)
		}
		res = defaultVal
	}

	return res, nil
}

var providerSurvey = []*survey.Question{
	{
		Name:   "name",
		Prompt: &survey.Input{Message: "Enter the name of your provider:"},
	},
	{
		Name:   "namespace",
		Prompt: &survey.Input{Message: "Enter the namespace of your provider:"},
	},
}

func (p *Plural) handleClusterBootstrap(c *cli.Context) error {
	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}

	attrs := gqlclient.ClusterAttributes{Name: c.String("name")}
	if c.String("handle") != "" {
		attrs.Handle = lo.ToPtr(c.String("handle"))
	}

	if c.IsSet("tag") {
		attrs.Tags = lo.Map(c.StringSlice("tag"), func(tag string, index int) *gqlclient.TagAttributes {
			tags := strings.Split(tag, "=")
			if len(tags) == 2 {
				return &gqlclient.TagAttributes{
					Name:  tags[0],
					Value: tags[1],
				}
			}
			return nil
		})
		attrs.Tags = lo.Filter(attrs.Tags, func(t *gqlclient.TagAttributes, ind int) bool { return t != nil })
	}

	existing, err := p.ConsoleClient.CreateCluster(attrs)
	if err != nil {
		return err
	}

	if existing.CreateCluster.DeployToken == nil {
		return fmt.Errorf("could not fetch deploy token from cluster")
	}

	deployToken := *existing.CreateCluster.DeployToken
	url := fmt.Sprintf("%s/ext/gql", p.ConsoleClient.Url())
	utils.Highlight("instaling agent on %s with url %s and initial deploy token %s\n", c.String("name"), p.ConsoleClient.Url(), deployToken)
	return p.doInstallOperator(url, deployToken)
}

func (p *Plural) handleCreateCluster(c *cli.Context) error {
	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}
	name := c.Args().Get(0)
	attr := gqlclient.ClusterAttributes{
		Name: name,
	}
	if c.String("handle") != "" {
		attr.Handle = lo.ToPtr(c.String("handle"))
	}
	if c.String("version") != "" {
		attr.Version = lo.ToPtr(c.String("version"))
	}

	providerList, err := p.ConsoleClient.ListProviders()
	if err != nil {
		return err
	}
	providerNames := []string{}
	providerMap := map[string]string{}
	cloudProviders := []string{}
	for _, prov := range providerList.ClusterProviders.Edges {
		providerNames = append(providerNames, prov.Node.Name)
		providerMap[prov.Node.Name] = prov.Node.ID
		cloudProviders = append(cloudProviders, prov.Node.Cloud)
	}

	existingProv := containers.ToSet[string](cloudProviders)
	availableProv := containers.ToSet[string](availableProviders)
	toCreate := availableProv.Difference(existingProv)
	createNewProvider := "Create New Provider"

	if toCreate.Len() != 0 {
		providerNames = append(providerNames, createNewProvider)
	}

	prompt := &survey.Select{
		Message: "Select one of the following providers:",
		Options: providerNames,
	}
	provider := ""
	if err := survey.AskOne(prompt, &provider, survey.WithValidator(survey.Required)); err != nil {
		return err
	}
	if provider != createNewProvider {
		utils.Success("Using provider %s\n", provider)
		id := providerMap[provider]
		attr.ProviderID = &id
	} else {

		clusterProv, err := p.handleCreateProvider(toCreate.List())
		if err != nil {
			return err
		}
		if clusterProv == nil {
			utils.Success("All supported providers are created\n")
			return nil
		}
		utils.Success("Provider %s created successfully\n", clusterProv.CreateClusterProvider.Name)
		attr.ProviderID = &clusterProv.CreateClusterProvider.ID
		provider = clusterProv.CreateClusterProvider.Cloud
	}

	ca, err := cd.AskCloudSettings(provider)
	if err != nil {
		return err
	}
	attr.CloudSettings = ca

	existing, err := p.ConsoleClient.CreateCluster(attr)
	if err != nil {
		return err
	}
	if existing == nil {
		return fmt.Errorf("couldn't create cluster")
	}
	return nil
}

func (p *Plural) handleInstallControlPlane(c *cli.Context) error {
	conf := config.Read()
	vals, err := cd.CreateControlPlane(conf)
	if err != nil {
		return err
	}

	fmt.Print("\n\n")
	utils.Highlight("===> writing values.secret.yaml, you should keep this in a secure location for future helm upgrades\n\n")
	if err := os.WriteFile("values.secret.yaml", []byte(vals), 0644); err != nil {
		return err
	}

	fmt.Println("After confirming everything looks correct in values.secret.yaml, run the following command to install:")
	utils.Highlight("helm upgrade --install --create-namespace -f values.secret.yaml console -n plrl-console")
	return nil
}

func (p *Plural) handleCreateProvider(existingProviders []string) (*gqlclient.CreateClusterProvider, error) {
	provider := ""
	var resp struct {
		Name      string
		Namespace string
	}
	if err := survey.Ask(providerSurvey, &resp); err != nil {
		return nil, err
	}

	prompt := &survey.Select{
		Message: "Select one of the following providers:",
		Options: existingProviders,
	}
	if err := survey.AskOne(prompt, &provider, survey.WithValidator(survey.Required)); err != nil {
		return nil, err
	}

	cps, err := cd.AskCloudProviderSettings(provider)
	if err != nil {
		return nil, err
	}

	providerAttr := gqlclient.ClusterProviderAttributes{
		Name:          resp.Name,
		Namespace:     &resp.Namespace,
		Cloud:         &provider,
		CloudSettings: cps,
	}
	clusterProv, err := p.ConsoleClient.CreateProvider(providerAttr)
	if err != nil {
		return nil, err
	}
	if clusterProv == nil {
		return nil, fmt.Errorf("provider was not created properly")
	}
	return clusterProv, nil
}
