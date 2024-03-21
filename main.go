// main.go

package main

import (
	"APR/console"
	"APR/routes"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.DebugMode)
	router = gin.Default()
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static/")
	router.Static("/assets", "./images/")
	routes.InitialiseRoutes(router)
	// Start serving the application
	go router.Run("127.0.0.1:8080")
	go console.Console()
	console.Shutdown = make(chan os.Signal, 1)
	signal.Notify(console.Shutdown, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-console.Shutdown
	println("Starting Shutdown")
}
