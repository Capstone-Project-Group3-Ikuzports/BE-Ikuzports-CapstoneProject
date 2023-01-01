package delivery

import (
	"ikuzports/features/auth"
	"ikuzports/features/user"
)

type UserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type GoogleRequest struct {
	Email   string `json:"email" form:"email"`
	Name    string `json:"name" form:"name"`
	Picture string `json:"picture" form:"picture"`
}

func ToCore(userReq UserRequest) auth.Core {
	userCore := auth.Core{
		Email:    userReq.Email,
		Password: userReq.Password,
	}
	return userCore
}

func ToCoreGoogle(googleReq GoogleRequest) user.GoogleCore {
	googleCore := user.GoogleCore{
		Email:   googleReq.Email,
		Name:    googleReq.Name,
		Picture: googleReq.Picture,
	}
	return googleCore
}
