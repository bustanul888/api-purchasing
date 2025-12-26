package db

import (
	"fmt"
	"log"
	"os"
	"task-be/app/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connection() *gorm.DB {
	godotenv.Load()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func init() {
	db := Connection()
	err := db.AutoMigrate(
		model.Supplier{},
		model.Item{},
		model.Purchasing{},
		model.PurchasingDetail{},
		model.User{},
	)
	if err != nil {
		panic(err)
	}
}
