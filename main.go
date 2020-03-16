package main

import (
	"os"
)

func main() {

	// this sucks
	journalConfigFile, err := journalConfigFileLogic(JournalConfigFile{}, os.Stdout)
	if err != nil {
		return
	}

	menuActions := &MenuActionsImpl{
		&ShellCommandExecuter{},
		journalConfigFile,
	}

	elapsedTimeChecker := &ElapsedTimeChecker{
		&RealLastNaggingTimeFileReader{},
		&RealLastNaggingTimeFileWriter{},
		&RealCurrentTimeGetter{},
	}

	shjournalnagger(os.Stdout, os.Stdin, journalConfigFile, menuActions, elapsedTimeChecker)
}
