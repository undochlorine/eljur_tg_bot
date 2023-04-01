package generate

import (
	"strconv"
	"tgbot/internal/types"
)

func GridInt(amountOfBlocks int, columns int, callbackDataPrefix string) types.InlineKeyboardMarkup {
	var rows int
	var rest int = 0
	if amountOfBlocks%columns == 0 {
		rows = amountOfBlocks / columns
	} else {
		rows = (amountOfBlocks / columns) + 1
		rest = amountOfBlocks % columns
	}
	inlineKeyboard := make([][]types.InlineKeyboardButton, rows)
	for i := 0; i < rows; i++ {
		if i == rows-1 && rest != 0 {
			inlineKeyboard[i] = make([]types.InlineKeyboardButton, rest)
		} else {
			inlineKeyboard[i] = make([]types.InlineKeyboardButton, columns)
		}
		for j := 0; j < columns; j++ {
			if i == rows-1 && rest != 0 && j >= rest {
				break
			}
			inlineKeyboard[i][j] = types.InlineKeyboardButton{
				Text:         strconv.Itoa(i*(columns) + j + 1),
				CallbackData: callbackDataPrefix + strconv.Itoa(i*(columns)+j+1) + ";",
			}
		}
	}
	callbackDataToGetBack := callbackDataPrefix[:len(callbackDataPrefix)-1]
	for {
		if callbackDataToGetBack[len(callbackDataToGetBack)-1] == ';' {
			break
		}
		callbackDataToGetBack = callbackDataToGetBack[:len(callbackDataToGetBack)-1]
	}
	inlineKeyboard = append(inlineKeyboard, []types.InlineKeyboardButton{
		{
			Text:         "⏪ Сменить предмет",
			CallbackData: callbackDataToGetBack,
		},
	})
	kbd := types.InlineKeyboardMarkup{
		InlineKeyboard: inlineKeyboard,
	}
	return kbd
}

func GridIntervalString(interval [][]string, columns int, callbackDataPrefix string) types.InlineKeyboardMarkup {
	amountOfBlocks := len(interval)
	var rows int
	var rest int = 0
	if amountOfBlocks%columns == 0 {
		rows = amountOfBlocks / columns
	} else {
		rows = (amountOfBlocks / columns) + 1
		rest = amountOfBlocks % columns
	}
	inlineKeyboard := make([][]types.InlineKeyboardButton, rows)
	for i := 0; i < rows; i++ {
		if i == rows-1 && rest != 0 {
			inlineKeyboard[i] = make([]types.InlineKeyboardButton, rest)
		} else {
			inlineKeyboard[i] = make([]types.InlineKeyboardButton, columns)
		}
		for j := 0; j < columns; j++ {
			if i == rows-1 && rest != 0 && j >= rest {
				break
			}
			currentInterval := interval[i*columns+j][0] + "-" + interval[i*columns+j][1]
			inlineKeyboard[i][j] = types.InlineKeyboardButton{
				Text:         currentInterval,
				CallbackData: callbackDataPrefix + currentInterval + ";",
			}
		}
	}
	kbd := types.InlineKeyboardMarkup{
		InlineKeyboard: inlineKeyboard,
	}
	return kbd
}
