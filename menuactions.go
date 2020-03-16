package main

import (
	"os"
	"os/exec"
)

//MenuActions does stuff
type MenuActions interface {
	openJournal(int)
	openConfigFile()
}

// MenuActionsImpl sends commands to edit journals
type MenuActionsImpl struct {
	commandExecuter   CommandExecuter
	journalConfigFile JournalConfigFile
}

// CommandExecuter Accepts commands that will actually run in a shell or whatever
type CommandExecuter interface {
	executeCommand(string)
}

func (j *MenuActionsImpl) openJournal(journalNum int) {
	path := j.journalConfigFile.Journals[journalNum-1].Path
	editor := j.journalConfigFile.Editor
	j.commandExecuter.executeCommand(editor + " " + path)
}

func (j *MenuActionsImpl) openConfigFile() {
	editor := j.journalConfigFile.Editor
	j.commandExecuter.executeCommand(editor + " " + journalConfigFilename)
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
