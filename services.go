package main

import (
	"strconv"
	"errors"
	"strings"
)

func GenerateGridInt(amountOfBlocks int, columns int, callbackDataPrefix string) (InlineKeyboardMarkup) {
	var rows int
	var rest int = 0
	if amountOfBlocks%columns == 0 {
		rows = amountOfBlocks / columns
	} else {
		rows = (amountOfBlocks / columns) + 1
		rest = amountOfBlocks%columns
	}
	inlineKeyboard := make([][]InlineKeyboardButton, rows)
	for i := 0; i < rows; i++ {
		if i == rows-1 && rest != 0 {
			inlineKeyboard[i] = make([]InlineKeyboardButton, rest)
		} else {
			inlineKeyboard[i] = make([]InlineKeyboardButton, columns)
		}
		for j := 0; j < columns; j++ {
			if i == rows-1 && rest != 0 && j >= rest {
				break
			}
			inlineKeyboard[i][j] = InlineKeyboardButton{
				Text: strconv.Itoa(i*(columns)+j+1),
				CallbackData: callbackDataPrefix+strconv.Itoa(i*(columns)+j+1)+";",
			}
		}
	}
	kbd := InlineKeyboardMarkup{
		InlineKeyboard: inlineKeyboard,
	}
	return kbd
}

func GenerateGridIntervalString(interval [][]string, columns int, callbackDataPrefix string) (InlineKeyboardMarkup) {
	amountOfBlocks := len(interval)
	var rows int
	var rest int = 0
	if amountOfBlocks%columns == 0 {
		rows = amountOfBlocks / columns
	} else {
		rows = (amountOfBlocks / columns) + 1
		rest = amountOfBlocks%columns
	}
	inlineKeyboard := make([][]InlineKeyboardButton, rows)
	for i := 0; i < rows; i++ {
		if i == rows-1 && rest != 0 {
			inlineKeyboard[i] = make([]InlineKeyboardButton, rest)
		} else {
			inlineKeyboard[i] = make([]InlineKeyboardButton, columns)
		}
		for j := 0; j < columns; j++ {
			if i == rows-1 && rest != 0 && j >= rest {
				break
			}
			currentInterval := interval[i*columns + j][0] + "-" + interval[i*columns + j][1]
			inlineKeyboard[i][j] = InlineKeyboardButton{
				Text: currentInterval,
				CallbackData: callbackDataPrefix+currentInterval+";",
			}
		}
	}
	kbd := InlineKeyboardMarkup{
		InlineKeyboard: inlineKeyboard,
	}
	return kbd
}

func findInterval(tasksBlock [][]string, str string) (int, error) {
	targetStart, targetEnd := "", ""
	facedDash := false
	for _, c := range str {
		if c != '-' && !facedDash {
			targetStart += string(c)
		} else if c == '-' {
			facedDash = true
		} else {
			targetEnd += string(c)
		}
	}
	if !facedDash {
		return -1, errors.New("Invalid inputed string")
	}
	for i, block := range tasksBlock {
		if block[0] == targetStart && block[1] == targetEnd {
			return i, nil
		}
	}
	return -1, nil
}

func handleTaskQuery(botUrl string, chatId int, subject, deepData, blockWithPostfix string, block [][]string) (error) {
	if deepData == blockWithPostfix {
	    kbd := GenerateGridIntervalString(block, 3, "library;ceko;"+subject+blockWithPostfix)
	    err := SendInlineKeyboard(botUrl, chatId, "Номер упражнения:", kbd)
	    if err != nil {
	        return err
	    }
	} else {
	    deepData = strings.Replace(deepData, blockWithPostfix, "", 1)
	    taskIndex, findErr := findInterval(block, strings.Replace(deepData, ";", "", -1))
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
	        if taskIndex != len(block) - 1 {
	            urlBatchToSend = append(urlBatchToSend, block[taskIndex + 1][2])
	        }
	        err := SendMessage(botUrl, chatId, "Вот выбранное упражнение и соседние с ним:")
	        if err != nil {
	            return err
	        }
	        err = sendPhotosGroup(botUrl, chatId, urlBatchToSend)
	        if err != nil {
	            return err
	        }
	        kbd := GenerateGridIntervalString(block, 3, "library;ceko;"+subject+blockWithPostfix)
	        err = SendInlineKeyboard(botUrl, chatId, "Упражнение выполнено? Выберите следующее:", kbd)
	        if err != nil {
	            return err
	        }
	    }
	}
	return nil
}