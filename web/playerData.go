package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/duncanleo/smite-tracker/db"
	"github.com/duncanleo/smite-tracker/model"
	"github.com/gin-gonic/gin"
)

func playerDatas(c *gin.Context) {
	var playerDatas []model.PlayerData
	idStr := c.Query("player_id")
	if len(idStr) == 0 {
		c.JSON(http.StatusBadRequest, response{
			Status: false,
			Error:  "missing parameters",
		})
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, response{
			Status: false,
			Error:  "bad parameters",
		})
		return
	}

	err = db.DB.
		Set("gorm:auto_preload", true).
		Find(&playerDatas, model.PlayerData{PlayerID: uint(id)}).
		Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, response{
			Status: false,
			Error:  "Error fetching data",
		})
		return
	}
	c.JSON(http.StatusOK, response{
		Status: true,
		Data:   playerDatas,
	})
}
