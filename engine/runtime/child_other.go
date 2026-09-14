//go:build !linux

package runtime

func Child(args []string) {
	panic("container runtime requires Linux")
}
