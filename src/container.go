package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

func child() {
	if err := syscall.Sethostname([]byte("container")); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inside isolated UTS namespace!")
	cmd := exec.Command("/bin/sh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "child" {
		child()
		return
	}

	cmd := exec.Command("/proc/self/exe", "child")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWNET,
	}

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
