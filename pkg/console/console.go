package console

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	consoleclient "github.com/pluralsh/console-client-go"
)

type consoleClient struct {
	ctx    context.Context
	client consoleclient.ConsoleClient
	url    string
	token  string
}

func (c *consoleClient) Client() consoleclient.ConsoleClient {
	return c.client
}

type ConsoleClient interface {
	Client() consoleclient.ConsoleClient
	Url() string
	Token() string
	ListClusters() (*consoleclient.ListClusters, error)
	GetCluster(clusterId, clusterName *string) (*consoleclient.ClusterFragment, error)
	GetDeployToken(clusterId, clusterName *string) (string, error)
	UpdateCluster(id string, attr consoleclient.ClusterUpdateAttributes) (*consoleclient.UpdateCluster, error)
	DeleteCluster(id string) error
	DetachCluster(id string) error
	ListClusterServices(clusterId, handle *string) ([]*consoleclient.ServiceDeploymentEdgeFragment, error)
	CreateRepository(url string, privateKey, passphrase, username, password *string) (*consoleclient.CreateGitRepository, error)
	ListRepositories() (*consoleclient.ListGitRepositories, error)
	UpdateRepository(id string, attrs consoleclient.GitAttributes) (*consoleclient.UpdateGitRepository, error)
	CreateClusterService(clusterId, clusterName *string, attr consoleclient.ServiceDeploymentAttributes) (*consoleclient.ServiceDeploymentExtended, error)
	UpdateClusterService(serviceId, serviceName, clusterName *string, attributes consoleclient.ServiceUpdateAttributes) (*consoleclient.ServiceDeploymentExtended, error)
	CloneService(clusterId string, serviceId, serviceName, clusterName *string, attributes consoleclient.ServiceCloneAttributes) (*consoleclient.ServiceDeploymentFragment, error)
	GetClusterService(serviceId, serviceName, clusterName *string) (*consoleclient.ServiceDeploymentExtended, error)
	DeleteClusterService(serviceId string) (*consoleclient.DeleteServiceDeployment, error)
	ListProviders() (*consoleclient.ListProviders, error)
	CreateProviderCredentials(name string, attr consoleclient.ProviderCredentialAttributes) (*consoleclient.CreateProviderCredential, error)
	DeleteProviderCredentials(id string) (*consoleclient.DeleteProviderCredential, error)
	SavePipeline(name string, attrs consoleclient.PipelineAttributes) (*consoleclient.PipelineFragment, error)
	CreateCluster(attributes consoleclient.ClusterAttributes) (*consoleclient.CreateCluster, error)
	CreateProvider(attr consoleclient.ClusterProviderAttributes) (*consoleclient.CreateClusterProvider, error)
	MyCluster() (*consoleclient.MyCluster, error)
	SaveServiceContext(name string, attributes consoleclient.ServiceContextAttributes) (*consoleclient.ServiceContextFragment, error)
	GetServiceContext(name string) (*consoleclient.ServiceContextFragment, error)
	KickClusterService(serviceId, serviceName, clusterName *string) (*consoleclient.ServiceDeploymentExtended, error)
	ListNotificationSinks(after *string, first *int64) (*consoleclient.ListNotificationSinks_NotificationSinks, error)
	CreateNotificationSinks(attr consoleclient.NotificationSinkAttributes) (*consoleclient.NotificationSinkFragment, error)
	UpdateDeploymentSettings(attr consoleclient.DeploymentSettingsAttributes) (*consoleclient.UpdateDeploymentSettings, error)
	GetGlobalSettings() (*consoleclient.DeploymentSettingsFragment, error)
}

type authedTransport struct {
	token   string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Token "+t.token)
	return t.wrapped.RoundTrip(req)
}

func NewConsoleClient(token, url string) (ConsoleClient, error) {
	httpClient := http.Client{
		Transport: &authedTransport{
			token:   token,
			wrapped: http.DefaultTransport,
		},
	}

	return &consoleClient{
		url:    url,
		token:  token,
		client: consoleclient.NewClient(&httpClient, NormalizeUrl(url), nil),
		ctx:    context.Background(),
	}, nil
}

func NormalizeUrl(url string) string {
	if !strings.HasPrefix(url, "https://") {
		url = fmt.Sprintf("https://%s", url)
	}
	if !strings.HasSuffix(url, "/gql") {
		url = fmt.Sprintf("%s/gql", url)
	}

	return url
}

func (c *consoleClient) Url() string {
	return c.url
}

func (c *consoleClient) Token() string {
	return c.token
}
