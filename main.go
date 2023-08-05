package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
<<<<<<< HEAD
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
=======
	"strings"
>>>>>>> fdfc0ad737b8413bd91d13335b59fbd93680ced1
)

var roman = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
	"XL":   40,
	"L":    50,
	"XC":   90,
	"C":    100,
}

func readInputFromConsole() interface{} {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите ваше выражение: ")
	scanner.Scan()
	input := scanner.Text()
	substrings := strings.Split(input, " ")
	if numberOne, err := strconv.Atoi(substrings[0]); err == nil {
		if numberTwo, err := strconv.Atoi(substrings[2]); err == nil {
			switch substrings[1] {
			case "+":
				return numberOne + numberTwo
			case "-":
				return numberOne - numberTwo
			case "*":
				return numberOne * numberTwo
			case "/":
				if numberTwo == 0 {
					fmt.Println("Сорьки, на ноль не умею делить")
					return 0
				}
				return numberOne / numberTwo
			default:
				return "Не понял что за оператор для расчёта"
			}
		}
		return "Второе число ты ввел не арабскими цифрами, сори пока"
	}
<<<<<<< HEAD
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
	log.Println(" — " + color.RedString("WARNING") + " >>> " + "Ошибка демаршалирования: ", err)
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
	log.Println(" — " + color.RedString("WARNING") + " >>> " + "Ошибка демаршалирования: ", err)
		return
	}

	for _, msg := range data.Messages {
		if len(data.Messages) == 1 {
			if msg.SingleSendMessage.Content.Type == "BOT_BLOCKED" || msg.SingleSendMessage.Content.Type == "BOT_UNBLOCKED" {
				AddConversationTag(id)
=======
	result := romanCalculate(input)
	if val, ok := result.(int); ok {
		return arabicToRoman(val)
	} else {
		return result
	}
}

func romanCalculate(input string) interface{} {
	substrings := strings.Split(input, " ")
	numberOne := roman[substrings[0]]
	operator := substrings[1]
	numberTwo := roman[substrings[2]]
	if _, err := strconv.Atoi(substrings[0]); err != nil {
		if _, err := strconv.Atoi(substrings[2]); err != nil {
			switch operator {
			case "+":
				return numberOne + numberTwo
			case "-":
				if numberOne < numberTwo {
					return "Вывод ошибки, так как в римской системе нет отрицательных чисел."
				} else {
					return numberOne - numberTwo
				}
			case "*":
				return numberOne * numberTwo
			case "/":
				if numberOne < numberTwo {
					return "Вывод ошибки, так как в римской системе нет значения ноль."
				} else {
					return numberOne / numberTwo
				}
			default:
				return "Не понял что ты от меня хочешь, я такие данные не умею считать"
>>>>>>> fdfc0ad737b8413bd91d13335b59fbd93680ced1
			}
		}
	}
	return "Ошибка: оба числа должны быть в римской системе"
}

func arabicToRoman(arabic int) string {
	romanMap := []struct {
		arabic int
		roman  string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var roman string

	for _, mapping := range romanMap {
		for arabic >= mapping.arabic {
			roman += mapping.roman
			arabic -= mapping.arabic
		}
	}
<<<<<<< HEAD

	payload, err := json.Marshal(requestTag)
	if err != nil {
	log.Println(" — " + color.RedString("WARNING") + " >>> " + "Ошибка при преобразовании в JSON: ", err)
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
	log.Println(" — " + color.RedString("WARNING") + " >>> " + "Ошибка демаршалирования: ", err)
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
	log.Println(" — " + color.GreenString("CLOSED") + " >>> "+"CONVERSATION: "+ id +" | TAG: " + requestTag)
}

func main() {
	log.Println(" — " + color.MagentaString("<<< SCRIPT RUN >>>"))
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
		log.Println(" — " + color.RedString("WARNING") + " >>> " + "Ошибка преобразования строки в число: ", err)
	}
	
	for range time.Tick(time.Second * time.Duration(interval)) {
		GETConversations()
	}
=======
	return roman
}

func main() {
	fmt.Println("Привет, я калькулятор, давай посчитаем что нибудь!")
	input := readInputFromConsole()
	fmt.Println(input)
>>>>>>> fdfc0ad737b8413bd91d13335b59fbd93680ced1
}
