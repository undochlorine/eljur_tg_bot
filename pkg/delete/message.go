package delete

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Message(botUrl string, chatId int, messageId int) error {
	var messageToDelete struct {
		ChatId    int `json:"chat_id"`
		MessageId int `json:"message_id"`
	}
	messageToDelete.ChatId = chatId
	messageToDelete.MessageId = messageId
	buf, err := json.Marshal(messageToDelete)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/deleteMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}
