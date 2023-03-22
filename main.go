package main

import (
	"encoding/json"
	"log"
	"fmt"
	"strings"
	"strconv"
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

			upd, _ := json.MarshalIndent(update, "", "-- ")
			fmt.Println(string( upd ))
			fmt.Println("---------------------------")

			chatId := update.CallbackQuery.Message.Chat.ChatId
			if update.Message.Sticker.FileId != "" {
				fmt.Println("received a sticker")
				err = SendSticker(botUrl, update)
				if err != nil {
					log.Println(err)
				}
			} else if len(update.Message.Entities) > 0 {
				if update.Message.Entities[0].Type == "bot_command" {
					fmt.Println("received a command")
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
					case data[:len("library;")] == "library;": {
						if data == "library;" {
							kbd := InlineKeyboardMarkup{
								InlineKeyboard: [][]InlineKeyboardButton{
									{
										InlineKeyboardButton{Text: "ЦЭКО", CallbackData: "library;ceko;"},
									},
								},
							}
							err = SendInlineKeyboard(botUrl, chatId, "Доступные ресурсы:", kbd)
							if err != nil {
								log.Println(err)
							}
						} else {
							deepData := strings.Replace(data, "library;", "", -1)
							switch {
								case deepData[0:len("ceko;")] == "ceko;": {
									if deepData == "ceko;" {
										kbd := InlineKeyboardMarkup{
											InlineKeyboard: [][]InlineKeyboardButton{
												{
													InlineKeyboardButton{Text: "Математика", CallbackData: "library;ceko;math;"},
												},
												{
													InlineKeyboardButton{Text: "⏪ Сменить ресурс", CallbackData: "library;"},

												},
											},
										}
										err = SendInlineKeyboard(botUrl, chatId, "Выберете предмет:", kbd)
										if err != nil {
											log.Println(err)
										}
									} else {
										deepData = strings.Replace(deepData, "ceko;", "", -1)
										switch {
											case deepData[0:len("math;")] == "math;": {
												if deepData == "math;" {
													kbd := GenerateGridInt(21, 4, "library;ceko;math;")
													err = SendInlineKeyboard(botUrl, chatId, "Номера доступных заданий:", kbd)
													if err != nil {
														log.Println(err)
													}
												} else {
													deepData = strings.Replace(deepData, "math;", "", -1)
													requestingBlock := ""
													for {
														if deepData[len(requestingBlock)] == ';' {
															break
														} else {
															requestingBlock += string(deepData[len(requestingBlock)])
														}
													}
													requestingBlockIndex, err := strconv.Atoi(requestingBlock)
													requestingBlockIndex--
													if err != nil {
														fmt.Sprintf("wrong test req: %v", err)
													} else {
														err = handleTaskQuery(
															botUrl,
															chatId,
															"math;",
															deepData,
															requestingBlock+";",
															mathBlocks[requestingBlockIndex],
														)
														if err != nil {
															log.Println(err)
														}
													}
												}
											}
											default:
												fmt.Println("Unknown callback query:")
										}
									}
								}
							}
						}
					}
					default:
						fmt.Println("Unknown callback query:")
				}
			} else {
				fmt.Println("received a plain message")
				err = SendMessage(botUrl, update.Message.Chat.ChatId, "...?")
				if err != nil {
					log.Println(err)
				}
			}
			offset = update.UpdateId + 1
		}
	}
}