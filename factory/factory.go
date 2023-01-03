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

	clubActivityDelivery "ikuzports/features/clubActivity/delivery"
	clubActivityRepo "ikuzports/features/clubActivity/repository"
	clubActivityService "ikuzports/features/clubActivity/service"

	productDelivery "ikuzports/features/product/delivery"
	productRepo "ikuzports/features/product/repository"
	productService "ikuzports/features/product/service"

	galeryDelivery "ikuzports/features/galery/delivery"
	galeryRepo "ikuzports/features/galery/repository"
	galeryService "ikuzports/features/galery/service"

	chatDelivery "ikuzports/features/chat/delivery"
	chatRepo "ikuzports/features/chat/repository"
	chatService "ikuzports/features/chat/service"

	productImageDelivery "ikuzports/features/productImage/delivery"
	productImageRepo "ikuzports/features/productImage/repository"
	productImageService "ikuzports/features/productImage/service"

	itemCategoryDelivery "ikuzports/features/itemCategory/delivery"
	itemCategoryRepo "ikuzports/features/itemCategory/repository"
	itemCategoryService "ikuzports/features/itemCategory/service"

	transactionDelivery "ikuzports/features/transaction/delivery"
	transactionRepo "ikuzports/features/transaction/repository"
	transactionService "ikuzports/features/transaction/service"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB, googleOauthConfig *oauth2.Config) {
	userRepoFactory := userRepo.New(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory, userRepoFactory)
	authDelivery.New(authServiceFactory, e, googleOauthConfig, userServiceFactory)

	clubMemberRepoFactory := clubMemberRepo.New(db)

	clubRepoFactory := clubRepo.New(db)
	clubServiceFactory := clubService.New(clubRepoFactory, clubMemberRepoFactory)
	clubDelivery.New(clubServiceFactory, e)

	clubMemberServiceFactory := clubMemberService.New(clubMemberRepoFactory, clubRepoFactory)
	clubMemberDelivery.New(clubMemberServiceFactory, e)

	eventRepoFactory := eventRepo.New(db)

	participantRepoFactory := participantRepo.New(db)
	participantServiceFactory := participantService.New(participantRepoFactory, eventRepoFactory)
	participantDelivery.New(participantServiceFactory, e)

	eventServiceFactory := eventService.New(eventRepoFactory, participantRepoFactory)
	eventDelivery.New(eventServiceFactory, userServiceFactory, e, googleOauthConfig)

	categoryRepoFactory := categoryRepo.New(db)
	categoryServiceFactory := categoryService.New(categoryRepoFactory)
	categoryDelivery.New(categoryServiceFactory, e)

	productRepoFactory := productRepo.New(db)
	productServiceFactory := productService.New(productRepoFactory)
	productDelivery.New(productServiceFactory, e)

	clubActivityRepoFactory := clubActivityRepo.New(db)
	clubActivityServiceFactory := clubActivityService.New(clubActivityRepoFactory, clubRepoFactory)
	clubActivityDelivery.New(clubActivityServiceFactory, e)

	galeryRepoFactory := galeryRepo.New(db)
	galeryServiceFactory := galeryService.New(galeryRepoFactory, clubRepoFactory)
	galeryDelivery.New(galeryServiceFactory, e)

	chatRepoFactory := chatRepo.New(db)
	chatServiceFactory := chatService.New(chatRepoFactory)
	chatDelivery.New(chatServiceFactory, e)

	productImageRepoFactory := productImageRepo.New(db)
	productImageServiceFactory := productImageService.New(productImageRepoFactory)
	productImageDelivery.New(productImageServiceFactory, e)

	itemCategoryRepoFactory := itemCategoryRepo.New(db)
	itemCategoryServiceFactory := itemCategoryService.New(itemCategoryRepoFactory)
	itemCategoryDelivery.New(itemCategoryServiceFactory, e)

	transactionRepoFactory := transactionRepo.New(db)
	transactionServiceFactory := transactionService.New(transactionRepoFactory, productRepoFactory, userRepoFactory)
	transactionDelivery.New(transactionServiceFactory, e)
}
