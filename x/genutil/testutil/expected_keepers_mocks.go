// Code generated by MockGen. DO NOT EDIT.
// Source: x/genutil/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	json "encoding/json"
	reflect "reflect"

	types "github.com/cometbft/cometbft/abci/types"
	codec "github.com/shapeshift/cosmos-sdk/codec"
	types0 "github.com/shapeshift/cosmos-sdk/types"
	types1 "github.com/shapeshift/cosmos-sdk/x/auth/types"
	exported "github.com/shapeshift/cosmos-sdk/x/bank/exported"
	gomock "github.com/golang/mock/gomock"
)

// MockStakingKeeper is a mock of StakingKeeper interface.
type MockStakingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockStakingKeeperMockRecorder
}

// MockStakingKeeperMockRecorder is the mock recorder for MockStakingKeeper.
type MockStakingKeeperMockRecorder struct {
	mock *MockStakingKeeper
}

// NewMockStakingKeeper creates a new mock instance.
func NewMockStakingKeeper(ctrl *gomock.Controller) *MockStakingKeeper {
	mock := &MockStakingKeeper{ctrl: ctrl}
	mock.recorder = &MockStakingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakingKeeper) EXPECT() *MockStakingKeeperMockRecorder {
	return m.recorder
}

// ApplyAndReturnValidatorSetUpdates mocks base method.
func (m *MockStakingKeeper) ApplyAndReturnValidatorSetUpdates(arg0 types0.Context) ([]types.ValidatorUpdate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyAndReturnValidatorSetUpdates", arg0)
	ret0, _ := ret[0].([]types.ValidatorUpdate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplyAndReturnValidatorSetUpdates indicates an expected call of ApplyAndReturnValidatorSetUpdates.
func (mr *MockStakingKeeperMockRecorder) ApplyAndReturnValidatorSetUpdates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyAndReturnValidatorSetUpdates", reflect.TypeOf((*MockStakingKeeper)(nil).ApplyAndReturnValidatorSetUpdates), arg0)
}

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// IterateAccounts mocks base method.
func (m *MockAccountKeeper) IterateAccounts(ctx types0.Context, process func(types1.AccountI) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateAccounts", ctx, process)
}

// IterateAccounts indicates an expected call of IterateAccounts.
func (mr *MockAccountKeeperMockRecorder) IterateAccounts(ctx, process interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateAccounts", reflect.TypeOf((*MockAccountKeeper)(nil).IterateAccounts), ctx, process)
}

// NewAccount mocks base method.
func (m *MockAccountKeeper) NewAccount(arg0 types0.Context, arg1 types1.AccountI) types1.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAccount", arg0, arg1)
	ret0, _ := ret[0].(types1.AccountI)
	return ret0
}

// NewAccount indicates an expected call of NewAccount.
func (mr *MockAccountKeeperMockRecorder) NewAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAccount", reflect.TypeOf((*MockAccountKeeper)(nil).NewAccount), arg0, arg1)
}

// SetAccount mocks base method.
func (m *MockAccountKeeper) SetAccount(arg0 types0.Context, arg1 types1.AccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAccount", arg0, arg1)
}

// SetAccount indicates an expected call of SetAccount.
func (mr *MockAccountKeeperMockRecorder) SetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetAccount), arg0, arg1)
}

// MockGenesisAccountsIterator is a mock of GenesisAccountsIterator interface.
type MockGenesisAccountsIterator struct {
	ctrl     *gomock.Controller
	recorder *MockGenesisAccountsIteratorMockRecorder
}

// MockGenesisAccountsIteratorMockRecorder is the mock recorder for MockGenesisAccountsIterator.
type MockGenesisAccountsIteratorMockRecorder struct {
	mock *MockGenesisAccountsIterator
}

// NewMockGenesisAccountsIterator creates a new mock instance.
func NewMockGenesisAccountsIterator(ctrl *gomock.Controller) *MockGenesisAccountsIterator {
	mock := &MockGenesisAccountsIterator{ctrl: ctrl}
	mock.recorder = &MockGenesisAccountsIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenesisAccountsIterator) EXPECT() *MockGenesisAccountsIteratorMockRecorder {
	return m.recorder
}

// IterateGenesisAccounts mocks base method.
func (m *MockGenesisAccountsIterator) IterateGenesisAccounts(cdc *codec.LegacyAmino, appGenesis map[string]json.RawMessage, cb func(types1.AccountI) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateGenesisAccounts", cdc, appGenesis, cb)
}

// IterateGenesisAccounts indicates an expected call of IterateGenesisAccounts.
func (mr *MockGenesisAccountsIteratorMockRecorder) IterateGenesisAccounts(cdc, appGenesis, cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateGenesisAccounts", reflect.TypeOf((*MockGenesisAccountsIterator)(nil).IterateGenesisAccounts), cdc, appGenesis, cb)
}

// MockGenesisBalancesIterator is a mock of GenesisBalancesIterator interface.
type MockGenesisBalancesIterator struct {
	ctrl     *gomock.Controller
	recorder *MockGenesisBalancesIteratorMockRecorder
}

// MockGenesisBalancesIteratorMockRecorder is the mock recorder for MockGenesisBalancesIterator.
type MockGenesisBalancesIteratorMockRecorder struct {
	mock *MockGenesisBalancesIterator
}

// NewMockGenesisBalancesIterator creates a new mock instance.
func NewMockGenesisBalancesIterator(ctrl *gomock.Controller) *MockGenesisBalancesIterator {
	mock := &MockGenesisBalancesIterator{ctrl: ctrl}
	mock.recorder = &MockGenesisBalancesIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenesisBalancesIterator) EXPECT() *MockGenesisBalancesIteratorMockRecorder {
	return m.recorder
}

// IterateGenesisBalances mocks base method.
func (m *MockGenesisBalancesIterator) IterateGenesisBalances(cdc codec.JSONCodec, appGenesis map[string]json.RawMessage, cb func(exported.GenesisBalance) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateGenesisBalances", cdc, appGenesis, cb)
}

// IterateGenesisBalances indicates an expected call of IterateGenesisBalances.
func (mr *MockGenesisBalancesIteratorMockRecorder) IterateGenesisBalances(cdc, appGenesis, cb interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateGenesisBalances", reflect.TypeOf((*MockGenesisBalancesIterator)(nil).IterateGenesisBalances), cdc, appGenesis, cb)
}
