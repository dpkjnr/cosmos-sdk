// Code generated by MockGen. DO NOT EDIT.
// Source: baseapp/abci_utils.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	crypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockValidatorStore is a mock of ValidatorStore interface.
type MockValidatorStore struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorStoreMockRecorder
}

// MockValidatorStoreMockRecorder is the mock recorder for MockValidatorStore.
type MockValidatorStoreMockRecorder struct {
	mock *MockValidatorStore
}

// NewMockValidatorStore creates a new mock instance.
func NewMockValidatorStore(ctrl *gomock.Controller) *MockValidatorStore {
	mock := &MockValidatorStore{ctrl: ctrl}
	mock.recorder = &MockValidatorStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatorStore) EXPECT() *MockValidatorStoreMockRecorder {
	return m.recorder
}

// GetPubKeyByConsAddr mocks base method.
func (m *MockValidatorStore) GetPubKeyByConsAddr(arg0 context.Context, arg1 types.ConsAddress) (crypto.PublicKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPubKeyByConsAddr", arg0, arg1)
	ret0, _ := ret[0].(crypto.PublicKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPubKeyByConsAddr indicates an expected call of GetPubKeyByConsAddr.
func (mr *MockValidatorStoreMockRecorder) GetPubKeyByConsAddr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPubKeyByConsAddr", reflect.TypeOf((*MockValidatorStore)(nil).GetPubKeyByConsAddr), arg0, arg1)
}

// MockGasTx is a mock of GasTx interface.
type MockGasTx struct {
	ctrl     *gomock.Controller
	recorder *MockGasTxMockRecorder
}

// MockGasTxMockRecorder is the mock recorder for MockGasTx.
type MockGasTxMockRecorder struct {
	mock *MockGasTx
}

// NewMockGasTx creates a new mock instance.
func NewMockGasTx(ctrl *gomock.Controller) *MockGasTx {
	mock := &MockGasTx{ctrl: ctrl}
	mock.recorder = &MockGasTxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGasTx) EXPECT() *MockGasTxMockRecorder {
	return m.recorder
}

// GetGas mocks base method.
func (m *MockGasTx) GetGas() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGas")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetGas indicates an expected call of GetGas.
func (mr *MockGasTxMockRecorder) GetGas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGas", reflect.TypeOf((*MockGasTx)(nil).GetGas))
}

// MockProposalTxVerifier is a mock of ProposalTxVerifier interface.
type MockProposalTxVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockProposalTxVerifierMockRecorder
}

// MockProposalTxVerifierMockRecorder is the mock recorder for MockProposalTxVerifier.
type MockProposalTxVerifierMockRecorder struct {
	mock *MockProposalTxVerifier
}

// NewMockProposalTxVerifier creates a new mock instance.
func NewMockProposalTxVerifier(ctrl *gomock.Controller) *MockProposalTxVerifier {
	mock := &MockProposalTxVerifier{ctrl: ctrl}
	mock.recorder = &MockProposalTxVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProposalTxVerifier) EXPECT() *MockProposalTxVerifierMockRecorder {
	return m.recorder
}

// PrepareProposalVerifyTx mocks base method.
func (m *MockProposalTxVerifier) PrepareProposalVerifyTx(arg0 types.Tx) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareProposalVerifyTx", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareProposalVerifyTx indicates an expected call of PrepareProposalVerifyTx.
func (mr *MockProposalTxVerifierMockRecorder) PrepareProposalVerifyTx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareProposalVerifyTx", reflect.TypeOf((*MockProposalTxVerifier)(nil).PrepareProposalVerifyTx), arg0)
}

// ProcessProposalVerifyTx mocks base method.
func (m *MockProposalTxVerifier) ProcessProposalVerifyTx(arg0 []byte) (types.Tx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessProposalVerifyTx", arg0)
	ret0, _ := ret[0].(types.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessProposalVerifyTx indicates an expected call of ProcessProposalVerifyTx.
func (mr *MockProposalTxVerifierMockRecorder) ProcessProposalVerifyTx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessProposalVerifyTx", reflect.TypeOf((*MockProposalTxVerifier)(nil).ProcessProposalVerifyTx), arg0)
}
