package service

import (
	"errors"
	"ikuzports/features/user"
	"ikuzports/utils/helper"
	"ikuzports/utils/thirdparty"

	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
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
func (service *userService) Create(input user.Core, c echo.Context) (err error) {
	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	// validasi email harus unik
	data, errFindEmail := service.userRepository.FindUser(input.Email)

	// helper.LogDebug("\n\n\n find email input  ", input.Email)
	// helper.LogDebug("\n\n\n find email data  ", data.Email)

	if data.Email == input.Email {
		return errors.New("Email " + input.Email + " already exist. Please pick another email.")
	}

	if errFindEmail != nil && !strings.Contains(errFindEmail.Error(), "found") {
		return helper.ServiceErrorMsg(errFindEmail)
	}

	// upload foto
	file, _ := c.FormFile("file")
	if file != nil {
		res, err := thirdparty.UploadProfile(c)
		if err != nil {
			return errors.New("Failed. Cannot Upload Data.")
		}
		log.Print(res)
		input.ProfileImageUrl = res
	} else {
		input.ProfileImageUrl = "https://www.hostpapa.com/knowledgebase/wp-content/uploads/2018/04/1-13.png"
	}

	bytePass, errEncrypt := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if errEncrypt != nil {
		log.Error(errEncrypt.Error())
		return helper.ServiceErrorMsg(err)
	}

	input.Password = string(bytePass)

	input.Role = "User"
	errCreate := service.userRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

// GetAll implements user.ServiceInterface
func (service *userService) GetAll(query string) (data []user.Core, err error) {

	if query == "" {
		data, err = service.userRepository.GetAll()
	} else {
		data, err = service.userRepository.GetAllWithSearch(query)
	}

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

func (service *userService) Update(input user.Core, id int, c echo.Context) error {
	// validasi input
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	if input.Password != "" {
		generate, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
		input.Password = string(generate)
	}

	// validasi user dgn id path param, apakah ada datanya di database
	_, errFindId := service.userRepository.GetById(id)
	if errFindId != nil {
		log.Error(errFindId.Error())
		return helper.ServiceErrorMsg(errFindId)
	}

	// validasi email harus unik pas update, kalau email nya sama dgn punya dia gpp
	data, errFindEmail := service.userRepository.FindUser(input.Email)

	// helper.LogDebug("\n\n\n find email input  ", input.Email, "--- id ", id)
	// helper.LogDebug("\n\n\n find email data  ", data.Email, "--- id ", data.ID)

	if (data.Email == input.Email) && (data.ID != uint(id)) {
		return errors.New("Failed. Email " + input.Email + " already exist at other user. Please pick another email.")
	}

	if errFindEmail != nil && !strings.Contains(errFindEmail.Error(), "found") {
		return helper.ServiceErrorMsg(errFindEmail)
	}

	// upload foto
	file, _ := c.FormFile("file")
	if file != nil {
		res, err := thirdparty.UploadProfile(c)
		if err != nil {
			return errors.New("Failed. Cannot Upload Data.")
		}
		log.Print(res)
		input.ProfileImageUrl = res
	} else {
		input.ProfileImageUrl = "https://www.hostpapa.com/knowledgebase/wp-content/uploads/2018/04/1-13.png"
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
	// validasi user dgn id path param, apakah ada datanya di database
	_, errFindId := service.userRepository.GetById(id)
	if errFindId != nil {
		log.Error(errFindId.Error())
		return helper.ServiceErrorMsg(errFindId)
	}

	// proses
	err := service.userRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}
