package handle

import (
	"errors"
	"strings"
	"tgbot/internal/types"
	"tgbot/pkg/send"
	"tgbot/pkg/service/find"
	"tgbot/pkg/service/generate"
)

func TaskQuery(botUrl string, chatId int, subject, deepData, blockWithPostfix string, block [][]string) error {
	if deepData == blockWithPostfix {
		kbd := generate.GridIntervalString(block, 3, "library;ceko;"+subject+blockWithPostfix)
		kbd.InlineKeyboard = append(
			kbd.InlineKeyboard,
			[]types.InlineKeyboardButton{
				{
					Text:         "⏪ Сменить задание",
					CallbackData: "library;ceko;" + subject,
				},
			},
		)
		err := send.InlineKeyboard(botUrl, chatId, "Номер упражнения:", kbd)
		if err != nil {
			return err
		}
	} else {
		deepData = strings.Replace(deepData, blockWithPostfix, "", 1)
		taskIndex, findErr := find.Interval(block, strings.Replace(deepData, ";", "", -1))
		if findErr != nil {
			return findErr
		} else if taskIndex == -1 {
			return errors.New("Unknown callback query of task")
		} else {
			var urlBatchToSend = []string{
				block[taskIndex][2],
			}
			if taskIndex != 0 {
				urlBatchToSend = append([]string{
					block[taskIndex-1][2],
				}, block[taskIndex][2])
			}
			if taskIndex != len(block)-1 {
				urlBatchToSend = append(urlBatchToSend, block[taskIndex+1][2])
			}
			_, err := send.Message(botUrl, chatId, "Вот выбранное упражнение и соседние с ним:")
			if err != nil {
				return err
			}
			err = send.PhotosGroup(botUrl, chatId, urlBatchToSend)
			if err != nil {
				return err
			}
			kbd := generate.GridIntervalString(block, 3, "library;ceko;"+subject+blockWithPostfix)
			kbd.InlineKeyboard = append(
				kbd.InlineKeyboard,
				[]types.InlineKeyboardButton{
					{
						Text:         "⏪ Сменить задание",
						CallbackData: "library;ceko;" + subject,
					},
				},
			)
			err = send.InlineKeyboard(botUrl, chatId, "Упражнение выполнено? Выберите следующее:", kbd)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
