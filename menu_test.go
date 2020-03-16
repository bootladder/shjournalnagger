package main

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RenderMenu(t *testing.T) {
	journalsFilename := "test-journals.yaml"
	f, _ := os.Open(journalsFilename)
	journalsReader := bufio.NewReader(f)
	journals, _ := parseJournals(journalsReader)

	menuStr := renderMenu(journals)

	assert.Contains(t, menuStr, "My Journal")
	assert.Contains(t, menuStr, "Other Journal")

	assert.Contains(t, menuStr, "1")
	assert.Contains(t, menuStr, "2")

	assert.Contains(t, menuStr, "[q] : Exit")
	assert.Contains(t, menuStr, "[e] : Edit Config")

	assert.Contains(t, menuStr, "Select an Option")
}
