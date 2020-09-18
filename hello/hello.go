package main

import (
	"dg/utils"
	"log"
	"os"
)

func init() {
	log.Println("in init")
}

// main function, where it all starts
func main() {
	//var s = "bla"
	var b = os.Getenv("PROCESSOR_IDENTIFIER")
	log.Println("GO test>> ", b, "PID=", os.Getpid(), "xx")

	log.Println("blue and red")

	type User struct {
		name     string
		email    string
		adminInt bool
	}

	david := User{name: "David James", adminInt: true}

	log.Println("user is", utils.StringEcho("bla"))
	log.Println("user is", utils.ToStringStruct(david))
	log.Println("user is", david)
}
