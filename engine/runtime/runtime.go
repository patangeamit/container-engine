package runtime

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"golang.org/x/sys/unix"
)

func Child() {
	fmt.Println("Contianer initialized")
	Must(unix.Unshare(unix.CLONE_NEWNET | unix.CLONE_NEWUTS))
	cmd := exec.Command("/bin/sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	Must(cmd.Run())
	fmt.Println("Container terminated")
}

func Run() {
	if runtime.GOOS != "linux" {
		fmt.Printf("OS (%v) not supported for containers. Exiting.\n", runtime.GOOS)
		return
	}
	fmt.Println("Running a container")
	cmd := exec.Command("/proc/self/exe", "--child")
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
