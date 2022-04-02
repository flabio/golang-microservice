package main

import (
	"test/app"
	"test/logger"
)

func main() {
	logger.Info("Starting our application...")
	app.Start()
}
