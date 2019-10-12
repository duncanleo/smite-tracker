package web

import (
	"log"
	"net/http"

	"github.com/duncanleo/smite-tracker/db"
	"github.com/duncanleo/smite-tracker/model"
	"github.com/gin-gonic/gin"
)

func gods(c *gin.Context) {
	var gods []model.God
	err := db.DB.
		Set("gorm:auto_preload", true).
		Find(&gods).
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
		Data:   gods,
	})
}
