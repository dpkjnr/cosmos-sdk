// Code generated by MockGen. DO NOT EDIT.
// Source: x/staking/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"

	stakingv1beta1 "cosmossdk.io/api/cosmos/staking/v1beta1"
	address "cosmossdk.io/core/address"
	math "cosmossdk.io/math"
	crypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/staking/types"
	gomock "github.com/golang/mock/gomock"
)

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

// AddressCodec mocks base method.
func (m *MockAccountKeeper) AddressCodec() address.Codec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddressCodec")
	ret0, _ := ret[0].(address.Codec)
	return ret0
}

// AddressCodec indicates an expected call of AddressCodec.
func (mr *MockAccountKeeperMockRecorder) AddressCodec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddressCodec", reflect.TypeOf((*MockAccountKeeper)(nil).AddressCodec))
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(arg0 context.Context, arg1 types.AccAddress) types.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(types.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), arg0, arg1)
}

// GetModuleAccount mocks base method.
func (m *MockAccountKeeper) GetModuleAccount(arg0 context.Context, arg1 string) types.ModuleAccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAccount", arg0, arg1)
	ret0, _ := ret[0].(types.ModuleAccountI)
	return ret0
}

// GetModuleAccount indicates an expected call of GetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) GetModuleAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAccount), arg0, arg1)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(arg0 string) types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", arg0)
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), arg0)
}

// IterateAccounts mocks base method.
func (m *MockAccountKeeper) IterateAccounts(arg0 context.Context, arg1 func(types.AccountI) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateAccounts", arg0, arg1)
}

// IterateAccounts indicates an expected call of IterateAccounts.
func (mr *MockAccountKeeperMockRecorder) IterateAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateAccounts", reflect.TypeOf((*MockAccountKeeper)(nil).IterateAccounts), arg0, arg1)
}

// SetModuleAccount mocks base method.
func (m *MockAccountKeeper) SetModuleAccount(arg0 context.Context, arg1 types.ModuleAccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetModuleAccount", arg0, arg1)
}

// SetModuleAccount indicates an expected call of SetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) SetModuleAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetModuleAccount), arg0, arg1)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// BurnCoins mocks base method.
func (m *MockBankKeeper) BurnCoins(arg0 context.Context, arg1 string, arg2 types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BurnCoins", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// BurnCoins indicates an expected call of BurnCoins.
func (mr *MockBankKeeperMockRecorder) BurnCoins(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BurnCoins", reflect.TypeOf((*MockBankKeeper)(nil).BurnCoins), arg0, arg1, arg2)
}

// DelegateCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) DelegateCoinsFromAccountToModule(arg0 context.Context, arg1 types.AccAddress, arg2 string, arg3 types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelegateCoinsFromAccountToModule", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelegateCoinsFromAccountToModule indicates an expected call of DelegateCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) DelegateCoinsFromAccountToModule(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelegateCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).DelegateCoinsFromAccountToModule), arg0, arg1, arg2, arg3)
}

// GetAllBalances mocks base method.
func (m *MockBankKeeper) GetAllBalances(arg0 context.Context, arg1 types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBalances", arg0, arg1)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// GetAllBalances indicates an expected call of GetAllBalances.
func (mr *MockBankKeeperMockRecorder) GetAllBalances(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBalances", reflect.TypeOf((*MockBankKeeper)(nil).GetAllBalances), arg0, arg1)
}

// GetBalance mocks base method.
func (m *MockBankKeeper) GetBalance(arg0 context.Context, arg1 types.AccAddress, arg2 string) types.Coin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", arg0, arg1, arg2)
	ret0, _ := ret[0].(types.Coin)
	return ret0
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockBankKeeperMockRecorder) GetBalance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockBankKeeper)(nil).GetBalance), arg0, arg1, arg2)
}

// GetSupply mocks base method.
func (m *MockBankKeeper) GetSupply(arg0 context.Context, arg1 string) types.Coin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupply", arg0, arg1)
	ret0, _ := ret[0].(types.Coin)
	return ret0
}

// GetSupply indicates an expected call of GetSupply.
func (mr *MockBankKeeperMockRecorder) GetSupply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupply", reflect.TypeOf((*MockBankKeeper)(nil).GetSupply), arg0, arg1)
}

