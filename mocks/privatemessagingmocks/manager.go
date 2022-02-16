// Code generated by mockery v1.0.0. DO NOT EDIT.

package privatemessagingmocks

import (
	context "context"

	database "github.com/hyperledger/firefly/pkg/database"
	fftypes "github.com/hyperledger/firefly/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"

	sysmessaging "github.com/hyperledger/firefly/internal/sysmessaging"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// EnsureLocalGroup provides a mock function with given fields: ctx, group
func (_m *Manager) EnsureLocalGroup(ctx context.Context, group *fftypes.Group) (bool, error) {
	ret := _m.Called(ctx, group)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Group) bool); ok {
		r0 = rf(ctx, group)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.Group) error); ok {
		r1 = rf(ctx, group)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroupByID provides a mock function with given fields: ctx, id
func (_m *Manager) GetGroupByID(ctx context.Context, id string) (*fftypes.Group, error) {
	ret := _m.Called(ctx, id)

	var r0 *fftypes.Group
	if rf, ok := ret.Get(0).(func(context.Context, string) *fftypes.Group); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroupsNS provides a mock function with given fields: ctx, ns, filter
func (_m *Manager) GetGroupsNS(ctx context.Context, ns string, filter database.AndFilter) ([]*fftypes.Group, *database.FilterResult, error) {
	ret := _m.Called(ctx, ns, filter)

	var r0 []*fftypes.Group
	if rf, ok := ret.Get(0).(func(context.Context, string, database.AndFilter) []*fftypes.Group); ok {
		r0 = rf(ctx, ns, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*fftypes.Group)
		}
	}

	var r1 *database.FilterResult
	if rf, ok := ret.Get(1).(func(context.Context, string, database.AndFilter) *database.FilterResult); ok {
		r1 = rf(ctx, ns, filter)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*database.FilterResult)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, database.AndFilter) error); ok {
		r2 = rf(ctx, ns, filter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewMessage provides a mock function with given fields: ns, msg
func (_m *Manager) NewMessage(ns string, msg *fftypes.MessageInOut) sysmessaging.MessageSender {
	ret := _m.Called(ns, msg)

	var r0 sysmessaging.MessageSender
	if rf, ok := ret.Get(0).(func(string, *fftypes.MessageInOut) sysmessaging.MessageSender); ok {
		r0 = rf(ns, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sysmessaging.MessageSender)
		}
	}

	return r0
}

// PrepareOperation provides a mock function with given fields: ctx, op
func (_m *Manager) PrepareOperation(ctx context.Context, op *fftypes.Operation) (*fftypes.PreparedOperation, error) {
	ret := _m.Called(ctx, op)

	var r0 *fftypes.PreparedOperation
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Operation) *fftypes.PreparedOperation); ok {
		r0 = rf(ctx, op)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.PreparedOperation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.Operation) error); ok {
		r1 = rf(ctx, op)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestReply provides a mock function with given fields: ctx, ns, request
func (_m *Manager) RequestReply(ctx context.Context, ns string, request *fftypes.MessageInOut) (*fftypes.MessageInOut, error) {
	ret := _m.Called(ctx, ns, request)

	var r0 *fftypes.MessageInOut
	if rf, ok := ret.Get(0).(func(context.Context, string, *fftypes.MessageInOut) *fftypes.MessageInOut); ok {
		r0 = rf(ctx, ns, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.MessageInOut)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *fftypes.MessageInOut) error); ok {
		r1 = rf(ctx, ns, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResolveInitGroup provides a mock function with given fields: ctx, msg
func (_m *Manager) ResolveInitGroup(ctx context.Context, msg *fftypes.Message) (*fftypes.Group, error) {
	ret := _m.Called(ctx, msg)

	var r0 *fftypes.Group
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Message) *fftypes.Group); ok {
		r0 = rf(ctx, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.Message) error); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunOperation provides a mock function with given fields: ctx, op
func (_m *Manager) RunOperation(ctx context.Context, op *fftypes.PreparedOperation) (bool, error) {
	ret := _m.Called(ctx, op)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.PreparedOperation) bool); ok {
		r0 = rf(ctx, op)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.PreparedOperation) error); ok {
		r1 = rf(ctx, op)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMessage provides a mock function with given fields: ctx, ns, in, waitConfirm
func (_m *Manager) SendMessage(ctx context.Context, ns string, in *fftypes.MessageInOut, waitConfirm bool) (*fftypes.Message, error) {
	ret := _m.Called(ctx, ns, in, waitConfirm)

	var r0 *fftypes.Message
	if rf, ok := ret.Get(0).(func(context.Context, string, *fftypes.MessageInOut, bool) *fftypes.Message); ok {
		r0 = rf(ctx, ns, in, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *fftypes.MessageInOut, bool) error); ok {
		r1 = rf(ctx, ns, in, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields:
func (_m *Manager) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
