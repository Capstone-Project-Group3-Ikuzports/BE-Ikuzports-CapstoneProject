package delivery

import "ikuzports/features/galery"

type InsertRequest struct {
	Url     string `json:"url" form:"url"`
	ClubID  uint   `json:"club_id" form:"club_id"`
	Caption string `json:"caption" form:"caption"`
}

func toCore(dataCore InsertRequest) galery.Core {
	return galery.Core{
		Url:     dataCore.Url,
		ClubID:  dataCore.ClubID,
		Caption: dataCore.Caption,
	}
}
