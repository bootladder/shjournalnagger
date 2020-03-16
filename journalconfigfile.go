package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var journalConfigFilename = "/home/steve/.shjournalnagger/journals.yaml"

//JournalConfigFile is the real ReadWriter.  It is a thin layer of UNTESTED code
type JournalConfigFile struct {
	Editor                 string `yaml:"editor"`
	NaggingIntervalSeconds int    `yaml:"naggingintervalseconds"`
	Journals               []Journal
}

//Journal is inside the config
type Journal struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

func (j JournalConfigFile) Read(p []byte) (n int, err error) {
	f, err := os.Open(journalConfigFilename)
	if err != nil {
		return 0, err
	}

	journalsReader := bufio.NewReader(f)
	journalsBuf, _ := ioutil.ReadAll(journalsReader)
	copy(p, journalsBuf)
	return len(journalsBuf), io.EOF
}

func (j JournalConfigFile) Write(p []byte) (n int, err error) {

	if _, err := os.Stat(journalConfigFilename); os.IsNotExist(err) {
		err = os.MkdirAll("/home/steve/.shjournalnagger", 0777)
		if err != nil {
			fmt.Printf("WRITING ERROR: %v", err)
		}

		err = ioutil.WriteFile(journalConfigFilename, p, 0777)
		if err != nil {
			fmt.Printf("WRITING ERROR: %v", err)
		}

	}

	return
}

func (j JournalConfigFile) parseJournals() JournalConfigFile { return JournalConfigFile{} }

// This function is tested
func journalConfigFileLogic(journalConfigFile io.ReadWriter,
	output io.Writer) (JournalConfigFile, error) {

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
		return JournalConfigFile{}, errors.New("Invalid Config File")
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
