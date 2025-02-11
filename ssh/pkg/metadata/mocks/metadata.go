// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	internalclient "github.com/shellhub-io/shellhub/pkg/api/internalclient"
	metadata "github.com/shellhub-io/shellhub/ssh/pkg/metadata"
	cryptossh "golang.org/x/crypto/ssh"

	mock "github.com/stretchr/testify/mock"

	models "github.com/shellhub-io/shellhub/pkg/models"

	ssh "github.com/gliderlabs/ssh"

	target "github.com/shellhub-io/shellhub/ssh/pkg/target"
)

// Metadata is an autogenerated mock type for the Metadata type
type Metadata struct {
	mock.Mock
}

// MaybeSetAPI provides a mock function with given fields: ctx, client
func (_m *Metadata) MaybeSetAPI(ctx ssh.Context, client internalclient.Client) internalclient.Client {
	ret := _m.Called(ctx, client)

	var r0 internalclient.Client
	if rf, ok := ret.Get(0).(func(ssh.Context, internalclient.Client) internalclient.Client); ok {
		r0 = rf(ctx, client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(internalclient.Client)
		}
	}

	return r0
}

// MaybeStoreAgentConn provides a mock function with given fields: ctx, client
func (_m *Metadata) MaybeStoreAgentConn(ctx ssh.Context, client *cryptossh.Client) *cryptossh.Client {
	ret := _m.Called(ctx, client)

	var r0 *cryptossh.Client
	if rf, ok := ret.Get(0).(func(ssh.Context, *cryptossh.Client) *cryptossh.Client); ok {
		r0 = rf(ctx, client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cryptossh.Client)
		}
	}

	return r0
}

// MaybeStoreDevice provides a mock function with given fields: ctx, lookup, api
func (_m *Metadata) MaybeStoreDevice(ctx ssh.Context, lookup map[string]string, api internalclient.Client) (*models.Device, []error) {
	ret := _m.Called(ctx, lookup, api)

	var r0 *models.Device
	var r1 []error
	if rf, ok := ret.Get(0).(func(ssh.Context, map[string]string, internalclient.Client) (*models.Device, []error)); ok {
		return rf(ctx, lookup, api)
	}
	if rf, ok := ret.Get(0).(func(ssh.Context, map[string]string, internalclient.Client) *models.Device); ok {
		r0 = rf(ctx, lookup, api)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	if rf, ok := ret.Get(1).(func(ssh.Context, map[string]string, internalclient.Client) []error); ok {
		r1 = rf(ctx, lookup, api)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]error)
		}
	}

	return r0, r1
}

