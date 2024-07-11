package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main(){
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	log.Default()
	message, err := greetings.Hello("Midas")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Midas", "test", "Qin", }
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
