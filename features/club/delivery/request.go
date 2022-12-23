package delivery

import "ikuzports/features/club"

type InsertRequest struct {
	Name           string `json:"name" form:"name"`
	Address        string `json:"address" form:"address"`
	City           string `json:"city" form:"city"`
	CategoryID     uint   `json:"category_id" form:"category_id"`
	Description    string `json:"description" form:"description"`
	Logo           string `json:"logo" form:"logo"`
	TermsCondition string `json:"terms_condition" form:"terms_condition"`
	MemberTotal    int    `json:"member_total" form:"member_total"`
	Rule           string `json:"rule" form:"rule"`
	Requirement    string `json:"requirement" form:"requirement"`
}

func toCore(clubInput InsertRequest) club.Core {
	clubCoreData := club.Core{
		Name:        clubInput.Name,
		Address:     clubInput.Address,
		City:        clubInput.City,
		CategoryID:  clubInput.CategoryID,
		Description: clubInput.Description,
		Logo:        clubInput.Logo,
		MemberTotal: clubInput.MemberTotal,
		Rule:        clubInput.Rule,
		Requirement: clubInput.Requirement,
	}
	return clubCoreData
}
