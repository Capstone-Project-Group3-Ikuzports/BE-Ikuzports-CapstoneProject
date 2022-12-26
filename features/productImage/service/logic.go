package service

import (
	"errors"
	"ikuzports/features/productImage"
	"ikuzports/utils/helper"
)

type productImageService struct {
	productImageRepository productImage.RepositoryInterface
}

func New(repo productImage.RepositoryInterface) productImage.ServiceInterface {
	return &productImageService{
		productImageRepository: repo,
	}
}

func (service *productImageService) Create(input productImage.ProductImageCore) (err error) {
	_, err = service.productImageRepository.Create(input)
	if err != nil {
		return errors.New("failed to insert data. error query")
	}
	return nil
}

func (service *productImageService) GetAll() (data []productImage.ProductImageCore, err error) {
	data, err = service.productImageRepository.GetAll()
	if err != nil {
		helper.LogDebug(err)
		return nil, helper.ServiceErrorMsg(err)
	}

	return data, nil
}

func (service *productImageService) GetByID(id int) (data productImage.ProductImageCore, err error) {
	data, err = service.productImageRepository.GetByID(id)
	if err != nil {
		helper.LogDebug(err)
		return productImage.ProductImageCore{}, helper.ServiceErrorMsg(err)
	}

	return data, nil
}
