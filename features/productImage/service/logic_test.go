package service

import (
	"errors"
	"ikuzports/features/productImage"
	"ikuzports/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.ProductImageRepo)

	t.Run("Success post image product", func(t *testing.T) {
		input := productImage.ProductImageCore{URL: "jpg", ProductID: 1}

		repo.On("Create", input).Return(1, nil).Once()
		srv := New(repo)
		err := srv.Create(input)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed post image product", func(t *testing.T) {
		input := productImage.ProductImageCore{URL: "jpg", ProductID: 1}

		repo.On("Create", input).Return(0, errors.New("failed")).Once()
		srv := New(repo)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	repo := new(mocks.ProductImageRepo)

	t.Run("Success get all product image", func(t *testing.T) {
		data := []productImage.ProductImageCore{{URL: "jpg", ProductID: 1}}

		repo.On("GetAll").Return(data, nil).Once()
		srv := New(repo)
		res, err := srv.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, data[0].URL, res[0].URL)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all product image", func(t *testing.T) {

		repo.On("GetAll").Return(nil, errors.New("failed")).Once()
		srv := New(repo)
		_, err := srv.GetAll()
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	repo := new(mocks.ProductImageRepo)

	t.Run("Success get product image by id", func(t *testing.T) {
		data := productImage.ProductImageCore{URL: "jpg", ProductID: 1}

		repo.On("GetByID", 1).Return(data, nil).Once()
		srv := New(repo)
		res, err := srv.GetByID(1)
		assert.NoError(t, err)
		assert.Equal(t, data.URL, res.URL)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get product image by id", func(t *testing.T) {

		repo.On("GetByID", 1).Return(productImage.ProductImageCore{}, errors.New("failed")).Once()
		srv := New(repo)
		_, err := srv.GetByID(1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
