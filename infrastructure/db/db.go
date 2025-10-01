package db

import (
	"fmt"
	"log"

	"github.com/nepile/event-scanner/infrastructure/config"
	"github.com/nepile/event-scanner/internal/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConfig.DBUser,
		config.AppConfig.DBPass,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBName,
	)

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error to connect database: %s", err)
	}

	DB.AutoMigrate(
		&entity.User{},
		&entity.Event{},
		&entity.Registrant{},
		&entity.Presence{},
	)

	log.Print("Success to connected & migrated database")
}
