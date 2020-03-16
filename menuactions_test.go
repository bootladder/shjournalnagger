package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockCommandExecuter struct {
	lastCommand string
}

func (m *MockCommandExecuter) executeCommand(cmd string) {
	m.lastCommand = cmd
}

func Test_openJournal_1_vim(t *testing.T) {
	mockCommandExecter := MockCommandExecuter{}

	journalConfigFile := getTestJournals()
	journalConfigFile.Editor = "vim"

	menuActions := MenuActionsImpl{&mockCommandExecter, journalConfigFile}
	menuActions.openJournal(1)

	assert.Equal(t, "vim "+journalConfigFile.Journals[0].Path, mockCommandExecter.lastCommand)
}

func Test_openJournal_2_nano(t *testing.T) {
	mockCommandExecter := MockCommandExecuter{}

	journalConfigFile := getTestJournals()
	journalConfigFile.Editor = "nano"

	menuActions := MenuActionsImpl{&mockCommandExecter, journalConfigFile}
	menuActions.openJournal(2)

	assert.Equal(t, "nano "+journalConfigFile.Journals[1].Path, mockCommandExecter.lastCommand)
}
