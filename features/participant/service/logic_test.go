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

func TestCreate(t *testing.T) {
	repo := new(mocks.ParticipantRepo)
	repo2 := new(mocks.EventRepo)

	t.Run("Success Create partisipant", func(t *testing.T) {
		data := participant.ParticipantCore{UserID: 3, EventID: 4, Status: "Participant"}
		input := event.EventCore{ID: 1, Name: "Main", UserID: 1, Address: "gor", City: "jakarta", CategoryID: 1, StartDate: time.Now(), EndDate: time.Now(), TotalParticipant: 20, ImageEvent: "jpg", Status: "Available", MaximumPeople: 30, Description: "main aja"}

		participant := participant.ParticipantCore{UserID: 1, EventID: 2, Status: "Organizer"}

		repo.On("FindMember", 4, 3).Return(participant, nil).Once()
		repo2.On("GetByID", 4).Return(input, nil).Once()
		repo.On("Create", data).Return(1, nil).Once()
		repo.On("UpdateParticipant", data).Return(1, nil).Once()

		srv := New(repo, repo2)
		err := srv.Create(data)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create participant, already join", func(t *testing.T) {
		participant := participant.ParticipantCore{UserID: 1, EventID: 2, Status: "Organizer"}

		repo.On("FindMember", 2, 1).Return(participant, nil).Once()

		srv := New(repo, repo2)
		err := srv.Create(participant)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	// t.Run("Failed Create participant", func(t *testing.T) {
	// 	data := participant.ParticipantCore{UserID: 5, EventID: 6, Status: "Available"}

	// 	participant := participant.ParticipantCore{UserID: 1, EventID: 2, Status: "Organizer"}

	// 	repo.On("FindMember", 6, 5).Return(participant, nil).Once()
	// 	repo2.On("GetByID", 6).Return(event.EventCore{}, errors.New("Failed")).Once()
	// 	srv := New(repo, repo2)
	// 	err := srv.Create(data)
	// 	assert.NotNil(t, err)
	// 	repo.AssertExpectations(t)
	// })

	t.Run("Failed post data", func(t *testing.T) {
		data := participant.ParticipantCore{UserID: 3, EventID: 4, Status: "Participant"}
		input := event.EventCore{ID: 1, Name: "Main", UserID: 1, Address: "gor", City: "jakarta", CategoryID: 1, StartDate: time.Now(), EndDate: time.Now(), TotalParticipant: 30, ImageEvent: "jpg", Status: "Available", MaximumPeople: 30, Description: "main aja"}

		participant := participant.ParticipantCore{UserID: 1, EventID: 2, Status: "Organizer"}

		repo.On("FindMember", 4, 3).Return(participant, nil).Once()
		repo2.On("GetByID", 4).Return(input, nil).Once()
		repo.On("UpdateStatus", data).Return(1, nil).Once()

		srv := New(repo, repo2)
		err := srv.Create(data)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed post data", func(t *testing.T) {
		data := participant.ParticipantCore{UserID: 3, EventID: 4, Status: "Participant"}
		input := event.EventCore{ID: 1, Name: "Main", UserID: 1, Address: "gor", City: "jakarta", CategoryID: 1, StartDate: time.Now(), EndDate: time.Now(), TotalParticipant: 30, ImageEvent: "jpg", Status: "Available", MaximumPeople: 30, Description: "main aja"}

		participant := participant.ParticipantCore{UserID: 1, EventID: 2, Status: "Organizer"}

		repo.On("FindMember", 4, 3).Return(participant, nil).Once()
		repo2.On("GetByID", 4).Return(input, nil).Once()
		repo.On("UpdateStatus", data).Return(0, errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Create(data)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed post data", func(t *testing.T) {
		data := participant.ParticipantCore{UserID: 3, EventID: 4, Status: "Participant"}
		input := event.EventCore{ID: 1, Name: "Main", UserID: 1, Address: "gor", City: "jakarta", CategoryID: 1, StartDate: time.Now(), EndDate: time.Now(), TotalParticipant: 9, ImageEvent: "jpg", Status: "Available", MaximumPeople: 30, Description: "main aja"}

		participant := participant.ParticipantCore{UserID: 1, EventID: 2, Status: "Organizer"}

		repo.On("FindMember", 4, 3).Return(participant, nil).Once()
		repo2.On("GetByID", 4).Return(input, nil).Once()
		// repo.On("UpdateStatus", data).Return(1, nil).Once()
		repo.On("Create", data).Return(0, errors.New("failed")).Once()

		srv := New(repo, repo2)
		err := srv.Create(data)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed post data", func(t *testing.T) {
		data := participant.ParticipantCore{UserID: 3, EventID: 4, Status: "Participant"}
		input := event.EventCore{ID: 1, Name: "Main", UserID: 1, Address: "gor", City: "jakarta", CategoryID: 1, StartDate: time.Now(), EndDate: time.Now(), TotalParticipant: 9, ImageEvent: "jpg", Status: "Available", MaximumPeople: 30, Description: "main aja"}

		participant := participant.ParticipantCore{UserID: 1, EventID: 2, Status: "Organizer"}

		repo.On("FindMember", 4, 3).Return(participant, nil).Once()
		repo2.On("GetByID", 4).Return(input, nil).Once()
		// repo.On("UpdateStatus", data).Return(1, nil).Once()
		repo.On("Create", data).Return(1, nil).Once()
		repo.On("UpdateParticipant", data).Return(0, errors.New("failed"))

		srv := New(repo, repo2)
		err := srv.Create(data)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
