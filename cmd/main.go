package main

import (
	"fmt"
	"os"
)

func main() {

	ok := os.Getenv("MODE")
	if ok == "" {
		ok = "server"
	}
	switch ok {

	case "server":

		serverRun()

	default:
		fmt.Println("Unknown mode. Exiting.")
	}
}
