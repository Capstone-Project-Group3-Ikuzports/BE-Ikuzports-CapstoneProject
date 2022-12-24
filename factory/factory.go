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

	eventDelivery "ikuzports/features/event/delivery"
	eventRepo "ikuzports/features/event/repository"
	eventService "ikuzports/features/event/service"

	participantDelivery "ikuzports/features/participant/delivery"
	participantRepo "ikuzports/features/participant/repository"
	participantService "ikuzports/features/participant/service"

	categoryDelivery "ikuzports/features/category/delivery"
	categoryRepo "ikuzports/features/category/repository"
	categoryService "ikuzports/features/category/service"

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

	participantRepoFactory := participantRepo.New(db)
	participantServiceFactory := participantService.New(participantRepoFactory)
	participantDelivery.New(participantServiceFactory, e)

	eventRepoFactory := eventRepo.New(db)
	eventServiceFactory := eventService.New(eventRepoFactory, participantRepoFactory)
	eventDelivery.New(eventServiceFactory, e)

	categoryRepoFactory := categoryRepo.New(db)
	categoryServiceFactory := categoryService.New(categoryRepoFactory)
	categoryDelivery.New(categoryServiceFactory, e)
}
