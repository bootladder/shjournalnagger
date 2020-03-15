package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AlwaysNaggingTimeChecker struct{}

func (a *AlwaysNaggingTimeChecker) isNaggingIntervalExpired(interval int) bool { return true }
func (a *AlwaysNaggingTimeChecker) updateLastNaggingTime()                     {}

func Test_ReadsLineOfUserInput(t *testing.T) {
	var writer bytes.Buffer
	reader := bytes.NewReader([]byte("user input line 1\nuser input line 2\n"))

	shjournalnagger(&writer, reader, JournalConfigFile{}, nil, &AlwaysNaggingTimeChecker{})

	length := reader.Len()
	if length != 0 {
		t.Fatalf("Did not read the user input")
	}
}

func Test_PrintsTheTopPrompt(t *testing.T) {
	var writer bytes.Buffer
	var reader bytes.Buffer
	shjournalnagger(&writer, &reader, JournalConfigFile{}, nil, &AlwaysNaggingTimeChecker{})

	if false == bytes.Contains(writer.Bytes(), []byte("Write something!!!")) {
		t.Fatalf("Does not Print the Top Prompt")
	}
}

func Test_PrintsMenu(t *testing.T) {
	var writer bytes.Buffer
	var reader bytes.Buffer
	journals := getTestJournals()

	shjournalnagger(&writer, &reader, journals, nil, &AlwaysNaggingTimeChecker{})
	assert.Contains(t, string(writer.Bytes()), "Select")
}

func Test_InvalidUserInput_PrintsMessage(t *testing.T) {
	var writer bytes.Buffer
	reader := bytes.NewReader([]byte("user input line 1\nuser input line 2\n"))

	shjournalnagger(&writer, reader, JournalConfigFile{}, nil, &AlwaysNaggingTimeChecker{})

	assert.Contains(t, string(writer.Bytes()), "Invalid Input")
}

type MockJournalOpener struct {
	called    bool
	lastValue int
}

func (m *MockJournalOpener) openJournal(i int) {
	m.called = true
	m.lastValue = i
}

func Test_Quit_QuitsQuietly(t *testing.T) {
	var writer bytes.Buffer
	reader := bytes.NewReader([]byte("q\n"))

	shjournalnagger(&writer, reader, JournalConfigFile{}, nil, &AlwaysNaggingTimeChecker{})
	assert.NotContains(t, string(writer.Bytes()), "Invalid Input")
}

func Test_N_Journals_AcceptsUserInput_1_to_N(t *testing.T) {
	journalConfigFile := JournalConfigFile{}
	journalConfigFile.Journals = make([]Journal, 4)
	journalConfigFile.Journals[0] = Journal{Name: "journal1", Path: "path"}
	journalConfigFile.Journals[1] = Journal{Name: "journal2", Path: "path"}
	journalConfigFile.Journals[2] = Journal{Name: "journal3", Path: "path"}
	journalConfigFile.Journals[3] = Journal{Name: "journal4", Path: "path"}

	var writer bytes.Buffer
	reader := bytes.NewReader([]byte("3\n"))
	mockJournalOpener := MockJournalOpener{}
	shjournalnagger(&writer, reader, journalConfigFile, &mockJournalOpener, &AlwaysNaggingTimeChecker{})
	assert.Equal(t, 3, mockJournalOpener.lastValue)
}

func Test_N_Journals_DoesNotAccept_0_Input(t *testing.T) {
	journalConfigFile := JournalConfigFile{}
	journalConfigFile.Journals = make([]Journal, 1)
	journalConfigFile.Journals[0] = Journal{Name: "journal1", Path: "path"}

	var writer bytes.Buffer
	reader := bytes.NewReader([]byte("0\n"))
	mockJournalOpener := MockJournalOpener{}
	shjournalnagger(&writer, reader, journalConfigFile, &mockJournalOpener, &AlwaysNaggingTimeChecker{})
	assert.Equal(t, false, mockJournalOpener.called)
}
