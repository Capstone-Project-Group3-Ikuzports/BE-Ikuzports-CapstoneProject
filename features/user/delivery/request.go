package delivery

import "ikuzports/features/user"

type InsertRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
	UserImage   string `json:"user_image" form:"user_image"`
	Gender      string `json:"gender" form:"gender"`
}

func toCore(userInput InsertRequest) user.Core {
	userCoreData := user.Core{
		Name:        userInput.Name,
		Email:       userInput.Email,
		Password:    userInput.Password,
		PhoneNumber: userInput.PhoneNumber,
		Address:     userInput.Address,
		UserImage:   userInput.UserImage,
		Gender:      userInput.Gender,
	}
	return userCoreData
}
