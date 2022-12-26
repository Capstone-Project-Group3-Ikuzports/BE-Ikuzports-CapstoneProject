package service

import (
	"errors"
	"ikuzports/features/club"
	"ikuzports/features/clubMember"
	"ikuzports/utils/helper"
)

type clubMemberService struct {
	clubMemberRepository clubMember.RepositoryInterface
	clubRepository       club.RepositoryInterface
	// validate       *validator.Validate
}

func New(repo clubMember.RepositoryInterface, repoClub club.RepositoryInterface) clubMember.ServiceInterface {
	return &clubMemberService{
		clubMemberRepository: repo,
		clubRepository:       repoClub,
		// validate:       validator.New(),
	}
}

// Create implements clubMember.ServiceInterface
func (service *clubMemberService) Create(input clubMember.Core) error {
	input.Status = "Requested"
	dataMember, errCr := service.clubMemberRepository.FindMember(int(input.ClubID), int(input.UserID))
	if errCr != nil {
		errCreate := service.clubMemberRepository.Create(input)
		if errCreate != nil {
			return errors.New("failed to insert data, error query")
		}
	}
	if dataMember.ClubID == input.ClubID && dataMember.UserID == input.UserID && dataMember.DeletedAt == input.DeletedAt {
		return errors.New(" failed to join, you are already join in this club")
	}
	errPut := service.clubMemberRepository.UpdateMember(int(input.ClubID))
	if errPut != nil {
		return errors.New("failed to update joined member, error query")
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

// Delete implements clubMember.ServiceInterface
func (service *clubMemberService) Delete(id int, userId int) error {
	data, err := service.clubMemberRepository.GetById(id)
	if err != nil {
		return helper.ServiceErrorMsg(err)
	}
	dataMember, errCr := service.clubRepository.GetStatus(int(data.ClubID), userId)
	if errCr != nil {
		return errors.New("error delete member. no data or you are not joined in this club")
	}
	if data.UserID != uint(userId) {
		if dataMember.Status != "Owner" {
			return errors.New("failed delete member, you are not the owner of the club")
		}
	}

	errCreate := service.clubMemberRepository.Delete(id)
	if errCreate != nil {
		return errors.New("failed to delete data, error query")
	}
	errPut := service.clubMemberRepository.UpdateMember(int(data.ClubID))
	if errPut != nil {
		return errors.New("failed to update joined member, error query")
	}
	return nil
}

// Update implements clubMember.ServiceInterface
func (service *clubMemberService) Update(input clubMember.Core, id int, userId int) error {
	input.Status = "Member"
	data, err := service.clubMemberRepository.GetById(id)
	if err != nil {
		return helper.ServiceErrorMsg(err)
	}
	dataMember, errCr := service.clubRepository.GetStatus(int(input.ClubID), userId)
	if errCr != nil {
		return errors.New("error delete member. no data or you are not joined in this club")
	}
	if data.UserID != uint(userId) {
		if dataMember.Status != "Owner" {
			return errors.New("failed update member, you are not the owner of the club")
		}
	}
	errCreate := service.clubMemberRepository.Update(input, id)
	if errCreate != nil {
		return errors.New("failed to delete data, error query")
	}
	return nil
}
