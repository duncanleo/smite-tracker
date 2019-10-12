package main

import (
	"github.com/duncanleo/smite-tracker/db"
	"github.com/duncanleo/smite-tracker/web"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db.DB.LogMode(gin.IsDebugging())
	web.StartServer()
}
