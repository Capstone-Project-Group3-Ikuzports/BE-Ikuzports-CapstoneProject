package delivery

import "ikuzports/features/auth"

type UserResponse struct {
	ID              uint   `json:"id"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Gender          string `json:"gender"`
	ProfileImageUrl string `json:"profile_image_url"`
	Role            string `json:"role"`
	Token           string `json:"token"`
}

func FromCore(dataCore auth.Core, token string) UserResponse {
	return UserResponse{
		ID:              dataCore.ID,
		FullName:        dataCore.FullName,
		Email:           dataCore.Email,
		Phone:           dataCore.Phone,
		Gender:          dataCore.Gender,
		ProfileImageUrl: dataCore.ProfileImageUrl,
		Role:            dataCore.Role,
		Token:           token,
	}
}
