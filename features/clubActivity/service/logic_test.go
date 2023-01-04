package service

import (
	"errors"
	"ikuzports/features/club"
	"ikuzports/features/clubActivity"
	"ikuzports/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.ClubActivityRepo)
	repo2 := new(mocks.ClubRepo)

	t.Run("Succes get all activity", func(t *testing.T) {
		input := []clubActivity.Core{{ClubID: 1, Name: "main", StartTime: "10:30", EndTime: "11:00", Day: "Kamis", Location: "gor", ActivityDetail: "on time"}}

		repo.On("GetAll").Return(input, nil).Once()

		srv := New(repo, repo2)
		res, err := srv.GetAll(0)
		assert.NoError(t, err)
		assert.Equal(t, input[0].ClubID, res[0].ClubID)

		repo.AssertExpectations(t)
	})
	t.Run("Succes get all activity", func(t *testing.T) {
		input := []clubActivity.Core{{ClubID: 1, Name: "main", StartTime: "10:30", EndTime: "11:00", Day: "Kamis", Location: "gor", ActivityDetail: "on time"}}

		repo.On("GetAllByClubId", uint(1)).Return(input, nil).Once()

		srv := New(repo, repo2)
		res, err := srv.GetAll(1)
		assert.NoError(t, err)
		assert.Equal(t, input[0].ClubID, res[0].ClubID)

		repo.AssertExpectations(t)
	})

	t.Run("Failed get all activity", func(t *testing.T) {

		repo.On("GetAll").Return(nil, errors.New("failed")).Once()

		srv := New(repo, repo2)
		_, err := srv.GetAll(0)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	repo := new(mocks.ClubActivityRepo)
	repo2 := new(mocks.ClubRepo)

	t.Run("Success Create activity", func(t *testing.T) {
		input := clubActivity.Core{ClubID: 1, Name: "main", StartTime: "10:30", EndTime: "11:00", Day: "Kamis", Location: "gor", ActivityDetail: "on time"}

		data2 := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}

		repo2.On("GetStatus", 1, 1).Return(data2, nil).Once()
		repo.On("Create", input, 1).Return(nil).Once()

		srv := New(repo, repo2)
		err := srv.Create(input, 1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed create activity, validation error", func(t *testing.T) {
		input := clubActivity.Core{ClubID: 1, StartTime: "10:30", EndTime: "11:00", Day: "Kamis", Location: "gor", ActivityDetail: "on time"}

		srv := New(repo, repo2)
		err := srv.Create(input, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed create activity, get status failed", func(t *testing.T) {
		input := clubActivity.Core{ClubID: 1, Name: "aqiz", StartTime: "10:30", EndTime: "11:00", Day: "Kamis", Location: "gor", ActivityDetail: "on time"}

		repo2.On("GetStatus", 1, 1).Return(club.Status{}, errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Create(input, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed create activity, not an owner", func(t *testing.T) {
		input := clubActivity.Core{ClubID: 1, Name: "aqiz", StartTime: "10:30", EndTime: "11:00", Day: "Kamis", Location: "gor", ActivityDetail: "on time"}

		data2 := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Member"}

		repo2.On("GetStatus", 1, 1).Return(data2, nil).Once()

		srv := New(repo, repo2)
		err := srv.Create(input, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed create activity", func(t *testing.T) {
		input := clubActivity.Core{ClubID: 1, Name: "aqiz", StartTime: "10:30", EndTime: "11:00", Day: "Kamis", Location: "gor", ActivityDetail: "on time"}

		data2 := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}

		repo2.On("GetStatus", 1, 1).Return(data2, nil).Once()
		repo.On("Create", input, 1).Return(errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Create(input, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := new(mocks.ClubActivityRepo)
	repo2 := new(mocks.ClubRepo)

	t.Run("Success update activity", func(t *testing.T) {
		input := clubActivity.Core{ClubID: 1, Name: "main", StartTime: "10:30", EndTime: "11:00", Day: "Kamis", Location: "gor", ActivityDetail: "on time"}

		data2 := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}

		repo2.On("GetStatus", 1, 1).Return(data2, nil).Once()
		repo.On("Update", input, 1).Return(nil).Once()

		srv := New(repo, repo2)
		err := srv.Update(input, 1, 1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed update activity, get status error", func(t *testing.T) {
		input := clubActivity.Core{ClubID: 1, Name: "aqiz", StartTime: "10:30", EndTime: "11:00", Day: "Kamis", Location: "gor", ActivityDetail: "on time"}

		repo2.On("GetStatus", 1, 1).Return(club.Status{}, errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Update(input, 1, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed update activity, not an owner", func(t *testing.T) {
		input := clubActivity.Core{ClubID: 1, Name: "aqiz", StartTime: "10:30", EndTime: "11:00", Day: "Kamis", Location: "gor", ActivityDetail: "on time"}

		data2 := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Member"}

		repo2.On("GetStatus", 1, 1).Return(data2, nil).Once()

		srv := New(repo, repo2)
		err := srv.Update(input, 1, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed update activity", func(t *testing.T) {
		input := clubActivity.Core{ClubID: 1, Name: "aqiz", StartTime: "10:30", EndTime: "11:00", Day: "Kamis", Location: "gor", ActivityDetail: "on time"}

		data2 := club.Status{ID: 1, UserID: 2, ClubID: 1, Status: "Owner"}

		repo2.On("GetStatus", 1, 1).Return(data2, nil).Once()
		repo.On("Update", input, 1).Return(errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Update(input, 1, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