// LockedCoins mocks base method.
func (m *MockBankKeeper) LockedCoins(arg0 context.Context, arg1 types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockedCoins", arg0, arg1)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// LockedCoins indicates an expected call of LockedCoins.
func (mr *MockBankKeeperMockRecorder) LockedCoins(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockedCoins", reflect.TypeOf((*MockBankKeeper)(nil).LockedCoins), arg0, arg1)
}

// SendCoinsFromModuleToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToModule(arg0 context.Context, arg1, arg2 string, arg3 types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToModule", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToModule indicates an expected call of SendCoinsFromModuleToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToModule(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToModule), arg0, arg1, arg2, arg3)
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(arg0 context.Context, arg1 types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", arg0, arg1)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), arg0, arg1)
}

// UndelegateCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) UndelegateCoinsFromModuleToAccount(arg0 context.Context, arg1 string, arg2 types.AccAddress, arg3 types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UndelegateCoinsFromModuleToAccount", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UndelegateCoinsFromModuleToAccount indicates an expected call of UndelegateCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) UndelegateCoinsFromModuleToAccount(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UndelegateCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).UndelegateCoinsFromModuleToAccount), arg0, arg1, arg2, arg3)
}

// MockValidatorSet is a mock of ValidatorSet interface.
type MockValidatorSet struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorSetMockRecorder
}

// MockValidatorSetMockRecorder is the mock recorder for MockValidatorSet.
type MockValidatorSetMockRecorder struct {
	mock *MockValidatorSet
}

// NewMockValidatorSet creates a new mock instance.
func NewMockValidatorSet(ctrl *gomock.Controller) *MockValidatorSet {
	mock := &MockValidatorSet{ctrl: ctrl}
	mock.recorder = &MockValidatorSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatorSet) EXPECT() *MockValidatorSetMockRecorder {
	return m.recorder
}

