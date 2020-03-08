package main

import (
	"bufio"
	"io"
)

func shjournalnagger(writer io.Writer, reader io.Reader) {

	writer.Write([]byte(defaultTopPrompt))
	writer.Write([]byte("hello"))

	userInputBuf := bufio.NewReader(reader)
	line, _ := userInputBuf.ReadBytes('\n')
	writer.Write(line)
}
