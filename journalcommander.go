package main

import (
	"os"
	"os/exec"
)

// JournalCommander sends commands to edit journals
type JournalCommander struct {
}

func (j JournalCommander) command(int) {
	bashCommand("vi")
}

func bashCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
