package service

import (
	"errors"
	"ikuzports/features/clubMember"
	"ikuzports/features/event"
	"ikuzports/features/product"
	"ikuzports/features/user"
	"ikuzports/utils/helper"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository user.RepositoryInterface
	validate       *validator.Validate
}

func New(repo user.RepositoryInterface) user.ServiceInterface {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

// Create implements user.ServiceInterface
func (service *userService) Create(input user.Core) (err error) {
	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	// validasi email harus unik
	data, errFindEmail := service.userRepository.FindUser(input.Email)

	if data.Email == input.Email {
		return errors.New("Email " + input.Email + " already exist. Please pick another email.")
	}

	if errFindEmail != nil && !strings.Contains(errFindEmail.Error(), "found") {
		return helper.ServiceErrorMsg(errFindEmail)
	}

	bytePass, errEncrypt := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if errEncrypt != nil {
		log.Error(errEncrypt.Error())
		return helper.ServiceErrorMsg(err)
	}

	input.Password = string(bytePass)
	errCreate := service.userRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

// GetAll implements user.ServiceInterface
func (service *userService) GetAll() (data []user.Core, err error) {

	data, err = service.userRepository.GetAll()

	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	if len(data) == 0 {
		helper.LogDebug("Get data success. No data.")
		return nil, errors.New("Get data success. No data.")
	}

	return data, err
}

func (service *userService) GetById(id int) (data user.Core, err error) {
	data, err = service.userRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		return user.Core{}, helper.ServiceErrorMsg(err)
	}

	if (user.Core{}) == data {
		helper.LogDebug("Get data success. No data.")
		return user.Core{}, errors.New("Get data success. No data.")
	}

	return data, err

}

func (service *userService) Update(input user.Core, id int) error {

	if input.Password != "" {
		generate, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
		input.Password = string(generate)
	}
	// validasi email harus unik pas update, kalau email nya sama dgn punya dia gpp
	data, errFindEmail := service.userRepository.FindUser(input.Email)
	if (data.Email == input.Email) && (data.ID != uint(id)) && input.Email != "" {
		return errors.New("Failed. Email " + input.Email + " already exist at other user. Please pick another email.")
	}

	if errFindEmail != nil && !strings.Contains(errFindEmail.Error(), "found") {
		return helper.ServiceErrorMsg(errFindEmail)
	}

	// proses
	err := service.userRepository.Update(input, id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

func (service *userService) Delete(id int) error {
	// proses
	err := service.userRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}

// GetClubs implements user.ServiceInterface
func (service *userService) GetClubs(id int) (data []clubMember.Core, err error) {
	data, err = service.userRepository.GetClubs(id)
	if err != nil {
		log.Error(err.Error())
		return data, helper.ServiceErrorMsg(err)
	}

	if len(data) == 0 {
		helper.LogDebug("Get data success. No data.")
		return nil, errors.New("Get data success. No data.")
	}

	return data, err
}

// GetEvents implements user.ServiceInterface
func (service *userService) GetEvents(id int) (data []event.EventCore, err error) {
	data, err = service.userRepository.GetEvents(id)
	if err != nil {
		log.Error(err.Error())
		return data, helper.ServiceErrorMsg(err)
	}

	if len(data) == 0 {
		helper.LogDebug("Get data success. No data.")
		return nil, errors.New("Get data success. No data.")
	}

	return data, err
}

// GetProducts implements user.ServiceInterface
func (service *userService) GetProducts(id int) (data []product.ProductCore, err error) {
	data, err = service.userRepository.GetProducts(id)
	if err != nil {
		log.Error(err.Error())
		return data, helper.ServiceErrorMsg(err)
	}

	if len(data) == 0 {
		helper.LogDebug("Get data success. No data.")
		return nil, errors.New("Get data success. No data.")
	}

	return data, err
}

// // GetTransactions implements user.ServiceInterface
// func (service *userService) GetTransactions(id int) (data []user.Transaction, err error) {
// 	data, err = service.userRepository.GetTransactions(id)
// 	if err != nil {
// 		log.Error(err.Error())
// 		return data, helper.ServiceErrorMsg(err)
// 	}

// 	if len(data) == 0 {
// 		helper.LogDebug("Get data success. No data.")
// 		return nil, errors.New("Get data success. No data.")
// 	}

// 	return data, err
// }
