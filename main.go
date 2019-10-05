package main

import (
	"github.com/duncanleo/brawl-scraper/db"
	"github.com/duncanleo/brawl-scraper/web"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db.DB.LogMode(gin.IsDebugging())
	web.StartServer()
}
