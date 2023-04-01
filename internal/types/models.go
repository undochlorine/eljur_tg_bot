package types

type Update struct {
	UpdateId      int           `json:"update_id,omitempty"`
	Message       Message       `json:"message,omitempty"`
	CallbackQuery CallbackQuery `json:"callback_query,omitempty"`
}

type Message struct {
	Chat              Chat                 `json:"chat,omitempty"`
	Text              string               `json:"text,omitempty"`
	Sticker           Sticker              `json:"sticker,omitempty"`
	WebAppData        WebAppData           `json:"web_app_data,omitempty"`
	Entities          []MessageEntity      `json:"entities,omitempty"`
	ReplyMarkupInline InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type MessageEntity struct {
	Type string `json:"type,omitempty"`
}

type WebAppData struct {
	Data       string `json:"data,omitempty"`
	ButtonText string `json:"button_text,omitempty"`
}

type CallbackQuery struct {
	Id           string  `json:"id,omitempty"`
	From         User    `json:"from,omitempty"`
	Message      Message `json:"message,omitempty"`
	ChatInstance string  `json:"chat_instance,omitempty"`
	Data         string  `json:"data,omitempty"`
}

type User struct {
	Id        int    `json:"id,omitempty"`
	IsBot     bool   `json:"is_bot,omitempty"`
	Firstname string `json:"first_name,omitempty"`
}

type Chat struct {
	ChatId int `json:"id,omitempty"`
}

type Sticker struct {
	FileId string `json:"file_id,omitempty"`
}

type RestResponse struct {
	Result []Update `json:"result,omitempty"`
}

type BotMessage struct {
	ChatId int    `json:"chat_id,omitempty"`
	Text   string `json:"text,omitempty"`
}

type BotMessageWithInlineKbd struct {
	ChatId            int                  `json:"chat_id,omitempty"`
	Text              string               `json:"text,omitempty"`
	ReplyMarkupInline InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

type BotSticker struct {
	ChatId  int    `json:"chat_id,omitempty"`
	Sticker string `json:"sticker,omitempty"`
}

type BotPhoto struct {
	ChatId int    `json:"chat_id,omitempty"`
	Photo  string `json:"photo,omitempty"`
}

type BotCommand struct {
	Command     string `json:"command,omitempty"`
	Description string `json:"description,omitempty"`
}

type BotCommands struct {
	Commands []BotCommand `json:"commands,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard,omitempty"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text,omitempty"`
	CallbackData string `json:"callback_data,omitempty"`
}
