package main

import (
	"fmt"
	"os"
)

var (
	USAGE_MSG           string = "Usage: main.go <run, config, apply>"
	UNKNOWN_COMMAND_MSG string = "Unknown command:"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(USAGE_MSG)
		return
	}
	switch os.Args[1] {
	case "run":
		fmt.Println("Running container...")
	case "config":
		fmt.Println("Cofiguring...")
	case "apply":
		fmt.Println("Applying...")
	default:
		fmt.Println(UNKNOWN_COMMAND_MSG, os.Args[1])
	}
}
