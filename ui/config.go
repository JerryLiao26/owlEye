package ui

import (
	"github.com/JerryLiao26/owlEye/helper"
	"github.com/zserge/lorca"
)

var configUIStatus = false

func ConfigUI() {
	if !configUIStatus && mainUIStatus {
		configHtml := helper.ReadAsHtml("./ui/html/config.html")

		ui, _ := lorca.New(configHtml, "", 480, 320)
		configUIStatus = true
		activeSubUI["config"] = ui // Append in UI map
		defer func() {
			_ = ui.Close()
			configUIStatus = false
			delete(activeSubUI, "config")
		}()

		// Wait for window close
		<-ui.Done()
		configUIStatus = false
		delete(activeSubUI, "config")
	}
}
