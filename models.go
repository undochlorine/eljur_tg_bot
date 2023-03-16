package main

type Update struct {
	UpdateId		int				`json:"update_id"`
	Message			Message 		`json:"message"`
	CallbackQuery	CallbackQuery	`json:"callback_query"`
}

type Message struct {
	Chat				Chat					`json:"chat"`
	Text				string					`json:"text"`
	Sticker 			Sticker					`json:"sticker"`
	WebAppData			WebAppData				`json:"web_app_data"`
	Entities			[]MessageEntity			`json:"entities"`
	ReplyMarkupInline	InlineKeyboardMarkup	`json:"reply_markup"`
}

type MessageEntity struct {
	Type	string	`json:"type"`
}

type WebAppData struct {
	Data 		string	`json:"data"`
	ButtonText	string	`json:"button_text"`
}

type CallbackQuery struct {
	Id				string 		`json:"id"`
	From 			User 		`json:"from"`
	Message 		Message 	`json:"message"`
	ChatInstance 	string 		`json:"chat_instance"`
	Data 			string		`json:"data"`
}

type User struct {
	Id 			int 	`json:"id"`
	IsBot 		bool 	`json:"is_bot"`
	Firstname 	string 	`json:"first_name"`
}

type Chat struct {
	ChatId	int 	`json:"id"`
}

type Sticker struct {
	FileId	string	`json:"file_id"`
}

type RestResponse struct {
	Result	[]Update 	`json:"result"`
}

type BotMessage struct {
	ChatId	int 	`json:"chat_id"`
	Text	string	`json:"text"`
}

type BotMessageWithInlineKbd struct {
	ChatId				int 					`json:"chat_id"`
	Text				string					`json:"text"`
	ReplyMarkupInline	InlineKeyboardMarkup	`json:"reply_markup"`
}

type BotSticker struct {
	ChatId	int 	`json:"chat_id"`
	Sticker string `json:"sticker"`
}

type BotPhoto struct {
	ChatId int 	`json:"chat_id"`
	Photo 	string 	`json:"photo"`
}

type BotCommand struct {
	Command 	string	`json:"command"`
	Description	string	`json:"description"`
}

type BotCommands struct {
	Commands 	[]BotCommand 	`json:"commands"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard 	[][]InlineKeyboardButton 	`json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text 			string		`json:"text"`
	CallbackData 	string		`json:"callback_data"`
}