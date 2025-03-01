package main

import (
	"fmt"
	"log"

	"github.com/google/go-cmp/cmp"
//	"github.com/rogerjd/greetings"
	"github.com/rogerjd/helloMod/morestrings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("abc"))
	fmt.Println(cmp.Diff("Hello world", "Hello world"))
	fmt.Println(quote.Go())

	// Set properties of the predefined logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

/*
	names := []string{"Gladys", "Samantha", "Darrin"}
//	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	//return

	for {
		var name string
		fmt.Print("enter a name: ")
		fmt.Scanln(&name)
		if name == "exit" {
			break
		}
//		msg, err := greetings.Hello(name)
		if err != nil {
			//log.Fatal(err)
			log.Println(err)
		} else {
			fmt.Println(msg)
		}
	}
*/
}
