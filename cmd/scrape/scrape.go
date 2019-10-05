package main

import (
	"log"
	"os"
	"strings"

	"github.com/duncanleo/brawl-scraper/brawl"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	brawlers := strings.Split(os.Getenv("BRAWLERS"), ",")
	log.Printf("Processing %d brawlers\n", len(brawlers))

	brawlClient := brawl.NewClient(os.Getenv("BRAWL_API_TOKEN"))

	for _, brawlerTag := range brawlers {
		x, err := brawlClient.GetPlayerData(brawlerTag)
		if err != nil {
			log.Println(err)
		}
		log.Printf("%+v\n", x)
	}
}
