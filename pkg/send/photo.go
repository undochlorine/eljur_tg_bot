package send

import (
	"bytes"
	"encoding/json"
	"net/http"

	"tgbot/internal/types"
)

func Photo(botUrl string, chatId int, url string) error {
	var botPhoto types.BotPhoto
	botPhoto.ChatId = chatId
	botPhoto.Photo = url
	buf, err := json.Marshal(botPhoto)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/sendPhoto", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}

func PhotosGroup(botUrl string, chatId int, urls []string) error {
	type InputMediaPhoto struct {
		Type  string `json:"type"`
		Media string `json:"media"`
	}

	type BotPhotos struct {
		ChatId int               `json:"chat_id"`
		Media  []InputMediaPhoto `json:"media"`
	}
	var botPhotos BotPhotos
	botPhotos.ChatId = chatId
	for _, u := range urls {
		botPhotos.Media = append(botPhotos.Media, InputMediaPhoto{
			Type:  "photo",
			Media: u,
		})
	}
	buf, err := json.MarshalIndent(botPhotos, "", " ")
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/sendMediaGroup", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}
