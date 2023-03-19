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
					case data[0:len("library;ceko;")] == "library;ceko;": {
						if data == "library;ceko;" {
							kbd := InlineKeyboardMarkup{
								InlineKeyboard: [][]InlineKeyboardButton{
									{
										InlineKeyboardButton{Text: "Математика", CallbackData: "library;ceko;math;"},
									},
								},
							}
							err = SendInlineKeyboard(botUrl, chatId, "Выберете предмет:", kbd)
							if err != nil {
								log.Println(err)
							}
						} else {
							deepData := strings.Replace(data, "library;ceko;", "", -1)
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
										switch {
											case deepData[0:len("1;")] == "1;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"1;",
													TasksOfBlock1Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("2;")] == "2;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"2;",
													TasksOfBlock2Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("3;")] == "3;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"3;",
													TasksOfBlock3Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("4;")] == "4;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"4;",
													TasksOfBlock4Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("5;")] == "5;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"5;",
													TasksOfBlock5Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("6;")] == "6;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"6;",
													TasksOfBlock6Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("7;")] == "7;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"7;",
													TasksOfBlock7Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("8;")] == "8;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"8;",
													TasksOfBlock8Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("9;")] == "9;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"9;",
													TasksOfBlock9Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("10;")] == "10;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"10;",
													TasksOfBlock10Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("11;")] == "11;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"11;",
													TasksOfBlock11Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("12;")] == "12;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"12;",
													TasksOfBlock12Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("13;")] == "13;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"13;",
													TasksOfBlock13Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("14;")] == "14;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"14;",
													TasksOfBlock14Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("15;")] == "15;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"15;",
													TasksOfBlock15Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("16;")] == "16;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"16;",
													TasksOfBlock16Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("17;")] == "17;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"17;",
													TasksOfBlock17Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("18;")] == "18;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"18;",
													TasksOfBlock18Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("19;")] == "19;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"19;",
													TasksOfBlock19Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("20;")] == "20;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"20;",
													TasksOfBlock20Math,
												)
												if err != nil {
													log.Println(err)
												}
											}
											case deepData[0:len("21;")] == "21;": {
												err = handleTaskQuery(
													botUrl,
													chatId,
													"math;",
													deepData,
													"21;",
													TasksOfBlock21Math,
												)
												if err != nil {
													log.Println(err)
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