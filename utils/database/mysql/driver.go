package mysql

import (
	"fmt"
	"ikuzports/config"
	chat "ikuzports/features/chat/repository"
	club "ikuzports/features/club/repository"
	clubActivity "ikuzports/features/clubActivity/repository"
	clubMember "ikuzports/features/clubMember/repository"
	event "ikuzports/features/event/repository"
	galery "ikuzports/features/galery/repository"
	product "ikuzports/features/product/repository"
	productImage "ikuzports/features/productImage/repository"
	transaction "ikuzports/features/transaction/repository"
	user "ikuzports/features/user/repository"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	migrateDB(db)

	return db
}

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&club.Club{})
	db.AutoMigrate(&chat.Chat{})
	db.AutoMigrate(&clubActivity.ClubActivity{})
	db.AutoMigrate(&club.Aggreement{})
	db.AutoMigrate(&galery.Galery{})
	db.AutoMigrate(&clubMember.ClubMember{})
	db.AutoMigrate(&event.Category{})
	db.AutoMigrate(&event.EventParticipant{})
	db.AutoMigrate(&event.Event{})
	db.AutoMigrate(&transaction.Transaction{})
	db.AutoMigrate(&product.ItemCategory{})
	db.AutoMigrate(&productImage.ProductImage{})
	db.AutoMigrate(&product.Product{})

}
