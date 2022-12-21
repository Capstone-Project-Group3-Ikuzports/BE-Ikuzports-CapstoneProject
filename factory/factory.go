package factory

import (
	authDelivery "ikuzports/features/auth/delivery"
	authRepo "ikuzports/features/auth/repository"
	authService "ikuzports/features/auth/service"

	userDelivery "ikuzports/features/user/delivery"
	userRepo "ikuzports/features/user/repository"
	userService "ikuzports/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)
}
