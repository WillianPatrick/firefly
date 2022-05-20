// Code generated by mockery v1.0.0. DO NOT EDIT.

package tokenmocks

import (
	context "context"

	config "github.com/hyperledger/firefly-common/pkg/config"

	core "github.com/hyperledger/firefly/pkg/core"

	fftypes "github.com/hyperledger/firefly-common/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"

	tokens "github.com/hyperledger/firefly/pkg/tokens"
)

// Plugin is an autogenerated mock type for the Plugin type
type Plugin struct {
	mock.Mock
}

// ActivateTokenPool provides a mock function with given fields: ctx, opID, pool
func (_m *Plugin) ActivateTokenPool(ctx context.Context, opID *fftypes.UUID, pool *core.TokenPool) (bool, error) {
	ret := _m.Called(ctx, opID, pool)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, *core.TokenPool) bool); ok {
		r0 = rf(ctx, opID, pool)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, *core.TokenPool) error); ok {
		r1 = rf(ctx, opID, pool)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BurnTokens provides a mock function with given fields: ctx, opID, poolLocator, burn
func (_m *Plugin) BurnTokens(ctx context.Context, opID *fftypes.UUID, poolLocator string, burn *core.TokenTransfer) error {
	ret := _m.Called(ctx, opID, poolLocator, burn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, string, *core.TokenTransfer) error); ok {
		r0 = rf(ctx, opID, poolLocator, burn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Capabilities provides a mock function with given fields:
func (_m *Plugin) Capabilities() *tokens.Capabilities {
	ret := _m.Called()

	var r0 *tokens.Capabilities
	if rf, ok := ret.Get(0).(func() *tokens.Capabilities); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tokens.Capabilities)
		}
	}

	return r0
}

// CreateTokenPool provides a mock function with given fields: ctx, opID, pool
func (_m *Plugin) CreateTokenPool(ctx context.Context, opID *fftypes.UUID, pool *core.TokenPool) (bool, error) {
	ret := _m.Called(ctx, opID, pool)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, *core.TokenPool) bool); ok {
		r0 = rf(ctx, opID, pool)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, *core.TokenPool) error); ok {
		r1 = rf(ctx, opID, pool)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields: ctx, name, _a2, callbacks
func (_m *Plugin) Init(ctx context.Context, name string, _a2 config.Section, callbacks tokens.Callbacks) error {
	ret := _m.Called(ctx, name, _a2, callbacks)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, config.Section, tokens.Callbacks) error); ok {
		r0 = rf(ctx, name, _a2, callbacks)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitConfig provides a mock function with given fields: _a0
func (_m *Plugin) InitConfig(_a0 config.KeySet) {
	_m.Called(_a0)
}

// MintTokens provides a mock function with given fields: ctx, opID, poolLocator, mint
func (_m *Plugin) MintTokens(ctx context.Context, opID *fftypes.UUID, poolLocator string, mint *core.TokenTransfer) error {
	ret := _m.Called(ctx, opID, poolLocator, mint)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, string, *core.TokenTransfer) error); ok {
		r0 = rf(ctx, opID, poolLocator, mint)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *Plugin) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *Plugin) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TokensApproval provides a mock function with given fields: ctx, opID, poolLocator, approval
func (_m *Plugin) TokensApproval(ctx context.Context, opID *fftypes.UUID, poolLocator string, approval *core.TokenApproval) error {
	ret := _m.Called(ctx, opID, poolLocator, approval)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, string, *core.TokenApproval) error); ok {
		r0 = rf(ctx, opID, poolLocator, approval)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TransferTokens provides a mock function with given fields: ctx, opID, poolLocator, transfer
func (_m *Plugin) TransferTokens(ctx context.Context, opID *fftypes.UUID, poolLocator string, transfer *core.TokenTransfer) error {
	ret := _m.Called(ctx, opID, poolLocator, transfer)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, string, *core.TokenTransfer) error); ok {
		r0 = rf(ctx, opID, poolLocator, transfer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
