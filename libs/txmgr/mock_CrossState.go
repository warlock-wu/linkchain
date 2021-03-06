// Code generated by mockery v1.0.0. DO NOT EDIT.

package txmgr

import mock "github.com/stretchr/testify/mock"
import types "github.com/lianxiangcloud/linkchain/types"

// MockCrossState is an autogenerated mock type for the CrossState type
type MockCrossState struct {
	mock.Mock
}

// AddSpecialTx provides a mock function with given fields: txs
func (_m *MockCrossState) AddSpecialTx(txs []types.Tx) {
	_m.Called(txs)
}

// GetTxValidators provides a mock function with given fields: height
func (_m *MockCrossState) GetTxValidators(height uint64) ([]*types.Validator, error) {
	ret := _m.Called(height)

	var r0 []*types.Validator
	if rf, ok := ret.Get(0).(func(uint64) []*types.Validator); ok {
		r0 = rf(height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Validator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveTxEntry provides a mock function with given fields: block, txsResult
func (_m *MockCrossState) SaveTxEntry(block *types.Block, txsResult *types.TxsResult) {
	_m.Called(block, txsResult)
}

// SaveTxValidators provides a mock function with given fields: height, validators
func (_m *MockCrossState) SaveTxValidators(height uint64, validators []*types.Validator) {
	_m.Called(height, validators)
}

// Sync provides a mock function with given fields:
func (_m *MockCrossState) Sync() {
	_m.Called()
}
