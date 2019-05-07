package ui

import (
	"github.com/zserge/lorca"

	"github.com/JerryLiao26/owlEye/helper"
)

func StartUI() {
	// Read index.html
	indexHtml := helper.ReadAsHtml("./ui/html/index.html")

	// Start window
	ui, _ := lorca.New(indexHtml, "", 800, 600)
	defer ui.Close()

	// Bind functions
	_ = ui.Bind("jumpInput", func() {
		// Read input.html
		inputHtml := helper.ReadAsHtml("./ui/html/input.html")

		_ = ui.Load(inputHtml)
	})

	// Wait for window close
	<-ui.Done()
}