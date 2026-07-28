package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// If this process is started with the argument "child",
	// we're already inside the new namespaces.
	if len(os.Args) > 1 && os.Args[1] == "child" {
		child()
		return
	}

	parent()
}

func parent() {
	// Re-execute this same binary.
	cmd := exec.Command("/proc/self/exe", "child")

	// Tell the kernel to create new namespaces for the child.
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNET | // New network namespace
			syscall.CLONE_NEWUTS | // New hostname/domain namespace
			syscall.CLONE_NEWIPC, // New System V/POSIX IPC namespace
	}

	// Connect stdin/stdout/stderr so we can interact with it.
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func child() {
	fmt.Println("Inside isolated namespaces!")

	// Change the hostname.
	// This only affects this UTS namespace.
	if err := syscall.Sethostname([]byte("pause-container")); err != nil {
		panic(err)
	}

	hostname, _ := os.Hostname()
	fmt.Println("Hostname:", hostname)

	// Print PID (still in the parent's PID namespace).
	fmt.Println("PID:", os.Getpid())

	fmt.Println("Sleeping forever...")

	// Simulate Kubernetes' pause container.
	// It simply stays alive to own the namespaces.
	for {
		syscall.Pause()
	}
}
