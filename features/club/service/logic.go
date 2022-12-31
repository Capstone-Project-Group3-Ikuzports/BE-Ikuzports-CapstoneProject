package service

import (
	"errors"
	"ikuzports/features/chat"
	"ikuzports/features/club"
	"ikuzports/features/clubActivity"
	"ikuzports/features/clubMember"
	"ikuzports/features/galery"
	"ikuzports/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type clubService struct {
	clubRepository       club.RepositoryInterface
	clubMemberRepository clubMember.RepositoryInterface
	validate             *validator.Validate
}

func New(repo club.RepositoryInterface, repoMember clubMember.RepositoryInterface) club.ServiceInterface {
	return &clubService{
		clubRepository:       repo,
		clubMemberRepository: repoMember,
		validate:             validator.New(),
	}
}

// Create implements club.ServiceInterface
func (service *clubService) Create(input club.Core, id int) error {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	errCreate := service.clubRepository.Create(input, id)
	if errCreate != nil {
		return errors.New("failed to insert data, error query")
	}

	clubID, errGet := service.clubRepository.GetLastID()
	if errGet != nil {
		return errors.New("error get last ID. error query")
	}
	dataMember := clubMember.Core{
		UserID: uint(id),
		ClubID: uint(clubID),
		Status: "Owner",
	}
	errCr := service.clubMemberRepository.Create(dataMember)
	if errCr != nil {
		return errors.New("error create club member. error query")
	}

	_, errUpdate := service.clubRepository.UpdateMember(clubID)
	if errUpdate != nil {
		return errors.New("error update joined member. error query")
	}

	return nil
}

// GetAll implements club.ServiceInterface
func (service *clubService) GetAll(queryName string, queryCity string, queryCategoryID int, queryPage int) (data []club.Core, page int, err error) {
	limit := 9
	offset := (queryPage - 1) * limit
	var jumlahData int

	if queryName == "" && queryCity == "" && queryCategoryID == 0 {
		data, jumlahData, err = service.clubRepository.GetAll(offset, limit)
	} else {
		data, jumlahData, err = service.clubRepository.GetAllWithSearch(queryName, queryCity, queryCategoryID, offset, limit)
	}

	if jumlahData%limit == 0 {
		page = jumlahData / limit
	} else {
		page = (jumlahData / limit) + 1
	}

	if err != nil {
		helper.LogDebug(err)
		return nil, 0, helper.ServiceErrorMsg(err)
	}

	return data, page, nil
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
func (service *clubService) Update(input club.Core, id int, userId int) error {

	dataMember, errCr := service.clubRepository.GetStatus(id, userId)
	if errCr != nil {
		return errors.New("error update club. no data")
	}
	if dataMember.Status != "Owner" {
		return errors.New("failed update data, you are not the owner of the club")
	}

	errUpdate := service.clubRepository.Update(input, id)
	if errUpdate != nil {
		return errors.New("failed update data, error query")
	}
	return nil

}

// Delete implements club.ServiceInterface
func (service *clubService) Delete(id int, userId int) error {

	dataMember, errCr := service.clubRepository.GetStatus(id, userId)
	if errCr != nil {
		helper.LogDebug("\n isi datamember = ", dataMember)
		return errors.New("error delete club. error query")
	}
	if dataMember.Status != "Owner" {
		return errors.New("failed delete data, you are not the owner of the club")
	}

	err := service.clubRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}

// GetMembers implements club.ServiceInterface
func (service *clubService) GetMembers(id int) (data []clubMember.Core, err error) {
	data, err = service.clubRepository.GetMembers(id)
	if err != nil {
		log.Error(err.Error())
		return data, helper.ServiceErrorMsg(err)
	}

	if len(data) == 0 {
		return nil, errors.New("Get data success. No data")
	}
	return data, err
}

// GetChats implements club.ServiceInterface
func (service *clubService) GetChats(id int) (data []chat.Core, err error) {
	data, err = service.clubRepository.GetChats(id)
	if err != nil {
		log.Error(err.Error())
		return data, helper.ServiceErrorMsg(err)
	}

	if len(data) == 0 {
		return nil, errors.New("Get data success. No data")
	}
	return data, err
}

// GetGaleries implements club.ServiceInterface
func (service *clubService) GetGaleries(id int) (data []galery.Core, err error) {
	data, err = service.clubRepository.GetGaleries(id)
	if err != nil {
		log.Error(err.Error())
		return data, helper.ServiceErrorMsg(err)
	}

	if len(data) == 0 {
		return nil, errors.New("Get data success. No data")
	}
	return data, err
}

// GetActivities implements club.ServiceInterface
func (service *clubService) GetActivities(id int) (data []clubActivity.Core, err error) {
	data, err = service.clubRepository.GetActivities(id)
	if err != nil {
		log.Error(err.Error())
		return data, helper.ServiceErrorMsg(err)
	}

	if len(data) == 0 {
		return nil, errors.New("Get data success. No data")
	}
	return data, err
}