// Delegation mocks base method.
func (m *MockValidatorSet) Delegation(arg0 context.Context, arg1 types.AccAddress, arg2 types.ValAddress) (types0.DelegationI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delegation", arg0, arg1, arg2)
	ret0, _ := ret[0].(types0.DelegationI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delegation indicates an expected call of Delegation.
func (mr *MockValidatorSetMockRecorder) Delegation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delegation", reflect.TypeOf((*MockValidatorSet)(nil).Delegation), arg0, arg1, arg2)
}

// GetPubKeyByConsAddr mocks base method.
func (m *MockValidatorSet) GetPubKeyByConsAddr(arg0 context.Context, arg1 types.ConsAddress) (crypto.PublicKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPubKeyByConsAddr", arg0, arg1)
	ret0, _ := ret[0].(crypto.PublicKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPubKeyByConsAddr indicates an expected call of GetPubKeyByConsAddr.
func (mr *MockValidatorSetMockRecorder) GetPubKeyByConsAddr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPubKeyByConsAddr", reflect.TypeOf((*MockValidatorSet)(nil).GetPubKeyByConsAddr), arg0, arg1)
}

// IterateBondedValidatorsByPower mocks base method.
func (m *MockValidatorSet) IterateBondedValidatorsByPower(arg0 context.Context, arg1 func(int64, types0.ValidatorI) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IterateBondedValidatorsByPower", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IterateBondedValidatorsByPower indicates an expected call of IterateBondedValidatorsByPower.
func (mr *MockValidatorSetMockRecorder) IterateBondedValidatorsByPower(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateBondedValidatorsByPower", reflect.TypeOf((*MockValidatorSet)(nil).IterateBondedValidatorsByPower), arg0, arg1)
}

// IterateLastValidators mocks base method.
func (m *MockValidatorSet) IterateLastValidators(arg0 context.Context, arg1 func(int64, types0.ValidatorI) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IterateLastValidators", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IterateLastValidators indicates an expected call of IterateLastValidators.
func (mr *MockValidatorSetMockRecorder) IterateLastValidators(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateLastValidators", reflect.TypeOf((*MockValidatorSet)(nil).IterateLastValidators), arg0, arg1)
}

// IterateValidators mocks base method.
func (m *MockValidatorSet) IterateValidators(arg0 context.Context, arg1 func(int64, types0.ValidatorI) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IterateValidators", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IterateValidators indicates an expected call of IterateValidators.
func (mr *MockValidatorSetMockRecorder) IterateValidators(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateValidators", reflect.TypeOf((*MockValidatorSet)(nil).IterateValidators), arg0, arg1)
}

// Jail mocks base method.
func (m *MockValidatorSet) Jail(arg0 context.Context, arg1 types.ConsAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Jail", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Jail indicates an expected call of Jail.
func (mr *MockValidatorSetMockRecorder) Jail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Jail", reflect.TypeOf((*MockValidatorSet)(nil).Jail), arg0, arg1)
}

// MaxValidators mocks base method.
func (m *MockValidatorSet) MaxValidators(arg0 context.Context) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxValidators", arg0)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaxValidators indicates an expected call of MaxValidators.
func (mr *MockValidatorSetMockRecorder) MaxValidators(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxValidators", reflect.TypeOf((*MockValidatorSet)(nil).MaxValidators), arg0)
}

// Slash mocks base method.
func (m *MockValidatorSet) Slash(arg0 context.Context, arg1 types.ConsAddress, arg2, arg3 int64, arg4 math.LegacyDec) (math.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Slash", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(math.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Slash indicates an expected call of Slash.
func (mr *MockValidatorSetMockRecorder) Slash(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Slash", reflect.TypeOf((*MockValidatorSet)(nil).Slash), arg0, arg1, arg2, arg3, arg4)
}

// SlashWithInfractionReason mocks base method.
func (m *MockValidatorSet) SlashWithInfractionReason(arg0 context.Context, arg1 types.ConsAddress, arg2, arg3 int64, arg4 math.LegacyDec, arg5 stakingv1beta1.Infraction) (math.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SlashWithInfractionReason", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(math.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SlashWithInfractionReason indicates an expected call of SlashWithInfractionReason.
func (mr *MockValidatorSetMockRecorder) SlashWithInfractionReason(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SlashWithInfractionReason", reflect.TypeOf((*MockValidatorSet)(nil).SlashWithInfractionReason), arg0, arg1, arg2, arg3, arg4, arg5)
}

// StakingTokenSupply mocks base method.
func (m *MockValidatorSet) StakingTokenSupply(arg0 context.Context) (math.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StakingTokenSupply", arg0)
	ret0, _ := ret[0].(math.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StakingTokenSupply indicates an expected call of StakingTokenSupply.
func (mr *MockValidatorSetMockRecorder) StakingTokenSupply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StakingTokenSupply", reflect.TypeOf((*MockValidatorSet)(nil).StakingTokenSupply), arg0)
}

// TotalBondedTokens mocks base method.
func (m *MockValidatorSet) TotalBondedTokens(arg0 context.Context) (math.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalBondedTokens", arg0)
	ret0, _ := ret[0].(math.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TotalBondedTokens indicates an expected call of TotalBondedTokens.
func (mr *MockValidatorSetMockRecorder) TotalBondedTokens(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalBondedTokens", reflect.TypeOf((*MockValidatorSet)(nil).TotalBondedTokens), arg0)
}

// Unjail mocks base method.
func (m *MockValidatorSet) Unjail(arg0 context.Context, arg1 types.ConsAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unjail", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unjail indicates an expected call of Unjail.
func (mr *MockValidatorSetMockRecorder) Unjail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unjail", reflect.TypeOf((*MockValidatorSet)(nil).Unjail), arg0, arg1)
}

// Validator mocks base method.
func (m *MockValidatorSet) Validator(arg0 context.Context, arg1 types.ValAddress) (types0.ValidatorI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validator", arg0, arg1)
	ret0, _ := ret[0].(types0.ValidatorI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Validator indicates an expected call of Validator.
func (mr *MockValidatorSetMockRecorder) Validator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validator", reflect.TypeOf((*MockValidatorSet)(nil).Validator), arg0, arg1)
}

// ValidatorByConsAddr mocks base method.
func (m *MockValidatorSet) ValidatorByConsAddr(arg0 context.Context, arg1 types.ConsAddress) (types0.ValidatorI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorByConsAddr", arg0, arg1)
	ret0, _ := ret[0].(types0.ValidatorI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorByConsAddr indicates an expected call of ValidatorByConsAddr.
func (mr *MockValidatorSetMockRecorder) ValidatorByConsAddr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorByConsAddr", reflect.TypeOf((*MockValidatorSet)(nil).ValidatorByConsAddr), arg0, arg1)
}

// MockDelegationSet is a mock of DelegationSet interface.
type MockDelegationSet struct {
	ctrl     *gomock.Controller
	recorder *MockDelegationSetMockRecorder
}

// MockDelegationSetMockRecorder is the mock recorder for MockDelegationSet.
type MockDelegationSetMockRecorder struct {
	mock *MockDelegationSet
}

// NewMockDelegationSet creates a new mock instance.
func NewMockDelegationSet(ctrl *gomock.Controller) *MockDelegationSet {
	mock := &MockDelegationSet{ctrl: ctrl}
	mock.recorder = &MockDelegationSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDelegationSet) EXPECT() *MockDelegationSetMockRecorder {
	return m.recorder
}

// GetValidatorSet mocks base method.
func (m *MockDelegationSet) GetValidatorSet() types0.ValidatorSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorSet")
	ret0, _ := ret[0].(types0.ValidatorSet)
	return ret0
}

// GetValidatorSet indicates an expected call of GetValidatorSet.
func (mr *MockDelegationSetMockRecorder) GetValidatorSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorSet", reflect.TypeOf((*MockDelegationSet)(nil).GetValidatorSet))
}

// IterateDelegations mocks base method.
func (m *MockDelegationSet) IterateDelegations(arg0 context.Context, arg1 types.AccAddress, arg2 func(int64, types0.DelegationI) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IterateDelegations", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// IterateDelegations indicates an expected call of IterateDelegations.
func (mr *MockDelegationSetMockRecorder) IterateDelegations(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateDelegations", reflect.TypeOf((*MockDelegationSet)(nil).IterateDelegations), arg0, arg1, arg2)
}

// MockStakingHooks is a mock of StakingHooks interface.
type MockStakingHooks struct {
	ctrl     *gomock.Controller
	recorder *MockStakingHooksMockRecorder
}

// MockStakingHooksMockRecorder is the mock recorder for MockStakingHooks.
type MockStakingHooksMockRecorder struct {
	mock *MockStakingHooks
}

// NewMockStakingHooks creates a new mock instance.
func NewMockStakingHooks(ctrl *gomock.Controller) *MockStakingHooks {
	mock := &MockStakingHooks{ctrl: ctrl}
	mock.recorder = &MockStakingHooksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakingHooks) EXPECT() *MockStakingHooksMockRecorder {
	return m.recorder
}

// AfterDelegationModified mocks base method.
func (m *MockStakingHooks) AfterDelegationModified(arg0 context.Context, arg1 types.AccAddress, arg2 types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterDelegationModified", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterDelegationModified indicates an expected call of AfterDelegationModified.
func (mr *MockStakingHooksMockRecorder) AfterDelegationModified(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterDelegationModified", reflect.TypeOf((*MockStakingHooks)(nil).AfterDelegationModified), arg0, arg1, arg2)
}

// AfterUnbondingInitiated mocks base method.
func (m *MockStakingHooks) AfterUnbondingInitiated(arg0 context.Context, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterUnbondingInitiated", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterUnbondingInitiated indicates an expected call of AfterUnbondingInitiated.
func (mr *MockStakingHooksMockRecorder) AfterUnbondingInitiated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterUnbondingInitiated", reflect.TypeOf((*MockStakingHooks)(nil).AfterUnbondingInitiated), arg0, arg1)
}

// AfterValidatorBeginUnbonding mocks base method.
func (m *MockStakingHooks) AfterValidatorBeginUnbonding(arg0 context.Context, arg1 types.ConsAddress, arg2 types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterValidatorBeginUnbonding", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterValidatorBeginUnbonding indicates an expected call of AfterValidatorBeginUnbonding.
func (mr *MockStakingHooksMockRecorder) AfterValidatorBeginUnbonding(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterValidatorBeginUnbonding", reflect.TypeOf((*MockStakingHooks)(nil).AfterValidatorBeginUnbonding), arg0, arg1, arg2)
}

// AfterValidatorBonded mocks base method.
func (m *MockStakingHooks) AfterValidatorBonded(arg0 context.Context, arg1 types.ConsAddress, arg2 types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterValidatorBonded", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterValidatorBonded indicates an expected call of AfterValidatorBonded.
func (mr *MockStakingHooksMockRecorder) AfterValidatorBonded(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterValidatorBonded", reflect.TypeOf((*MockStakingHooks)(nil).AfterValidatorBonded), arg0, arg1, arg2)
}

// AfterValidatorCreated mocks base method.
func (m *MockStakingHooks) AfterValidatorCreated(arg0 context.Context, arg1 types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterValidatorCreated", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterValidatorCreated indicates an expected call of AfterValidatorCreated.
func (mr *MockStakingHooksMockRecorder) AfterValidatorCreated(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterValidatorCreated", reflect.TypeOf((*MockStakingHooks)(nil).AfterValidatorCreated), arg0, arg1)
}

// AfterValidatorRemoved mocks base method.
func (m *MockStakingHooks) AfterValidatorRemoved(arg0 context.Context, arg1 types.ConsAddress, arg2 types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterValidatorRemoved", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterValidatorRemoved indicates an expected call of AfterValidatorRemoved.
func (mr *MockStakingHooksMockRecorder) AfterValidatorRemoved(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterValidatorRemoved", reflect.TypeOf((*MockStakingHooks)(nil).AfterValidatorRemoved), arg0, arg1, arg2)
}

// BeforeDelegationCreated mocks base method.
func (m *MockStakingHooks) BeforeDelegationCreated(arg0 context.Context, arg1 types.AccAddress, arg2 types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeDelegationCreated", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeforeDelegationCreated indicates an expected call of BeforeDelegationCreated.
func (mr *MockStakingHooksMockRecorder) BeforeDelegationCreated(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeDelegationCreated", reflect.TypeOf((*MockStakingHooks)(nil).BeforeDelegationCreated), arg0, arg1, arg2)
}

// BeforeDelegationRemoved mocks base method.
func (m *MockStakingHooks) BeforeDelegationRemoved(arg0 context.Context, arg1 types.AccAddress, arg2 types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeDelegationRemoved", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeforeDelegationRemoved indicates an expected call of BeforeDelegationRemoved.
func (mr *MockStakingHooksMockRecorder) BeforeDelegationRemoved(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeDelegationRemoved", reflect.TypeOf((*MockStakingHooks)(nil).BeforeDelegationRemoved), arg0, arg1, arg2)
}

// BeforeDelegationSharesModified mocks base method.
func (m *MockStakingHooks) BeforeDelegationSharesModified(arg0 context.Context, arg1 types.AccAddress, arg2 types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeDelegationSharesModified", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeforeDelegationSharesModified indicates an expected call of BeforeDelegationSharesModified.
func (mr *MockStakingHooksMockRecorder) BeforeDelegationSharesModified(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeDelegationSharesModified", reflect.TypeOf((*MockStakingHooks)(nil).BeforeDelegationSharesModified), arg0, arg1, arg2)
}

// BeforeValidatorModified mocks base method.
func (m *MockStakingHooks) BeforeValidatorModified(arg0 context.Context, arg1 types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeValidatorModified", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeforeValidatorModified indicates an expected call of BeforeValidatorModified.
func (mr *MockStakingHooksMockRecorder) BeforeValidatorModified(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeValidatorModified", reflect.TypeOf((*MockStakingHooks)(nil).BeforeValidatorModified), arg0, arg1)
}

// BeforeValidatorSlashed mocks base method.
func (m *MockStakingHooks) BeforeValidatorSlashed(arg0 context.Context, arg1 types.ValAddress, arg2 math.LegacyDec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeforeValidatorSlashed", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// BeforeValidatorSlashed indicates an expected call of BeforeValidatorSlashed.
func (mr *MockStakingHooksMockRecorder) BeforeValidatorSlashed(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeforeValidatorSlashed", reflect.TypeOf((*MockStakingHooks)(nil).BeforeValidatorSlashed), arg0, arg1, arg2)
}
