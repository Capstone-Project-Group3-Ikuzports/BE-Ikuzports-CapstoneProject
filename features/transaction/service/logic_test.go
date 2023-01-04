package service

import (
	"errors"
	"ikuzports/features/transaction"
	"ikuzports/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	repo := new(mocks.TransactionRepo)
	repo2 := new(mocks.ProductRepo)
	repo3 := new(mocks.UserRepo)

	t.Run("Success post transaction", func(t *testing.T) {
		input := transaction.TransactionCore{UserID: 1, TotalPrice: 10000, TotalQuantity: 1, ProductID: 1, TransactionID: "123123", StatusPayment: "paid", VirtualAccount: "12321", TransactionTime: "2022-09-08", OrderID: "2"}

		repo.On("Update", input).Return(1, nil).Once()
		srv := New(repo, repo2, repo3)
		err := srv.Update(input)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed post transaction", func(t *testing.T) {
		input := transaction.TransactionCore{UserID: 1, TotalPrice: 10000, TotalQuantity: 1, ProductID: 1, TransactionID: "123123", StatusPayment: "paid", VirtualAccount: "12321", TransactionTime: "2022-09-08", OrderID: "2"}

		repo.On("Update", input).Return(0, errors.New("failed")).Once()
		srv := New(repo, repo2, repo3)
		err := srv.Update(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	repo := new(mocks.TransactionRepo)
	repo2 := new(mocks.ProductRepo)
	repo3 := new(mocks.UserRepo)

	t.Run("Success get all transaction", func(t *testing.T) {
		data := []transaction.TransactionCore{{UserID: 1, TotalPrice: 10000, TotalQuantity: 1, ProductID: 1, TransactionID: "123123", StatusPayment: "paid", VirtualAccount: "12321", TransactionTime: "2022-09-08", OrderID: "2"}}

		repo.On("GetAll").Return(data, nil).Once()
		srv := New(repo, repo2, repo3)
		res, err := srv.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, data[0].UserID, res[0].UserID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all transaction", func(t *testing.T) {

		repo.On("GetAll").Return(nil, errors.New("failed")).Once()
		srv := New(repo, repo2, repo3)
		_, err := srv.GetAll()
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetByID(t *testing.T) {
	repo := new(mocks.TransactionRepo)
	repo2 := new(mocks.ProductRepo)
	repo3 := new(mocks.UserRepo)

	t.Run("Success get transaction by id", func(t *testing.T) {
		data := transaction.TransactionCore{UserID: 1, TotalPrice: 10000, TotalQuantity: 1, ProductID: 1, TransactionID: "123123", StatusPayment: "paid", VirtualAccount: "12321", TransactionTime: "2022-09-08", OrderID: "2"}

		repo.On("GetByID", 1).Return(data, nil).Once()
		srv := New(repo, repo2, repo3)
		res, err := srv.GetByID(1)
		assert.NoError(t, err)
		assert.Equal(t, data.TotalPrice, res.TotalPrice)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get transaction by id", func(t *testing.T) {

		repo.On("GetByID", 1).Return(transaction.TransactionCore{}, errors.New("failed")).Once()
		srv := New(repo, repo2, repo3)
		_, err := srv.GetByID(1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

// func TestCreate(t *testing.T) {
// 	repo := new(mocks.TransactionRepo)
// 	repo2 := new(mocks.ProductRepo)
// 	repo3 := new(mocks.UserRepo)

// t.Run("Success create transaction", func(t *testing.T) {

// 	input := transaction.TransactionCore{UserID: 1, TotalPrice: 10000, TotalQuantity: 1, ProductID: 1, TransactionID: "", StatusPayment: "", VirtualAccount: "12321", TransactionTime: "2022-09-08", OrderID: "orderID"}

// 	dataProduct := product.ProductCore{Name: "product", Price: 123, Description: "uk 32", UserID: 1, ItemCategoryID: 1, City: "jkt"}

// 	dataUser := user.Core{ID: 1, Name: "jojo", Email: "jojo", Password: "jojo", PhoneNumber: "089213912", Address: "disana", UserImage: "jpg", Gender: "Male"}

// 	repo2.On("GetByID", int(input.ProductID)).Return(dataProduct, nil).Once()
// 	repo3.On("GetById", int(input.UserID)).Return(dataUser, nil).Once()
// 	repo.On("Create", input).Return(1, nil).Once()

// 	srv := New(repo, repo2, repo3)
// 	_, err := srv.Create(input)
// 	assert.NoError(t, err)
// 	repo.AssertExpectations(t)
// })

// t.Run("Failed post transaction", func(t *testing.T) {
// 	input := transaction.TransactionCore{UserID: 1, TotalPrice: 10000, TotalQuantity: 1, ProductID: 1, TransactionID: "", StatusPayment: "", VirtualAccount: "12321", TransactionTime: "2022-09-08", OrderID: "orderID"}

// 	repo2.On("GetByID", 1).Return(product.ProductCore{}, errors.New("failed")).Once()
// 	srv := New(repo, repo2, repo3)
// 	_, err := srv.Create(input)
// 	assert.NotNil(t, err)
// 	repo.AssertExpectations(t)
// })

// t.Run("Failed post transaction", func(t *testing.T) {
// 	input := transaction.TransactionCore{UserID: 1, TotalPrice: 10000, TotalQuantity: 1, ProductID: 1, TransactionID: "", StatusPayment: "", VirtualAccount: "12321", TransactionTime: "2022-09-08", OrderID: "orderID"}

// 	dataProduct := product.ProductCore{Name: "product", Price: 123, Description: "uk 32", UserID: 1, ItemCategoryID: 1, City: "jkt"}

// 	repo2.On("GetByID", 1).Return(dataProduct, nil).Once()
// 	repo3.On("GetById", 1).Return(user.Core{}, errors.New("failed"))
// 	srv := New(repo, repo2, repo3)
// 	_, err := srv.Create(input)
// 	assert.NotNil(t, err)
// 	repo.AssertExpectations(t)
// })

// t.Run("Failed post transaction", func(t *testing.T) {
// 	input := transaction.TransactionCore{UserID: 1, TotalPrice: 10000, TotalQuantity: 1, ProductID: 1, TransactionID: "", StatusPayment: "", VirtualAccount: "12321", TransactionTime: "2022-09-08", OrderID: "orderID"}

// 	dataProduct := product.ProductCore{Name: "product", Price: 123, Description: "uk 32", UserID: 1, ItemCategoryID: 1, City: "jkt"}

// 	dataUser := user.Core{ID: 1, Name: "jojo", Email: "jojo", Password: "jojo", PhoneNumber: "089213912", Address: "disana", UserImage: "jpg", Gender: "Male"}

// 	repo2.On("GetByID", 1).Return(dataProduct, nil).Once()
// 	repo3.On("GetById", 1).Return(dataUser, nil).Once()
// 	repo.On("Create", input).Return(0, errors.New("failed")).Once()
// 	srv := New(repo, repo2, repo3)
// 	_, err := srv.Create(input)
// 	assert.NotNil(t, err)
// 	repo.AssertExpectations(t)
// })

// }
