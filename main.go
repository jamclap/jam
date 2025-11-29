package main

import (
	"log"
	"os"

	"github.com/jamclap/jam/jam/run"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	path := os.Args[1]
	r := run.NewRunner()
	err := r.RunPath(path)
	if err != nil {
		log.Panic(err)
	}
}
