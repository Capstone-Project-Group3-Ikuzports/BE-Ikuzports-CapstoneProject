package service

import (
	"errors"
	"fmt"
	"ikuzports/features/auth"
	"ikuzports/features/user"
	"ikuzports/middlewares"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	authData       auth.RepositoryInterface
	userRepository user.RepositoryInterface
	validate       *validator.Validate
}

func New(data auth.RepositoryInterface, userData user.RepositoryInterface) auth.ServiceInterface {
	return &authService{
		authData:       data,
		userRepository: userData,
		validate:       validator.New(),
	}
}

func (service *authService) Login(dataCore auth.Core) (auth.Core, string, error) {

	if errValidate := service.validate.Struct(dataCore); errValidate != nil {
		log.Error(errValidate.Error())
		return auth.Core{}, "", errors.New("Failed to login, error validate input, please check your input")
	}

	result, errLogin := service.authData.FindUser(dataCore.Email)
	if errLogin != nil {
		log.Error(errLogin.Error())
		if strings.Contains(errLogin.Error(), "table") {
			return auth.Core{}, "", errors.New("failed to login, error on request, please contact your administrator")
		} else if strings.Contains(errLogin.Error(), "found") {
			return auth.Core{}, "", errors.New("failed to login, email not found, please check password again")
		} else {
			return auth.Core{}, "", errors.New("failed to login, other error, please contact your administrator")
		}
	}

	errCheckPass := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(dataCore.Password))
	fmt.Println("Data Core = ", dataCore)
	fmt.Println("Result = ", result)
	if errCheckPass != nil {
		log.Error(errCheckPass.Error())
		return auth.Core{}, "", errors.New("failed to login, password didn't match, please check password again")
	}

	token, errToken := middlewares.CreateToken(int(result.ID), result.Name)
	if errToken != nil {
		log.Error(errToken.Error())
		return auth.Core{}, "", errors.New("failed to login, error on generate token, please check password again")
	}

	return result, token, nil
}

func (service *authService) LoginGoogle(input user.GoogleCore) (data auth.Core, token string, err error) {
	result, errLogin := service.authData.FindUser(input.Email)
	if errLogin != nil {
		log.Error(errLogin.Error())
		if strings.Contains(errLogin.Error(), "table") {
			return auth.Core{}, "", errors.New("failed to login, error on request, please contact your administrator")
		} else if strings.Contains(errLogin.Error(), "found") {
			userCore := user.Core{
				Name:      input.Name,
				Email:     input.Email,
				UserImage: input.Picture,
			}

			errCr := service.userRepository.Create(userCore)
			if errCr != nil {
				log.Error(errCr.Error())
			} else {
				res, _ := service.authData.FindUser(userCore.Email)
				token, errToken := middlewares.CreateToken(int(res.ID), res.Name)
				if errToken != nil {
					log.Error(errToken.Error())
					return auth.Core{}, "", errors.New("failed to login, error on generate token, please check password again")
				}

				return res, token, nil
			}
		} else {
			return auth.Core{}, "", errors.New("failed to login, other error, please contact your administrator")
		}
	}
	token, errToken := middlewares.CreateToken(int(result.ID), result.Name)
	if errToken != nil {
		log.Error(errToken.Error())
		return auth.Core{}, "", errors.New("failed to login, error on generate token, please check password again")
	}

	return result, token, nil

}
