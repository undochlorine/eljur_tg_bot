package send

import (
	"bytes"
	"encoding/json"
	"net/http"
	"tgbot/internal/types"
)

func Message(botUrl string, chatId int, message string) (int, error) {
	var botMessage types.BotMessage
	botMessage.ChatId = chatId
	botMessage.Text = message
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return 0, err
	}
	resp, err := http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var responseMap map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseMap)
	if err != nil {
		return 0, err
	}
	messageId := int(responseMap["result"].(map[string]interface{})["message_id"].(float64))

	return messageId, nil
}
