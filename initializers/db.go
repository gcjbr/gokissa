package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dbUrl := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	DB = db

	if err != nil {
		panic("🔴 Failed to connect to DB!")
	}
}

func SyncDb() {
	DB.AutoMigrate()
}
