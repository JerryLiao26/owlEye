package helper

import "github.com/gen2brain/beeep"

func ShowNotify(title string, message string) {
	_ = beeep.Notify(title, message, "")
}

func ShowAlert(title string, message string) {
	_ = beeep.Alert(title, message, "")
}
