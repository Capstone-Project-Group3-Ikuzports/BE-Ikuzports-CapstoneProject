package service

import (
	"errors"
	"ikuzports/features/category"
)

type categoryService struct {
	categoryRepository category.RepositoryInterface
}

func New(repo category.RepositoryInterface) category.ServiceInterface {
	return &categoryService{
		categoryRepository: repo,
	}
}

func (service *categoryService) GetAll() (data []category.CategoryCore, err error) {
	data, err = service.categoryRepository.GetAll()
	if err != nil {
		return nil, errors.New("error read data")
	}

	return data, nil
}
