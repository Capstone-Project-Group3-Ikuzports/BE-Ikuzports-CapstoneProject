package service

import (
	"errors"
	"ikuzports/features/product"
	"ikuzports/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.ProductRepo)

	t.Run("Succes get all clubs", func(t *testing.T) {
		data := []product.ProductCore{{Name: "product", Price: 123, Description: "uk 32", UserID: 1, ItemCategoryID: 1, City: "jkt"}}

		repo.On("GetAllFilter", 1, "name", "city", 9, 0).Return(data, 9, nil).Once()

		srv := New(repo)
		res, _, err := srv.GetAll(1, "name", "city", 1)
		assert.NoError(t, err)
		assert.Equal(t, data[0].Name, res[0].Name)
		repo.AssertExpectations(t)
	})

	t.Run("Succes get all clubs", func(t *testing.T) {
		data := []product.ProductCore{{Name: "product", Price: 123, Description: "uk 32", UserID: 1, ItemCategoryID: 1, City: "jkt"}}

		repo.On("GetAll", 9, 9).Return(data, 10, nil).Once()

		srv := New(repo)
		res, _, err := srv.GetAll(0, "", "", 2)
		assert.NoError(t, err)
		assert.Equal(t, data[0].Name, res[0].Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all clubs", func(t *testing.T) {
		repo.On("GetAllFilter", 1, "name", "city", 9, 0).Return(nil, 0, errors.New("faield")).Once()

		srv := New(repo)
		_, _, err := srv.GetAll(1, "name", "city", 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	repo := new(mocks.ProductRepo)

	t.Run("Success Create product", func(t *testing.T) {
		data := product.ProductCore{Name: "product", Price: 123, Description: "uk 32", UserID: 1, ItemCategoryID: 1, City: "jkt"}
		repo.On("Create", data).Return(1, nil).Once()
		srv := New(repo)
		err := srv.Create(data)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create product", func(t *testing.T) {
		data := product.ProductCore{Name: "product", Price: 123, Description: "uk 32", UserID: 1, ItemCategoryID: 1, City: "jkt"}
		repo.On("Create", data).Return(0, errors.New("failed to insert data, error query")).Once()
		srv := New(repo)
		err := srv.Create(data)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create product, validate error", func(t *testing.T) {
		data := product.ProductCore{Price: 123, Description: "uk 32", UserID: 1, ItemCategoryID: 1, City: "jkt"}

		srv := New(repo)
		err := srv.Create(data)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.ProductRepo)

	t.Run("Success get product", func(t *testing.T) {
		input := product.ProductCore{Name: "product", Price: 123, Description: "uk 32", UserID: 1, ItemCategoryID: 1, City: "jkt"}

		repo.On("GetByID", 1).Return(input, nil).Once()
		srv := New(repo)
		response, err := srv.GetByID(1)

		assert.NoError(t, err)
		assert.Equal(t, input.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get product by id", func(t *testing.T) {
		inputRepo := product.ProductCore{Name: "product", Price: 123, Description: "uk 32", UserID: 1, ItemCategoryID: 1, City: "jkt"}

		repo.On("GetByID", 1).Return(product.ProductCore{}, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetByID(1)

		assert.NotNil(t, err)
		assert.NotEqual(t, inputRepo.Price, response.Price)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := new(mocks.ProductRepo)

	t.Run("Success update product", func(t *testing.T) {
		input := product.ProductCore{Name: "product", Price: 123, Description: "uk 32", UserID: 1, ItemCategoryID: 1, City: "jkt"}

		repo.On("Update", 1, input).Return(1, nil).Once()
		srv := New(repo)
		err := srv.Update(1, input)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed update product by id", func(t *testing.T) {
		input := product.ProductCore{Name: "product", Price: 123, Description: "uk 32", UserID: 1, ItemCategoryID: 1, City: "jkt"}

		repo.On("Update", 1, input).Return(0, errors.New("failed")).Once()
		srv := New(repo)
		err := srv.Update(1, input)

		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.ProductRepo)

	t.Run("Success Delete product", func(t *testing.T) {

		repo.On("Delete", 1).Return(1, nil).Once()
		srv := New(repo)
		err := srv.Delete(1)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Delete product by id", func(t *testing.T) {

		repo.On("Delete", 1).Return(0, errors.New("failed")).Once()
		srv := New(repo)
		err := srv.Delete(1)

		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetImages(t *testing.T) {
	repo := new(mocks.ProductRepo)

	t.Run("Success get images product", func(t *testing.T) {
		input := []product.ProductImage{{ID: 1, URL: "jpg"}}

		repo.On("GetImages", 1).Return(input, nil).Once()
		srv := New(repo)
		response, err := srv.GetImages(1)

		assert.NoError(t, err)
		assert.Equal(t, input[0].URL, response[0].URL)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get images product", func(t *testing.T) {

		repo.On("GetImages", 1).Return(nil, errors.New("failed")).Once()
		srv := New(repo)
		_, err := srv.GetImages(1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
