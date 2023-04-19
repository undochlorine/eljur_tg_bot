package send

import (
	"log"
	"tgbot/pkg/edit"
)

func WaitMock(botUrl string, chatId int, timeout chan struct{}) {
	msgId, err := Message(botUrl, chatId, "Ожидайте:")
	if err != nil {
		return
	}
	select {
	case <-timeout:
		err = edit.MessageText(botUrl, chatId, msgId, "Готово! Файл поврежден:")
		if err != nil {
			log.Println(err)
		}
		return
	}
}
