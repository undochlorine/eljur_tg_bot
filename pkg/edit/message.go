package edit

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func MessageText(botUrl string, chatId int, messageId int, newText string) error {
	var messageToEdit struct {
		ChatId    int    `json:"chat_id"`
		MessageId int    `json:"message_id"`
		Text      string `json:"text"`
	}
	messageToEdit.ChatId = chatId
	messageToEdit.MessageId = messageId
	messageToEdit.Text = newText
	buf, err := json.Marshal(messageToEdit)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/editMessageText", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}
