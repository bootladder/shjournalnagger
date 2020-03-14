package main

import (
	"bufio"
	"io"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

//Journals hello
type Journals struct {
	Dontcare string `yaml:"dontcare"`
	Journals []struct {
		Name string `yaml:"name"`
		Path string `yaml:"path"`
	}
}

func parseJournals(reader io.Reader) (Journals, error) {

	journalsReader := bufio.NewReader(reader)
	journalsBuf, _ := ioutil.ReadAll(journalsReader)

	journals := Journals{}

	err := yaml.Unmarshal([]byte(journalsBuf), &journals)
	if err != nil {
		return Journals{}, err
	}

	return journals, nil
}

func parseJournalsBytes(b []byte) (Journals, error) {

	journals := Journals{}

	err := yaml.Unmarshal(b, &journals)
	if err != nil {
		return Journals{}, err
	}

	return journals, nil
}
