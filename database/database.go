package database

import (
	"github.com/1Nelsonel/BlogPost/model"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=nelsonel01 password=0101 dbname=blogpost port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic("Error connecting to database")
	}

	log.Info("Connection success")

	db.AutoMigrate(new(model.Blog))

	DBConn = db
}