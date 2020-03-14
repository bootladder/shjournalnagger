package main

import (
	"os"
)

func main() {

	journalConfigFile := JournalConfigFile{}
	journals, err := journalConfigFileLogic(journalConfigFile, os.Stdout)
	if err != nil {
		return
	}

	shjournalnagger(os.Stdout, os.Stdin, journals, JournalCommander{})
}
