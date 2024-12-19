// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	client "github.com/pluralsh/console/go/client"

	mock "github.com/stretchr/testify/mock"
)

// ConsoleClient is an autogenerated mock type for the ConsoleClient type
type ConsoleClient struct {
	mock.Mock
}

// AgentUrl provides a mock function with given fields: id
func (_m *ConsoleClient) AgentUrl(id string) (string, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for AgentUrl")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CloneService provides a mock function with given fields: clusterId, serviceId, serviceName, clusterName, attributes
func (_m *ConsoleClient) CloneService(clusterId string, serviceId *string, serviceName *string, clusterName *string, attributes client.ServiceCloneAttributes) (*client.ServiceDeploymentFragment, error) {
	ret := _m.Called(clusterId, serviceId, serviceName, clusterName, attributes)

	if len(ret) == 0 {
		panic("no return value specified for CloneService")
	}

	var r0 *client.ServiceDeploymentFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *string, *string, *string, client.ServiceCloneAttributes) (*client.ServiceDeploymentFragment, error)); ok {
		return rf(clusterId, serviceId, serviceName, clusterName, attributes)
	}
	if rf, ok := ret.Get(0).(func(string, *string, *string, *string, client.ServiceCloneAttributes) *client.ServiceDeploymentFragment); ok {
		r0 = rf(clusterId, serviceId, serviceName, clusterName, attributes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ServiceDeploymentFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *string, *string, *string, client.ServiceCloneAttributes) error); ok {
		r1 = rf(clusterId, serviceId, serviceName, clusterName, attributes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateCluster provides a mock function with given fields: attributes
func (_m *ConsoleClient) CreateCluster(attributes client.ClusterAttributes) (*client.CreateCluster, error) {
	ret := _m.Called(attributes)

	if len(ret) == 0 {
		panic("no return value specified for CreateCluster")
	}

	var r0 *client.CreateCluster
	var r1 error
	if rf, ok := ret.Get(0).(func(client.ClusterAttributes) (*client.CreateCluster, error)); ok {
		return rf(attributes)
	}
	if rf, ok := ret.Get(0).(func(client.ClusterAttributes) *client.CreateCluster); ok {
		r0 = rf(attributes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.CreateCluster)
		}
	}

	if rf, ok := ret.Get(1).(func(client.ClusterAttributes) error); ok {
		r1 = rf(attributes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateClusterService provides a mock function with given fields: clusterId, clusterName, attr
func (_m *ConsoleClient) CreateClusterService(clusterId *string, clusterName *string, attr client.ServiceDeploymentAttributes) (*client.ServiceDeploymentExtended, error) {
	ret := _m.Called(clusterId, clusterName, attr)

	if len(ret) == 0 {
		panic("no return value specified for CreateClusterService")
	}

	var r0 *client.ServiceDeploymentExtended
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, *string, client.ServiceDeploymentAttributes) (*client.ServiceDeploymentExtended, error)); ok {
		return rf(clusterId, clusterName, attr)
	}
	if rf, ok := ret.Get(0).(func(*string, *string, client.ServiceDeploymentAttributes) *client.ServiceDeploymentExtended); ok {
		r0 = rf(clusterId, clusterName, attr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ServiceDeploymentExtended)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, *string, client.ServiceDeploymentAttributes) error); ok {
		r1 = rf(clusterId, clusterName, attr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateNotificationSinks provides a mock function with given fields: attr
func (_m *ConsoleClient) CreateNotificationSinks(attr client.NotificationSinkAttributes) (*client.NotificationSinkFragment, error) {
	ret := _m.Called(attr)

	if len(ret) == 0 {
		panic("no return value specified for CreateNotificationSinks")
	}

	var r0 *client.NotificationSinkFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(client.NotificationSinkAttributes) (*client.NotificationSinkFragment, error)); ok {
		return rf(attr)
	}
	if rf, ok := ret.Get(0).(func(client.NotificationSinkAttributes) *client.NotificationSinkFragment); ok {
		r0 = rf(attr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.NotificationSinkFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(client.NotificationSinkAttributes) error); ok {
		r1 = rf(attr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePipelineContext provides a mock function with given fields: id, attrs
func (_m *ConsoleClient) CreatePipelineContext(id string, attrs client.PipelineContextAttributes) (*client.PipelineContextFragment, error) {
	ret := _m.Called(id, attrs)

	if len(ret) == 0 {
		panic("no return value specified for CreatePipelineContext")
	}

	var r0 *client.PipelineContextFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(string, client.PipelineContextAttributes) (*client.PipelineContextFragment, error)); ok {
		return rf(id, attrs)
	}
	if rf, ok := ret.Get(0).(func(string, client.PipelineContextAttributes) *client.PipelineContextFragment); ok {
		r0 = rf(id, attrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.PipelineContextFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(string, client.PipelineContextAttributes) error); ok {
		r1 = rf(id, attrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProvider provides a mock function with given fields: attr
func (_m *ConsoleClient) CreateProvider(attr client.ClusterProviderAttributes) (*client.CreateClusterProvider, error) {
	ret := _m.Called(attr)

	if len(ret) == 0 {
		panic("no return value specified for CreateProvider")
	}

	var r0 *client.CreateClusterProvider
	var r1 error
	if rf, ok := ret.Get(0).(func(client.ClusterProviderAttributes) (*client.CreateClusterProvider, error)); ok {
		return rf(attr)
	}
	if rf, ok := ret.Get(0).(func(client.ClusterProviderAttributes) *client.CreateClusterProvider); ok {
		r0 = rf(attr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.CreateClusterProvider)
		}
	}

	if rf, ok := ret.Get(1).(func(client.ClusterProviderAttributes) error); ok {
		r1 = rf(attr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProviderCredentials provides a mock function with given fields: name, attr
func (_m *ConsoleClient) CreateProviderCredentials(name string, attr client.ProviderCredentialAttributes) (*client.CreateProviderCredential, error) {
	ret := _m.Called(name, attr)

	if len(ret) == 0 {
		panic("no return value specified for CreateProviderCredentials")
	}

	var r0 *client.CreateProviderCredential
	var r1 error
	if rf, ok := ret.Get(0).(func(string, client.ProviderCredentialAttributes) (*client.CreateProviderCredential, error)); ok {
		return rf(name, attr)
	}
	if rf, ok := ret.Get(0).(func(string, client.ProviderCredentialAttributes) *client.CreateProviderCredential); ok {
		r0 = rf(name, attr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.CreateProviderCredential)
		}
	}

	if rf, ok := ret.Get(1).(func(string, client.ProviderCredentialAttributes) error); ok {
		r1 = rf(name, attr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePullRequest provides a mock function with given fields: id, branch, context
func (_m *ConsoleClient) CreatePullRequest(id string, branch *string, context *string) (*client.PullRequestFragment, error) {
	ret := _m.Called(id, branch, context)

	if len(ret) == 0 {
		panic("no return value specified for CreatePullRequest")
	}

	var r0 *client.PullRequestFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *string, *string) (*client.PullRequestFragment, error)); ok {
		return rf(id, branch, context)
	}
	if rf, ok := ret.Get(0).(func(string, *string, *string) *client.PullRequestFragment); ok {
		r0 = rf(id, branch, context)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.PullRequestFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *string, *string) error); ok {
		r1 = rf(id, branch, context)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRepository provides a mock function with given fields: url, privateKey, passphrase, username, password
func (_m *ConsoleClient) CreateRepository(url string, privateKey *string, passphrase *string, username *string, password *string) (*client.CreateGitRepository, error) {
	ret := _m.Called(url, privateKey, passphrase, username, password)

	if len(ret) == 0 {
		panic("no return value specified for CreateRepository")
	}

	var r0 *client.CreateGitRepository
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *string, *string, *string, *string) (*client.CreateGitRepository, error)); ok {
		return rf(url, privateKey, passphrase, username, password)
	}
	if rf, ok := ret.Get(0).(func(string, *string, *string, *string, *string) *client.CreateGitRepository); ok {
		r0 = rf(url, privateKey, passphrase, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.CreateGitRepository)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *string, *string, *string, *string) error); ok {
		r1 = rf(url, privateKey, passphrase, username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCluster provides a mock function with given fields: id
func (_m *ConsoleClient) DeleteCluster(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteCluster")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteClusterService provides a mock function with given fields: serviceId
func (_m *ConsoleClient) DeleteClusterService(serviceId string) (*client.DeleteServiceDeployment, error) {
	ret := _m.Called(serviceId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteClusterService")
	}

	var r0 *client.DeleteServiceDeployment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*client.DeleteServiceDeployment, error)); ok {
		return rf(serviceId)
	}
	if rf, ok := ret.Get(0).(func(string) *client.DeleteServiceDeployment); ok {
		r0 = rf(serviceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.DeleteServiceDeployment)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(serviceId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProviderCredentials provides a mock function with given fields: id
func (_m *ConsoleClient) DeleteProviderCredentials(id string) (*client.DeleteProviderCredential, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProviderCredentials")
	}

	var r0 *client.DeleteProviderCredential
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*client.DeleteProviderCredential, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *client.DeleteProviderCredential); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.DeleteProviderCredential)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DetachCluster provides a mock function with given fields: id
func (_m *ConsoleClient) DetachCluster(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DetachCluster")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ExtUrl provides a mock function with no fields
func (_m *ConsoleClient) ExtUrl() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ExtUrl")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetCluster provides a mock function with given fields: clusterId, clusterName
func (_m *ConsoleClient) GetCluster(clusterId *string, clusterName *string) (*client.ClusterFragment, error) {
	ret := _m.Called(clusterId, clusterName)

	if len(ret) == 0 {
		panic("no return value specified for GetCluster")
	}

	var r0 *client.ClusterFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, *string) (*client.ClusterFragment, error)); ok {
		return rf(clusterId, clusterName)
	}
	if rf, ok := ret.Get(0).(func(*string, *string) *client.ClusterFragment); ok {
		r0 = rf(clusterId, clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ClusterFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, *string) error); ok {
		r1 = rf(clusterId, clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClusterService provides a mock function with given fields: serviceId, serviceName, clusterName
func (_m *ConsoleClient) GetClusterService(serviceId *string, serviceName *string, clusterName *string) (*client.ServiceDeploymentExtended, error) {
	ret := _m.Called(serviceId, serviceName, clusterName)

	if len(ret) == 0 {
		panic("no return value specified for GetClusterService")
	}

	var r0 *client.ServiceDeploymentExtended
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, *string, *string) (*client.ServiceDeploymentExtended, error)); ok {
		return rf(serviceId, serviceName, clusterName)
	}
	if rf, ok := ret.Get(0).(func(*string, *string, *string) *client.ServiceDeploymentExtended); ok {
		r0 = rf(serviceId, serviceName, clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ServiceDeploymentExtended)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, *string, *string) error); ok {
		r1 = rf(serviceId, serviceName, clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeployToken provides a mock function with given fields: clusterId, clusterName
func (_m *ConsoleClient) GetDeployToken(clusterId *string, clusterName *string) (string, error) {
	ret := _m.Called(clusterId, clusterName)

	if len(ret) == 0 {
		panic("no return value specified for GetDeployToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, *string) (string, error)); ok {
		return rf(clusterId, clusterName)
	}
	if rf, ok := ret.Get(0).(func(*string, *string) string); ok {
		r0 = rf(clusterId, clusterName)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*string, *string) error); ok {
		r1 = rf(clusterId, clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGlobalSettings provides a mock function with no fields
func (_m *ConsoleClient) GetGlobalSettings() (*client.DeploymentSettingsFragment, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetGlobalSettings")
	}

	var r0 *client.DeploymentSettingsFragment
	var r1 error
	if rf, ok := ret.Get(0).(func() (*client.DeploymentSettingsFragment, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *client.DeploymentSettingsFragment); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.DeploymentSettingsFragment)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPipelineContext provides a mock function with given fields: id
func (_m *ConsoleClient) GetPipelineContext(id string) (*client.PipelineContextFragment, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetPipelineContext")
	}

	var r0 *client.PipelineContextFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*client.PipelineContextFragment, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *client.PipelineContextFragment); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.PipelineContextFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPrAutomationByName provides a mock function with given fields: name
func (_m *ConsoleClient) GetPrAutomationByName(name string) (*client.PrAutomationFragment, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetPrAutomationByName")
	}

	var r0 *client.PrAutomationFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*client.PrAutomationFragment, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *client.PrAutomationFragment); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.PrAutomationFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProject provides a mock function with given fields: name
func (_m *ConsoleClient) GetProject(name string) (*client.ProjectFragment, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetProject")
	}

	var r0 *client.ProjectFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*client.ProjectFragment, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *client.ProjectFragment); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ProjectFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRepository provides a mock function with given fields: id
func (_m *ConsoleClient) GetRepository(id string) (*client.GetGitRepository, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetRepository")
	}

	var r0 *client.GetGitRepository
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*client.GetGitRepository, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *client.GetGitRepository); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.GetGitRepository)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServiceContext provides a mock function with given fields: name
func (_m *ConsoleClient) GetServiceContext(name string) (*client.ServiceContextFragment, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetServiceContext")
	}

	var r0 *client.ServiceContextFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*client.ServiceContextFragment, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *client.ServiceContextFragment); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ServiceContextFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KickClusterService provides a mock function with given fields: serviceId, serviceName, clusterName
func (_m *ConsoleClient) KickClusterService(serviceId *string, serviceName *string, clusterName *string) (*client.ServiceDeploymentExtended, error) {
	ret := _m.Called(serviceId, serviceName, clusterName)

	if len(ret) == 0 {
		panic("no return value specified for KickClusterService")
	}

	var r0 *client.ServiceDeploymentExtended
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, *string, *string) (*client.ServiceDeploymentExtended, error)); ok {
		return rf(serviceId, serviceName, clusterName)
	}
	if rf, ok := ret.Get(0).(func(*string, *string, *string) *client.ServiceDeploymentExtended); ok {
		r0 = rf(serviceId, serviceName, clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ServiceDeploymentExtended)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, *string, *string) error); ok {
		r1 = rf(serviceId, serviceName, clusterName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClusterServices provides a mock function with given fields: clusterId, handle
func (_m *ConsoleClient) ListClusterServices(clusterId *string, handle *string) ([]*client.ServiceDeploymentEdgeFragment, error) {
	ret := _m.Called(clusterId, handle)

	if len(ret) == 0 {
		panic("no return value specified for ListClusterServices")
	}

	var r0 []*client.ServiceDeploymentEdgeFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, *string) ([]*client.ServiceDeploymentEdgeFragment, error)); ok {
		return rf(clusterId, handle)
	}
	if rf, ok := ret.Get(0).(func(*string, *string) []*client.ServiceDeploymentEdgeFragment); ok {
		r0 = rf(clusterId, handle)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*client.ServiceDeploymentEdgeFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, *string) error); ok {
		r1 = rf(clusterId, handle)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClusters provides a mock function with no fields
func (_m *ConsoleClient) ListClusters() (*client.ListClusters, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListClusters")
	}

	var r0 *client.ListClusters
	var r1 error
	if rf, ok := ret.Get(0).(func() (*client.ListClusters, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *client.ListClusters); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ListClusters)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNotificationSinks provides a mock function with given fields: after, first
func (_m *ConsoleClient) ListNotificationSinks(after *string, first *int64) (*client.ListNotificationSinks_NotificationSinks, error) {
	ret := _m.Called(after, first)

	if len(ret) == 0 {
		panic("no return value specified for ListNotificationSinks")
	}

	var r0 *client.ListNotificationSinks_NotificationSinks
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, *int64) (*client.ListNotificationSinks_NotificationSinks, error)); ok {
		return rf(after, first)
	}
	if rf, ok := ret.Get(0).(func(*string, *int64) *client.ListNotificationSinks_NotificationSinks); ok {
		r0 = rf(after, first)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ListNotificationSinks_NotificationSinks)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, *int64) error); ok {
		r1 = rf(after, first)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProviders provides a mock function with no fields
func (_m *ConsoleClient) ListProviders() (*client.ListProviders, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListProviders")
	}

	var r0 *client.ListProviders
	var r1 error
	if rf, ok := ret.Get(0).(func() (*client.ListProviders, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *client.ListProviders); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ListProviders)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRepositories provides a mock function with no fields
func (_m *ConsoleClient) ListRepositories() (*client.ListGitRepositories, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListRepositories")
	}

	var r0 *client.ListGitRepositories
	var r1 error
	if rf, ok := ret.Get(0).(func() (*client.ListGitRepositories, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *client.ListGitRepositories); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ListGitRepositories)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListStackRuns provides a mock function with given fields: stackID
func (_m *ConsoleClient) ListStackRuns(stackID string) (*client.ListStackRuns, error) {
	ret := _m.Called(stackID)

	if len(ret) == 0 {
		panic("no return value specified for ListStackRuns")
	}

	var r0 *client.ListStackRuns
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*client.ListStackRuns, error)); ok {
		return rf(stackID)
	}
	if rf, ok := ret.Get(0).(func(string) *client.ListStackRuns); ok {
		r0 = rf(stackID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ListStackRuns)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(stackID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MyCluster provides a mock function with no fields
func (_m *ConsoleClient) MyCluster() (*client.MyCluster, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for MyCluster")
	}

	var r0 *client.MyCluster
	var r1 error
	if rf, ok := ret.Get(0).(func() (*client.MyCluster, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *client.MyCluster); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.MyCluster)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SavePipeline provides a mock function with given fields: name, attrs
func (_m *ConsoleClient) SavePipeline(name string, attrs client.PipelineAttributes) (*client.PipelineFragmentMinimal, error) {
	ret := _m.Called(name, attrs)

	if len(ret) == 0 {
		panic("no return value specified for SavePipeline")
	}

	var r0 *client.PipelineFragmentMinimal
	var r1 error
	if rf, ok := ret.Get(0).(func(string, client.PipelineAttributes) (*client.PipelineFragmentMinimal, error)); ok {
		return rf(name, attrs)
	}
	if rf, ok := ret.Get(0).(func(string, client.PipelineAttributes) *client.PipelineFragmentMinimal); ok {
		r0 = rf(name, attrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.PipelineFragmentMinimal)
		}
	}

	if rf, ok := ret.Get(1).(func(string, client.PipelineAttributes) error); ok {
		r1 = rf(name, attrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveServiceContext provides a mock function with given fields: name, attributes
func (_m *ConsoleClient) SaveServiceContext(name string, attributes client.ServiceContextAttributes) (*client.ServiceContextFragment, error) {
	ret := _m.Called(name, attributes)

	if len(ret) == 0 {
		panic("no return value specified for SaveServiceContext")
	}

	var r0 *client.ServiceContextFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(string, client.ServiceContextAttributes) (*client.ServiceContextFragment, error)); ok {
		return rf(name, attributes)
	}
	if rf, ok := ret.Get(0).(func(string, client.ServiceContextAttributes) *client.ServiceContextFragment); ok {
		r0 = rf(name, attributes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ServiceContextFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(string, client.ServiceContextAttributes) error); ok {
		r1 = rf(name, attributes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Token provides a mock function with no fields
func (_m *ConsoleClient) Token() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Token")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UpdateCluster provides a mock function with given fields: id, attr
func (_m *ConsoleClient) UpdateCluster(id string, attr client.ClusterUpdateAttributes) (*client.UpdateCluster, error) {
	ret := _m.Called(id, attr)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCluster")
	}

	var r0 *client.UpdateCluster
	var r1 error
	if rf, ok := ret.Get(0).(func(string, client.ClusterUpdateAttributes) (*client.UpdateCluster, error)); ok {
		return rf(id, attr)
	}
	if rf, ok := ret.Get(0).(func(string, client.ClusterUpdateAttributes) *client.UpdateCluster); ok {
		r0 = rf(id, attr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.UpdateCluster)
		}
	}

	if rf, ok := ret.Get(1).(func(string, client.ClusterUpdateAttributes) error); ok {
		r1 = rf(id, attr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateClusterService provides a mock function with given fields: serviceId, serviceName, clusterName, attributes
func (_m *ConsoleClient) UpdateClusterService(serviceId *string, serviceName *string, clusterName *string, attributes client.ServiceUpdateAttributes) (*client.ServiceDeploymentExtended, error) {
	ret := _m.Called(serviceId, serviceName, clusterName, attributes)

	if len(ret) == 0 {
		panic("no return value specified for UpdateClusterService")
	}

	var r0 *client.ServiceDeploymentExtended
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, *string, *string, client.ServiceUpdateAttributes) (*client.ServiceDeploymentExtended, error)); ok {
		return rf(serviceId, serviceName, clusterName, attributes)
	}
	if rf, ok := ret.Get(0).(func(*string, *string, *string, client.ServiceUpdateAttributes) *client.ServiceDeploymentExtended); ok {
		r0 = rf(serviceId, serviceName, clusterName, attributes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.ServiceDeploymentExtended)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, *string, *string, client.ServiceUpdateAttributes) error); ok {
		r1 = rf(serviceId, serviceName, clusterName, attributes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateDeploymentSettings provides a mock function with given fields: attr
func (_m *ConsoleClient) UpdateDeploymentSettings(attr client.DeploymentSettingsAttributes) (*client.UpdateDeploymentSettings, error) {
	ret := _m.Called(attr)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDeploymentSettings")
	}

	var r0 *client.UpdateDeploymentSettings
	var r1 error
	if rf, ok := ret.Get(0).(func(client.DeploymentSettingsAttributes) (*client.UpdateDeploymentSettings, error)); ok {
		return rf(attr)
	}
	if rf, ok := ret.Get(0).(func(client.DeploymentSettingsAttributes) *client.UpdateDeploymentSettings); ok {
		r0 = rf(attr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.UpdateDeploymentSettings)
		}
	}

	if rf, ok := ret.Get(1).(func(client.DeploymentSettingsAttributes) error); ok {
		r1 = rf(attr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRepository provides a mock function with given fields: id, attrs
func (_m *ConsoleClient) UpdateRepository(id string, attrs client.GitAttributes) (*client.UpdateGitRepository, error) {
	ret := _m.Called(id, attrs)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRepository")
	}

	var r0 *client.UpdateGitRepository
	var r1 error
	if rf, ok := ret.Get(0).(func(string, client.GitAttributes) (*client.UpdateGitRepository, error)); ok {
		return rf(id, attrs)
	}
	if rf, ok := ret.Get(0).(func(string, client.GitAttributes) *client.UpdateGitRepository); ok {
		r0 = rf(id, attrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.UpdateGitRepository)
		}
	}

	if rf, ok := ret.Get(1).(func(string, client.GitAttributes) error); ok {
		r1 = rf(id, attrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Url provides a mock function with no fields
func (_m *ConsoleClient) Url() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Url")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewConsoleClient creates a new instance of ConsoleClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewConsoleClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *ConsoleClient {
	mock := &ConsoleClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
