// Code generated by mockery v2.36.1. DO NOT EDIT.

package mock_envoy

import (
	api "github.com/envoyproxy/envoy/contrib/golang/common/go/api"
	mock "github.com/stretchr/testify/mock"
)

// FilterCallbackHandler is an autogenerated mock type for the FilterCallbackHandler type
type FilterCallbackHandler struct {
	mock.Mock
}

type FilterCallbackHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *FilterCallbackHandler) EXPECT() *FilterCallbackHandler_Expecter {
	return &FilterCallbackHandler_Expecter{mock: &_m.Mock}
}

// ClearRouteCache provides a mock function with given fields:
func (_m *FilterCallbackHandler) ClearRouteCache() {
	_m.Called()
}

// FilterCallbackHandler_ClearRouteCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ClearRouteCache'
type FilterCallbackHandler_ClearRouteCache_Call struct {
	*mock.Call
}

// ClearRouteCache is a helper method to define mock.On call
func (_e *FilterCallbackHandler_Expecter) ClearRouteCache() *FilterCallbackHandler_ClearRouteCache_Call {
	return &FilterCallbackHandler_ClearRouteCache_Call{Call: _e.mock.On("ClearRouteCache")}
}

func (_c *FilterCallbackHandler_ClearRouteCache_Call) Run(run func()) *FilterCallbackHandler_ClearRouteCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FilterCallbackHandler_ClearRouteCache_Call) Return() *FilterCallbackHandler_ClearRouteCache_Call {
	_c.Call.Return()
	return _c
}

func (_c *FilterCallbackHandler_ClearRouteCache_Call) RunAndReturn(run func()) *FilterCallbackHandler_ClearRouteCache_Call {
	_c.Call.Return(run)
	return _c
}

// Continue provides a mock function with given fields: _a0
func (_m *FilterCallbackHandler) Continue(_a0 api.StatusType) {
	_m.Called(_a0)
}

// FilterCallbackHandler_Continue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Continue'
type FilterCallbackHandler_Continue_Call struct {
	*mock.Call
}

// Continue is a helper method to define mock.On call
//   - _a0 api.StatusType
func (_e *FilterCallbackHandler_Expecter) Continue(_a0 interface{}) *FilterCallbackHandler_Continue_Call {
	return &FilterCallbackHandler_Continue_Call{Call: _e.mock.On("Continue", _a0)}
}

func (_c *FilterCallbackHandler_Continue_Call) Run(run func(_a0 api.StatusType)) *FilterCallbackHandler_Continue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.StatusType))
	})
	return _c
}

func (_c *FilterCallbackHandler_Continue_Call) Return() *FilterCallbackHandler_Continue_Call {
	_c.Call.Return()
	return _c
}

func (_c *FilterCallbackHandler_Continue_Call) RunAndReturn(run func(api.StatusType)) *FilterCallbackHandler_Continue_Call {
	_c.Call.Return(run)
	return _c
}

// GetProperty provides a mock function with given fields: key
func (_m *FilterCallbackHandler) GetProperty(key string) (string, error) {
	ret := _m.Called(key)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterCallbackHandler_GetProperty_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProperty'
type FilterCallbackHandler_GetProperty_Call struct {
	*mock.Call
}

// GetProperty is a helper method to define mock.On call
//   - key string
func (_e *FilterCallbackHandler_Expecter) GetProperty(key interface{}) *FilterCallbackHandler_GetProperty_Call {
	return &FilterCallbackHandler_GetProperty_Call{Call: _e.mock.On("GetProperty", key)}
}

func (_c *FilterCallbackHandler_GetProperty_Call) Run(run func(key string)) *FilterCallbackHandler_GetProperty_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *FilterCallbackHandler_GetProperty_Call) Return(_a0 string, _a1 error) *FilterCallbackHandler_GetProperty_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FilterCallbackHandler_GetProperty_Call) RunAndReturn(run func(string) (string, error)) *FilterCallbackHandler_GetProperty_Call {
	_c.Call.Return(run)
	return _c
}

// Log provides a mock function with given fields: level, msg
func (_m *FilterCallbackHandler) Log(level api.LogType, msg string) {
	_m.Called(level, msg)
}

// FilterCallbackHandler_Log_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Log'
type FilterCallbackHandler_Log_Call struct {
	*mock.Call
}

// Log is a helper method to define mock.On call
//   - level api.LogType
//   - msg string
func (_e *FilterCallbackHandler_Expecter) Log(level interface{}, msg interface{}) *FilterCallbackHandler_Log_Call {
	return &FilterCallbackHandler_Log_Call{Call: _e.mock.On("Log", level, msg)}
}

func (_c *FilterCallbackHandler_Log_Call) Run(run func(level api.LogType, msg string)) *FilterCallbackHandler_Log_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(api.LogType), args[1].(string))
	})
	return _c
}

func (_c *FilterCallbackHandler_Log_Call) Return() *FilterCallbackHandler_Log_Call {
	_c.Call.Return()
	return _c
}

func (_c *FilterCallbackHandler_Log_Call) RunAndReturn(run func(api.LogType, string)) *FilterCallbackHandler_Log_Call {
	_c.Call.Return(run)
	return _c
}

// LogLevel provides a mock function with given fields:
func (_m *FilterCallbackHandler) LogLevel() api.LogType {
	ret := _m.Called()

	var r0 api.LogType
	if rf, ok := ret.Get(0).(func() api.LogType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(api.LogType)
	}

	return r0
}

