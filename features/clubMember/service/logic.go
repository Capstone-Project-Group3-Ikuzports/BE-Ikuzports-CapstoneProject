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
func (service *clubMemberService) Create(input clubMember.Core) error {
	input.Status = "Member"
	dataMember, errCr := service.clubMemberRepository.FindMember(int(input.ClubID), int(input.UserID))
	if errCr != nil {
		errCreate := service.clubMemberRepository.Create(input)
		if errCreate != nil {
			return errors.New("failed to insert data, error query")
		}
	}
	if dataMember.ClubID == input.ClubID && dataMember.UserID == input.UserID {
		return errors.New(" failed to join, you are already join in this club")
	}
	errPut := service.clubMemberRepository.UpdateMember(int(input.ClubID))
	if errPut != nil {
		return errors.New("failed to update data, error query")
	}
	return nil
}

// GetAll implements clubMember.ServiceInterface
func (service *clubMemberService) GetAll(queryStatus string) (data []clubMember.Core, err error) {
	if queryStatus == "" {
		data, err = service.clubMemberRepository.GetAll()
	} else {
		data, err = service.clubMemberRepository.GetAllByStatus(queryStatus)
	}
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
