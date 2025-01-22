package tags

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"DeleteBotBlocked/internal/config"

	"github.com/fatih/color"
)

type RequestTag struct {
	TagName string `json:"tagName"`
}

type RequestCancelConversation struct {
	Status  string `json:"status"`
	AgentID string `json:"agentId"`
}

func AddConversationTag(id string, globalKey string) {
	url := "https://dm9epl.api.infobip.com/ccaas/1/conversations/" + id + "/tags"

	requestTag := RequestTag{
		TagName: config.GlobalTag,
	}

	payload, err := json.Marshal(requestTag)
	if err != nil {
		log.Println(" — "+color.RedString("WARNING")+" >>> "+"Error converting to JSON: ", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", globalKey)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	CancelConversation(id, requestTag.TagName, globalKey)
}

func CancelConversation(id string, requestTag string, globalKey string) {
	requestCancelConversation := RequestCancelConversation{
		Status:  "CLOSED",
		AgentID: config.GlobalAgentId,
	}
	payload, err := json.Marshal(requestCancelConversation)
	if err != nil {
		log.Println(" — "+color.RedString("WARNING")+" >>> "+"Unmarshalling error: ", err)
	}
	url := "https://dm9epl.api.infobip.com/ccaas/1/conversations/" + id
	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payload))

	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", globalKey)

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()
	log.Println(" — " + color.GreenString("CLOSED") + " >>> " + "CONVERSATION: " + id + " | TAG: " + requestTag)
}
