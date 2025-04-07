package main

import (
	"log"
	"os"

	"github.com/mholm/wlog/internal/cli"
)

func main() {
	app, err := cli.NewCLI()
	if err != nil {
		log.Fatalf("Failed to initialize CLI: %v", err)
	}

	if err := app.Execute(); err != nil {
		os.Exit(1)
	}
} 