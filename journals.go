package main

import (
	"bufio"
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func parseJournals(reader io.Reader) (JournalConfigFile, error) {

	journalsReader := bufio.NewReader(reader)
	journalsBuf, _ := ioutil.ReadAll(journalsReader)

	journals := JournalConfigFile{}

	err := yaml.Unmarshal([]byte(journalsBuf), &journals)
	if err != nil {
		return JournalConfigFile{}, err
	}

	return journals, nil
}

func parseJournalsBytes(b []byte) (JournalConfigFile, error) {

	journals := JournalConfigFile{}

	err := yaml.Unmarshal(b, &journals)
	if err != nil {
		return JournalConfigFile{}, err
	}

	return journals, nil
}
