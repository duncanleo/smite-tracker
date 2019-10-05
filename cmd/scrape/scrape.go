package main

import (
	"log"
	"os"
	"strings"

	"github.com/duncanleo/brawl-scraper/brawl"
	"github.com/duncanleo/brawl-scraper/db"
	"github.com/duncanleo/brawl-scraper/model"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	brawlers := strings.Split(os.Getenv("BRAWLERS"), ",")
	log.Printf("Processing %d brawlers\n", len(brawlers))

	brawlClient := brawl.NewClient(os.Getenv("BRAWL_API_TOKEN"))

	for _, brawlerTag := range brawlers {
		playerStats, err := brawlClient.GetPlayerData(brawlerTag)
		if err != nil {
			log.Println(err)
			continue
		}
		// Upsert player object
		var player model.Player
		err = db.DB.
			Where(model.Player{Tag: brawlerTag}).
			Assign(model.Player{
				Name:      playerStats.Name,
				NameColor: playerStats.NameColor,
				Tag:       brawlerTag,
			}).
			FirstOrCreate(&player).
			Error
		if err != nil {
			log.Println(err)
			continue
		}
		// Create player data
		var playerData = model.PlayerData{
			Player:               player,
			TrophyCount:          playerStats.Trophies,
			ExpLevel:             playerStats.ExpLevel,
			ExpPoints:            playerStats.ExpPoints,
			ThreeV3Victories:     playerStats.ThreeVs3Victories,
			SoloVictories:        playerStats.SoloVictories,
			DuoVictories:         playerStats.DuoVictories,
			BestRoboRumbleTime:   playerStats.BestRoboRumbleTime,
			BestTimeAsBigBrawler: playerStats.BestTimeAsBigBrawler,
		}
		err = db.DB.Create(&playerData).Error
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
