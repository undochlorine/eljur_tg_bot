package get

import (
	"math/rand"
	"tgbot/assets/stickers"
)

func RandomSticker() string {
	var i = rand.Intn(len(stickers.Stickers))
	return stickers.Stickers[i]
}
