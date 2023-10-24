// Code generated by mockery v2.32.0. DO NOT EDIT.

package mock_envoy

import mock "github.com/stretchr/testify/mock"

// RequestHeaderMap is an autogenerated mock type for the RequestHeaderMap type
type RequestHeaderMap struct {
	mock.Mock
}

type RequestHeaderMap_Expecter struct {
	mock *mock.Mock
}

func (_m *RequestHeaderMap) EXPECT() *RequestHeaderMap_Expecter {
	return &RequestHeaderMap_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: key, value
func (_m *RequestHeaderMap) Add(key string, value string) {
	_m.Called(key, value)
}

// RequestHeaderMap_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type RequestHeaderMap_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - key string
//   - value string
func (_e *RequestHeaderMap_Expecter) Add(key interface{}, value interface{}) *RequestHeaderMap_Add_Call {
	return &RequestHeaderMap_Add_Call{Call: _e.mock.On("Add", key, value)}
}

func (_c *RequestHeaderMap_Add_Call) Run(run func(key string, value string)) *RequestHeaderMap_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *RequestHeaderMap_Add_Call) Return() *RequestHeaderMap_Add_Call {
	_c.Call.Return()
	return _c
}

func (_c *RequestHeaderMap_Add_Call) RunAndReturn(run func(string, string)) *RequestHeaderMap_Add_Call {
	_c.Call.Return(run)
	return _c
}

// ByteSize provides a mock function with given fields:
func (_m *RequestHeaderMap) ByteSize() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// RequestHeaderMap_ByteSize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ByteSize'
type RequestHeaderMap_ByteSize_Call struct {
	*mock.Call
}

// ByteSize is a helper method to define mock.On call
func (_e *RequestHeaderMap_Expecter) ByteSize() *RequestHeaderMap_ByteSize_Call {
	return &RequestHeaderMap_ByteSize_Call{Call: _e.mock.On("ByteSize")}
}

func (_c *RequestHeaderMap_ByteSize_Call) Run(run func()) *RequestHeaderMap_ByteSize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RequestHeaderMap_ByteSize_Call) Return(_a0 uint64) *RequestHeaderMap_ByteSize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RequestHeaderMap_ByteSize_Call) RunAndReturn(run func() uint64) *RequestHeaderMap_ByteSize_Call {
	_c.Call.Return(run)
	return _c
}

// Del provides a mock function with given fields: key
func (_m *RequestHeaderMap) Del(key string) {
	_m.Called(key)
}

// RequestHeaderMap_Del_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Del'
type RequestHeaderMap_Del_Call struct {
	*mock.Call
}

// Del is a helper method to define mock.On call
//   - key string
func (_e *RequestHeaderMap_Expecter) Del(key interface{}) *RequestHeaderMap_Del_Call {
	return &RequestHeaderMap_Del_Call{Call: _e.mock.On("Del", key)}
}

func (_c *RequestHeaderMap_Del_Call) Run(run func(key string)) *RequestHeaderMap_Del_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RequestHeaderMap_Del_Call) Return() *RequestHeaderMap_Del_Call {
	_c.Call.Return()
	return _c
}

func (_c *RequestHeaderMap_Del_Call) RunAndReturn(run func(string)) *RequestHeaderMap_Del_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: key
func (_m *RequestHeaderMap) Get(key string) (string, bool) {
	ret := _m.Called(key)

	var r0 string
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (string, bool)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// RequestHeaderMap_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type RequestHeaderMap_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - key string
func (_e *RequestHeaderMap_Expecter) Get(key interface{}) *RequestHeaderMap_Get_Call {
	return &RequestHeaderMap_Get_Call{Call: _e.mock.On("Get", key)}
}

func (_c *RequestHeaderMap_Get_Call) Run(run func(key string)) *RequestHeaderMap_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RequestHeaderMap_Get_Call) Return(_a0 string, _a1 bool) *RequestHeaderMap_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RequestHeaderMap_Get_Call) RunAndReturn(run func(string) (string, bool)) *RequestHeaderMap_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetRaw provides a mock function with given fields: name
func (_m *RequestHeaderMap) GetRaw(name string) string {
	ret := _m.Called(name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RequestHeaderMap_GetRaw_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRaw'
type RequestHeaderMap_GetRaw_Call struct {
	*mock.Call
}

// GetRaw is a helper method to define mock.On call
//   - name string
func (_e *RequestHeaderMap_Expecter) GetRaw(name interface{}) *RequestHeaderMap_GetRaw_Call {
	return &RequestHeaderMap_GetRaw_Call{Call: _e.mock.On("GetRaw", name)}
}

func (_c *RequestHeaderMap_GetRaw_Call) Run(run func(name string)) *RequestHeaderMap_GetRaw_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RequestHeaderMap_GetRaw_Call) Return(_a0 string) *RequestHeaderMap_GetRaw_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RequestHeaderMap_GetRaw_Call) RunAndReturn(run func(string) string) *RequestHeaderMap_GetRaw_Call {
	_c.Call.Return(run)
	return _c
}

