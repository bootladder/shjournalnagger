package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func shjournalnagger(
	writer io.Writer,
	reader io.Reader,
	journalConfigFile JournalConfigFile,
	journalOpener JournalOpener) {

	writer.Write([]byte(defaultTopPrompt))
	writer.Write([]byte(renderMenu(journalConfigFile)))

	userInputBuf := bufio.NewReader(reader)
	inputBytesWithNewline, err := userInputBuf.ReadBytes('\n')

	if err != nil {
		writer.Write([]byte("Quitting"))
		return
	}

	line := strings.TrimSpace(string(inputBytesWithNewline))

	if string(line) == "q" {
		return
	}

	number, err := strconv.Atoi((line))
	if err == nil {
		if number > 0 && number <= len(journalConfigFile.Journals) {
			journalOpener.openJournal(number)
			return
		}
	}

	writer.Write([]byte("Invalid Input\n"))
}
