package main

import (
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
}
