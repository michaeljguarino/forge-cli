package edge

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	gqlclient "github.com/pluralsh/console/go/client"
	"github.com/pluralsh/plural-cli/pkg/client"
	"github.com/pluralsh/plural-cli/pkg/console/errors"
	"github.com/pluralsh/plural-cli/pkg/utils"
	"github.com/samber/lo"
	"github.com/urfave/cli"
	"helm.sh/helm/v3/pkg/action"
	"k8s.io/apimachinery/pkg/util/wait"
)

var consoleToken string
var consoleURL string

type Plural struct {
	client.Plural
	HelmConfiguration *action.Configuration
}

func init() {
	consoleToken = ""
	consoleURL = ""
}

func Command(clients client.Plural, helmConfiguration *action.Configuration) cli.Command {
	return cli.Command{
		Name:        "edge",
		Usage:       "manage edge clusters",
		Subcommands: Commands(clients, helmConfiguration),
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "token",
				Usage:       "console token",
				EnvVar:      "PLURAL_CONSOLE_TOKEN",
				Destination: &consoleToken,
			},
			cli.StringFlag{
				Name:        "url",
				Usage:       "console url address",
				EnvVar:      "PLURAL_CONSOLE_URL",
				Destination: &consoleURL,
			},
		},
		Category: "Edge",
	}
}

func Commands(clients client.Plural, helmConfiguration *action.Configuration) []cli.Command {
	p := Plural{
		HelmConfiguration: helmConfiguration,
		Plural:            clients,
	}
	return []cli.Command{
		{
			Name:   "bootstrap",
			Action: p.handleEdgeBootstrap,
			Usage:  "registers edge cluster and installs agent onto it using the current kubeconfig",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:     "machine-id",
					Usage:    "the unique id of the edge device on which this cluster runs",
					Required: true,
				},
				cli.StringFlag{
					Name:     "project",
					Usage:    "the project this cluster will belong to",
					Required: false,
				},
			},
		},
	}
}

func (p *Plural) handleEdgeBootstrap(c *cli.Context) error {
	machineID := c.String("machine-id")
	project := c.String("project")

	if err := p.InitConsoleClient(consoleToken, consoleURL); err != nil {
		return err
	}

	registrationAttributes, err := p.getClusterRegistrationAttributes(machineID, project)
	if err != nil {
		return err
	}

	utils.Highlight("registering new cluster on %s machine\n", machineID)
	_, err = p.ConsoleClient.CreateClusterRegistration(*registrationAttributes) // TODO: Handle the case when it already exists, i.e. after reboot.
	if err != nil {
		return err
	}

	utils.Highlight("waiting for registration to be completed\n")
	var complete bool
	var registration *gqlclient.ClusterRegistrationFragment
	_ = wait.PollUntilContextCancel(context.Background(), 30*time.Second, true, func(_ context.Context) (done bool, err error) {
		complete, registration = p.ConsoleClient.IsClusterRegistrationComplete(machineID)
		return complete, nil
	})

	clusterAttributes, err := p.getClusterAttributes(registration)
	if err != nil {
		return err
	}

	utils.Highlight("creating %s cluster\n", registration.Name)
	cluster, err := p.ConsoleClient.CreateCluster(*clusterAttributes)
	if err != nil {
		if errors.Like(err, "handle") {
			handle := lo.ToPtr(clusterAttributes.Name)
			if clusterAttributes.Handle != nil {
				handle = clusterAttributes.Handle
			}
			return p.ReinstallOperator(c, nil, handle)
		}
		return err
	}

	if cluster.CreateCluster.DeployToken == nil {
		return fmt.Errorf("could not fetch deploy token from cluster")
	}

	url := p.ConsoleClient.ExtUrl()
	if agentUrl, err := p.ConsoleClient.AgentUrl(cluster.CreateCluster.ID); err == nil {
		url = agentUrl
	}

	utils.Highlight("installing agent on %s cluster with %s URL\n", registration.Name, p.ConsoleClient.Url())
	return p.DoInstallOperator(url, *cluster.CreateCluster.DeployToken, "")
}

func (p *Plural) getClusterRegistrationAttributes(machineID, project string) (*gqlclient.ClusterRegistrationCreateAttributes, error) {
	attributes := gqlclient.ClusterRegistrationCreateAttributes{MachineID: machineID}

	if project != "" {
		p, err := p.ConsoleClient.GetProject(project)
		if err != nil {
			return nil, err
		}
		if p == nil {
			return nil, fmt.Errorf("cannot find %s project", project)
		}
		attributes.ProjectID = lo.ToPtr(p.ID)
	}

	return &attributes, nil
}

func (p *Plural) getClusterAttributes(registration *gqlclient.ClusterRegistrationFragment) (*gqlclient.ClusterAttributes, error) {
	attributes := gqlclient.ClusterAttributes{
		Name:   registration.Name,
		Handle: &registration.Handle,
	}

	if registration.Tags != nil {
		attributes.Tags = lo.Map(registration.Tags, func(tag *gqlclient.ClusterTags, index int) *gqlclient.TagAttributes {
			if tag == nil {
				return nil
			}

			return &gqlclient.TagAttributes{
				Name:  tag.Name,
				Value: tag.Value,
			}
		})
	}

	if registration.Metadata != nil {
		metadata, err := json.Marshal(registration.Metadata)
		if err != nil {
			return nil, err
		}
		attributes.Metadata = lo.ToPtr(string(metadata))
	}

	if registration.Project != nil {
		attributes.ProjectID = &registration.Project.ID
	}

	return &attributes, nil
}
