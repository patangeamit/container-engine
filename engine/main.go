package main

import (
	"os"

	"github.com/patangeamit/container-engine/cli"
	containerRuntime "github.com/patangeamit/container-engine/runtime"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--child" {
		containerRuntime.Child()
		return
	}
	cli.Cli(os.Args)
}
