package main

import (
	"github.com/JerryLiao26/owlEye/server"
	"github.com/JerryLiao26/owlEye/ui"
)

func main() {
	// Init server in goroutine
	go server.StartServer()

	// Init UI
	ui.StartUI()
}
