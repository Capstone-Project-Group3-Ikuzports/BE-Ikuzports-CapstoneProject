package delivery

import (
	"ikuzports/features/clubActivity"
)

type ActivityResponse struct {
	ID             uint   `json:"id" form:"id"`
	ClubID         uint   `json:"club_id" form:"club_id"`
	Name           string `json:"name" form:"name"`
	StartTime      string `json:"start_time" form:"start_time"`
	EndTime        string `json:"end_time" form:"end_time"`
	Day            string `json:"day" form:"day"`
	Location       string `json:"location" form:"location"`
	ActivityDetail string `json:"activity_detail" form:"activity_detail"`
}

func fromCore(dataCore clubActivity.Core) ActivityResponse {
	return ActivityResponse{
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

func fromCoreList(dataCore []clubActivity.Core) []ActivityResponse {
	var dataResponse []ActivityResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
