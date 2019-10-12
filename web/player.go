package web

import (
	"log"
	"net/http"

	"github.com/duncanleo/smite-tracker/db"
	"github.com/duncanleo/smite-tracker/model"
	"github.com/gin-gonic/gin"
)

func players(c *gin.Context) {
	var players []model.Player
	err := db.DB.
		Set("gorm:auto_preload", true).
		Find(&players).
		Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response{
			Status: false,
			Error:  "Error fetching data",
		})
	}
	c.JSON(http.StatusOK, response{
		Status: true,
		Data:   players,
	})
}
