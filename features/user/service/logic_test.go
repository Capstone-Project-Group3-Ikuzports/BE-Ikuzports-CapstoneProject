package service

// import (
// 	"errors"
// 	"ikuzports/features/user"
// 	"ikuzports/mocks"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestGetAll(t *testing.T) {
// 	repo := new(mocks.UserRepo)
// 	t.Run("Success get all user", func(t *testing.T) {
// 		inputRepo := []user.Core{{ID: 1, Name: "alta", Email: "alta", Password: "alta", PhoneNumber: "089213912", Address: "disana", UserImage: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg", Gender: "Male"}}
// 		repo.On("GetAll").Return(inputRepo, nil).Once()
// 		srv := New(repo)
// 		response, err := srv.GetAll()
// 		assert.NoError(t, err)
// 		assert.Equal(t, inputRepo[0].Name, response[0].Name)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Failed get all", func(t *testing.T) {
// 		repo.On("GetAll").Return(nil, errors.New("failed")).Once()
// 		srv := New(repo)
// 		response, err := srv.GetAll()
// 		assert.NotNil(t, err)
// 		assert.Nil(t, response)
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestGetById(t *testing.T) {
// 	repo := new(mocks.UserRepo)
// 	t.Run("Success get user", func(t *testing.T) {
// 		inputRepo := user.Core{ID: 1, Name: "alta", Email: "alta", Password: "alta", PhoneNumber: "089213912", Address: "disana", UserImage: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg", Gender: "Male"}
// 		repo.On("GetById", 1).Return(inputRepo, nil).Once()
// 		srv := New(repo)
// 		response, err := srv.GetById(1)
// 		assert.NoError(t, err)
// 		assert.Equal(t, inputRepo.Name, response.Name)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Failed get by id", func(t *testing.T) {
// 		inputRepo := user.Core{ID: 1, Name: "alta", Email: "alta", Password: "alta", PhoneNumber: "089213912", Address: "disana", UserImage: "https://preview.keenthemes.com/metronic-v4/theme/assets/pages/media/profile/profile_user.jpg", Gender: "Male"}
// 		repo.On("GetById", 1).Return(user.Core{}, errors.New("failed")).Once()
// 		srv := New(repo)
// 		response, err := srv.GetById(1)
// 		assert.NotNil(t, err)
// 		assert.NotEqual(t, inputRepo.Name, response.Name)
// 		repo.AssertExpectations(t)
// 	})
// }

// func TestDelete(t *testing.T) {
// 	repo := new(mocks.UserRepo)
// 	t.Run("Success delete User", func(t *testing.T) {
// 		repo.On("Delete", 1).Return(nil).Once()
// 		srv := New(repo)
// 		err := srv.Delete(1)
// 		assert.NoError(t, err)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("Failed delete User", func(t *testing.T) {
// 		repo.On("Delete", 1).Return(errors.New("failed")).Once()
// 		srv := New(repo)
// 		err := srv.Delete(1)
// 		assert.NotNil(t, err)
// 		repo.AssertExpectations(t)
// 	})
// }
