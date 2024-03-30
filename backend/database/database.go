package database

import (
	"fmt"
	"log"
	"os"

	"github.com/domenicwalther/rubikon/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=aws-0-eu-central-1.pooler.supabase.com user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database.", err)
		os.Exit(2)
	}

	log.Println("Connected to database.")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Migrating database...")
	db.AutoMigrate(&models.Skin{}, &models.User{}, &models.Streak{})
	DB = Dbinstance{Db: db}
}
