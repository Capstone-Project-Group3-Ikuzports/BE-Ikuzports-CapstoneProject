package service

import (
	"errors"
	"ikuzports/features/event"
	"ikuzports/features/participant"
	"ikuzports/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	repo := new(mocks.EventRepo)
	repo2 := new(mocks.ParticipantRepo)

	t.Run("Success get event by id", func(t *testing.T) {
		input := event.EventCore{ID: 1, Name: "Main", UserID: 1, Address: "gor", City: "jakarta", CategoryID: 1, StartDate: time.Now(), EndDate: time.Now(), TotalParticipant: 10, ImageEvent: "jpg", Status: "available", MaximumPeople: 30, Description: "main aja"}

		repo.On("GetByID", 1).Return(input, nil).Once()
		srv := New(repo, repo2)
		response, err := srv.GetByID(1)

		assert.NoError(t, err)
		assert.Equal(t, input.Name, response.Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get event by id", func(t *testing.T) {

		repo.On("GetByID", 1).Return(event.EventCore{}, errors.New("failed")).Once()
		srv := New(repo, repo2)
		_, err := srv.GetByID(1)

		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.EventRepo)
	repo2 := new(mocks.ParticipantRepo)

	t.Run("Success delete event", func(t *testing.T) {
		repo.On("Delete", 1).Return(1, nil).Once()
		srv := New(repo, repo2)
		err := srv.Delete(1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Delete User", func(t *testing.T) {
		repo.On("Delete", 1).Return(0, errors.New("error delete data")).Once()
		srv := New(repo, repo2)
		err := srv.Delete(1)
		assert.NotNil(t, err)
		assert.Equal(t, "error delete data", err.Error())
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	repo := new(mocks.EventRepo)
	repo2 := new(mocks.ParticipantRepo)

	t.Run("Success get all event", func(t *testing.T) {
		input := []event.EventCore{{ID: 1, Name: "Main", UserID: 1, Address: "gor", City: "jakarta", CategoryID: 1, StartDate: time.Now(), EndDate: time.Now(), TotalParticipant: 10, ImageEvent: "jpg", Status: "available", MaximumPeople: 30, Description: "main aja"}}

		repo.On("GetDate").Return(input, nil).Once()
		repo.On("GetAllFilter", 15, 15, 1, "city", "status").Return(input, nil).Once()

		srv := New(repo, repo2)
		res, err := srv.GetAll(1, 2, "city", "status")
		assert.NoError(t, err)
		assert.Equal(t, input[0].Name, res[0].Name)
		repo.AssertExpectations(t)
	})

	t.Run("Success get all event", func(t *testing.T) {
		input := []event.EventCore{{ID: 1, Name: "Main", UserID: 1, Address: "gor", City: "jakarta", CategoryID: 1, StartDate: time.Now(), EndDate: time.Now(), TotalParticipant: 10, ImageEvent: "jpg", Status: "available", MaximumPeople: 30, Description: "main aja"}}

		repo.On("GetDate").Return(input, nil).Once()
		repo.On("GetAll", 15, 15).Return(input, nil).Once()

		srv := New(repo, repo2)
		res, err := srv.GetAll(0, 2, "", "")
		assert.NoError(t, err)
		assert.Equal(t, input[0].Name, res[0].Name)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all event, err get date", func(t *testing.T) {
		repo.On("GetDate").Return(nil, errors.New("failed")).Once()
		srv := New(repo, repo2)
		_, err := srv.GetAll(1, 2, "city", "status")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all event, err get date", func(t *testing.T) {
		layout := "2006-01-02"
		startDate, _ := time.Parse(layout, "2022-08-09")

		input := []event.EventCore{{ID: 1, Name: "Main", UserID: 1, Address: "gor", City: "jakarta", CategoryID: 1, StartDate: startDate, EndDate: time.Now(), TotalParticipant: 10, ImageEvent: "jpg", Status: "available", MaximumPeople: 30, Description: "main aja"}}

		repo.On("GetDate").Return(input, nil).Once()
		repo.On("UpdateStatus", 1, "Not Available").Return(1, nil).Once()
		repo.On("GetAll", 15, 15).Return(nil, errors.New("failed")).Once()

		srv := New(repo, repo2)
		_, err := srv.GetAll(0, 2, "", "")
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	repo := new(mocks.EventRepo)
	repo2 := new(mocks.ParticipantRepo)

	t.Run("Success create event", func(t *testing.T) {
		layout := "2006-01-02"
		startDate, _ := time.Parse(layout, "2022-08-09")

		input := event.EventCore{ID: 1, Name: "Main", UserID: 1, Address: "gor", City: "jakarta", CategoryID: 1, StartDate: startDate, EndDate: time.Now(), TotalParticipant: 10, ImageEvent: "jpg", Status: "Available", MaximumPeople: 30, Description: "main aja"}

		participant := participant.ParticipantCore{UserID: 1, EventID: 1, Status: "Organizer"}

		repo.On("Create", input).Return(1, nil).Once()
		repo.On("GetLastID").Return(1, nil).Once()
		repo2.On("Create", participant).Return(1, nil).Once()
		repo.On("UpdateTotal", 1).Return(1, nil).Once()

		srv := New(repo, repo2)
		err := srv.Create(input)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed create event, some field empty", func(t *testing.T) {
		input := event.EventCore{ID: 1, UserID: 1, Address: "gor", City: "jakarta", CategoryID: 1, StartDate: time.Now(), EndDate: time.Now(), TotalParticipant: 10, ImageEvent: "jpg", Status: "Available", MaximumPeople: 30, Description: "main aja"}

		srv := New(repo, repo2)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed create event, start time after end time", func(t *testing.T) {
		layout := "2006-01-02"
		endDate, _ := time.Parse(layout, "2022-08-09")

		input := event.EventCore{ID: 1, Name: "main", UserID: 1, Address: "gor", City: "jakarta", CategoryID: 1, StartDate: time.Now(), EndDate: endDate, TotalParticipant: 10, ImageEvent: "jpg", Status: "Available", MaximumPeople: 30, Description: "main aja"}

		srv := New(repo, repo2)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed create event", func(t *testing.T) {
		input := event.EventCore{ID: 1, Name: "Main", UserID: 1, Address: "gor", City: "jakarta", CategoryID: 1, StartDate: time.Now(), EndDate: time.Now(), TotalParticipant: 10, ImageEvent: "jpg", Status: "Available", MaximumPeople: 30, Description: "main aja"}

		repo.On("Create", input).Return(0, errors.New("failed")).Once()
		srv := New(repo, repo2)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get last id", func(t *testing.T) {
		input := event.EventCore{ID: 1, Name: "Main", UserID: 1, Address: "gor", City: "jakarta", CategoryID: 1, StartDate: time.Now(), EndDate: time.Now(), TotalParticipant: 10, ImageEvent: "jpg", Status: "Available", MaximumPeople: 30, Description: "main aja"}

		repo.On("Create", input).Return(1, nil).Once()
		repo.On("GetLastID").Return(0, errors.New("failed")).Once()
		srv := New(repo, repo2)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed create participant", func(t *testing.T) {
		input := event.EventCore{ID: 1, Name: "Main", UserID: 1, Address: "gor", City: "jakarta", CategoryID: 1, StartDate: time.Now(), EndDate: time.Now(), TotalParticipant: 10, ImageEvent: "jpg", Status: "Available", MaximumPeople: 30, Description: "main aja"}

		participant := participant.ParticipantCore{UserID: 1, EventID: 1, Status: "Organizer"}

		repo.On("Create", input).Return(1, nil).Once()
		repo.On("GetLastID").Return(1, nil).Once()
		repo2.On("Create", participant).Return(0, errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed create participant", func(t *testing.T) {
		input := event.EventCore{ID: 1, Name: "Main", UserID: 1, Address: "gor", City: "jakarta", CategoryID: 1, StartDate: time.Now(), EndDate: time.Now(), TotalParticipant: 10, ImageEvent: "jpg", Status: "Available", MaximumPeople: 30, Description: "main aja"}

		participant := participant.ParticipantCore{UserID: 1, EventID: 1, Status: "Organizer"}

		repo.On("Create", input).Return(1, nil).Once()
		repo.On("GetLastID").Return(1, nil).Once()
		repo2.On("Create", participant).Return(1, nil).Once()
		repo.On("UpdateTotal", 1).Return(0, errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
