package service

import (
	"errors"
	"ikuzports/features/chat"
	"ikuzports/utils/helper"
)

type chatService struct {
	chatRepository chat.RepositoryInterface
	// validate       *validator.Validate
}

func New(repo chat.RepositoryInterface) chat.ServiceInterface {
	return &chatService{
		chatRepository: repo,
		// validate:       validator.New(),
	}
}

// Create implements chat.ServiceInterface
func (service *chatService) Create(input chat.Core) error {

	errUpdate := service.chatRepository.Create(input)
	if errUpdate != nil {
		return errors.New("failed update data, error query")
	}
	return nil
}

// Delete implements chat.ServiceInterface
func (service *chatService) Delete(id int, userId int) error {
	data, err := service.chatRepository.GetById(id)
	if err != nil {
		return helper.ServiceErrorMsg(err)
	}
	if data.UserID != uint(userId) {
		return errors.New("failed delete chat, you are not the owner of the chat")
	}
	errCreate := service.chatRepository.Delete(id)
	if errCreate != nil {
		return errors.New("failed to delete chat, error query")
	}
	return nil
}

// GetAll implements chat.ServiceInterface
func (service *chatService) GetAll() (data []chat.Core, err error) {
	data, err = service.chatRepository.GetAll()
	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	return data, nil
}
