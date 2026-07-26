package main

import (
	"fmt"
	"log"

	"king.com/greeting"
)

func main() {
	log.SetPrefix("greeting:")
	log.SetFlags(0)
	message, err := greeting.Hello("")
	if err != nil {
		// log.Fatal(err)
	}
	fmt.Println(message)

	msg, _ := greeting.Hello("Dreamer")
	fmt.Println(msg)
}