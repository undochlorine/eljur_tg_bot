package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"fmt"
	"bytes"
	"strconv"
)

func main() {
	botAPI := "https://api.telegram.org/bot"
	botUrl := botAPI + BOT_TOKEN

	offset := 0

	for {
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			log.Println("Somethin' went wrong: ", err.Error())
		}
		for _, update := range updates {
			if update.Message.Sticker.FileId != "" {
				fmt.Println("received a sticker")
				err = sendSticker(botUrl, update)
			} else {
				fmt.Println("received a plain message")
				err = respond(botUrl, update)
			}
			offset = update.UpdateId + 1
			if err != nil {
				log.Println(err)
			}
		}
		upds, _ := json.Marshal(updates)
		fmt.Println(string( upds ))
	}
}

func getUpdates(botUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

func respond(botUrl string, update Update) (error) {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = update.Message.Text
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl + "/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

func sendSticker(botUrl string, update Update) (error) {
	var botSticker BotSticker
	botSticker.ChatId = update.Message.Chat.ChatId
	botSticker.Sticker = GetRandSticker()
	buf, err := json.Marshal(botSticker)
	if err != nil {
		return err
	}
	fmt.Println(bytes.NewBuffer(buf))
	_, err = http.Post(botUrl + "/sendSticker", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	return nil
}