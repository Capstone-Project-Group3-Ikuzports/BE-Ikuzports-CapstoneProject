package repository

import (
	"ikuzports/features/galery"

	"gorm.io/gorm"
)

type Galery struct {
	gorm.Model
	Url     string
	ClubID  uint
	Caption string
}

type Club struct {
	gorm.Model
	Name   string
	Galery []Galery
}

func fromCore(dataCore galery.Core) Galery {
	return Galery{
		Url:     dataCore.Url,
		ClubID:  dataCore.ClubID,
		Caption: dataCore.Caption,
	}
}
func (dataCore *Galery) toCore() galery.Core {
	return galery.Core{
		ID:      dataCore.ID,
		Url:     dataCore.Url,
		ClubID:  dataCore.ClubID,
		Caption: dataCore.Caption,
	}
}

func toCoreList(dataModel []Galery) []galery.Core {
	var dataCore []galery.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
