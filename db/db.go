package db

import (
	"log"
	"os"

	"github.com/duncanleo/brawl-scraper/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/joho/godotenv/autoload"
)

var (
	// DB Database
	DB *gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Panic(err)
	}
	DB.AutoMigrate(&model.Player{}, &model.PlayerData{}, &model.Brawler{})
}
