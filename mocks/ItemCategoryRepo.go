// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	itemcategory "ikuzports/features/itemCategory"

	mock "github.com/stretchr/testify/mock"
)

// ItemCategoryRepo is an autogenerated mock type for the RepositoryInterface type
type ItemCategoryRepo struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *ItemCategoryRepo) GetAll() ([]itemcategory.ItemCategoryCore, error) {
	ret := _m.Called()

	var r0 []itemcategory.ItemCategoryCore
	if rf, ok := ret.Get(0).(func() []itemcategory.ItemCategoryCore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]itemcategory.ItemCategoryCore)
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

type mockConstructorTestingTNewItemCategoryRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewItemCategoryRepo creates a new instance of ItemCategoryRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewItemCategoryRepo(t mockConstructorTestingTNewItemCategoryRepo) *ItemCategoryRepo {
	mock := &ItemCategoryRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}