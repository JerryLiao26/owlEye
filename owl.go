package main

import (
	"github.com/JerryLiao26/owlEye/helper"
	"github.com/JerryLiao26/owlEye/ui"
)

func main() {
	// Test server in goroutine
	go helper.TestServer()

	// Init UI
	ui.StartUI()
}
