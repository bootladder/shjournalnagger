package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func shjournalnagger(
	userOutputWriter io.Writer,
	userInputReader io.Reader,
	journalConfigFile JournalConfigFile,
	menuActions MenuActions,
	naggingIntervalTracker NaggingIntervalTracker) {

	if false == naggingIntervalTracker.isNaggingIntervalExpired(journalConfigFile.NaggingIntervalSeconds) {
		return
	}

	naggingIntervalTracker.updateLastNaggingTime()

	userOutputWriter.Write([]byte(defaultTopPrompt))
	userOutputWriter.Write([]byte(renderMenu(journalConfigFile)))

	userInputBuf := bufio.NewReader(userInputReader)
	inputBytesWithNewline, err := userInputBuf.ReadBytes('\n')

	if err != nil {
		userOutputWriter.Write([]byte("Quitting"))
		return
	}

	line := strings.TrimSpace(string(inputBytesWithNewline))

	if string(line) == "q" {
		return
	}

	number, err := strconv.Atoi((line))
	if err == nil {
		if number > 0 && number <= len(journalConfigFile.Journals) {
			menuActions.openJournal(number)
			return
		}
	}

	if string(line) == "e" {
		menuActions.openConfigFile()
		return
	}

	userOutputWriter.Write([]byte("Invalid Input\n"))
}
