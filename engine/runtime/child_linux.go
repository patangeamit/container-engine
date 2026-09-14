//go:build linux

package runtime

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"go.yaml.in/yaml/v3"
	"golang.org/x/sys/unix"
)

func Child(args []string) {

	flagSet := flag.NewFlagSet("child", flag.ExitOnError)
	configFile := flagSet.String("f", "config.yaml", "Configuration file for containers")
	flagSet.Parse(args[1:])

	var config Config
	data, err := os.ReadFile(*configFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	Must(yaml.Unmarshal([]byte(data), &config))

	unshare(&config.Namespaces)

	if config.Namespaces.Hostname {
		unix.Sethostname([]byte(config.Name))
	}

	cmd := exec.Command(config.EntryPoint[0], config.EntryPoint[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	fmt.Println("Contianer initialized")

	Must(cmd.Start())
	Must(cmd.Wait())

	fmt.Println("Container terminated")
}

func unshare(ns *Namespaces) {
	var flag int
	if ns.Hostname {
		flag |= unix.CLONE_NEWUTS
	}
	if ns.Network {
		flag |= unix.CLONE_NEWNET
	}
	if ns.Ipc {
		flag |= unix.CLONE_NEWIPC
	}
	if ns.Process {
		flag |= unix.CLONE_NEWPID
	}
	if ns.Time {
		flag |= unix.CLONE_NEWTIME
	}
	if ns.Mount {
		flag |= unix.CLONE_NEWNS
	}
	if ns.Cgroup {
		flag |= unix.CLONE_NEWCGROUP
	}
	if ns.User {
		flag |= unix.CLONE_NEWUSER
	}
	Must(unix.Unshare(flag))
}
