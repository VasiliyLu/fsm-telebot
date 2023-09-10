// Code generated by mockery v2.33.2. DO NOT EDIT.

package fsm

import (
	mock "github.com/stretchr/testify/mock"
	telebot "gopkg.in/telebot.v3"
)

// MockContext is an autogenerated mock type for the Context type
type MockContext struct {
	mock.Mock
}

type MockContext_Expecter struct {
	mock *mock.Mock
}

func (_m *MockContext) EXPECT() *MockContext_Expecter {
	return &MockContext_Expecter{mock: &_m.Mock}
}

// Bot provides a mock function with given fields:
func (_m *MockContext) Bot() *telebot.Bot {
	ret := _m.Called()

	var r0 *telebot.Bot
	if rf, ok := ret.Get(0).(func() *telebot.Bot); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*telebot.Bot)
		}
	}

	return r0
}

// MockContext_Bot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Bot'
type MockContext_Bot_Call struct {
	*mock.Call
}

// Bot is a helper method to define mock.On call
func (_e *MockContext_Expecter) Bot() *MockContext_Bot_Call {
	return &MockContext_Bot_Call{Call: _e.mock.On("Bot")}
}

func (_c *MockContext_Bot_Call) Run(run func()) *MockContext_Bot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_Bot_Call) Return(_a0 *telebot.Bot) *MockContext_Bot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Bot_Call) RunAndReturn(run func() *telebot.Bot) *MockContext_Bot_Call {
	_c.Call.Return(run)
	return _c
}

// Finish provides a mock function with given fields: deleteData
func (_m *MockContext) Finish(deleteData bool) error {
	ret := _m.Called(deleteData)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(deleteData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockContext_Finish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Finish'
type MockContext_Finish_Call struct {
	*mock.Call
}

// Finish is a helper method to define mock.On call
//   - deleteData bool
func (_e *MockContext_Expecter) Finish(deleteData interface{}) *MockContext_Finish_Call {
	return &MockContext_Finish_Call{Call: _e.mock.On("Finish", deleteData)}
}

func (_c *MockContext_Finish_Call) Run(run func(deleteData bool)) *MockContext_Finish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *MockContext_Finish_Call) Return(_a0 error) *MockContext_Finish_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Finish_Call) RunAndReturn(run func(bool) error) *MockContext_Finish_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: key, to
func (_m *MockContext) Get(key string, to interface{}) error {
	ret := _m.Called(key, to)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(key, to)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockContext_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockContext_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - key string
//   - to interface{}
func (_e *MockContext_Expecter) Get(key interface{}, to interface{}) *MockContext_Get_Call {
	return &MockContext_Get_Call{Call: _e.mock.On("Get", key, to)}
}

func (_c *MockContext_Get_Call) Run(run func(key string, to interface{})) *MockContext_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}))
	})
	return _c
}

func (_c *MockContext_Get_Call) Return(_a0 error) *MockContext_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Get_Call) RunAndReturn(run func(string, interface{}) error) *MockContext_Get_Call {
	_c.Call.Return(run)
	return _c
}

// MustGet provides a mock function with given fields: key, to
func (_m *MockContext) MustGet(key string, to interface{}) {
	_m.Called(key, to)
}

// MockContext_MustGet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MustGet'
type MockContext_MustGet_Call struct {
	*mock.Call
}

// MustGet is a helper method to define mock.On call
//   - key string
//   - to interface{}
func (_e *MockContext_Expecter) MustGet(key interface{}, to interface{}) *MockContext_MustGet_Call {
	return &MockContext_MustGet_Call{Call: _e.mock.On("MustGet", key, to)}
}

func (_c *MockContext_MustGet_Call) Run(run func(key string, to interface{})) *MockContext_MustGet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}))
	})
	return _c
}

func (_c *MockContext_MustGet_Call) Return() *MockContext_MustGet_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockContext_MustGet_Call) RunAndReturn(run func(string, interface{})) *MockContext_MustGet_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: state
func (_m *MockContext) Set(state State) error {
	ret := _m.Called(state)

	var r0 error
	if rf, ok := ret.Get(0).(func(State) error); ok {
		r0 = rf(state)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockContext_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type MockContext_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - state State
func (_e *MockContext_Expecter) Set(state interface{}) *MockContext_Set_Call {
	return &MockContext_Set_Call{Call: _e.mock.On("Set", state)}
}

func (_c *MockContext_Set_Call) Run(run func(state State)) *MockContext_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(State))
	})
	return _c
}

func (_c *MockContext_Set_Call) Return(_a0 error) *MockContext_Set_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Set_Call) RunAndReturn(run func(State) error) *MockContext_Set_Call {
	_c.Call.Return(run)
	return _c
}

// State provides a mock function with given fields:
func (_m *MockContext) State() (State, error) {
	ret := _m.Called()

	var r0 State
	var r1 error
	if rf, ok := ret.Get(0).(func() (State, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() State); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(State)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockContext_State_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'State'
type MockContext_State_Call struct {
	*mock.Call
}

// State is a helper method to define mock.On call
func (_e *MockContext_Expecter) State() *MockContext_State_Call {
	return &MockContext_State_Call{Call: _e.mock.On("State")}
}

func (_c *MockContext_State_Call) Run(run func()) *MockContext_State_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContext_State_Call) Return(_a0 State, _a1 error) *MockContext_State_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockContext_State_Call) RunAndReturn(run func() (State, error)) *MockContext_State_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: key, data
func (_m *MockContext) Update(key string, data interface{}) error {
	ret := _m.Called(key, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}) error); ok {
		r0 = rf(key, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockContext_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockContext_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - key string
//   - data interface{}
func (_e *MockContext_Expecter) Update(key interface{}, data interface{}) *MockContext_Update_Call {
	return &MockContext_Update_Call{Call: _e.mock.On("Update", key, data)}
}

func (_c *MockContext_Update_Call) Run(run func(key string, data interface{})) *MockContext_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}))
	})
	return _c
}

func (_c *MockContext_Update_Call) Return(_a0 error) *MockContext_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContext_Update_Call) RunAndReturn(run func(string, interface{}) error) *MockContext_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockContext creates a new instance of MockContext. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockContext(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockContext {
	mock := &MockContext{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}