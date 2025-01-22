package main

import (
	"log"
	"strconv"
	"time"

	"DeleteBotBlocked/internal/config"
	"DeleteBotBlocked/internal/conversations"

	"github.com/fatih/color"
)

func main() {
	log.Println(" — " + color.MagentaString("<<< SCRIPT RUN >>>"))
	config.LoadConfig()

	interval, err := strconv.Atoi(config.GlobalInterval)
	if err != nil {
		log.Println(" — "+color.RedString("WARNING")+" >>> "+"Error converting string to number: ", err)
	}

	for range time.Tick(time.Second * time.Duration(interval)) {
		conversations.GETConversations(config.GlobalKey)
	}
}
