package service

import (
	"errors"
	itemcategory "ikuzports/features/itemCategory"
)

type itemCategoryService struct {
	itemCategoryRepository itemcategory.RepositoryInterface
}

func New(repo itemcategory.RepositoryInterface) itemcategory.ServiceInterface {
	return &itemCategoryService{
		itemCategoryRepository: repo,
	}
}

func (service *itemCategoryService) GetAll() (data []itemcategory.ItemCategoryCore, err error) {
	data, err = service.itemCategoryRepository.GetAll()
	if err != nil {
		return nil, errors.New("error read data")
	}

	return data, nil
}
