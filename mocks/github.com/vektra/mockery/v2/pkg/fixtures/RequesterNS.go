// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// RequesterNS is an autogenerated mock type for the RequesterNS type
type RequesterNS struct {
	mock.Mock
}

type RequesterNS_Expecter struct {
	mock *mock.Mock
}

func (_m *RequesterNS) EXPECT() *RequesterNS_Expecter {
	return &RequesterNS_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: path
func (_m *RequesterNS) Get(path string) (http.Response, error) {
	ret := _m.Called(path)

	var r0 http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (http.Response, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) http.Response); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(http.Response)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequesterNS_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type RequesterNS_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - path string
func (_e *RequesterNS_Expecter) Get(path interface{}) *RequesterNS_Get_Call {
	return &RequesterNS_Get_Call{Call: _e.mock.On("Get", path)}
}

func (_c *RequesterNS_Get_Call) Run(run func(path string)) *RequesterNS_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RequesterNS_Get_Call) Return(_a0 http.Response, _a1 error) *RequesterNS_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RequesterNS_Get_Call) RunAndReturn(run func(string) (http.Response, error)) *RequesterNS_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewRequesterNS creates a new instance of RequesterNS. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequesterNS(t interface {
	mock.TestingT
	Cleanup(func())
}) *RequesterNS {
	mock := &RequesterNS{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
