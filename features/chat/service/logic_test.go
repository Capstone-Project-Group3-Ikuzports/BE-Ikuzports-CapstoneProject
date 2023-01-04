package service

import (
	"errors"
	"ikuzports/features/chat"
	"ikuzports/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.ChatRepo)

	t.Run("Succes get all chat", func(t *testing.T) {
		chat := []chat.Core{{ID: 1, UserID: 1, ClubID: 1, Message: "Halo"}}

		repo.On("GetAll").Return(chat, nil).Once()
		srv := New(repo)

		response, err := srv.GetAll()
		assert.NoError(t, err)
		assert.Equal(t, chat[0].Message, response[0].Message)
		repo.AssertExpectations(t)
	})

	t.Run("Failed get all chat", func(t *testing.T) {
		repo.On("GetAll").Return(nil, errors.New("failed")).Once()
		srv := New(repo)

		response, err := srv.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, response)
		repo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	repo := new(mocks.ChatRepo)
	t.Run("Success Create chat", func(t *testing.T) {
		input := chat.Core{ID: 1, UserID: 1, ClubID: 1, Message: "Halo"}
		repo.On("Create", input).Return(nil).Once()
		srv := New(repo)
		err := srv.Create(input)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Create chat", func(t *testing.T) {
		input := chat.Core{ID: 1, UserID: 1, ClubID: 1, Message: "Halo"}
		repo.On("Create", input).Return(errors.New("failed to insert data, error query")).Once()
		srv := New(repo)
		err := srv.Create(input)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.ChatRepo)
	t.Run("Succes Delete chat", func(t *testing.T) {
		input := chat.Core{ID: 1, UserID: 1, ClubID: 1, Message: "Halo"}
		repo.On("GetById", 1).Return(input, nil).Once()
		repo.On("Delete", 1).Return(nil).Once()

		srv := New(repo)
		err := srv.Delete(1, 1)
		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete chat", func(t *testing.T) {
		repo.On("GetById", 1).Return(chat.Core{}, errors.New("failed")).Once()
		srv := New(repo)
		err := srv.Delete(1, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Failed Delete chat", func(t *testing.T) {
		input := chat.Core{ID: 1, UserID: 1, ClubID: 1, Message: "Halo"}
		repo.On("GetById", 1).Return(input, nil).Once()

		srv := New(repo)
		err := srv.Delete(1, 4)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed Delete chat", func(t *testing.T) {
		input := chat.Core{ID: 1, UserID: 1, ClubID: 1, Message: "Halo"}
		repo.On("GetById", 1).Return(input, nil).Once()
		repo.On("Delete", 1).Return(errors.New("failed")).Once()

		srv := New(repo)
		err := srv.Delete(1, 1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})

}
