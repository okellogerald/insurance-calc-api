package db

import (
	"insurance-calc-api/logs"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	i, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))
	if err != nil {
		panic("unable to connect to database")
	} else {
		DB = i
		syncDB()
		logs.Info("successfully connected to database")
	}
}

func syncDB() {
}
