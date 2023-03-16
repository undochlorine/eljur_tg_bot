package main

import (
	"strconv"
	"fmt"
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
			fmt.Println(i*columns + j)
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