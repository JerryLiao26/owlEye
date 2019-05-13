package ui

import (
	"github.com/JerryLiao26/owlEye/helper"
	"github.com/zserge/lorca"
)

// Package global flag
var mainUIStatus = false

// Map of sub-UIs
var activeSubUI = make(map[string]lorca.UI)

func StartUI() {
	if !mainUIStatus {
		// Read index.html
		indexHtml := helper.ReadAsHtml("./ui/html/index.html")

		// Start window
		ui, _ := lorca.New(indexHtml, "", 800, 600)
		mainUIStatus = true
		defer func() {
			_ = ui.Close()
			mainUIStatus = false
			closeAllSubUI()
		}()

		// Bind functions
		_ = ui.Bind("jumpInput", func() {
			// Read input.html
			inputHtml := helper.ReadAsHtml("./ui/html/input.html")

			_ = ui.Load(inputHtml)
		})

		_ = ui.Bind("jumpPhoto", func() {
			photoHtml := helper.ReadAsHtml("./ui/html/photo.html")

			_ = ui.Load(photoHtml)
		})

		_ = ui.Bind("showConfig", func() {
			ConfigUI()
		})

		_ = ui.Bind("exit", func() {
			_ = ui.Close()
			mainUIStatus = false
			closeAllSubUI()
		})

		// Wait for window close
		<-ui.Done()
		mainUIStatus = false
		closeAllSubUI()
	}
}

func closeAllSubUI() {
	for _, ui := range activeSubUI {
		_ = ui.Close()
	}
}
