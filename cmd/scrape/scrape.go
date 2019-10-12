package main

import (
	"os"

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

}
