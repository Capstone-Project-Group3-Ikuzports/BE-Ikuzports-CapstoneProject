package delivery

import (
	_chat "ikuzports/features/chat"
	_chatRes "ikuzports/features/chat/delivery"
	"ikuzports/features/club"
	"ikuzports/features/clubActivity"
	_activity "ikuzports/features/clubActivity/delivery"
	_members "ikuzports/features/clubMember"
	"ikuzports/features/galery"
	_galery "ikuzports/features/galery/delivery"
)

type ClubResponse struct {
	ID           uint   `json:"id" form:"id"`
	Name         string `json:"name" form:"name"`
	Address      string `json:"address" form:"address"`
	City         string `json:"city" form:"city"`
	CategoryID   uint   `json:"category_id" form:"category_id"`
	CategoryName string `json:"category_name" form:"category_name"`
	Description  string `json:"description" form:"description"`
	Logo         string `json:"logo" form:"logo"`
	JoinedMember uint   `json:"joined_member" form:"joined_member"`
	MemberTotal  uint   `json:"member_total" form:"member_total"`
	Rule         string `json:"rule" form:"rule"`
	Requirement  string `json:"requirement" form:"requirement"`
}

type MemberResponse struct {
	ID              uint   `json:"id" form:"id"`
	UserID          uint   `json:"user_id" form:"user_id"`
	UserName        string `json:"name" form:"name"`
	UserPhoneNumber string `json:"phone_number" form:"phone_number"`
	UserImage       string `json:"user_image" form:"user_image"`
	ClubID          uint   `json:"club_id" form:"club_id"`
	Status          string `json:"status" form:"status"`
}

func fromCore(dataCore club.Core) ClubResponse {
	return ClubResponse{
		ID:           dataCore.ID,
		Name:         dataCore.Name,
		Address:      dataCore.Address,
		City:         dataCore.City,
		CategoryID:   dataCore.Category.ID,
		CategoryName: dataCore.Category.Name,
		Description:  dataCore.Description,
		Logo:         dataCore.Logo,
		JoinedMember: dataCore.JoinedMember,
		MemberTotal:  dataCore.MemberTotal,
		Rule:         dataCore.Rule,
		Requirement:  dataCore.Requirement,
	}
}

func fromCoreList(dataCore []club.Core) []ClubResponse {
	var dataResponse []ClubResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

func fromCoreMember(dataCore _members.Core) MemberResponse {
	return MemberResponse{
		ID:              dataCore.ID,
		UserID:          dataCore.User.ID,
		ClubID:          dataCore.ClubID,
		UserName:        dataCore.User.Name,
		UserPhoneNumber: dataCore.User.Phone_number,
		UserImage:       dataCore.User.UserImage,
		Status:          dataCore.Status,
	}
}
func fromCoreMemberList(dataCore []_members.Core) []MemberResponse {
	var dataResponse []MemberResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreMember(v))
	}
	return dataResponse
}

func fromCoreChat(dataCore _chat.Core) _chatRes.ChatResponse {
	return _chatRes.ChatResponse{
		ID:        dataCore.ID,
		UserID:    dataCore.UserID,
		UserName:  dataCore.User.Name,
		ClubID:    dataCore.ClubID,
		Message:   dataCore.Message,
		CreatedAt: dataCore.CreatedAt,
	}
}
func fromCoreChatList(dataCore []_chat.Core) []_chatRes.ChatResponse {
	var dataResponse []_chatRes.ChatResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreChat(v))
	}
	return dataResponse
}

func fromCoreGalery(dataCore galery.Core) _galery.GaleryResponse {
	return _galery.GaleryResponse{
		ID:      dataCore.ID,
		Url:     dataCore.Url,
		ClubID:  dataCore.ClubID,
		Caption: dataCore.Caption,
	}
}
func fromCoreGaleryList(dataCore []galery.Core) []_galery.GaleryResponse {
	var dataResponse []_galery.GaleryResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreGalery(v))
	}
	return dataResponse
}

func fromCoreActivity(dataCore clubActivity.Core) _activity.ActivityResponse {
	return _activity.ActivityResponse{
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
func fromCoreActivityList(dataCore []clubActivity.Core) []_activity.ActivityResponse {
	var dataResponse []_activity.ActivityResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCoreActivity(v))
	}
	return dataResponse
}
