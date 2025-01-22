package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var GlobalKey string
var GlobalAgentId string
var GlobalTag string
var GlobalInterval string
var GlobalURL string

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	GlobalKey = os.Getenv("KEY")
	GlobalAgentId = os.Getenv("AGENT_ID")
	GlobalTag = os.Getenv("TAG")
	GlobalInterval = os.Getenv("INTERVAL")
	GlobalURL = os.Getenv("URL")
}