// MaybeStoreEstablished provides a mock function with given fields: ctx, value
func (_m *Metadata) MaybeStoreEstablished(ctx ssh.Context, value bool) bool {
	ret := _m.Called(ctx, value)

	var r0 bool
	if rf, ok := ret.Get(0).(func(ssh.Context, bool) bool); ok {
		r0 = rf(ctx, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MaybeStoreFingerprint provides a mock function with given fields: ctx, value
func (_m *Metadata) MaybeStoreFingerprint(ctx ssh.Context, value string) string {
	ret := _m.Called(ctx, value)

	var r0 string
	if rf, ok := ret.Get(0).(func(ssh.Context, string) string); ok {
		r0 = rf(ctx, value)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MaybeStoreLookup provides a mock function with given fields: ctx, tag, api
func (_m *Metadata) MaybeStoreLookup(ctx ssh.Context, tag *target.Target, api internalclient.Client) (map[string]string, error) {
	ret := _m.Called(ctx, tag, api)

	var r0 map[string]string
	var r1 error
	if rf, ok := ret.Get(0).(func(ssh.Context, *target.Target, internalclient.Client) (map[string]string, error)); ok {
		return rf(ctx, tag, api)
	}
	if rf, ok := ret.Get(0).(func(ssh.Context, *target.Target, internalclient.Client) map[string]string); ok {
		r0 = rf(ctx, tag, api)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	if rf, ok := ret.Get(1).(func(ssh.Context, *target.Target, internalclient.Client) error); ok {
		r1 = rf(ctx, tag, api)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaybeStoreSSHID provides a mock function with given fields: ctx, value
func (_m *Metadata) MaybeStoreSSHID(ctx ssh.Context, value string) string {
	ret := _m.Called(ctx, value)

	var r0 string
	if rf, ok := ret.Get(0).(func(ssh.Context, string) string); ok {
		r0 = rf(ctx, value)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MaybeStoreTarget provides a mock function with given fields: ctx, sshid
func (_m *Metadata) MaybeStoreTarget(ctx ssh.Context, sshid string) (*target.Target, error) {
	ret := _m.Called(ctx, sshid)

	var r0 *target.Target
	var r1 error
	if rf, ok := ret.Get(0).(func(ssh.Context, string) (*target.Target, error)); ok {
		return rf(ctx, sshid)
	}
	if rf, ok := ret.Get(0).(func(ssh.Context, string) *target.Target); ok {
		r0 = rf(ctx, sshid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*target.Target)
		}
	}

	if rf, ok := ret.Get(1).(func(ssh.Context, string) error); ok {
		r1 = rf(ctx, sshid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RestoreAPI provides a mock function with given fields: ctx
func (_m *Metadata) RestoreAPI(ctx ssh.Context) internalclient.Client {
	ret := _m.Called(ctx)

	var r0 internalclient.Client
	if rf, ok := ret.Get(0).(func(ssh.Context) internalclient.Client); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(internalclient.Client)
		}
	}

	return r0
}

// RestoreAgentConn provides a mock function with given fields: ctx
func (_m *Metadata) RestoreAgentConn(ctx ssh.Context) *cryptossh.Client {
	ret := _m.Called(ctx)

	var r0 *cryptossh.Client
	if rf, ok := ret.Get(0).(func(ssh.Context) *cryptossh.Client); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cryptossh.Client)
		}
	}

	return r0
}

// RestoreAuthenticationMethod provides a mock function with given fields: ctx
func (_m *Metadata) RestoreAuthenticationMethod(ctx ssh.Context) metadata.AuthenticationMethod {
	ret := _m.Called(ctx)

	var r0 metadata.AuthenticationMethod
	if rf, ok := ret.Get(0).(func(ssh.Context) metadata.AuthenticationMethod); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(metadata.AuthenticationMethod)
	}

	return r0
}

// RestoreDevice provides a mock function with given fields: ctx
func (_m *Metadata) RestoreDevice(ctx ssh.Context) *models.Device {
	ret := _m.Called(ctx)

	var r0 *models.Device
	if rf, ok := ret.Get(0).(func(ssh.Context) *models.Device); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	return r0
}

// RestoreEstablished provides a mock function with given fields: ctx
func (_m *Metadata) RestoreEstablished(ctx ssh.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(ssh.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RestoreFingerprint provides a mock function with given fields: ctx
func (_m *Metadata) RestoreFingerprint(ctx ssh.Context) string {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(ssh.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RestoreLookup provides a mock function with given fields: ctx
func (_m *Metadata) RestoreLookup(ctx ssh.Context) map[string]string {
	ret := _m.Called(ctx)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(ssh.Context) map[string]string); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// RestorePassword provides a mock function with given fields: ctx
func (_m *Metadata) RestorePassword(ctx ssh.Context) string {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(ssh.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RestoreRequest provides a mock function with given fields: ctx
func (_m *Metadata) RestoreRequest(ctx ssh.Context) string {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(ssh.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RestoreTarget provides a mock function with given fields: ctx
func (_m *Metadata) RestoreTarget(ctx ssh.Context) *target.Target {
	ret := _m.Called(ctx)

	var r0 *target.Target
	if rf, ok := ret.Get(0).(func(ssh.Context) *target.Target); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*target.Target)
		}
	}

	return r0
}

// StoreAuthenticationMethod provides a mock function with given fields: ctx, method
func (_m *Metadata) StoreAuthenticationMethod(ctx ssh.Context, method metadata.AuthenticationMethod) {
	_m.Called(ctx, method)
}

// StorePassword provides a mock function with given fields: ctx, value
func (_m *Metadata) StorePassword(ctx ssh.Context, value string) {
	_m.Called(ctx, value)
}

// StoreRequest provides a mock function with given fields: ctx, value
func (_m *Metadata) StoreRequest(ctx ssh.Context, value string) {
	_m.Called(ctx, value)
}

// NewMetadata creates a new instance of Metadata. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMetadata(t interface {
	mock.TestingT
	Cleanup(func())
}) *Metadata {
	mock := &Metadata{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
