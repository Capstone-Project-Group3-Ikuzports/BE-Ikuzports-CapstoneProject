package service

import (
	"ikuzports/features/product"
	"ikuzports/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type productService struct {
	productRepository product.RepositoryInterface
	validate          *validator.Validate
}

func New(repo product.RepositoryInterface) product.ServiceInterface {
	return &productService{
		productRepository: repo,
		validate:          validator.New(),
	}
}

func (service *productService) GetAll(queryItemCategoryID int, queryCity, queryName string, queryPage int) (data []product.ProductCore, page int, err error) {
	limit := 12
	offset := (queryPage - 1) * limit
	var jumlahData int

	if queryName == "" && queryCity == "" && queryItemCategoryID == 0 {
		data, jumlahData, err = service.productRepository.GetAll(limit, offset)
	} else {
		data, jumlahData, err = service.productRepository.GetAllFilter(queryItemCategoryID, queryCity, queryName, limit, offset)
	}

	if err != nil {
		helper.LogDebug(err)
		return nil, 0, helper.ServiceErrorMsg(err)
	}

	if jumlahData%limit == 0 {
		page = jumlahData / limit
	} else {
		page = (jumlahData / limit) + 1
	}

	return data, page, nil
}

func (service *productService) Create(input product.ProductCore) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	_, err = service.productRepository.Create(input)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}

	return nil
}

func (service *productService) GetByID(id int) (data product.ProductCore, err error) {
	data, err = service.productRepository.GetByID(id)
	if err != nil {
		log.Error(err.Error())
		return product.ProductCore{}, helper.ServiceErrorMsg(err)
	}

	return data, nil
}

func (service *productService) Update(id int, input product.ProductCore) (err error) {
	_, err = service.productRepository.Update(id, input)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}

func (service *productService) Delete(id int) (err error) {
	_, err = service.productRepository.Delete(id)
	if err != nil {
		log.Error(err.Error())
		return helper.ServiceErrorMsg(err)
	}
	return nil
}

// GetImages implements product.ServiceInterface
func (service *productService) GetImages(id int) (data []product.ProductImage, err error) {
	data, err = service.productRepository.GetImages(id)
	if err != nil {
		log.Error(err.Error())
		return data, helper.ServiceErrorMsg(err)
	}

	return data, err
}
