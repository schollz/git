package main

import (
	"os"
	"os/exec"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		if args[1] == "pull" {
			args = append(args[:2], append([]string{"--depth=1"}, args[2:]...)...)
		} else if args[1] == "clone" {
			args = append(args[:2], append([]string{"--depth=1", "--shallow-submodules", "--single-branch"}, args[2:]...)...)
		}
	}
	args[0] = "/usr/bin/git"
	cmd := exec.Command(args[0])
	if len(args) > 0 {
		cmd = exec.Command(args[0], args[1:]...)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}
