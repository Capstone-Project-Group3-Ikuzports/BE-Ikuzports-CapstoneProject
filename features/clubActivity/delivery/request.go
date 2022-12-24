package delivery

import (
	"ikuzports/features/clubActivity"
	"time"
)

type InsertRequest struct {
	ClubID         uint   `json:"club_id" form:"club_id"`
	Name           string `json:"name" form:"name"`
	StartTime      string `json:"start_time" form:"start_time"`
	EndTime        string `json:"end_time" form:"end_time"`
	Day            string `json:"day" form:"day"`
	Location       string `json:"location" form:"location"`
	ActivityDetail string `json:"activity_detail" form:"activity_detail"`
}

func toCore(dataCore InsertRequest) clubActivity.Core {
	return clubActivity.Core{
		ClubID:         dataCore.ClubID,
		Name:           dataCore.Name,
		StartTime:      timeParsing(dataCore.StartTime),
		EndTime:        timeParsing(dataCore.EndTime),
		Day:            dataCore.Day,
		Location:       dataCore.Location,
		ActivityDetail: dataCore.ActivityDetail,
	}
}

func timeParsing(year string) (dateParsed time.Time) {
	layoutFormat := "2006-01-02 15:04:05"

	dateParsed, _ = time.Parse(layoutFormat, year)

	return dateParsed
}
