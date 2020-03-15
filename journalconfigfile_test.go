package main

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockReadWriter struct {
	shouldFailReading            bool
	shouldReadInvalidFileContent bool
	bytesWritten                 []byte
}

func (m MockReadWriter) Read(b []byte) (int, error) {
	if m.shouldFailReading {
		return 0, errors.New("fail")
	}
	if m.shouldReadInvalidFileContent {
		var dummyString = []byte("blah blha bad ymal file")
		copy(b, dummyString)
		return len(dummyString), io.EOF
	}

	journalBytes, _ := ioutil.ReadFile("test-journals.yaml")
	copy(b, journalBytes)

	return len(journalBytes), io.EOF
}
func (m *MockReadWriter) Write(b []byte) (int, error) {

	m.bytesWritten = b
	return 0, nil
}

func Test_JournalConfigFileLogic_CannotRead_WritesDefaultConfig_PrintsMessage(t *testing.T) {

	var mockReadWriter MockReadWriter
	mockReadWriter.shouldFailReading = true

	var writer bytes.Buffer
	journalConfigFileLogic(&mockReadWriter, &writer)

	assert.Equal(t, defaultJournalConfig, string(mockReadWriter.bytesWritten))
	assert.Contains(t, string(writer.Bytes()), noJournalConfigErrorMessage)
}

func Test_JournalConfigFileLogic_InvalidFileContent_PrintsMessage_ReturnsError(t *testing.T) {

	var mockReadWriter MockReadWriter
	mockReadWriter.shouldFailReading = false
	mockReadWriter.shouldReadInvalidFileContent = true

	var writer bytes.Buffer
	_, err := journalConfigFileLogic(&mockReadWriter, &writer)

	assert.Contains(t, string(writer.Bytes()), invalidJournalConfigErrorMessage)
	assert.NotNil(t, err)
}

func Test_JournalConfigFileLogic_HappyPath_LoadsJournals(t *testing.T) {
	var mockReadWriter MockReadWriter
	mockReadWriter.shouldFailReading = false
	mockReadWriter.shouldReadInvalidFileContent = false

	var writer bytes.Buffer
	journals, err := journalConfigFileLogic(&mockReadWriter, &writer)

	assert.Equal(t, "Other Journal", journals.Journals[1].Name)
	assert.Nil(t, err)
}

func Test_SanityCheck(t *testing.T) {

	var mockReadWriter MockReadWriter
	mockReadWriter.shouldFailReading = false
	mockReadWriter.shouldReadInvalidFileContent = false

	ioutil.ReadAll(mockReadWriter)
	//fmt.Println(string(b))
	//fmt.Println("WAT")
}
