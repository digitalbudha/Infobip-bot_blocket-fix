package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var GlobalKey string = ""
var GlobalAgentId string = ""
var GlobalTag string = ""
var GlobalInterval string = ""

type Conversation struct {
	ID           string     `json:"id"`
	Topic        string     `json:"topic"`
	Summary      string     `json:"summary"`
	Status       string     `json:"status"`
	Priority     string     `json:"priority"`
	QueueID      string     `json:"queueId"`
	AgentID      string     `json:"agentId"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	ClosedAt     *time.Time `json:"closedAt,omitempty"`
	PendingSince time.Time  `json:"pendingSince"`
	FormID       string     `json:"formId"`
}

type Pagination struct {
	TotalItems int    `json:"totalItems"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	OrderBy    string `json:"orderBy"`
}

type Data struct {
	Conversations []Conversation `json:"conversations"`
	Pagination    Pagination     `json:"pagination"`
}

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
	Messages   []Message  `json:"messages"`
	Pagination Pagination `json:"pagination"`
}

type RequestTag struct {
	TagName string `json:"tagName"`
}

type RequestCancelConversation struct {
	Status  string `json:"status"`
	AgentID string `json:"agentId"`
}

func GETConversations() {
	url := "https://dm9epl.api.infobip.com/ccaas/1/conversations?status=OPEN&limit=999"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
	}
	req.Header.Add("Authorization", GlobalKey)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	// Распаковка JSON-данных в структуру Data
	var data Data
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Println("Ошибка демаршалирования:", err)
	}

	for _, conv := range data.Conversations {
		GETConversationsMessages(conv.ID)
	}
}

func GETConversationsMessages(id string) {
	url := "https://dm9epl.api.infobip.com/ccaas/1/conversations/" + id + "/messages?limit=2"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Authorization", GlobalKey)

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
		log.Println("Ошибка демаршалирования:", err)
		return
	}

	for _, msg := range data.Messages {
		if len(data.Messages) == 1 {
			if msg.SingleSendMessage.Content.Type == "BOT_BLOCKED" {
				AddConversationTag(id)
			}
		}
	}
}

func AddConversationTag(id string) {
	url := "https://dm9epl.api.infobip.com/ccaas/1/conversations/" + id + "/tags"

	requestTag := RequestTag{
		TagName: GlobalTag,
	}

	payload, err := json.Marshal(requestTag)
	if err != nil {
		log.Println("Ошибка при преобразовании в JSON:", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", GlobalKey)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	CancelConversation(id, requestTag.TagName)
}

func CancelConversation(id string, requestTag string) {
	requestCancelConversation := RequestCancelConversation{
		Status:  "CLOSED",
		AgentID: GlobalAgentId,
	}
	payload, err := json.Marshal(requestCancelConversation)
	if err != nil {
		log.Println("Error marshalling: ", err)
	}
	url := "https://dm9epl.api.infobip.com/ccaas/1/conversations/" + id
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))

	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", GlobalKey)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	log.Printf("Закрыл: %s c тегом: %s\n", id, requestTag)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка при загрузке файла .env:", err)
	}
	GlobalKey = os.Getenv("KEY")
	GlobalAgentId = os.Getenv("AGENT_ID")
	GlobalTag = os.Getenv("TAG")
	GlobalInterval = os.Getenv("INTERVAL")
	interval, err := strconv.Atoi(GlobalInterval)
	if err != nil {
		log.Println("Ошибка в преобразовании строки в число:", err)
	}
	for range time.Tick(time.Second * time.Duration(interval)) {
		GETConversations()
	}
}