// Host provides a mock function with given fields:
func (_m *RequestHeaderMap) Host() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RequestHeaderMap_Host_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Host'
type RequestHeaderMap_Host_Call struct {
	*mock.Call
}

// Host is a helper method to define mock.On call
func (_e *RequestHeaderMap_Expecter) Host() *RequestHeaderMap_Host_Call {
	return &RequestHeaderMap_Host_Call{Call: _e.mock.On("Host")}
}

func (_c *RequestHeaderMap_Host_Call) Run(run func()) *RequestHeaderMap_Host_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RequestHeaderMap_Host_Call) Return(_a0 string) *RequestHeaderMap_Host_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RequestHeaderMap_Host_Call) RunAndReturn(run func() string) *RequestHeaderMap_Host_Call {
	_c.Call.Return(run)
	return _c
}

// Method provides a mock function with given fields:
func (_m *RequestHeaderMap) Method() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RequestHeaderMap_Method_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Method'
type RequestHeaderMap_Method_Call struct {
	*mock.Call
}

// Method is a helper method to define mock.On call
func (_e *RequestHeaderMap_Expecter) Method() *RequestHeaderMap_Method_Call {
	return &RequestHeaderMap_Method_Call{Call: _e.mock.On("Method")}
}

func (_c *RequestHeaderMap_Method_Call) Run(run func()) *RequestHeaderMap_Method_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RequestHeaderMap_Method_Call) Return(_a0 string) *RequestHeaderMap_Method_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RequestHeaderMap_Method_Call) RunAndReturn(run func() string) *RequestHeaderMap_Method_Call {
	_c.Call.Return(run)
	return _c
}

// Path provides a mock function with given fields:
func (_m *RequestHeaderMap) Path() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RequestHeaderMap_Path_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Path'
type RequestHeaderMap_Path_Call struct {
	*mock.Call
}

// Path is a helper method to define mock.On call
func (_e *RequestHeaderMap_Expecter) Path() *RequestHeaderMap_Path_Call {
	return &RequestHeaderMap_Path_Call{Call: _e.mock.On("Path")}
}

func (_c *RequestHeaderMap_Path_Call) Run(run func()) *RequestHeaderMap_Path_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RequestHeaderMap_Path_Call) Return(_a0 string) *RequestHeaderMap_Path_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RequestHeaderMap_Path_Call) RunAndReturn(run func() string) *RequestHeaderMap_Path_Call {
	_c.Call.Return(run)
	return _c
}

// Protocol provides a mock function with given fields:
func (_m *RequestHeaderMap) Protocol() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RequestHeaderMap_Protocol_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Protocol'
type RequestHeaderMap_Protocol_Call struct {
	*mock.Call
}

// Protocol is a helper method to define mock.On call
func (_e *RequestHeaderMap_Expecter) Protocol() *RequestHeaderMap_Protocol_Call {
	return &RequestHeaderMap_Protocol_Call{Call: _e.mock.On("Protocol")}
}

func (_c *RequestHeaderMap_Protocol_Call) Run(run func()) *RequestHeaderMap_Protocol_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RequestHeaderMap_Protocol_Call) Return(_a0 string) *RequestHeaderMap_Protocol_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RequestHeaderMap_Protocol_Call) RunAndReturn(run func() string) *RequestHeaderMap_Protocol_Call {
	_c.Call.Return(run)
	return _c
}

// Range provides a mock function with given fields: f
func (_m *RequestHeaderMap) Range(f func(string, string) bool) {
	_m.Called(f)
}

// RequestHeaderMap_Range_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Range'
type RequestHeaderMap_Range_Call struct {
	*mock.Call
}

// Range is a helper method to define mock.On call
//   - f func(string , string) bool
func (_e *RequestHeaderMap_Expecter) Range(f interface{}) *RequestHeaderMap_Range_Call {
	return &RequestHeaderMap_Range_Call{Call: _e.mock.On("Range", f)}
}

func (_c *RequestHeaderMap_Range_Call) Run(run func(f func(string, string) bool)) *RequestHeaderMap_Range_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(string, string) bool))
	})
	return _c
}

