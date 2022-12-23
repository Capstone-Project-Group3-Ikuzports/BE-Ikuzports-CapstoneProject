package service

import (
	"errors"
	"ikuzports/features/clubMember"
	"ikuzports/utils/helper"
)

type clubMemberService struct {
	clubMemberRepository clubMember.RepositoryInterface
	// validate       *validator.Validate
}

func New(repo clubMember.RepositoryInterface) clubMember.ServiceInterface {
	return &clubMemberService{
		clubMemberRepository: repo,
		// validate:       validator.New(),
	}
}

// Create implements clubMember.ServiceInterface
func (service *clubMemberService) Create(input clubMember.Core, id int) error {
	errCreate := service.clubMemberRepository.Create(input, id)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// GetAll implements clubMember.ServiceInterface
func (service *clubMemberService) GetAll() (data []clubMember.Core, err error) {
	data, err = service.clubMemberRepository.GetAll()
	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	return data, nil
}

// GetById implements clubMember.ServiceInterface
func (service *clubMemberService) GetById(id int) (data clubMember.Core, err error) {
	data, err = service.clubMemberRepository.GetById(id)
	if err != nil {
		return clubMember.Core{}, helper.ServiceErrorMsg(err)
	}
	return data, err
}