// FilterCallbackHandler_LogLevel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LogLevel'
type FilterCallbackHandler_LogLevel_Call struct {
	*mock.Call
}

// LogLevel is a helper method to define mock.On call
func (_e *FilterCallbackHandler_Expecter) LogLevel() *FilterCallbackHandler_LogLevel_Call {
	return &FilterCallbackHandler_LogLevel_Call{Call: _e.mock.On("LogLevel")}
}

func (_c *FilterCallbackHandler_LogLevel_Call) Run(run func()) *FilterCallbackHandler_LogLevel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FilterCallbackHandler_LogLevel_Call) Return(_a0 api.LogType) *FilterCallbackHandler_LogLevel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FilterCallbackHandler_LogLevel_Call) RunAndReturn(run func() api.LogType) *FilterCallbackHandler_LogLevel_Call {
	_c.Call.Return(run)
	return _c
}

// RecoverPanic provides a mock function with given fields:
func (_m *FilterCallbackHandler) RecoverPanic() {
	_m.Called()
}

// FilterCallbackHandler_RecoverPanic_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RecoverPanic'
type FilterCallbackHandler_RecoverPanic_Call struct {
	*mock.Call
}

// RecoverPanic is a helper method to define mock.On call
func (_e *FilterCallbackHandler_Expecter) RecoverPanic() *FilterCallbackHandler_RecoverPanic_Call {
	return &FilterCallbackHandler_RecoverPanic_Call{Call: _e.mock.On("RecoverPanic")}
}

func (_c *FilterCallbackHandler_RecoverPanic_Call) Run(run func()) *FilterCallbackHandler_RecoverPanic_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FilterCallbackHandler_RecoverPanic_Call) Return() *FilterCallbackHandler_RecoverPanic_Call {
	_c.Call.Return()
	return _c
}

func (_c *FilterCallbackHandler_RecoverPanic_Call) RunAndReturn(run func()) *FilterCallbackHandler_RecoverPanic_Call {
	_c.Call.Return(run)
	return _c
}

// SendLocalReply provides a mock function with given fields: responseCode, bodyText, headers, grpcStatus, details
func (_m *FilterCallbackHandler) SendLocalReply(responseCode int, bodyText string, headers map[string][]string, grpcStatus int64, details string) {
	_m.Called(responseCode, bodyText, headers, grpcStatus, details)
}

// FilterCallbackHandler_SendLocalReply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendLocalReply'
type FilterCallbackHandler_SendLocalReply_Call struct {
	*mock.Call
}

// SendLocalReply is a helper method to define mock.On call
//   - responseCode int
//   - bodyText string
//   - headers map[string][]string
//   - grpcStatus int64
//   - details string
func (_e *FilterCallbackHandler_Expecter) SendLocalReply(responseCode interface{}, bodyText interface{}, headers interface{}, grpcStatus interface{}, details interface{}) *FilterCallbackHandler_SendLocalReply_Call {
	return &FilterCallbackHandler_SendLocalReply_Call{Call: _e.mock.On("SendLocalReply", responseCode, bodyText, headers, grpcStatus, details)}
}

func (_c *FilterCallbackHandler_SendLocalReply_Call) Run(run func(responseCode int, bodyText string, headers map[string][]string, grpcStatus int64, details string)) *FilterCallbackHandler_SendLocalReply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(string), args[2].(map[string][]string), args[3].(int64), args[4].(string))
	})
	return _c
}

func (_c *FilterCallbackHandler_SendLocalReply_Call) Return() *FilterCallbackHandler_SendLocalReply_Call {
	_c.Call.Return()
	return _c
}

func (_c *FilterCallbackHandler_SendLocalReply_Call) RunAndReturn(run func(int, string, map[string][]string, int64, string)) *FilterCallbackHandler_SendLocalReply_Call {
	_c.Call.Return(run)
	return _c
}

// StreamInfo provides a mock function with given fields:
func (_m *FilterCallbackHandler) StreamInfo() api.StreamInfo {
	ret := _m.Called()

	var r0 api.StreamInfo
	if rf, ok := ret.Get(0).(func() api.StreamInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(api.StreamInfo)
		}
	}

	return r0
}

// FilterCallbackHandler_StreamInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StreamInfo'
type FilterCallbackHandler_StreamInfo_Call struct {
	*mock.Call
}

// StreamInfo is a helper method to define mock.On call
func (_e *FilterCallbackHandler_Expecter) StreamInfo() *FilterCallbackHandler_StreamInfo_Call {
	return &FilterCallbackHandler_StreamInfo_Call{Call: _e.mock.On("StreamInfo")}
}

func (_c *FilterCallbackHandler_StreamInfo_Call) Run(run func()) *FilterCallbackHandler_StreamInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FilterCallbackHandler_StreamInfo_Call) Return(_a0 api.StreamInfo) *FilterCallbackHandler_StreamInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FilterCallbackHandler_StreamInfo_Call) RunAndReturn(run func() api.StreamInfo) *FilterCallbackHandler_StreamInfo_Call {
	_c.Call.Return(run)
	return _c
}

// NewFilterCallbackHandler creates a new instance of FilterCallbackHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFilterCallbackHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *FilterCallbackHandler {
	mock := &FilterCallbackHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
