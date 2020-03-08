package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JournalsParser(t *testing.T) {
	journalsFilename := "test-journals.yaml"
	f, _ := os.Open(journalsFilename)
	journalsReader := bufio.NewReader(f)
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
