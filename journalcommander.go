package main

import (
	"os"
	"os/exec"
)

//JournalOpener opens a journal
type JournalOpener interface {
	openJournal(int)
}

// JournalCommander sends commands to edit journals
type JournalCommander struct {
	commandExecuter   CommandExecuter
	journalConfigFile JournalConfigFile
}

// CommandExecuter Accepts commands that will actually run in a shell or whatever
type CommandExecuter interface {
	executeCommand(string)
}

func (j *JournalCommander) openJournal(journalNum int) {
	path := j.journalConfigFile.Journals[journalNum-1].Path
	editor := j.journalConfigFile.Editor
	j.commandExecuter.executeCommand(editor + " " + path)
}

// ShellCommandExecuter  runs commands in the shell
type ShellCommandExecuter struct {
}

func (b *ShellCommandExecuter) executeCommand(command string) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
