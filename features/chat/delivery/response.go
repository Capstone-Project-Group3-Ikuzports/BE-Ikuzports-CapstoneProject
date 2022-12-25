package delivery

import (
	"ikuzports/features/chat"
	"time"
)

type ChatResponse struct {
	ID        uint      `json:"id" form:"id"`
	UserID    uint      `json:"user_id" form:"user_id"`
	UserName  string    `json:"user_name" form:"user_name"`
	ClubID    uint      `json:"club_id" form:"club_id"`
	Message   string    `json:"message" form:"message"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

func fromCore(dataCore chat.Core) ChatResponse {
	return ChatResponse{
		ID:        dataCore.ID,
		UserID:    dataCore.UserID,
		UserName:  dataCore.User.Name,
		ClubID:    dataCore.ClubID,
		Message:   dataCore.Message,
		CreatedAt: dataCore.CreatedAt,
	}
}

func fromCoreList(dataCore []chat.Core) []ChatResponse {
	var dataResponse []ChatResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
