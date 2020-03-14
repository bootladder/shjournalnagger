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
	line, err := userInputBuf.ReadBytes('\n')

	if err != nil {
		writer.Write([]byte("Quitting"))
		return
	}

	switch line[0] {
	case '1':
		commander.command(1)
	case 'q':
		break
	default:
		writer.Write([]byte("Invalid Input"))
	}

}
