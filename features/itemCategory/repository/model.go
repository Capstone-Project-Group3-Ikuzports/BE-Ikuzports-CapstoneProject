package repository

import (
	itemcategory "ikuzports/features/itemCategory"

	"gorm.io/gorm"
)

type ItemCategory struct {
	gorm.Model
	Name string
}

func (dataModel *ItemCategory) toCore() itemcategory.ItemCategoryCore {
	return itemcategory.ItemCategoryCore{
		ID:   dataModel.ID,
		Name: dataModel.Name,
	}
}

func toCoreList(dataModel []ItemCategory) []itemcategory.ItemCategoryCore {
	var dataCore []itemcategory.ItemCategoryCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
