package repository

import (
	"ikuzports/features/chat"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	UserID  uint
	User    User
	ClubID  uint
	Message string
}

type User struct {
	gorm.Model
	Name string
	Chat []Chat
}

func fromCore(dataCore chat.Core) Chat {
	return Chat{
		ClubID:  dataCore.ClubID,
		UserID:  dataCore.UserID,
		Message: dataCore.Message,
	}
}
func (dataCore *Chat) toCore() chat.Core {
	return chat.Core{
		ID:     dataCore.ID,
		UserID: dataCore.UserID,
		User: chat.User{
			ID:   dataCore.User.ID,
			Name: dataCore.User.Name,
		},
		ClubID:    dataCore.ClubID,
		Message:   dataCore.Message,
		CreatedAt: dataCore.CreatedAt,
	}
}

func toCoreList(dataModel []Chat) []chat.Core {
	var dataCore []chat.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
