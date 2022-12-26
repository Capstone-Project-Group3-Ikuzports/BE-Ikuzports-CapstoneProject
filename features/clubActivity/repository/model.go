package repository

import (
	"ikuzports/features/clubActivity"

	"gorm.io/gorm"
)

type ClubActivity struct {
	gorm.Model
	ClubID         uint
	Name           string
	StartTime      string
	EndTime        string
	Day            string
	Location       string
	ActivityDetail string
}

func fromCore(dataCore clubActivity.Core) ClubActivity {
	return ClubActivity{
		ClubID:         dataCore.ClubID,
		Name:           dataCore.Name,
		StartTime:      dataCore.StartTime,
		EndTime:        dataCore.EndTime,
		Day:            dataCore.Day,
		Location:       dataCore.Location,
		ActivityDetail: dataCore.ActivityDetail,
	}
}
func (dataCore *ClubActivity) toCore() clubActivity.Core {
	return clubActivity.Core{
		ID:             dataCore.ID,
		ClubID:         dataCore.ClubID,
		Name:           dataCore.Name,
		StartTime:      dataCore.StartTime,
		EndTime:        dataCore.EndTime,
		Day:            dataCore.Day,
		Location:       dataCore.Location,
		ActivityDetail: dataCore.ActivityDetail,
	}
}

func toCoreList(dataModel []ClubActivity) []clubActivity.Core {
	var dataCore []clubActivity.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
