// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	product "ikuzports/features/product"

	mock "github.com/stretchr/testify/mock"
)

// ProductRepo is an autogenerated mock type for the RepositoryInterface type
type ProductRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *ProductRepo) Create(input product.ProductCore) (int, error) {
	ret := _m.Called(input)

	var r0 int
	if rf, ok := ret.Get(0).(func(product.ProductCore) int); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(product.ProductCore) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *ProductRepo) Delete(id int) (int, error) {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: limit, offset
func (_m *ProductRepo) GetAll(limit int, offset int) ([]product.ProductCore, int, error) {
	ret := _m.Called(limit, offset)

	var r0 []product.ProductCore
	if rf, ok := ret.Get(0).(func(int, int) []product.ProductCore); ok {
		r0 = rf(limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.ProductCore)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int) int); ok {
		r1 = rf(limit, offset)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int) error); ok {
		r2 = rf(limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAllFilter provides a mock function with given fields: queryItemCategoryID, queryCity, queryName, offet, limit
func (_m *ProductRepo) GetAllFilter(queryItemCategoryID int, queryCity string, queryName string, offet int, limit int) ([]product.ProductCore, int, error) {
	ret := _m.Called(queryItemCategoryID, queryCity, queryName, offet, limit)

	var r0 []product.ProductCore
	if rf, ok := ret.Get(0).(func(int, string, string, int, int) []product.ProductCore); ok {
		r0 = rf(queryItemCategoryID, queryCity, queryName, offet, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.ProductCore)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, string, string, int, int) int); ok {
		r1 = rf(queryItemCategoryID, queryCity, queryName, offet, limit)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, string, string, int, int) error); ok {
		r2 = rf(queryItemCategoryID, queryCity, queryName, offet, limit)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetByID provides a mock function with given fields: id
func (_m *ProductRepo) GetByID(id int) (product.ProductCore, error) {
	ret := _m.Called(id)

	var r0 product.ProductCore
	if rf, ok := ret.Get(0).(func(int) product.ProductCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(product.ProductCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImages provides a mock function with given fields: id
func (_m *ProductRepo) GetImages(id int) ([]product.ProductImage, error) {
	ret := _m.Called(id)

	var r0 []product.ProductImage
	if rf, ok := ret.Get(0).(func(int) []product.ProductImage); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.ProductImage)
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

// Update provides a mock function with given fields: id, input
func (_m *ProductRepo) Update(id int, input product.ProductCore) (int, error) {
	ret := _m.Called(id, input)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, product.ProductCore) int); ok {
		r0 = rf(id, input)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, product.ProductCore) error); ok {
		r1 = rf(id, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductRepo creates a new instance of ProductRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductRepo(t mockConstructorTestingTNewProductRepo) *ProductRepo {
	mock := &ProductRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
