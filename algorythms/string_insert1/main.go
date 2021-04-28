package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gordonbondon/exercises/algorythms/string_insert1/check"
)

func main() {
	var stringFlag string
	var resultFlag string

	flag.StringVar(&stringFlag, "string", "", "string to try inserting")
	flag.StringVar(&resultFlag, "result", "", "expected result")
	flag.Parse()

	if stringFlag == "" {
		log.Fatalf("string required")
	}

	if resultFlag == "" {
		log.Fatalf("result required")
	}

	fmt.Println(check.Check(stringFlag, resultFlag))
}
