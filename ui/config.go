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

		// Init
		jsString := "init(\"" + helper.TextServerPath + "\", \"" + helper.TextServerPort + "\", \"" + helper.ImageServerPath + "\", \"" + helper.ImageServerPort + "\")"
		ui.Eval(jsString)

		// Bind functions
		_ = ui.Bind("save", func(tspa string, tspo string, ispa string, ispo string) {
			helper.TextServerPath = tspa
			helper.TextServerPort = tspo
			helper.ImageServerPath = ispa
			helper.ImageServerPort = ispo

			// Save config
			helper.SaveConf()

			// Test server in goroutine
			go helper.TestServer()

			_ = ui.Close()
		})

		// Wait for window close
		<-ui.Done()
		configUIStatus = false
		delete(activeSubUI, "config")
	}
}
