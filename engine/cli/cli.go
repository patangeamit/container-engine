package cli

import (
	"fmt"

	"github.com/patangeamit/container-engine/runtime"
)

var (
	USAGE_MSG           string = "Usage: main.go <run, config, apply>"
	UNKNOWN_COMMAND_MSG string = "Unknown command:"
)

func Cli(args []string) {
	if len(args) < 2 {
		fmt.Println(USAGE_MSG)
		return
	}
	switch args[1] {
	case "run":
		runtime.Run()
	case "config":
		fmt.Println("Cofiguring...")
	case "apply":
		fmt.Println("Applying...")
	case "list":
		fmt.Printf("%-20s %-12s %-15s\n", "NAME", "STATUS", "NAMESPACES")
		fmt.Printf("%-20s %-12s %-15s\n", "nginx-turbo", "stopped", "IPC, NET")
	default:
		fmt.Println(UNKNOWN_COMMAND_MSG, args[1])
	}
}
