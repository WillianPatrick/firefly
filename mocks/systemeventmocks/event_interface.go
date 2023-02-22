// Code generated by mockery v2.20.0. DO NOT EDIT.

package systemeventmocks

import (
	system "github.com/hyperledger/firefly/internal/events/system"
	mock "github.com/stretchr/testify/mock"
)

// EventInterface is an autogenerated mock type for the EventInterface type
type EventInterface struct {
	mock.Mock
}

// AddSystemEventListener provides a mock function with given fields: ns, el
func (_m *EventInterface) AddSystemEventListener(ns string, el system.EventListener) error {
	ret := _m.Called(ns, el)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, system.EventListener) error); ok {
		r0 = rf(ns, el)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewEventInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewEventInterface creates a new instance of EventInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEventInterface(t mockConstructorTestingTNewEventInterface) *EventInterface {
	mock := &EventInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
