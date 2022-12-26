package service

import (
	"errors"
	"ikuzports/features/club"
	"ikuzports/features/galery"
	"ikuzports/utils/helper"

	"github.com/go-playground/validator/v10"
)

type galeryService struct {
	galeryRepository galery.RepositoryInterface
	clubRepository   club.RepositoryInterface
	validate         *validator.Validate
}

func New(repo galery.RepositoryInterface, repoClub club.RepositoryInterface) galery.ServiceInterface {
	return &galeryService{
		galeryRepository: repo,
		clubRepository:   repoClub,
		validate:         validator.New(),
	}
}

// GetAll implements galery.ServiceInterface
func (service *galeryService) GetAll() (data []galery.Core, err error) {
	data, err = service.galeryRepository.GetAll()
	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	return data, nil
}

// Create implements galery.ServiceInterface
func (service *galeryService) Create(input galery.Core, id int) error {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	dataMember, errCr := service.clubRepository.GetStatus(int(input.ClubID), id)
	if errCr != nil {
		return errors.New("error create galery. no data or you are not joined in this club")
	}
	helper.LogDebug(dataMember)
	if dataMember.Status != "Owner" {
		return errors.New("failed create galery, you are not the owner of the club")
	}
	errCreate := service.galeryRepository.Create(input, id)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// GetById implements galery.ServiceInterface
func (service *galeryService) GetById(id int) (data galery.Core, err error) {
	data, err = service.galeryRepository.GetById(id)
	if err != nil {
		return data, helper.ServiceErrorMsg(err)
	}
	return data, err
}

// Delete implements galery.ServiceInterface
func (service *galeryService) Delete(id int, userId int) error {
	data, err := service.galeryRepository.GetById(id)
	if err != nil {
		return helper.ServiceErrorMsg(err)
	}
	dataMember, errCr := service.clubRepository.GetStatus(int(data.ClubID), userId)
	if errCr != nil {
		return errors.New("error delete image. no data or you are not joined in this club")
	}
	if dataMember.Status != "Owner" {
		return errors.New("failed delete image, you are not the owner of the club")
	}
	errCreate := service.galeryRepository.Delete(id)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// Update implements galery.ServiceInterface
func (service *galeryService) Update(input galery.Core, id int, userId int) error {
	data, err := service.galeryRepository.GetById(id)
	if err != nil {
		return helper.ServiceErrorMsg(err)
	}
	dataMember, errCr := service.clubRepository.GetStatus(int(data.ClubID), userId)
	if errCr != nil {
		return errors.New("error update club. no data or you are not joined in this club")
	}
	if dataMember.Status != "Owner" {
		return errors.New("failed upload image, you are not the owner of the club")
	}
	input.ClubID = data.ClubID
	errCreate := service.galeryRepository.Update(input, id)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}
