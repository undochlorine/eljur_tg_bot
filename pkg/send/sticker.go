package send

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"tgbot/internal/types"
	"tgbot/pkg/service/get"
)

func Sticker(botUrl string, update types.Update) error {
	var botSticker types.BotSticker
	botSticker.ChatId = update.Message.Chat.ChatId
	botSticker.Sticker = get.RandomSticker()
	buf, err := json.Marshal(botSticker)
	if err != nil {
		return err
	}
	fmt.Println(bytes.NewBuffer(buf))
	_, err = http.Post(botUrl+"/sendSticker", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	return nil
}
