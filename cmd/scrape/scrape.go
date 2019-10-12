package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/duncanleo/go-smite-api/smite"
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
	client := smite.Client{
		DevID:   os.Getenv("SMITE_DEV_ID"),
		AuthKey: os.Getenv("SMITE_AUTH_KEY"),
	}

	session, err := client.CreateSession()
	if err != nil {
		log.Panic(err)
	}

	// Create/Update all gods
	gods, err := client.GetGods(session.SessionID)
	if err != nil {
		log.Panic(err)
	}
	for _, god := range gods {
		var godModel = model.God{
			GameID:  god.ID,
			Name:    god.Name,
			CardURL: god.GodCardURL,
			IconURL: god.GodIconURL,
		}
		err = db.DB.
			FirstOrCreate(&godModel, model.God{GameID: god.ID}).
			Error
		if err != nil {
			log.Println("Error saving/creating god", err)
		} else {
			log.Printf("Created/Updated god '%s'\n", god.Name)
		}
	}

	for _, playerName := range strings.Split(os.Getenv("PLAYERS"), ",") {
		log.Printf("Processing player '%s'\n", playerName)
		playerIDResults, err := client.GetPlayerIDByName(session.SessionID, playerName)
		if err != nil {
			log.Println("Error fetching player ID using name", err)
			continue
		}
		if len(playerIDResults) < 1 {
			log.Println("Error fetching player ID using name")
			continue
		}
		playerResults, err := client.GetPlayer(session.SessionID, playerIDResults[0].PlayerID)
		if err != nil {
			log.Println("Error fetching player data using ID", err)
			continue
		}
		if len(playerResults) < 1 {
			log.Println("Error fetching player data using ID")
			continue
		}

		// Create player model
		var player model.Player
		err = db.DB.
			Where(model.Player{Name: playerName}).
			Assign(model.Player{
				Name:   playerName,
				GameID: playerIDResults[0].PlayerID,
				Status: playerResults[0].PersonalStatusMessage,
			}).
			FirstOrCreate(&player).
			Error

		if err != nil {
			log.Println(err)
			continue
		}

		// Get god ranks
		godRanks, err := client.GetPlayerGodRanks(session.SessionID, playerIDResults[0].PlayerID)
		if err != nil {
			log.Println("Error fetching player god ranks", err)
			continue
		}

		var tempGodWorshippers = 0
		var topGodData = godRanks[0]

		for _, godData := range godRanks {
			if godData.Worshippers > tempGodWorshippers {
				topGodData = godData
				tempGodWorshippers = godData.Worshippers
			}
		}

		topGodGameID, _ := strconv.Atoi(topGodData.GodID)

		var topGodModel model.God
		err = db.DB.
			Where(model.God{GameID: topGodGameID}).
			Find(&topGodModel).
			Error

		// Create player data
		var playerData = model.PlayerData{
			Player:            player,
			HoursPlayed:       playerResults[0].HoursPlayed,
			Leaves:            playerResults[0].Leaves,
			Level:             playerResults[0].Level,
			Losses:            playerResults[0].Losses,
			MasteryLevel:      playerResults[0].MasteryLevel,
			Wins:              playerResults[0].Wins,
			TotalAchievements: playerResults[0].TotalAchievements,
			TotalWorshippers:  playerResults[0].TotalWorshippers,
			TopGod:            topGodModel,
			TopGodRank:        topGodData.Rank,
			TopGodWorshippers: topGodData.Worshippers,
		}

		err = db.DB.Create(&playerData).Error
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
