package service

import (
	"errors"
	"ikuzports/features/club"
	"ikuzports/features/galery"
	"ikuzports/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.GaleryRepo)
	repo2 := new(mocks.ClubRepo)

	t.Run("Success Get All Galeries", func(t *testing.T) {
		data := []galery.Core{{ID: 1, Url: "jpg", ClubID: 2, Caption: "foto"}}

		repo.On("GetAll").Return(data, nil).Once()
		srv := New(repo, repo2)
		res, err := srv.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, data[0].Url, res[0].Url)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get All Galeries", func(t *testing.T) {

		repo.On("GetAll").Return(nil, errors.New("failed")).Once()
		srv := New(repo, repo2)
		_, err := srv.GetAll()
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	repo := new(mocks.GaleryRepo)
	repo2 := new(mocks.ClubRepo)

	t.Run("Success Create image", func(t *testing.T) {
		data := galery.Core{ID: 1, Url: "jpg", ClubID: 2, Caption: "foto"}

		data2 := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}

		repo2.On("GetStatus", 2, 1).Return(data2, nil).Once()
		repo.On("Create", data, 1).Return(nil).Once()

		srv := New(repo, repo2)
		err := srv.Create(data, 1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed create galery, validation error", func(t *testing.T) {
		data := galery.Core{ID: 1, ClubID: 2, Caption: "foto"}

		srv := New(repo, repo2)
		err := srv.Create(data, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed create galery, get status failed", func(t *testing.T) {
		data := galery.Core{ID: 1, Url: "jpg", ClubID: 2, Caption: "foto"}

		repo2.On("GetStatus", 2, 1).Return(club.Status{}, errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Create(data, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed create galery, not an owner", func(t *testing.T) {
		data := galery.Core{ID: 1, Url: "jpg", ClubID: 2, Caption: "foto"}

		data2 := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Member"}

		repo2.On("GetStatus", 2, 1).Return(data2, nil).Once()

		srv := New(repo, repo2)
		err := srv.Create(data, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed create galery", func(t *testing.T) {
		data := galery.Core{ID: 1, Url: "jpg", ClubID: 2, Caption: "foto"}

		data2 := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}

		repo2.On("GetStatus", 2, 1).Return(data2, nil).Once()
		repo.On("Create", data, 1).Return(errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Create(data, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.GaleryRepo)
	repo2 := new(mocks.ClubRepo)

	t.Run("Success Get Image by id", func(t *testing.T) {
		data := galery.Core{ID: 1, Url: "jpg", ClubID: 2, Caption: "foto"}

		repo.On("GetById", 1).Return(data, nil).Once()
		srv := New(repo, repo2)
		res, err := srv.GetById(1)
		assert.NoError(t, err)
		assert.Equal(t, data.Url, res.Url)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Get Image by id", func(t *testing.T) {

		repo.On("GetById", 1).Return(galery.Core{}, errors.New("failed")).Once()
		srv := New(repo, repo2)
		_, err := srv.GetById(1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.GaleryRepo)
	repo2 := new(mocks.ClubRepo)

	id := 1
	userId := 2
	t.Run("Success delete galeries", func(t *testing.T) {
		data := galery.Core{ID: 1, Url: "jpg", ClubID: 2, Caption: "foto"}
		data2 := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}

		repo.On("GetById", 1).Return(data, nil).Once()
		repo2.On("GetStatus", 2, 2).Return(data2, nil).Once()
		repo.On("Delete", 1).Return(nil).Once()

		srv := New(repo, repo2)
		err := srv.Delete(id, userId)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("failed delete galeries", func(t *testing.T) {

		repo.On("GetById", 1).Return(galery.Core{}, errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Delete(id, userId)
		assert.NotNil(t, err)
		// assert.NotEqual(t, inputRepo.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("failed delete galeries", func(t *testing.T) {
		data := galery.Core{ID: 1, Url: "jpg", ClubID: 2, Caption: "foto"}
		repo.On("GetById", 1).Return(data, nil).Once()
		repo2.On("GetStatus", 2, 2).Return(club.Status{}, errors.New("error delete image. no data")).Once()

		srv := New(repo, repo2)
		err := srv.Delete(id, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed delete galeries", func(t *testing.T) {
		data2 := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Member"}
		data := galery.Core{ID: 1, Url: "jpg", ClubID: 2, Caption: "foto"}

		repo.On("GetById", 1).Return(data, nil).Once()
		repo2.On("GetStatus", 2, 2).Return(data2, nil).Once()

		srv := New(repo, repo2)
		err := srv.Delete(id, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed delete galeries", func(t *testing.T) {
		data2 := club.Status{ID: 1, UserID: 2, ClubID: 3, Status: "Owner"}
		data := galery.Core{ID: 1, Url: "jpg", ClubID: 2, Caption: "foto"}

		repo.On("GetById", 1).Return(data, nil).Once()
		repo2.On("GetStatus", 2, 2).Return(data2, nil).Once()
		repo.On("Delete", 1).Return(errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Delete(id, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := new(mocks.GaleryRepo)
	repo2 := new(mocks.ClubRepo)

	id := 1
	userId := 2
	t.Run("Success update foto", func(t *testing.T) {
		data := galery.Core{ID: 1, Url: "jpg", ClubID: 2, Caption: "foto"}
		data2 := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}

		repo.On("GetById", 1).Return(data, nil).Once()
		repo2.On("GetStatus", 2, 2).Return(data2, nil).Once()
		repo.On("Update", data, 1).Return(nil).Once()

		srv := New(repo, repo2)
		err := srv.Update(data, id, userId)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("failed update galeries", func(t *testing.T) {
		data := galery.Core{ID: 1, Url: "jpg", ClubID: 2, Caption: "foto"}
		repo.On("GetById", 1).Return(galery.Core{}, errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Update(data, id, userId)
		assert.NotNil(t, err)
		// assert.NotEqual(t, inputRepo.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("failed update galeries", func(t *testing.T) {
		data := galery.Core{ID: 1, Url: "jpg", ClubID: 2, Caption: "foto"}
		repo.On("GetById", 1).Return(data, nil).Once()
		repo2.On("GetStatus", 2, 2).Return(club.Status{}, errors.New("error update image. no data")).Once()

		srv := New(repo, repo2)
		err := srv.Update(data, id, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed update galeries", func(t *testing.T) {
		data2 := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Member"}
		data := galery.Core{ID: 1, Url: "jpg", ClubID: 2, Caption: "foto"}

		repo.On("GetById", 1).Return(data, nil).Once()
		repo2.On("GetStatus", 2, 2).Return(data2, nil).Once()

		srv := New(repo, repo2)
		err := srv.Update(data, id, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("failed update galeries", func(t *testing.T) {
		data2 := club.Status{ID: 1, UserID: 2, ClubID: 3, Status: "Owner"}
		data := galery.Core{ID: 1, Url: "jpg", ClubID: 2, Caption: "foto"}

		repo.On("GetById", 1).Return(data, nil).Once()
		repo2.On("GetStatus", 2, 2).Return(data2, nil).Once()
		repo.On("Update", data, 1).Return(errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Update(data, id, userId)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
