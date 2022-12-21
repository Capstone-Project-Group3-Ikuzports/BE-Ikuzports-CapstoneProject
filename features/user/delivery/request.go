package delivery

import "ikuzports/features/user"

type InsertRequest struct {
	FullName        string `json:"full_name" form:"full_name"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	Phone           string `json:"phone" form:"phone"`
	Gender          string `json:"gender" form:"gender"`
	ProfileImageUrl string `json:"profile_image_url" form:"profile_image_url"`
}

type UpdateRequest struct {
	ID              uint   `json:"id" form:"id"`
	FullName        string `json:"full_name" form:"full_name"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	Phone           string `json:"phone" form:"phone"`
	Gender          string `json:"gender" form:"gender"`
	ProfileImageUrl string `json:"profile_image_url" form:"profile_image_url"`
}

func toCore(i interface{}) user.Core {
	switch i.(type) {
	case InsertRequest:
		cnv := i.(InsertRequest)
		return user.Core{
			FullName:        cnv.FullName,
			Email:           cnv.Email,
			Password:        cnv.Password,
			Phone:           cnv.Phone,
			Gender:          cnv.Gender,
			ProfileImageUrl: cnv.ProfileImageUrl,
		}

	case UpdateRequest:
		cnv := i.(UpdateRequest)
		return user.Core{
			ID:              cnv.ID,
			FullName:        cnv.FullName,
			Email:           cnv.Email,
			Password:        cnv.Password,
			Phone:           cnv.Phone,
			Gender:          cnv.Gender,
			ProfileImageUrl: cnv.ProfileImageUrl,
		}
	}

	return user.Core{}
}
