package service

import (
	"errors"
	"ikuzports/features/auth"
	"ikuzports/features/user"
	"ikuzports/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	repo := new(mocks.AuthRepo)
	repo2 := new(mocks.UserRepo)

	// t.Run("Success login", func(t *testing.T) {
	// 	data := auth.Core{Name: "jojo", Email: "jojo@gmail.com", Password: "qwerty", PhoneNumber: "08999", Address: "disana", City: "jakarta", UserImage: "jpg", Gender: "male", Biodata: "a"}

	// 	repo.On("FindUser", data.Email).Return(data, nil).Once()
	// 	srv := New(repo, repo2)
	// 	res, _, err := srv.Login(data)
	// 	assert.NoError(t, err)
	// 	assert.Equal(t, data.Password, res.Password)
	// 	repo.AssertExpectations(t)
	// })

	t.Run("Failed login, field empty", func(t *testing.T) {
		data := auth.Core{Password: "Jojo123!", PhoneNumber: "08999", Address: "disana", City: "jakarta", UserImage: "jpg", Gender: "male", Biodata: "a"}

		srv := New(repo, repo2)
		_, _, err := srv.Login(data)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed login, cant find user", func(t *testing.T) {
		data := auth.Core{Name: "jojo", Email: "jojo@gmail.com", Password: "Jojo123!", PhoneNumber: "08999", Address: "disana", City: "jakarta", UserImage: "jpg", Gender: "male", Biodata: "a"}
		repo.On("FindUser", data.Email).Return(auth.Core{}, errors.New("table")).Once()

		srv := New(repo, repo2)
		_, _, err := srv.Login(data)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed login, not found", func(t *testing.T) {
		data := auth.Core{Name: "jojo", Email: "jojo@gmail.com", Password: "Jojo123!", PhoneNumber: "08999", Address: "disana", City: "jakarta", UserImage: "jpg", Gender: "male", Biodata: "a"}
		repo.On("FindUser", data.Email).Return(auth.Core{}, errors.New("found")).Once()

		srv := New(repo, repo2)
		_, _, err := srv.Login(data)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed login", func(t *testing.T) {
		data := auth.Core{Name: "jojo", Email: "jojo@gmail.com", Password: "Jojo123!", PhoneNumber: "08999", Address: "disana", City: "jakarta", UserImage: "jpg", Gender: "male", Biodata: "a"}
		repo.On("FindUser", data.Email).Return(auth.Core{}, errors.New("failed")).Once()

		srv := New(repo, repo2)
		_, _, err := srv.Login(data)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed login", func(t *testing.T) {
		data := auth.Core{Name: "jojo", Email: "jojo@gmail.com", Password: "qwerty", PhoneNumber: "08999", Address: "disana", City: "jakarta", UserImage: "jpg", Gender: "male", Biodata: "a"}

		repo.On("FindUser", data.Email).Return(data, nil).Once()
		srv := New(repo, repo2)
		res, _, err := srv.Login(data)
		assert.NotNil(t, err)
		assert.NotEqual(t, data.Password, res.Password)
		repo.AssertExpectations(t)
	})

}

func TestLoginGoogle(t *testing.T) {
	repo := new(mocks.AuthRepo)
	repo2 := new(mocks.UserRepo)

	t.Run("Failed login, cant find user", func(t *testing.T) {

		input := user.GoogleCore{Email: "jojo@gmail.com", Name: "jojo", Picture: "jpg"}
		repo.On("FindUser", input.Email).Return(auth.Core{}, errors.New("table")).Once()

		srv := New(repo, repo2)
		_, _, err := srv.LoginGoogle(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed login, cant find user", func(t *testing.T) {
		// data := user.Core{Name: "jojo", Email: "jojo@gmail.com", UserImage: "jpg", LoginMethod: "Google"}

		input := user.GoogleCore{Email: "jojo@gmail.com", Name: "jojo", Picture: "jpg"}

		repo.On("FindUser", input.Email).Return(auth.Core{}, errors.New("failed")).Once()

		srv := New(repo, repo2)
		_, _, err := srv.LoginGoogle(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Success login", func(t *testing.T) {
		data := user.Core{Name: "jojo", Email: "jojo@gmail.com", UserImage: "jpg", LoginMethod: "Google"}

		data1 := auth.Core{Name: "jojo", Email: "jojo@gmail.com", Password: "qwerty", PhoneNumber: "08999", Address: "disana", City: "jakarta", UserImage: "jpg", Gender: "male", Biodata: "a"}

		input := user.GoogleCore{Email: "jojo@gmail.com", Name: "jojo", Picture: "jpg"}

		repo.On("FindUser", input.Email).Return(auth.Core{}, errors.New("found")).Once()
		repo2.On("Create", data).Return(nil).Once()
		repo.On("FindUser", data.Email).Return(data1, nil)

		srv := New(repo, repo2)
		_, _, err := srv.LoginGoogle(input)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	// t.Run("Success Login google", func(t *testing.T) {
	// 	data := auth.Core{Name: "jojo", Email: "jojo@gmail.com", Password: "Qwerty123!", PhoneNumber: "08999", Address: "disana", City: "jakarta", UserImage: "jpg", Gender: "male", Biodata: "a"}

	// 	repo.On("FindUser", data.Email).Return(data, nil).Once()

	// 	srv := New(repo, repo2)
	// 	_, _, err := srv.Login(data)
	// 	assert.NoError(t, err)
	// 	assert.NotEqual(t, "", data.Name)
	// 	repo.AssertExpectations(t)
	// })
}
