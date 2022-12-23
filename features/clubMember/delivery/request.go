package delivery

import "ikuzports/features/clubMember"

type MemberRequest struct {
	ID     uint   `json:"id" form:"id"`
	UserID uint   `json:"user_id" form:"user_id"`
	ClubID uint   `json:"club_id" form:"club_id"`
	Status string `json:"status" form:"status"`
}

func toCore(dataCore MemberRequest) clubMember.Core {
	return clubMember.Core{
		ID:     dataCore.ID,
		UserID: dataCore.UserID,
		ClubID: dataCore.ClubID,
		Status: dataCore.Status,
	}
}
