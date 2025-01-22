package messages

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"DeleteBotBlocked/internal/config"
	"DeleteBotBlocked/internal/models"
	"DeleteBotBlocked/internal/tags"

	"github.com/fatih/color"
)

type Content struct {
	Text           string      `json:"text"`
	ShowUrlPreview interface{} `json:"showUrlPreview"`
}

type ContentType struct {
	Type string `json:"type"`
}

type SingleSendMessage struct {
	From struct {
		BotId string `json:"botId"`
		Type  string `json:"type"`
	} `json:"from"`
	To struct {
		ChatId string `json:"chatId"`
		Type   string `json:"type"`
	} `json:"to"`
	Content     ContentType `json:"content"`
	ParseMode   interface{} `json:"parseMode"`
	ReplyMarkup interface{} `json:"replyMarkup"`
	Channel     string      `json:"channel"`
	Direction   string      `json:"direction"`
}

type Message struct {
	ID                string            `json:"id"`
	Channel           string            `json:"channel"`
	From              string            `json:"from"`
	To                string            `json:"to"`
	Direction         string            `json:"direction"`
	ConversationID    string            `json:"conversationId"`
	CreatedAt         time.Time         `json:"createdAt"`
	UpdatedAt         time.Time         `json:"updatedAt"`
	Content           Content           `json:"content"`
	SingleSendMessage SingleSendMessage `json:"singleSendMessage"`
	ContentType       string            `json:"contentType"`
}

type ConversationData struct {
	Messages   []Message         `json:"messages"`
	Pagination models.Pagination `json:"pagination"`
}

func GETConversationsMessages(id string, globalKey string) {
	url := config.GlobalURL + "/" + id + "/messages?limit=2"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Authorization", globalKey)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}

	var data ConversationData
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Println(" — "+color.RedString("WARNING")+" >>> "+"Unmarshalling error: ", err)
		return
	}

	for _, msg := range data.Messages {
		if len(data.Messages) == 1 {
			if msg.SingleSendMessage.Content.Type == "BOT_BLOCKED" || msg.SingleSendMessage.Content.Type == "BOT_UNBLOCKED" {
				tags.AddConversationTag(id, globalKey)
			}
		}
	}
}