func (_c *RequestHeaderMap_Range_Call) Return() *RequestHeaderMap_Range_Call {
	_c.Call.Return()
	return _c
}

func (_c *RequestHeaderMap_Range_Call) RunAndReturn(run func(func(string, string) bool)) *RequestHeaderMap_Range_Call {
	_c.Call.Return(run)
	return _c
}

// RangeWithCopy provides a mock function with given fields: f
func (_m *RequestHeaderMap) RangeWithCopy(f func(string, string) bool) {
	_m.Called(f)
}

// RequestHeaderMap_RangeWithCopy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RangeWithCopy'
type RequestHeaderMap_RangeWithCopy_Call struct {
	*mock.Call
}

// RangeWithCopy is a helper method to define mock.On call
//   - f func(string , string) bool
func (_e *RequestHeaderMap_Expecter) RangeWithCopy(f interface{}) *RequestHeaderMap_RangeWithCopy_Call {
	return &RequestHeaderMap_RangeWithCopy_Call{Call: _e.mock.On("RangeWithCopy", f)}
}

func (_c *RequestHeaderMap_RangeWithCopy_Call) Run(run func(f func(string, string) bool)) *RequestHeaderMap_RangeWithCopy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(string, string) bool))
	})
	return _c
}

func (_c *RequestHeaderMap_RangeWithCopy_Call) Return() *RequestHeaderMap_RangeWithCopy_Call {
	_c.Call.Return()
	return _c
}

func (_c *RequestHeaderMap_RangeWithCopy_Call) RunAndReturn(run func(func(string, string) bool)) *RequestHeaderMap_RangeWithCopy_Call {
	_c.Call.Return(run)
	return _c
}

// Scheme provides a mock function with given fields:
func (_m *RequestHeaderMap) Scheme() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RequestHeaderMap_Scheme_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Scheme'
type RequestHeaderMap_Scheme_Call struct {
	*mock.Call
}

// Scheme is a helper method to define mock.On call
func (_e *RequestHeaderMap_Expecter) Scheme() *RequestHeaderMap_Scheme_Call {
	return &RequestHeaderMap_Scheme_Call{Call: _e.mock.On("Scheme")}
}

func (_c *RequestHeaderMap_Scheme_Call) Run(run func()) *RequestHeaderMap_Scheme_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RequestHeaderMap_Scheme_Call) Return(_a0 string) *RequestHeaderMap_Scheme_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RequestHeaderMap_Scheme_Call) RunAndReturn(run func() string) *RequestHeaderMap_Scheme_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: key, value
func (_m *RequestHeaderMap) Set(key string, value string) {
	_m.Called(key, value)
}

// RequestHeaderMap_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type RequestHeaderMap_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - key string
//   - value string
func (_e *RequestHeaderMap_Expecter) Set(key interface{}, value interface{}) *RequestHeaderMap_Set_Call {
	return &RequestHeaderMap_Set_Call{Call: _e.mock.On("Set", key, value)}
}

func (_c *RequestHeaderMap_Set_Call) Run(run func(key string, value string)) *RequestHeaderMap_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *RequestHeaderMap_Set_Call) Return() *RequestHeaderMap_Set_Call {
	_c.Call.Return()
	return _c
}

func (_c *RequestHeaderMap_Set_Call) RunAndReturn(run func(string, string)) *RequestHeaderMap_Set_Call {
	_c.Call.Return(run)
	return _c
}

// Values provides a mock function with given fields: key
func (_m *RequestHeaderMap) Values(key string) []string {
	ret := _m.Called(key)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// RequestHeaderMap_Values_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Values'
type RequestHeaderMap_Values_Call struct {
	*mock.Call
}

// Values is a helper method to define mock.On call
//   - key string
func (_e *RequestHeaderMap_Expecter) Values(key interface{}) *RequestHeaderMap_Values_Call {
	return &RequestHeaderMap_Values_Call{Call: _e.mock.On("Values", key)}
}

func (_c *RequestHeaderMap_Values_Call) Run(run func(key string)) *RequestHeaderMap_Values_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RequestHeaderMap_Values_Call) Return(_a0 []string) *RequestHeaderMap_Values_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RequestHeaderMap_Values_Call) RunAndReturn(run func(string) []string) *RequestHeaderMap_Values_Call {
	_c.Call.Return(run)
	return _c
}

// NewRequestHeaderMap creates a new instance of RequestHeaderMap. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRequestHeaderMap(t interface {
	mock.TestingT
	Cleanup(func())
}) *RequestHeaderMap {
	mock := &RequestHeaderMap{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
