package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	var message = "Hello 🗺️!"
	return emoji.Sprint(message)
}
