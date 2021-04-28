package main

import (
	"log"

	"github.com/gordonbondon/exercises/takehome/webcrawler1/internal/cmd"
)

// main runs webcrawler1 CLI
func main() {
	if err := cmd.Run(); err != nil {
		log.Fatalf("%v", err)
	}
}
