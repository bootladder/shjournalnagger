package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var journalsFilename = "/home/steve/.shjournalnagger/journals.yaml"

//JournalConfigFile is the real ReadWriter.  It is a thin layer of UNTESTED code
type JournalConfigFile struct {
	journals Journals
}

func (j JournalConfigFile) Read(p []byte) (n int, err error) {
	f, err := os.Open(journalsFilename)
	if err != nil {
		return 0, err
	}

	journalsReader := bufio.NewReader(f)
	journalsBuf, _ := ioutil.ReadAll(journalsReader)
	copy(p, journalsBuf)
	return 0, err
}
func (j JournalConfigFile) Write(p []byte) (n int, err error) {

	if _, err := os.Stat(journalsFilename); os.IsNotExist(err) {
		err = os.MkdirAll("/home/steve/.shjournalnagger", 0777)
		if err != nil {
			fmt.Printf("WRITING ERROR: %v", err)
		}

		err = ioutil.WriteFile(journalsFilename, p, 0777)
		if err != nil {
			fmt.Printf("WRITING ERROR: %v", err)
		}

	}

	return
}
func (j JournalConfigFile) parseJournals() Journals { return Journals{} }

// This function is tested
func journalConfigFileLogic(journalConfigFile io.ReadWriter,
	output io.Writer) (Journals, error) {

	// Open and read the file
	//b := make([]byte, 1000	)
	//_, err := journalConfigFile.Read(b)
	b, err := ioutil.ReadAll(journalConfigFile)

	// if doesnt exist, create it
	if err != nil {
		output.Write([]byte(noJournalConfigErrorMessage))
		journalConfigFile.Write([]byte(defaultJournalConfig))
		_, err = journalConfigFile.Read(b)
	}

	// if invalid, print error message
	journals, err := parseJournalsBytes(b)
	if err != nil {
		output.Write([]byte(invalidJournalConfigErrorMessage))
		return Journals{}, errors.New("Invalid Config File")
	}
	return journals, err
}

var defaultJournalConfig = `
---
dontcare: Hello
journals:
  - name: My Journal
    path: /tmp/journals/myjournal
  - name: Other Journal
    path: /tmp/journals/otherjournal

`

var noJournalConfigErrorMessage = `
No config file.  Creating a default config file at 
~/.shjournalnagger/journals.yaml
`

var invalidJournalConfigErrorMessage = `
Invalid config file.  Quitting.
Either fix the config file or delete it and a new one will be created next time
`
