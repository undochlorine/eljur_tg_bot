package send

import (
	"bytes"
	"encoding/json"
	"net/http"
	"tgbot/internal/types"
)

func InlineKeyboard(botUrl string, chatId int, message string, kbd types.InlineKeyboardMarkup) error {
	var messageToSend types.BotMessageWithInlineKbd
	messageToSend.ChatId = chatId
	messageToSend.Text = message
	messageToSend.ReplyMarkupInline = kbd
	buf, err := json.Marshal(messageToSend)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}
