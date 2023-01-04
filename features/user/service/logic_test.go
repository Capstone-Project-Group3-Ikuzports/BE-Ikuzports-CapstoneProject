package service

import (
	"errors"
	"ikuzports/features/clubMember"
	"ikuzports/features/event"
	"ikuzports/features/transaction"
	"ikuzports/features/user"
	"ikuzports/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.UserRepo)
	t.Run("Success get all user", func(t *testing.T) {
		inputRepo := []user.Core{{ID: 1, Name: "alta", Email: "alta", Password: "alta", PhoneNumber: "089213912", Address: "disana", UserImage: "jpg", Gender: "Male"}}
		repo.On("GetAll").Return(inputRepo, nil).Once()
		srv := New(repo)
		response, err := srv.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].Name, response[0].Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.UserRepo)
	t.Run("Success get user", func(t *testing.T) {
		inputRepo := user.Core{ID: 1, Name: "jojo", Email: "jojo", Password: "jojo", PhoneNumber: "089213912", Address: "disana", UserImage: "jpg", Gender: "Male"}
		repo.On("GetById", 1).Return(inputRepo, nil).Once()
		srv := New(repo)
		response, err := srv.GetById(1)
		assert.NoError(t, err)
		assert.Equal(t, inputRepo.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get by id", func(t *testing.T) {
		inputRepo := user.Core{ID: 1, Name: "alta", Email: "alta", Password: "alta", PhoneNumber: "089213912", Address: "disana", UserImage: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg", Gender: "Male"}
		repo.On("GetById", 1).Return(user.Core{}, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetById(1)
		assert.NotNil(t, err)
		assert.NotEqual(t, inputRepo.Name, response.Name)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.UserRepo)
	t.Run("Success delete User", func(t *testing.T) {
		repo.On("Delete", 1).Return(nil).Once()
		srv := New(repo)
		err := srv.Delete(1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete User", func(t *testing.T) {
		repo.On("Delete", 1).Return(errors.New("failed")).Once()
		srv := New(repo)
		err := srv.Delete(1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetClubs(t *testing.T) {
	repo := new(mocks.UserRepo)
	t.Run("Success get all user clubs", func(t *testing.T) {
		inputRepo := []clubMember.Core{{ID: 1, UserID: 1, ClubID: 1, Status: "member"}}
		repo.On("GetClubs", 1).Return(inputRepo, nil).Once()
		srv := New(repo)
		response, err := srv.GetClubs(1)
		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].ID, response[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all clubs", func(t *testing.T) {
		repo.On("GetClubs", 1).Return(nil, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetClubs(1)
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

func TestGetProducts(t *testing.T) {
	repo := new(mocks.UserRepo)
	t.Run("Success get all user products", func(t *testing.T) {
		inputRepo := []user.ProductCore{{ID: 1, Name: "sepatu spec", Price: 12000, Description: "ukuran 30", UserID: 1, ItemCategoryID: 1, City: "Jakarta"}}
		repo.On("GetProducts", 1).Return(inputRepo, nil).Once()
		srv := New(repo)
		response, err := srv.GetProducts(1)
		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].ID, response[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all products", func(t *testing.T) {
		repo.On("GetProducts", 1).Return(nil, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetProducts(1)
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

func TestGetEvents(t *testing.T) {
	repo := new(mocks.UserRepo)
	t.Run("Success get all user events", func(t *testing.T) {
		inputRepo := []event.EventCore{{ID: 1, Name: "main bola", UserID: 1, Address: "disana", City: "jakarta", CategoryID: 1, StartDate: time.Now(), EndDate: time.Now(), TotalParticipant: 10, ImageEvent: "jpg", Status: "avail", MaximumPeople: 20, Description: "just join"}}
		repo.On("GetEvents", 1).Return(inputRepo, nil).Once()
		srv := New(repo)
		response, err := srv.GetEvents(1)
		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].ID, response[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all events", func(t *testing.T) {
		repo.On("GetEvents", 1).Return(nil, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetEvents(1)
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

func TestGetTransactions(t *testing.T) {
	repo := new(mocks.UserRepo)
	t.Run("Success get all user transactions", func(t *testing.T) {
		inputRepo := []transaction.TransactionCore{{ID: 1, UserID: 1, TotalPrice: 12000, TotalQuantity: 1, ProductID: 1, TransactionID: "1223", StatusPayment: "belum dibayar", VirtualAccount: "121212", TransactionTime: "2022-10-12", OrderID: "1"}}
		repo.On("GetTransactions", 1).Return(inputRepo, nil).Once()
		srv := New(repo)
		response, err := srv.GetTransactions(1)
		assert.NoError(t, err)
		assert.Equal(t, inputRepo[0].ID, response[0].ID)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all Transactions", func(t *testing.T) {
		repo.On("GetTransactions", 1).Return(nil, errors.New("failed")).Once()
		srv := New(repo)
		response, err := srv.GetTransactions(1)
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	repo := new(mocks.UserRepo)

	// t.Run("success create users", func(t *testing.T) {
	// 	input := user.Core{Name: "doni", Email: "doni@gmail.com", Password: "doniDoni123!", PhoneNumber: "089213912", Address: "disana", UserImage: "jpg", Gender: "Male"}

	// 	// data := user.Core{Name: "doni", Email: "dona@gmail.com", Password: "doniDoni123!", PhoneNumber: "089213912", Address: "disana", UserImage: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg", Gender: "Male"}
	// 	input.LoginMethod = "Regular"

	// 	repo.On("FindUser", input.Email).Return(user.Core{}, errors.New("failed")).Once()
	// 	repo.On("Create", input).Return(nil).Once()

	// 	srv := New(repo)
	// 	err := srv.Create(input)
	// 	assert.Nil(t, err)
	// 	repo.AssertExpectations(t)
	// })
	t.Run("failed create users, name empty", func(t *testing.T) {
		input := user.Core{Email: "doni", Password: "jojo", PhoneNumber: "089213912", Address: "disana", UserImage: "jpg", Gender: "Male"}

		srv := New(repo)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed create users, password not strong enough", func(t *testing.T) {

		input := user.Core{Name: "doni", Email: "doni@gmail.com", Password: "jojo", PhoneNumber: "089213912", Address: "disana", UserImage: "jpg", Gender: "Male"}

		srv := New(repo)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed create users, email has been used", func(t *testing.T) {

		input := user.Core{Name: "doni", Password: "jojoJojo123!", Email: "dona@gmail.com", PhoneNumber: "089213912", Address: "disana", UserImage: "jpg", Gender: "Male"}

		data := user.Core{Name: "doni", Email: "dona@gmail.com", Password: "doniDoni123!", PhoneNumber: "089213912", Address: "disana", UserImage: "jpg", Gender: "Male"}

		repo.On("FindUser", input.Email).Return(data, nil).Once()

		srv := New(repo)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed create users, email syntax error", func(t *testing.T) {

		input := user.Core{Name: "doni", Password: "jojoJojo123!", Email: "adoni@gmailcom", PhoneNumber: "089213912", Address: "disana", UserImage: "jpg", Gender: "Male"}

		srv := New(repo)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := new(mocks.UserRepo)

	t.Run("success update users", func(t *testing.T) {
		input := user.Core{Name: "doni", Email: "", Password: "", PhoneNumber: "089213912", Address: "disana", UserImage: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg", Gender: "Male"}

		// data := user.Core{ID: 1, Name: "doni", Email: "dona@gmail.com", Password: "doniDoni123!", PhoneNumber: "089213912", Address: "disana", UserImage: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg", Gender: "Male"}

		repo.On("FindUser", input.Email).Return(user.Core{}, errors.New("failed")).Once()
		repo.On("Update", input, 1).Return(nil).Once()
		srv := New(repo)
		err := srv.Update(input, 1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("failed update users, error validate", func(t *testing.T) {

		input := user.Core{Name: "doni", Email: "doni", Password: "jojo", PhoneNumber: "089213912", Address: "disana", UserImage: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg", Gender: "Male"}

		srv := New(repo)
		err := srv.Update(input, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed update users", func(t *testing.T) {

		input := user.Core{Name: "doni", Email: "doni", Password: "jojo", PhoneNumber: "089213912", Address: "disana", UserImage: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg", Gender: "Male"}
		generate, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
		input.Password = string(generate)

		srv := New(repo)
		err := srv.Update(input, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed update users", func(t *testing.T) {

		input := user.Core{Name: "doni", Email: "doni@gmail.com", Password: "jojoJojo123!", PhoneNumber: "089213912", Address: "disana", UserImage: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg", Gender: "Male"}

		data := user.Core{Name: "doni", Email: "doni@gmail.com", Password: "jojoJojo123!", PhoneNumber: "089213912", Address: "disana", UserImage: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg", Gender: "Male"}

		generate, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
		input.Password = string(generate)

		repo.On("FindUser", input.Email).Return(data, nil).Once()

		srv := New(repo)
		err := srv.Update(input, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed update users", func(t *testing.T) {

		input := user.Core{Name: "doni", Email: "", Password: "", PhoneNumber: "089213912", Address: "disana", UserImage: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg", Gender: "Male"}

		// data := user.Core{ID: 1, Name: "doni", Email: "dona@gmail.com", Password: "doniDoni123!", PhoneNumber: "089213912", Address: "disana", UserImage: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg", Gender: "Male"}

		repo.On("FindUser", input.Email).Return(user.Core{}, errors.New("failed")).Once()
		repo.On("Update", input, 1).Return(errors.New("failed")).Once()
		srv := New(repo)
		err := srv.Update(input, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
