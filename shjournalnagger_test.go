package main

import (
	"bytes"
	"testing"
)

func Test_WritesAnythingToWriter(t *testing.T) {
	var writer bytes.Buffer
	var reader bytes.Buffer
	shjournalnagger(&writer, &reader)

	if false == bytes.Contains(writer.Bytes(), []byte("hello")) {
		t.Fatalf("Does not contain hello")
	}
}

func Test_ReadsLineOfUserInput(t *testing.T) {
	var writer bytes.Buffer
	reader := bytes.NewReader([]byte("user input line 1\nuser input line 2\n"))

	shjournalnagger(&writer, reader)

	length := reader.Len()
	if length != 0 {
		t.Fatalf("Did not read the user input")
	}
}

func Test_PrintsTheTopPrompt(t *testing.T) {
	var writer bytes.Buffer
	var reader bytes.Buffer
	shjournalnagger(&writer, &reader)

	if false == bytes.Contains(writer.Bytes(), []byte("something")) {
		t.Fatalf("Does not Print the Top Prompt")
	}
}

func Test_PrintsMenu(t *testing.T) {
	var writer bytes.Buffer
	var reader bytes.Buffer
	shjournalnagger(&writer, &reader)
}
