package solution

import "github.com/kyokomi/emoji"

func GetMessage() string {
	var message = "Hello 🗺️ !"
	return emoji.Sprint(message)
}
