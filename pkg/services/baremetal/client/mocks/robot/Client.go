// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/syself/hrobot-go/models"

	v1beta1 "github.com/syself/cluster-api-provider-hetzner/api/v1beta1"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// DeleteBootRescue provides a mock function with given fields: id
func (_m *Client) DeleteBootRescue(id int) (*models.Rescue, error) {
	ret := _m.Called(id)

	var r0 *models.Rescue
	if rf, ok := ret.Get(0).(func(int) *models.Rescue); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Rescue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBMServer provides a mock function with given fields: _a0
func (_m *Client) GetBMServer(_a0 int) (*models.Server, error) {
	ret := _m.Called(_a0)

	var r0 *models.Server
	if rf, ok := ret.Get(0).(func(int) *models.Server); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Server)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBootRescue provides a mock function with given fields: id
func (_m *Client) GetBootRescue(id int) (*models.Rescue, error) {
	ret := _m.Called(id)

	var r0 *models.Rescue
	if rf, ok := ret.Get(0).(func(int) *models.Rescue); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Rescue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReboot provides a mock function with given fields: _a0
func (_m *Client) GetReboot(_a0 int) (*models.Reset, error) {
	ret := _m.Called(_a0)

	var r0 *models.Reset
	if rf, ok := ret.Get(0).(func(int) *models.Reset); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Reset)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListBMServers provides a mock function with given fields:
func (_m *Client) ListBMServers() ([]models.Server, error) {
	ret := _m.Called()

	var r0 []models.Server
	if rf, ok := ret.Get(0).(func() []models.Server); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Server)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListSSHKeys provides a mock function with given fields:
func (_m *Client) ListSSHKeys() ([]models.Key, error) {
	ret := _m.Called()

	var r0 []models.Key
	if rf, ok := ret.Get(0).(func() []models.Key); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RebootBMServer provides a mock function with given fields: _a0, _a1
func (_m *Client) RebootBMServer(_a0 int, _a1 v1beta1.RebootType) (*models.ResetPost, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *models.ResetPost
	if rf, ok := ret.Get(0).(func(int, v1beta1.RebootType) *models.ResetPost); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ResetPost)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, v1beta1.RebootType) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetBMServerName provides a mock function with given fields: _a0, _a1
func (_m *Client) SetBMServerName(_a0 int, _a1 string) (*models.Server, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *models.Server
	if rf, ok := ret.Get(0).(func(int, string) *models.Server); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Server)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetBootRescue provides a mock function with given fields: id, fingerprint
func (_m *Client) SetBootRescue(id int, fingerprint string) (*models.Rescue, error) {
	ret := _m.Called(id, fingerprint)

	var r0 *models.Rescue
	if rf, ok := ret.Get(0).(func(int, string) *models.Rescue); ok {
		r0 = rf(id, fingerprint)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Rescue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(id, fingerprint)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetSSHKey provides a mock function with given fields: name, publickey
func (_m *Client) SetSSHKey(name string, publickey string) (*models.Key, error) {
	ret := _m.Called(name, publickey)

	var r0 *models.Key
	if rf, ok := ret.Get(0).(func(string, string) *models.Key); ok {
		r0 = rf(name, publickey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, publickey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateCredentials provides a mock function with given fields:
func (_m *Client) ValidateCredentials() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
