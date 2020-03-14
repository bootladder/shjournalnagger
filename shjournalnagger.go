package main

import (
	"bufio"
	"io"
)

func shjournalnagger(
	writer io.Writer,
	reader io.Reader,
	journals Journals,
	commander Commander) {

	writer.Write([]byte(defaultTopPrompt))

	menuStr := renderMenu(journals)
	writer.Write([]byte(menuStr))

	userInputBuf := bufio.NewReader(reader)
	line, _ := userInputBuf.ReadBytes('\n')

	if string(line) != "1\n" {
		writer.Write([]byte("Invalid Input"))
	} else {
		commander.command(1)
	}

}
