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
	err := run.RunPath(path)
	// b, err := os.ReadFile(path)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// source := string(b)
	// 	err := run.Run(`
	// package main

	// func main() {
	// 	println("Hi!")
	// }
	// 	`)
	if err != nil {
		log.Panic(err)
	}
}
