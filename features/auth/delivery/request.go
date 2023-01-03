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

type TokenRequest struct {
	AccessToken string `json:"access_token" form:"access_token"`
	Authuser    string `json:"authuser" form:"authuser"`
	ExpiresIn   int    `json:"expires_in" form:"expires_in"`
	Prompt      string `json:"prompt" form:"prompt"`
	Scope       string `json:"scope" form:"scope"`
	TokenType   string `json:"token_type" form:"token_type"`
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
