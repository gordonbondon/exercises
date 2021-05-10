package main

import (
	"log"

	"github.com/gordonbondon/exercises/takehome/cron_parser1/internal/cmd"
)

// main runs cron_parser1 CLI
func main() {
	if err := cmd.Run(); err != nil {
		log.Fatalf("%v", err)
	}
}
