// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	gqlclient "github.com/pluralsh/console-client-go"
	mock "github.com/stretchr/testify/mock"
)

// ConsoleClient is an autogenerated mock type for the ConsoleClient type
type ConsoleClient struct {
	mock.Mock
}

// CreateCluster provides a mock function with given fields: attributes
func (_m *ConsoleClient) CreateCluster(attributes gqlclient.ClusterAttributes) (*gqlclient.CreateCluster, error) {
	ret := _m.Called(attributes)

	var r0 *gqlclient.CreateCluster
	var r1 error
	if rf, ok := ret.Get(0).(func(gqlclient.ClusterAttributes) (*gqlclient.CreateCluster, error)); ok {
		return rf(attributes)
	}
	if rf, ok := ret.Get(0).(func(gqlclient.ClusterAttributes) *gqlclient.CreateCluster); ok {
		r0 = rf(attributes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlclient.CreateCluster)
		}
	}

	if rf, ok := ret.Get(1).(func(gqlclient.ClusterAttributes) error); ok {
		r1 = rf(attributes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateClusterService provides a mock function with given fields: clusterId, clusterName, attr
func (_m *ConsoleClient) CreateClusterService(clusterId *string, clusterName *string, attr gqlclient.ServiceDeploymentAttributes) (*gqlclient.ServiceDeploymentFragment, error) {
	ret := _m.Called(clusterId, clusterName, attr)

	var r0 *gqlclient.ServiceDeploymentFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, *string, gqlclient.ServiceDeploymentAttributes) (*gqlclient.ServiceDeploymentFragment, error)); ok {
		return rf(clusterId, clusterName, attr)
	}
	if rf, ok := ret.Get(0).(func(*string, *string, gqlclient.ServiceDeploymentAttributes) *gqlclient.ServiceDeploymentFragment); ok {
		r0 = rf(clusterId, clusterName, attr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlclient.ServiceDeploymentFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, *string, gqlclient.ServiceDeploymentAttributes) error); ok {
		r1 = rf(clusterId, clusterName, attr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProvider provides a mock function with given fields: attr
func (_m *ConsoleClient) CreateProvider(attr gqlclient.ClusterProviderAttributes) (*gqlclient.CreateClusterProvider, error) {
	ret := _m.Called(attr)

	var r0 *gqlclient.CreateClusterProvider
	var r1 error
	if rf, ok := ret.Get(0).(func(gqlclient.ClusterProviderAttributes) (*gqlclient.CreateClusterProvider, error)); ok {
		return rf(attr)
	}
	if rf, ok := ret.Get(0).(func(gqlclient.ClusterProviderAttributes) *gqlclient.CreateClusterProvider); ok {
		r0 = rf(attr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlclient.CreateClusterProvider)
		}
	}

	if rf, ok := ret.Get(1).(func(gqlclient.ClusterProviderAttributes) error); ok {
		r1 = rf(attr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateProviderCredentials provides a mock function with given fields: name, attr
func (_m *ConsoleClient) CreateProviderCredentials(name string, attr gqlclient.ProviderCredentialAttributes) (*gqlclient.CreateProviderCredential, error) {
	ret := _m.Called(name, attr)

	var r0 *gqlclient.CreateProviderCredential
	var r1 error
	if rf, ok := ret.Get(0).(func(string, gqlclient.ProviderCredentialAttributes) (*gqlclient.CreateProviderCredential, error)); ok {
		return rf(name, attr)
	}
	if rf, ok := ret.Get(0).(func(string, gqlclient.ProviderCredentialAttributes) *gqlclient.CreateProviderCredential); ok {
		r0 = rf(name, attr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlclient.CreateProviderCredential)
		}
	}

	if rf, ok := ret.Get(1).(func(string, gqlclient.ProviderCredentialAttributes) error); ok {
		r1 = rf(name, attr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRepository provides a mock function with given fields: url, privateKey, passphrase, username, password
func (_m *ConsoleClient) CreateRepository(url string, privateKey *string, passphrase *string, username *string, password *string) (*gqlclient.CreateGitRepository, error) {
	ret := _m.Called(url, privateKey, passphrase, username, password)

	var r0 *gqlclient.CreateGitRepository
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *string, *string, *string, *string) (*gqlclient.CreateGitRepository, error)); ok {
		return rf(url, privateKey, passphrase, username, password)
	}
	if rf, ok := ret.Get(0).(func(string, *string, *string, *string, *string) *gqlclient.CreateGitRepository); ok {
		r0 = rf(url, privateKey, passphrase, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlclient.CreateGitRepository)
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

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteClusterService provides a mock function with given fields: serviceId
func (_m *ConsoleClient) DeleteClusterService(serviceId string) (*gqlclient.DeleteServiceDeployment, error) {
	ret := _m.Called(serviceId)

	var r0 *gqlclient.DeleteServiceDeployment
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*gqlclient.DeleteServiceDeployment, error)); ok {
		return rf(serviceId)
	}
	if rf, ok := ret.Get(0).(func(string) *gqlclient.DeleteServiceDeployment); ok {
		r0 = rf(serviceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlclient.DeleteServiceDeployment)
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
func (_m *ConsoleClient) DeleteProviderCredentials(id string) (*gqlclient.DeleteProviderCredential, error) {
	ret := _m.Called(id)

	var r0 *gqlclient.DeleteProviderCredential
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*gqlclient.DeleteProviderCredential, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *gqlclient.DeleteProviderCredential); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlclient.DeleteProviderCredential)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCluster provides a mock function with given fields: clusterId, clusterName
func (_m *ConsoleClient) GetCluster(clusterId *string, clusterName *string) (*gqlclient.ClusterFragment, error) {
	ret := _m.Called(clusterId, clusterName)

	var r0 *gqlclient.ClusterFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, *string) (*gqlclient.ClusterFragment, error)); ok {
		return rf(clusterId, clusterName)
	}
	if rf, ok := ret.Get(0).(func(*string, *string) *gqlclient.ClusterFragment); ok {
		r0 = rf(clusterId, clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlclient.ClusterFragment)
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
func (_m *ConsoleClient) GetClusterService(serviceId *string, serviceName *string, clusterName *string) (*gqlclient.ServiceDeploymentExtended, error) {
	ret := _m.Called(serviceId, serviceName, clusterName)

	var r0 *gqlclient.ServiceDeploymentExtended
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, *string, *string) (*gqlclient.ServiceDeploymentExtended, error)); ok {
		return rf(serviceId, serviceName, clusterName)
	}
	if rf, ok := ret.Get(0).(func(*string, *string, *string) *gqlclient.ServiceDeploymentExtended); ok {
		r0 = rf(serviceId, serviceName, clusterName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlclient.ServiceDeploymentExtended)
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
func (_m *ConsoleClient) ListClusterServices(clusterId *string, handle *string) ([]*gqlclient.ServiceDeploymentEdgeFragment, error) {
	ret := _m.Called(clusterId, handle)

	var r0 []*gqlclient.ServiceDeploymentEdgeFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, *string) ([]*gqlclient.ServiceDeploymentEdgeFragment, error)); ok {
		return rf(clusterId, handle)
	}
	if rf, ok := ret.Get(0).(func(*string, *string) []*gqlclient.ServiceDeploymentEdgeFragment); ok {
		r0 = rf(clusterId, handle)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gqlclient.ServiceDeploymentEdgeFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, *string) error); ok {
		r1 = rf(clusterId, handle)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListClusters provides a mock function with given fields:
func (_m *ConsoleClient) ListClusters() (*gqlclient.ListClusters, error) {
	ret := _m.Called()

	var r0 *gqlclient.ListClusters
	var r1 error
	if rf, ok := ret.Get(0).(func() (*gqlclient.ListClusters, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *gqlclient.ListClusters); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlclient.ListClusters)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProviders provides a mock function with given fields:
func (_m *ConsoleClient) ListProviders() (*gqlclient.ListProviders, error) {
	ret := _m.Called()

	var r0 *gqlclient.ListProviders
	var r1 error
	if rf, ok := ret.Get(0).(func() (*gqlclient.ListProviders, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *gqlclient.ListProviders); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlclient.ListProviders)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRepositories provides a mock function with given fields:
func (_m *ConsoleClient) ListRepositories() (*gqlclient.ListGitRepositories, error) {
	ret := _m.Called()

	var r0 *gqlclient.ListGitRepositories
	var r1 error
	if rf, ok := ret.Get(0).(func() (*gqlclient.ListGitRepositories, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *gqlclient.ListGitRepositories); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlclient.ListGitRepositories)
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
func (_m *ConsoleClient) SavePipeline(name string, attrs gqlclient.PipelineAttributes) (*gqlclient.PipelineFragment, error) {
	ret := _m.Called(name, attrs)

	var r0 *gqlclient.PipelineFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(string, gqlclient.PipelineAttributes) (*gqlclient.PipelineFragment, error)); ok {
		return rf(name, attrs)
	}
	if rf, ok := ret.Get(0).(func(string, gqlclient.PipelineAttributes) *gqlclient.PipelineFragment); ok {
		r0 = rf(name, attrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlclient.PipelineFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(string, gqlclient.PipelineAttributes) error); ok {
		r1 = rf(name, attrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCluster provides a mock function with given fields: id, attr
func (_m *ConsoleClient) UpdateCluster(id string, attr gqlclient.ClusterUpdateAttributes) (*gqlclient.UpdateCluster, error) {
	ret := _m.Called(id, attr)

	var r0 *gqlclient.UpdateCluster
	var r1 error
	if rf, ok := ret.Get(0).(func(string, gqlclient.ClusterUpdateAttributes) (*gqlclient.UpdateCluster, error)); ok {
		return rf(id, attr)
	}
	if rf, ok := ret.Get(0).(func(string, gqlclient.ClusterUpdateAttributes) *gqlclient.UpdateCluster); ok {
		r0 = rf(id, attr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlclient.UpdateCluster)
		}
	}

	if rf, ok := ret.Get(1).(func(string, gqlclient.ClusterUpdateAttributes) error); ok {
		r1 = rf(id, attr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateClusterService provides a mock function with given fields: serviceId, serviceName, clusterName, attributes
func (_m *ConsoleClient) UpdateClusterService(serviceId *string, serviceName *string, clusterName *string, attributes gqlclient.ServiceUpdateAttributes) (*gqlclient.ServiceDeploymentFragment, error) {
	ret := _m.Called(serviceId, serviceName, clusterName, attributes)

	var r0 *gqlclient.ServiceDeploymentFragment
	var r1 error
	if rf, ok := ret.Get(0).(func(*string, *string, *string, gqlclient.ServiceUpdateAttributes) (*gqlclient.ServiceDeploymentFragment, error)); ok {
		return rf(serviceId, serviceName, clusterName, attributes)
	}
	if rf, ok := ret.Get(0).(func(*string, *string, *string, gqlclient.ServiceUpdateAttributes) *gqlclient.ServiceDeploymentFragment); ok {
		r0 = rf(serviceId, serviceName, clusterName, attributes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlclient.ServiceDeploymentFragment)
		}
	}

	if rf, ok := ret.Get(1).(func(*string, *string, *string, gqlclient.ServiceUpdateAttributes) error); ok {
		r1 = rf(serviceId, serviceName, clusterName, attributes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRepository provides a mock function with given fields: id, attrs
func (_m *ConsoleClient) UpdateRepository(id string, attrs gqlclient.GitAttributes) (*gqlclient.UpdateGitRepository, error) {
	ret := _m.Called(id, attrs)

	var r0 *gqlclient.UpdateGitRepository
	var r1 error
	if rf, ok := ret.Get(0).(func(string, gqlclient.GitAttributes) (*gqlclient.UpdateGitRepository, error)); ok {
		return rf(id, attrs)
	}
	if rf, ok := ret.Get(0).(func(string, gqlclient.GitAttributes) *gqlclient.UpdateGitRepository); ok {
		r0 = rf(id, attrs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gqlclient.UpdateGitRepository)
		}
	}

	if rf, ok := ret.Get(1).(func(string, gqlclient.GitAttributes) error); ok {
		r1 = rf(id, attrs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Url provides a mock function with given fields:
func (_m *ConsoleClient) Url() string {
	ret := _m.Called()

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
