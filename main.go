package main

import (
	"os"
)

func main() {

	// this sucks
	journals, err := journalConfigFileLogic(JournalConfigFile{}, os.Stdout)
	if err != nil {
		return
	}

	shjournalnagger(os.Stdout, os.Stdin, journals, &JournalCommander{&ShellCommandExecuter{}, journals})
}
