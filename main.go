package main

import (
	"os"

	"github.com/duncanleo/brawl-scraper/db"
	"github.com/duncanleo/brawl-scraper/web"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Fixie support
	if len(os.Getenv("FIXIE_URL")) > 0 {
		os.Setenv("HTTP_PROXY", os.Getenv("FIXIE_URL"))
		os.Setenv("HTTPS_PROXY", os.Getenv("FIXIE_URL"))
	}
	db.DB.LogMode(gin.IsDebugging())
	web.StartServer()
}
