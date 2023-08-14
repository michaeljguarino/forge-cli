// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	client_gokubernetes "k8s.io/client-go/kubernetes"

	mock "github.com/stretchr/testify/mock"

	rest "k8s.io/client-go/rest"

	v1 "k8s.io/api/core/v1"

	v1alpha1 "github.com/pluralsh/plural-operator/apis/platform/v1alpha1"

	vpnv1alpha1 "github.com/pluralsh/plural-operator/apis/vpn/v1alpha1"
)

// Kube is an autogenerated mock type for the Kube type
type Kube struct {
	mock.Mock
}

// Apply provides a mock function with given fields: path, force
func (_m *Kube) Apply(path string, force bool) error {
	ret := _m.Called(path, force)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool) error); ok {
		r0 = rf(path, force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FinalizeNamespace provides a mock function with given fields: namespace
func (_m *Kube) FinalizeNamespace(namespace string) error {
	ret := _m.Called(namespace)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(namespace)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetClient provides a mock function with given fields:
func (_m *Kube) GetClient() *client_gokubernetes.Clientset {
	ret := _m.Called()

	var r0 *client_gokubernetes.Clientset
	if rf, ok := ret.Get(0).(func() *client_gokubernetes.Clientset); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client_gokubernetes.Clientset)
		}
	}

	return r0
}

// GetRestClient provides a mock function with given fields:
func (_m *Kube) GetRestClient() *rest.RESTClient {
	ret := _m.Called()

	var r0 *rest.RESTClient
	if rf, ok := ret.Get(0).(func() *rest.RESTClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rest.RESTClient)
		}
	}

	return r0
}

// LogTail provides a mock function with given fields: namespace, name
func (_m *Kube) LogTail(namespace string, name string) (*v1alpha1.LogTail, error) {
	ret := _m.Called(namespace, name)

	var r0 *v1alpha1.LogTail
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*v1alpha1.LogTail, error)); ok {
		return rf(namespace, name)
	}
	if rf, ok := ret.Get(0).(func(string, string) *v1alpha1.LogTail); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.LogTail)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LogTailList provides a mock function with given fields: namespace
func (_m *Kube) LogTailList(namespace string) (*v1alpha1.LogTailList, error) {
	ret := _m.Called(namespace)

	var r0 *v1alpha1.LogTailList
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*v1alpha1.LogTailList, error)); ok {
		return rf(namespace)
	}
	if rf, ok := ret.Get(0).(func(string) *v1alpha1.LogTailList); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.LogTailList)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Node provides a mock function with given fields: name
func (_m *Kube) Node(name string) (*v1.Node, error) {
	ret := _m.Called(name)

	var r0 *v1.Node
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*v1.Node, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *v1.Node); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Node)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Nodes provides a mock function with given fields:
func (_m *Kube) Nodes() (*v1.NodeList, error) {
	ret := _m.Called()

	var r0 *v1.NodeList
	var r1 error
	if rf, ok := ret.Get(0).(func() (*v1.NodeList, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *v1.NodeList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.NodeList)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Proxy provides a mock function with given fields: namespace, name
func (_m *Kube) Proxy(namespace string, name string) (*v1alpha1.Proxy, error) {
	ret := _m.Called(namespace, name)

	var r0 *v1alpha1.Proxy
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*v1alpha1.Proxy, error)); ok {
		return rf(namespace, name)
	}
	if rf, ok := ret.Get(0).(func(string, string) *v1alpha1.Proxy); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.Proxy)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProxyList provides a mock function with given fields: namespace
func (_m *Kube) ProxyList(namespace string) (*v1alpha1.ProxyList, error) {
	ret := _m.Called(namespace)

	var r0 *v1alpha1.ProxyList
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*v1alpha1.ProxyList, error)); ok {
		return rf(namespace)
	}
	if rf, ok := ret.Get(0).(func(string) *v1alpha1.ProxyList); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1alpha1.ProxyList)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Secret provides a mock function with given fields: namespace, name
