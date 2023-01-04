// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	transaction "ikuzports/features/transaction"

	mock "github.com/stretchr/testify/mock"
)

// TransactionRepo is an autogenerated mock type for the RepositoryInterface type
type TransactionRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *TransactionRepo) Create(input transaction.TransactionCore) (int, error) {
	ret := _m.Called(input)

	var r0 int
	if rf, ok := ret.Get(0).(func(transaction.TransactionCore) int); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(transaction.TransactionCore) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *TransactionRepo) GetAll() ([]transaction.TransactionCore, error) {
	ret := _m.Called()

	var r0 []transaction.TransactionCore
	if rf, ok := ret.Get(0).(func() []transaction.TransactionCore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transaction.TransactionCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *TransactionRepo) GetByID(id int) (transaction.TransactionCore, error) {
	ret := _m.Called(id)

	var r0 transaction.TransactionCore
	if rf, ok := ret.Get(0).(func(int) transaction.TransactionCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(transaction.TransactionCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: input
func (_m *TransactionRepo) Update(input transaction.TransactionCore) (int, error) {
	ret := _m.Called(input)

	var r0 int
	if rf, ok := ret.Get(0).(func(transaction.TransactionCore) int); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(transaction.TransactionCore) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransactionRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionRepo creates a new instance of TransactionRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionRepo(t mockConstructorTestingTNewTransactionRepo) *TransactionRepo {
	mock := &TransactionRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
