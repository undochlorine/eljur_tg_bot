package send

import (
	"bytes"
	"encoding/json"
	"net/http"
	"tgbot/internal/types"
)

func Message(botUrl string, chatId int, message string) error {
	var botMessage types.BotMessage
	botMessage.ChatId = chatId
	botMessage.Text = message
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}
