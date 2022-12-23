package delivery

import (
	"ikuzports/features/clubMember"
)

type MemberResponse struct {
	ID     uint   `json:"id" form:"id"`
	UserID uint   `json:"user_id" form:"user_id"`
	ClubID uint   `json:"club_id" form:"club_id"`
	Status string `json:"status" form:"status"`
}

func fromCore(dataCore clubMember.Core) MemberResponse {
	return MemberResponse{
		ID:     dataCore.ID,
		UserID: dataCore.UserID,
		ClubID: dataCore.ClubID,
		Status: dataCore.Status,
	}
}

func fromCoreList(dataCore []clubMember.Core) []MemberResponse {
	var dataResponse []MemberResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
