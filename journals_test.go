package main

import (
	"bufio"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JournalsParser(t *testing.T) {
	journalsReader := getTestJournalsConfigReader()
	journals := parseJournals(journalsReader)
	name1 := journals.Journals[0].Name
	path1 := journals.Journals[0].Path
	name2 := journals.Journals[1].Name
	path2 := journals.Journals[1].Path

	assert.Equal(t, name1, "My Journal")
	assert.Equal(t, name2, "Other Journal")
	assert.Equal(t, path1, "/tmp/journals/myjournal")
	assert.Equal(t, path2, "/tmp/journals/otherjournal")
}

func getTestJournalsConfigReader() io.Reader {
	journalsFilename := "test-journals.yaml"
	f, _ := os.Open(journalsFilename)
	return bufio.NewReader(f)
}

func getTestJournals() Journals {
	journalsReader := getTestJournalsConfigReader()
	return parseJournals(journalsReader)
}
