package main

type Update struct {
	UpdateId	int		`json:"update_id"`
	Message		Message `json: "message"`
}

type Message struct {
	Chat	Chat	`json:"chat"`
	Text	string	`json:"text"`
	Sticker	Sticker	`json:"sticker"`
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

type BotSticker struct {
	ChatId	int 	`json:"chat_id"`
	Sticker string `json:"sticker"`
}