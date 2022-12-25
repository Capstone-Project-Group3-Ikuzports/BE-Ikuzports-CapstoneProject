package delivery

import (
	"ikuzports/features/galery"
)

type GaleryResponse struct {
	ID      uint   `json:"id" form:"id"`
	Url     string `json:"url" form:"url"`
	ClubID  uint   `json:"club_id" form:"club_id"`
	Caption string `json:"caption" form:"caption"`
}

func fromCore(dataCore galery.Core) GaleryResponse {
	return GaleryResponse{
		ID:      dataCore.ID,
		Url:     dataCore.Url,
		ClubID:  dataCore.ClubID,
		Caption: dataCore.Caption,
	}
}

func fromCoreList(dataCore []galery.Core) []GaleryResponse {
	var dataResponse []GaleryResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
