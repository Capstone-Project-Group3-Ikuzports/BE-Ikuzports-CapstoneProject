package delivery

import "ikuzports/features/club"

type ClubResponse struct {
	ID           uint   `json:"id" form:"id"`
	Name         string `json:"name" form:"name"`
	Address      string `json:"address" form:"address"`
	City         string `json:"city" form:"city"`
	CategoryID   uint   `json:"category_id" form:"category_id"`
	CategoryName string `json:"category_name" form:"category_name"`
	Description  string `json:"description" form:"description"`
	Logo         string `json:"logo" form:"logo"`
	JoinedMember int    `json:"joined_member" form:"joined_member"`
	MemberTotal  int    `json:"member_total" form:"member_total"`
	Rule         string `json:"rule" form:"rule"`
	Requirement  string `json:"requirement" form:"requirement"`
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
