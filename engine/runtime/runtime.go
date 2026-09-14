package runtime

import (
	"fmt"
	"os"
	"os/exec"
	goRuntime "runtime"
)

func Run() {
	if goRuntime.GOOS != "linux" {
		fmt.Printf("OS (%v) not supported for containers. Exiting.\n", goRuntime.GOOS)
		return
	}
	fmt.Println("Running a container")
	cmd := exec.Command(
		"/proc/self/exe",
		append([]string{"--child"}, os.Args[1:]...)...,
	)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	Must(cmd.Run())
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustRet(out any, err error) any {
	if err != nil {
		panic(err)
	} else {
		return out
	}
}
