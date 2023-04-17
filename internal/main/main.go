package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"tgbot/assets/ceko"
	"tgbot/internal/types"
	"tgbot/pkg/send"
	"tgbot/pkg/service/generate"
	"tgbot/pkg/service/handle"
)

func main() {

	BOT_TOKEN := os.Getenv("BOT_TOKEN")
	BOT_NAME := os.Getenv("BOT_NAME")
	if BOT_TOKEN == "" {
		log.Println("BOT_TOKEN env var is not set")
		secretJsonFile, err := os.Open("../../config/secret.json")
		if err != nil {
			log.Fatal("Json secret file wasn't found! Exit.")
		}
		fmt.Println("Successfully opened secret.json")
		type SecretJson struct {
			BOT_TOKEN string
		}

		secretJsonByteValue, _ := io.ReadAll(secretJsonFile)

		var secretJsonData SecretJson

		json.Unmarshal(secretJsonByteValue, &secretJsonData)

		BOT_TOKEN = secretJsonData.BOT_TOKEN

		secretJsonFile.Close()
	}
	if BOT_NAME == "" {
		log.Println("BOT_NAME env var is not set")
		secretJsonFile, err := os.Open("../../config/secret.json")
		if err != nil {
			log.Fatal("Json secret file wasn't found! Exit.")
		}
		fmt.Println("Successfully opened secret.json")
		type SecretJson struct {
			BOT_NAME string
		}

		secretJsonByteValue, _ := io.ReadAll(secretJsonFile)

		var secretJsonData SecretJson

		json.Unmarshal(secretJsonByteValue, &secretJsonData)

		BOT_NAME = secretJsonData.BOT_NAME

		secretJsonFile.Close()
	}

	fmt.Println("Bot has been launched")

	botAPI := "https://api.telegram.org/bot"
	botUrl := botAPI + BOT_TOKEN

	commands := types.BotCommands{
		Commands: []types.BotCommand{
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
			fmt.Println(string(upd))
			fmt.Println("---------------------------")

			chatId := update.CallbackQuery.Message.Chat.ChatId
			if update.Message.Sticker.FileId != "" {
				fmt.Println("received a sticker")
				err = send.Sticker(botUrl, update)
				if err != nil {
					log.Println(err)
				}
			} else if len(update.Message.Entities) > 0 {
				if update.Message.Entities[0].Type == "bot_command" {
					fmt.Println("received a command")
					// proceed command like /start@<bot_name>
					validCommand := true
					fmt.Println("full command: ", update.Message.Text)
					if strings.Contains(update.Message.Text, "@") {
						receivedBotName := ""
						for {
							if strings.Contains(receivedBotName, "@") || len(update.Message.Text) == 0 {
								break
							}
							receivedBotName = string(update.Message.Text[len(update.Message.Text)-1]) + receivedBotName
							update.Message.Text = update.Message.Text[0 : len(update.Message.Text)-1]
						}
						fmt.Println("rec full: ", receivedBotName)
						receivedBotName = receivedBotName[1:]
						fmt.Println("rec without @: ", receivedBotName)
						fmt.Println("bot_name: ", BOT_NAME)
						if receivedBotName != BOT_NAME {
							validCommand = false
						}
					}
					if validCommand {
						switch update.Message.Text[1:] {
						case "start":
							{
								err = send.Message(botUrl, update.Message.Chat.ChatId, "Welcome!")
								if err != nil {
									log.Println(err)
								}
							}
						case "library":
							{
								kbd := types.InlineKeyboardMarkup{
									InlineKeyboard: [][]types.InlineKeyboardButton{
										{
											types.InlineKeyboardButton{Text: "ЦЭКО", CallbackData: "library;ceko;"},
										},
									},
								}
								err = send.InlineKeyboard(botUrl, update.Message.Chat.ChatId, "Доступные ресурсы:", kbd)
								if err != nil {
									log.Println(err)
								}
							}
						case "feedback":
							{
								err = send.Message(botUrl, update.Message.Chat.ChatId, "Here is contact of the person who is willing to help you:\nhttps://t.me/undochlorine")
								if err != nil {
									log.Println(err)
								}
							}
						default:
							validCommand = false
						}
					}
					if !validCommand {
						err = send.Message(botUrl, update.Message.Chat.ChatId, "I'm not aware about this command, try later)")
						if err != nil {
							log.Println(err)
						}
					}
				}
			} else if update.CallbackQuery.Id != "" {
				fmt.Println("received a callback query")
				data := update.CallbackQuery.Data
				switch {
				case data[:len("library;")] == "library;":
					{
						if data == "library;" {
							kbd := types.InlineKeyboardMarkup{
								InlineKeyboard: [][]types.InlineKeyboardButton{
									{
										types.InlineKeyboardButton{Text: "ЦЭКО", CallbackData: "library;ceko;"},
									},
								},
							}
							err = send.InlineKeyboard(botUrl, chatId, "Доступные ресурсы:", kbd)
							if err != nil {
								log.Println(err)
							}
						} else {
							deepData := strings.Replace(data, "library;", "", -1)
							switch {
							case deepData[0:len("ceko;")] == "ceko;":
								{
									if deepData == "ceko;" {
										kbd := types.InlineKeyboardMarkup{
											InlineKeyboard: [][]types.InlineKeyboardButton{
												{
													types.InlineKeyboardButton{Text: "Математика", CallbackData: "library;ceko;math;"},
												},
												{
													types.InlineKeyboardButton{Text: "⏪ Сменить ресурс", CallbackData: "library;"},
												},
											},
										}
										err = send.InlineKeyboard(botUrl, chatId, "Выберете предмет:", kbd)
										if err != nil {
											log.Println(err)
										}
									} else {
										deepData = strings.Replace(deepData, "ceko;", "", -1)
										switch {
										case deepData[0:len("math;")] == "math;":
											{
												if deepData == "math;" {
													kbd := generate.GridInt(21, 4, "library;ceko;math;")
													err = send.InlineKeyboard(botUrl, chatId, "Номера доступных заданий:", kbd)
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
														err = handle.TaskQuery(
															botUrl,
															chatId,
															"math;",
															deepData,
															requestingBlock+";",
															ceko.MathBlocks[requestingBlockIndex],
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
				err = send.Message(botUrl, update.Message.Chat.ChatId, "...?")
				if err != nil {
					log.Println(err)
				}
			}
			offset = update.UpdateId + 1
		}
	}
}
