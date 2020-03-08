package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"

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

func parseJournals(reader io.Reader) Journals {

	journalsReader := bufio.NewReader(reader)
	journalsBuf, _ := ioutil.ReadAll(journalsReader)

	journals := Journals{}

	err := yaml.Unmarshal([]byte(journalsBuf), &journals)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", journals)

	return journals
}
