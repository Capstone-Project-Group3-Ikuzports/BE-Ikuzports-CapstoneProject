package delivery

import "ikuzports/features/chat"

type InsertRequest struct {
	UserID  uint   `json:"user_id" form:"user_id"`
	ClubID  uint   `json:"club_id" form:"club_id"`
	Message string `json:"message" form:"message"`
}

func toCore(dataCore InsertRequest) chat.Core {
	return chat.Core{
		UserID:  dataCore.UserID,
		ClubID:  dataCore.ClubID,
		Message: dataCore.Message,
	}
}
