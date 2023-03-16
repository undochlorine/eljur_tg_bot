package main

import (
	"encoding/json"
	"log"
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Bot has been launched")

	botAPI := "https://api.telegram.org/bot"
	botUrl := botAPI + BOT_TOKEN

	commands := BotCommands{
		Commands: []BotCommand{
			{Command: "/start", Description: "Simplify your life with our awesome bot"},
			{Command: "/library", Description: "Get list of resources to study"},
			{Command: "/feedback", Description: "Ask or report here"},
		},
	}

	err := SetCommands(botUrl, commands)
	if err != nil {
		fmt.Println("Failed to set bot's commands:")
		log.Println(err)
	}

	offset := 0

	for {
		updates, err := GetUpdates(botUrl, offset)
		if err != nil {
			log.Println("Somethin' went wrong: ", err.Error())
		}
		for _, update := range updates {
			if update.Message.Sticker.FileId != "" {
				fmt.Println("received a sticker")
				err = SendSticker(botUrl, update)
			} else if len(update.Message.Entities) > 0 {
				fmt.Println("received a command")
				if update.Message.Entities[0].Type == "bot_command" {
					switch update.Message.Text[1:] {
						case "start": {
							err = SendMessage(botUrl, update.Message.Chat.ChatId, "Welcome!")
							if err != nil {
								log.Println(err)
							}
						}
						case "library": {
							kbd := InlineKeyboardMarkup{
								InlineKeyboard: [][]InlineKeyboardButton{
									{
										InlineKeyboardButton{Text: "ЦЭКО", CallbackData: "library;ceko;"},
									},
								},
							}
							err = SendInlineKeyboard(botUrl, update.Message.Chat.ChatId, "Доступные ресурсы:", kbd)
							if err != nil {
								log.Println(err)
							}
						}
						case "feedback": {
							err = SendMessage(botUrl, update.Message.Chat.ChatId, "Here is contact of the person who is willing to help you:\nhttps://t.me/undochlorine")
							if err != nil {
								log.Println(err)
							}
						}
						default: {
							err = SendMessage(botUrl, update.Message.Chat.ChatId, "I'm not aware about this command, try later)")
							if err != nil {
								log.Println(err)
							}
						}
					}
				}
			} else if update.CallbackQuery.Id != "" {
				fmt.Println("received a callback query")
				data := update.CallbackQuery.Data
				switch {
					case data[0:len("library;ceko;")] == "library;ceko;": {
						if data == "library;ceko;" {
							kbd := InlineKeyboardMarkup{
								InlineKeyboard: [][]InlineKeyboardButton{
									{
										InlineKeyboardButton{Text: "Алгебра", CallbackData: "library;ceko;math;"},
									},
								},
							}
							err = SendInlineKeyboard(botUrl, update.CallbackQuery.Message.Chat.ChatId, "Предметы:", kbd)
							if err != nil {
								log.Println(err)
							}
						} else {
							deepData := strings.Replace(data, "library;ceko;", "", -1)
							switch {
								case deepData[0:len("math;")] == "math;": {
									if deepData == "math;" {
										kbd := GenerateGridInt(21, 4, "library;ceko;math;")
										err = SendInlineKeyboard(botUrl, update.CallbackQuery.Message.Chat.ChatId, "Задания:", kbd)
										if err != nil {
											log.Println(err)
										}
									} else {
										deepData = strings.Replace(deepData, "math;", "", -1)
										switch {
											case deepData[0:len("1;")] == "1;": {
												if deepData == "1;" {
													kbd := GenerateGridIntervalString(TasksOfBlock1, 3, "library;ceko;math;1;")
													err = SendInlineKeyboard(botUrl, update.CallbackQuery.Message.Chat.ChatId, "Номер упражнения:", kbd)
													if err != nil {
														log.Println(err)
													}
												} else {
													fmt.Println("^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*^*")
												}
												
											}
											case deepData[0:len("2;")] == "2;": {

											}
											case deepData[0:len("3;")] == "3;": {

											}
											case deepData[0:len("4;")] == "4;": {

											}
											case deepData[0:len("5;")] == "5;": {

											}
											case deepData[0:len("6;")] == "6;": {

											}
											case deepData[0:len("7;")] == "7;": {

											}
											case deepData[0:len("8;")] == "8;": {

											}
											case deepData[0:len("9;")] == "9;": {

											}
											case deepData[0:len("10;")] == "10;": {

											}
											case deepData[0:len("11;")] == "11;": {

											}
											case deepData[0:len("12;")] == "12;": {

											}
											case deepData[0:len("13;")] == "13;": {

											}
											case deepData[0:len("14;")] == "14;": {

											}
											case deepData[0:len("15;")] == "15;": {

											}
											case deepData[0:len("16;")] == "16;": {

											}
											case deepData[0:len("17;")] == "17;": {

											}
											case deepData[0:len("18;")] == "18;": {

											}
											case deepData[0:len("19;")] == "19;": {

											}
											case deepData[0:len("20;")] == "20;": {

											}
											case deepData[0:len("21;")] == "21;": {

											}
											default:
												fmt.Println("Unknown callback query:")
										}
									}
								}
								default:
									fmt.Println("Unknown callback query:")
							}
						}
					}
					default:
						fmt.Println("Unknown callback query:")
				}
			} else {
				fmt.Println("received a plain message")
				err = SendMessage(botUrl, update.Message.Chat.ChatId, "...?")
			}
			if err != nil {
				log.Println(err)
			}
			offset = update.UpdateId + 1
			upd, _ := json.MarshalIndent(update, "", "-- ")
			fmt.Println("---------------------------")
			fmt.Println(string( upd ))
		}
	}
}