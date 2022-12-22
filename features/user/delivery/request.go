package delivery

import "ikuzports/features/user"

type InsertRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
	City        string `json:"city" form:"city"`
	UserImage   string `json:"user_image" form:"user_image"`
	Gender      string `json:"gender" form:"gender"`
	Biodata     string `json:"biodata" form:"biodata"`
}

func toCore(userInput InsertRequest) user.Core {
	userCoreData := user.Core{
		Name:        userInput.Name,
		Email:       userInput.Email,
		Password:    userInput.Password,
		PhoneNumber: userInput.PhoneNumber,
		Address:     userInput.Address,
		City:        userInput.City,
		UserImage:   userInput.UserImage,
		Gender:      userInput.Gender,
		Biodata:     userInput.Biodata,
	}
	return userCoreData
}
