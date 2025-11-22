package main

import (
	"log"

	"github.com/jamclap/jam/jam/run"
)

func main() {
	err := run.Run(`
package main

func main() {
	println("Hi!")
}
	`)
	if err != nil {
		log.Panic(err)
	}
}
