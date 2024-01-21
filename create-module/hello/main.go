package main

import (
	"fmt"
	"log"

	"andreassundstrom/greetings"
)

func main() {
	message, err := greetings.Hello("Andreas")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(message)
}