func (_m *Kube) Secret(namespace string, name string) (*v1.Secret, error) {
	ret := _m.Called(namespace, name)

	var r0 *v1.Secret
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*v1.Secret, error)); ok {
		return rf(namespace, name)
	}
	if rf, ok := ret.Get(0).(func(string, string) *v1.Secret); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Secret)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WireguardPeer provides a mock function with given fields: namespace, name
func (_m *Kube) WireguardPeer(namespace string, name string) (*vpnv1alpha1.WireguardPeer, error) {
	ret := _m.Called(namespace, name)

	var r0 *vpnv1alpha1.WireguardPeer
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*vpnv1alpha1.WireguardPeer, error)); ok {
		return rf(namespace, name)
	}
	if rf, ok := ret.Get(0).(func(string, string) *vpnv1alpha1.WireguardPeer); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vpnv1alpha1.WireguardPeer)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WireguardPeerCreate provides a mock function with given fields: namespace, wireguardPeer
func (_m *Kube) WireguardPeerCreate(namespace string, wireguardPeer *vpnv1alpha1.WireguardPeer) (*vpnv1alpha1.WireguardPeer, error) {
	ret := _m.Called(namespace, wireguardPeer)

	var r0 *vpnv1alpha1.WireguardPeer
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *vpnv1alpha1.WireguardPeer) (*vpnv1alpha1.WireguardPeer, error)); ok {
		return rf(namespace, wireguardPeer)
	}
	if rf, ok := ret.Get(0).(func(string, *vpnv1alpha1.WireguardPeer) *vpnv1alpha1.WireguardPeer); ok {
		r0 = rf(namespace, wireguardPeer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vpnv1alpha1.WireguardPeer)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *vpnv1alpha1.WireguardPeer) error); ok {
		r1 = rf(namespace, wireguardPeer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WireguardPeerDelete provides a mock function with given fields: namespace, name
func (_m *Kube) WireguardPeerDelete(namespace string, name string) error {
	ret := _m.Called(namespace, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(namespace, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WireguardPeerList provides a mock function with given fields: namespace
func (_m *Kube) WireguardPeerList(namespace string) (*vpnv1alpha1.WireguardPeerList, error) {
	ret := _m.Called(namespace)

	var r0 *vpnv1alpha1.WireguardPeerList
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*vpnv1alpha1.WireguardPeerList, error)); ok {
		return rf(namespace)
	}
	if rf, ok := ret.Get(0).(func(string) *vpnv1alpha1.WireguardPeerList); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vpnv1alpha1.WireguardPeerList)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WireguardServer provides a mock function with given fields: namespace, name
func (_m *Kube) WireguardServer(namespace string, name string) (*vpnv1alpha1.WireguardServer, error) {
	ret := _m.Called(namespace, name)

	var r0 *vpnv1alpha1.WireguardServer
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*vpnv1alpha1.WireguardServer, error)); ok {
		return rf(namespace, name)
	}
	if rf, ok := ret.Get(0).(func(string, string) *vpnv1alpha1.WireguardServer); ok {
		r0 = rf(namespace, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vpnv1alpha1.WireguardServer)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(namespace, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WireguardServerList provides a mock function with given fields: namespace
func (_m *Kube) WireguardServerList(namespace string) (*vpnv1alpha1.WireguardServerList, error) {
	ret := _m.Called(namespace)

	var r0 *vpnv1alpha1.WireguardServerList
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*vpnv1alpha1.WireguardServerList, error)); ok {
		return rf(namespace)
	}
	if rf, ok := ret.Get(0).(func(string) *vpnv1alpha1.WireguardServerList); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*vpnv1alpha1.WireguardServerList)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewKube creates a new instance of Kube. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewKube(t interface {
	mock.TestingT
	Cleanup(func())
}) *Kube {
	mock := &Kube{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
