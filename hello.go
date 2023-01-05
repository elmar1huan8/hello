package main 
import (
	"fmt"
	"log"

	"example.com/greetings"
)


func main() {
	// Set properties of the predefined logger, including
	// the log entry prefix and a flag to diable printing
	// the time, source file, and a line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//Request a greeting message.
	message, err := greetings.Hello("Jordan")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
}
}	



