package service

import (
	"errors"
	"ikuzports/features/club"
	"ikuzports/features/clubMember"
	"ikuzports/utils/helper"
	"strconv"

	"github.com/labstack/gommon/log"
)

type clubService struct {
	clubRepository       club.RepositoryInterface
	clubMemberRepository clubMember.RepositoryInterface
	// validate       *validator.Validate
}

func New(repo club.RepositoryInterface, repoo clubMember.RepositoryInterface) club.ServiceInterface {
	return &clubService{
		clubRepository:       repo,
		clubMemberRepository: repoo,
	}
}

// Create implements club.ServiceInterface
func (service *clubService) Create(input club.Core, id int) error {
	// if errValidate := service.validate.Struct(input); errValidate != nil {
	// 	return errValidate
	// }
	errCreate := service.clubRepository.Create(input, id)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}
	return nil
}

// GetAll implements club.ServiceInterface
func (service *clubService) GetAll(queryName, queryCity, queryCategory string) (data []club.Core, err error) {
	queryCategoryID, err := strconv.Atoi(queryCategory)
	if err != nil {
		return nil, helper.ServiceErrorMsg(err)
	}
	if queryName == "" && queryCity == "" && queryCategoryID == 0 {
		data, err = service.clubRepository.GetAll()
	} else {
		data, err = service.clubRepository.GetAllWithSearch(queryName, queryCity, queryCategoryID)
	}

	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	return data, nil
}

// GetById implements club.ServiceInterface
func (service *clubService) GetById(id int) (data club.Core, err error) {
	data, err = service.clubRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return club.Core{}, helper.ServiceErrorMsg(err)
	}
	return data, err
}

// Update implements club.ServiceInterface
func (service *clubService) Update(input club.Core, id int) error {
	errUpdate := service.clubRepository.Update(input, id)
	if errUpdate != nil {
		return errors.New("failed update data, error query")
	}
	return nil

}

// Delete implements club.ServiceInterface
func (service *clubService) Delete(id int) error {
	err := service.clubRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}
