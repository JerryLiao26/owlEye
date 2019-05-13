package server

import (
	"context"
	"github.com/JerryLiao26/owlEye/helper"
	"github.com/kataras/iris"
	"time"
)

var runningServers []*iris.Application

/* Server operations */
func StartServer() {
	// Load config
	helper.LoadConf()

	// Init
	runningServers = make([]*iris.Application, 0)

	// Run server
	if (helper.ImageServerPath == helper.TextServerPath) && (helper.ImageServerPort == helper.TextServerPort) {
		// Server with default middleware
		app := iris.Default()

		// Handlers
		app.Get("/hello", helloHandler)
		app.Post("/text", textHandler)
		app.Post("/image", imageHandler)
		app.Post("/upload", uploadHandler)

		// Store in group
		runningServers = append(runningServers, app)

		// Run server
		serveString := helper.TextServerPath + ":" + helper.ImageServerPort
		_ = app.Run(iris.Addr(serveString))
	} else {
		// Apps
		textApp := iris.Default()
		imageApp := iris.Default()

		// Handlers
		textApp.Get("/hello", helloHandler)
		textApp.Post("/text", textHandler)

		imageApp.Get("/hello", helloHandler)
		imageApp.Post("/image", imageHandler)
		imageApp.Post("/upload", uploadHandler)

		// Store
		runningServers = append(runningServers, textApp)
		runningServers = append(runningServers, imageApp)

		// Run
		textServerString := helper.TextServerPath + ":" + helper.TextServerPort
		imageServerString := helper.ImageServerPath + ":" + helper.ImageServerPort
		go func() {
			_ = textApp.Run(iris.Addr(textServerString))
		}()
		go func() {
			_ = imageApp.Run(iris.Addr(imageServerString))
		}()
	}
}

func StopServers(restart bool) {
	// Timeout
	timeout := 1 * time.Second

	// Setup context
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Shutdown
	for _, app := range runningServers {
		_ = app.Shutdown(ctx)
	}

	// Wait for shutdown finish
	<-ctx.Done()

	if restart {
		StartServer()
	}
}

/* Handlers */
func helloHandler(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"status": true,
	})
}

func textHandler(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"status": true,
	})
}

func imageHandler(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"status": true,
	})
}

func uploadHandler(ctx iris.Context) {
	_, _ = ctx.JSON(iris.Map{
		"status": true,
	})
}
