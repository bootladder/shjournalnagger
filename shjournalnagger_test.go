package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReadsLineOfUserInput(t *testing.T) {
	var writer bytes.Buffer
	reader := bytes.NewReader([]byte("user input line 1\nuser input line 2\n"))

	shjournalnagger(&writer, reader, Journals{}, nil)

	length := reader.Len()
	if length != 0 {
		t.Fatalf("Did not read the user input")
	}
}

func Test_PrintsTheTopPrompt(t *testing.T) {
	var writer bytes.Buffer
	var reader bytes.Buffer
	shjournalnagger(&writer, &reader, Journals{}, nil)

	if false == bytes.Contains(writer.Bytes(), []byte("Write something!!!")) {
		t.Fatalf("Does not Print the Top Prompt")
	}
}

func Test_PrintsMenu(t *testing.T) {
	var writer bytes.Buffer
	var reader bytes.Buffer
	journals := getTestJournals()

	shjournalnagger(&writer, &reader, journals, nil)
	assert.Contains(t, string(writer.Bytes()), "Select")
}

func Test_InvalidUserInput_PrintsMessage(t *testing.T) {
	var writer bytes.Buffer
	reader := bytes.NewReader([]byte("user input line 1\nuser input line 2\n"))

	shjournalnagger(&writer, reader, Journals{}, nil)

	assert.Contains(t, string(writer.Bytes()), "Invalid Input")
}

type MockCommander struct {
	called    bool
	lastValue int
}

func (m *MockCommander) command(i int) {
	m.called = true
	m.lastValue = i
}

func Test_Journal1Selected_CallsCommander(t *testing.T) {
	var writer bytes.Buffer
	reader := bytes.NewReader([]byte("1\n"))

	mockCommander := MockCommander{}
	shjournalnagger(&writer, reader, Journals{}, &mockCommander)

	assert.NotContains(t, string(writer.Bytes()), "Invalid Input")
	assert.Equal(t, 1, mockCommander.lastValue)
}

func Test_Quit_QuitsQuietly(t *testing.T) {
	var writer bytes.Buffer
	reader := bytes.NewReader([]byte("q\n"))

	shjournalnagger(&writer, reader, Journals{}, nil)
	assert.NotContains(t, string(writer.Bytes()), "Invalid Input")
}
