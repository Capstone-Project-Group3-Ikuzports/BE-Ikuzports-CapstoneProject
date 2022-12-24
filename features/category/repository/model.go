package repository

import (
	"ikuzports/features/category"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string
}

func (dataModel *Category) toCore() category.CategoryCore {
	return category.CategoryCore{
		ID:   dataModel.ID,
		Name: dataModel.Name,
	}
}

func toCoreList(dataModel []Category) []category.CategoryCore {
	var dataCore []category.CategoryCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
