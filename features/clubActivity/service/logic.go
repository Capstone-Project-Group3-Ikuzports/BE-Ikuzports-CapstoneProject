package service

import (
	"errors"
	"ikuzports/features/club"
	"ikuzports/features/clubActivity"
	"ikuzports/utils/helper"

	"github.com/go-playground/validator/v10"
)

type clubActivityService struct {
	clubActivityRepository clubActivity.RepositoryInterface
	clubRepository         club.RepositoryInterface
	validate               *validator.Validate
}

func New(repo clubActivity.RepositoryInterface, clubRepo club.RepositoryInterface) clubActivity.ServiceInterface {
	return &clubActivityService{
		clubActivityRepository: repo,
		clubRepository:         clubRepo,
		validate:               validator.New(),
	}
}

// Create implements clubActivity.ServiceInterface
func (service *clubActivityService) Create(input clubActivity.Core, idUser int) error {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	dataMember, errCr := service.clubRepository.GetStatus(int(input.ClubID), idUser)
	if errCr != nil {
		return errors.New(" error update club. no data or you have not joined this club")
	}
	if dataMember.Status != "Owner" {
		return errors.New(" failed create activity, you are not the owner of the club")
	}

	errUpdate := service.clubActivityRepository.Create(input, idUser)
	if errUpdate != nil {
		return errors.New("failed update data, error query")
	}
	return nil

}

// GetAll implements clubActivity.ServiceInterface
func (service *clubActivityService) GetAll(queryClubId int) (data []clubActivity.Core, err error) {
	if queryClubId == 0 {
		data, err = service.clubActivityRepository.GetAll()
	} else {
		data, err = service.clubActivityRepository.GetAllByClubId(uint(queryClubId))
	}

	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	return data, nil
}

// Update implements clubActivity.ServiceInterface
func (service *clubActivityService) Update(input clubActivity.Core, id int, userId int) error {

	dataMember, errCr := service.clubRepository.GetStatus(id, userId)
	if errCr != nil {
		return errors.New("error update club. no data")
	}
	if dataMember.Status != "Owner" {
		return errors.New("failed update data, you are not the owner of the club")
	}

	errUpdate := service.clubActivityRepository.Update(input, id)
	if errUpdate != nil {
		return errors.New("failed update data, error query")
	}
	return nil

}
