package mysql

import (
	"fmt"
	"ikuzports/config"
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
	db.AutoMigrate(&user.Club{})
	db.AutoMigrate(&user.Chat{})
	db.AutoMigrate(&user.ClubActivity{})
	db.AutoMigrate(&user.Aggreement{})
	db.AutoMigrate(&user.Galery{})
	db.AutoMigrate(&user.ClubMember{})
	db.AutoMigrate(&user.Category{})
	db.AutoMigrate(&user.EventParticipant{})
	db.AutoMigrate(&user.Event{})
	db.AutoMigrate(&user.Transaction{})
	db.AutoMigrate(&user.ItemCategory{})
	db.AutoMigrate(&user.ProductImage{})
	db.AutoMigrate(&user.Product{})

}
