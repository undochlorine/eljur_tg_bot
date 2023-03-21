package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"bytes"
)

func SendMessage(botUrl string, chatId int, message string) (error) {
	var botMessage BotMessage
	botMessage.ChatId = chatId
	botMessage.Text = message
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

func SendInlineKeyboard(botUrl string, chatId int, message string, kbd InlineKeyboardMarkup) (error) {
	var messageToSend BotMessageWithInlineKbd
	messageToSend.ChatId = chatId
	messageToSend.Text = message
	messageToSend.ReplyMarkupInline = kbd
	buf, err := json.Marshal(messageToSend)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl + "/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

func SendSticker(botUrl string, update Update) (error) {
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

func sendPhoto(botUrl string, chatId int, url string) (error) {
	var botPhoto BotPhoto
	botPhoto.ChatId = chatId
	botPhoto.Photo = url
	buf, err := json.Marshal(botPhoto)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl + "/sendPhoto", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

func sendPhotosGroup(botUrl string, chatId int, urls []string) (error) {
	type InputMediaPhoto struct {
		Type 	string		 	`json:"type"`
		Media 	string 			`json:"media"`
	}
	
	type BotPhotos struct {
		ChatId	int 					`json:"chat_id"`
		Media 	[]InputMediaPhoto 		`json:"media"`
	}
	var botPhotos BotPhotos
	botPhotos.ChatId = chatId
	for _, u := range urls {
		botPhotos.Media = append(botPhotos.Media, InputMediaPhoto{
			Type: "photo",
			Media: u,
		})
	}
	buf, err := json.MarshalIndent(botPhotos, "", " ")
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl + "/sendMediaGroup", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}