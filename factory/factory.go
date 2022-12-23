package factory

import (
	authDelivery "ikuzports/features/auth/delivery"
	authRepo "ikuzports/features/auth/repository"
	authService "ikuzports/features/auth/service"

	userDelivery "ikuzports/features/user/delivery"
	userRepo "ikuzports/features/user/repository"
	userService "ikuzports/features/user/service"

	clubDelivery "ikuzports/features/club/delivery"
	clubRepo "ikuzports/features/club/repository"
	clubService "ikuzports/features/club/service"

	clubMemberDelivery "ikuzports/features/clubMember/delivery"
	clubMemberRepo "ikuzports/features/clubMember/repository"
	clubMemberService "ikuzports/features/clubMember/service"

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

	clubMemberRepoFactory := clubMemberRepo.New(db)
	clubMemberServiceFactory := clubMemberService.New(clubMemberRepoFactory)
	clubMemberDelivery.New(clubMemberServiceFactory, e)

	clubRepoFactory := clubRepo.New(db)
	clubServiceFactory := clubService.New(clubRepoFactory, clubMemberRepoFactory)
	clubDelivery.New(clubServiceFactory, e)

}
