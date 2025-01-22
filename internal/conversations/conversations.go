package conversations

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"DeleteBotBlocked/internal/config"
	"DeleteBotBlocked/internal/messages"
	"DeleteBotBlocked/internal/models"

	"github.com/fatih/color"
)

type Data struct {
	Conversations []models.Conversation `json:"conversations"`
	Pagination    models.Pagination     `json:"pagination"`
}

func GETConversations(globalKey string) {
	url := config.GlobalURL + "?status=OPEN&limit=999"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", globalKey)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	var data Data
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Println(" — "+color.RedString("WARNING")+" >>> "+"Unmarshalling error: ", err)
	}

	for _, conv := range data.Conversations {
		messages.GETConversationsMessages(conv.ID, globalKey)
	}
}
