// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	clubMember "ikuzports/features/clubMember"
	event "ikuzports/features/event"

	mock "github.com/stretchr/testify/mock"

	transaction "ikuzports/features/transaction"

	user "ikuzports/features/user"
)

// UserRepo is an autogenerated mock type for the RepositoryInterface type
type UserRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *UserRepo) Create(input user.Core) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.Core) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *UserRepo) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindUser provides a mock function with given fields: email
func (_m *UserRepo) FindUser(email string) (user.Core, error) {
	ret := _m.Called(email)

	var r0 user.Core
	if rf, ok := ret.Get(0).(func(string) user.Core); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *UserRepo) GetAll() ([]user.Core, error) {
	ret := _m.Called()

	var r0 []user.Core
	if rf, ok := ret.Get(0).(func() []user.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.Core)
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

// GetById provides a mock function with given fields: id
func (_m *UserRepo) GetById(id int) (user.Core, error) {
	ret := _m.Called(id)

	var r0 user.Core
	if rf, ok := ret.Get(0).(func(int) user.Core); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClubs provides a mock function with given fields: id
func (_m *UserRepo) GetClubs(id int) ([]clubMember.Core, error) {
	ret := _m.Called(id)

	var r0 []clubMember.Core
	if rf, ok := ret.Get(0).(func(int) []clubMember.Core); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]clubMember.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEvents provides a mock function with given fields: id
func (_m *UserRepo) GetEvents(id int) ([]event.EventCore, error) {
	ret := _m.Called(id)

	var r0 []event.EventCore
	if rf, ok := ret.Get(0).(func(int) []event.EventCore); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]event.EventCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProducts provides a mock function with given fields: id
func (_m *UserRepo) GetProducts(id int) ([]user.ProductCore, error) {
	ret := _m.Called(id)

	var r0 []user.ProductCore
	if rf, ok := ret.Get(0).(func(int) []user.ProductCore); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.ProductCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactions provides a mock function with given fields: id
func (_m *UserRepo) GetTransactions(id int) ([]transaction.TransactionCore, error) {
	ret := _m.Called(id)

	var r0 []transaction.TransactionCore
	if rf, ok := ret.Get(0).(func(int) []transaction.TransactionCore); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]transaction.TransactionCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: input, id
func (_m *UserRepo) Update(input user.Core, id int) error {
	ret := _m.Called(input, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(user.Core, int) error); ok {
		r0 = rf(input, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepo creates a new instance of UserRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepo(t mockConstructorTestingTNewUserRepo) *UserRepo {
	mock := &UserRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
