package main

import (
	"dg/utils"
	"fmt"
	"log"
	"os"
)

// main function, where it all starts
func main() {
	//var s = "bla"
	var b = os.Getenv("PROCESSOR_IDENTIFIER")
	fmt.Println("GO test>> ", b, "PID=", os.Getpid(), "xx")

	log.Println("blue and red")

	type User struct {
		name     string
		email    string
		adminInt bool
	}

	david := User{name: "David James", adminInt: true}

	log.Println("user is", utils.StringEcho("bla"))
	log.Println("user is", utils.ToStringStruct(david))
}
