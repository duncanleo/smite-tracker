package main

import (
	"log"
	"os"
	"strings"

	"github.com/duncanleo/smite-tracker/brawl"
	"github.com/duncanleo/smite-tracker/db"
	"github.com/duncanleo/smite-tracker/model"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	// Fixie support
	if len(os.Getenv("FIXIE_URL")) > 0 {
		os.Setenv("HTTP_PROXY", os.Getenv("FIXIE_URL"))
		os.Setenv("HTTPS_PROXY", os.Getenv("FIXIE_URL"))
	}
}

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
		// Create brawlers' data
		var topBrawlerIndex = 0
		var tempBrawlerPower = 0
		var tempBrawlerRank = 0
		var tempBrawlerTrophies = 0
		var tempBrawlerStarPowerCount = 0
		for i, b := range playerStats.Brawlers {
			if b.Trophies > tempBrawlerTrophies &&
				b.Rank >= tempBrawlerRank &&
				b.Power >= tempBrawlerPower &&
				len(b.StarPowers) >= tempBrawlerStarPowerCount {
				topBrawlerIndex = i
				tempBrawlerPower = b.Power
				tempBrawlerRank = b.Rank
				tempBrawlerTrophies = b.Trophies
				tempBrawlerStarPowerCount = len(b.StarPowers)
			}
			var brawler model.Brawler
			err = db.DB.
				Where(model.Brawler{GameID: b.ID}).
				Assign(model.Brawler{
					GameID: b.ID,
					Name:   b.Name,
				}).
				FirstOrCreate(&brawler).
				Error
			if err != nil {
				log.Println(err)
				continue
			}
		}

		var topBrawler model.Brawler
		db.DB.First(&topBrawler, &model.Brawler{GameID: playerStats.Brawlers[topBrawlerIndex].ID})

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
			TopBrawler:           topBrawler,
		}
		err = db.DB.Create(&playerData).Error
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
