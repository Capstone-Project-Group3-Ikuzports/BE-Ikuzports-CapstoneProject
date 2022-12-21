package delivery

import "ikuzports/features/user"

type UserResponse struct {
	ID              uint   `json:"id"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Gender          string `json:"gender"`
	ProfileImageUrl string `json:"profile_image_url"`
	Role            string `json:"role"`
}

func fromCore(dataCore user.Core) UserResponse {
	return UserResponse{
		ID:              dataCore.ID,
		FullName:        dataCore.FullName,
		Email:           dataCore.Email,
		Phone:           dataCore.Phone,
		Gender:          dataCore.Gender,
		ProfileImageUrl: dataCore.ProfileImageUrl,
		Role:            dataCore.Role,
	}
}

func fromCoreList(dataCore []user.Core) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
