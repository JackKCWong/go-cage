package main

import (
	"os"
	"os/exec"
	"syscall"
	"github.com/JackKCWong/go-cage/cage"
)


func main() {
	newroot := os.Args[1]
	exe := os.Args[2:]

	// if err := cage.MountProc(newroot); err != nil {
	// 	panic(err)
	// }

	if err := cage.PivotRoot(newroot); err != nil {
		panic(err)
	}

	cmd := exec.Command(exe[0], exe[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS,
	}

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}