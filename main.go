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

	journalCommander := &JournalCommander{&ShellCommandExecuter{}, journalConfigFile}

	elapsedTimeChecker := &ElapsedTimeChecker{
		&RealLastNaggingTimeFileReader{},
		&RealLastNaggingTimeFileWriter{},
		&RealCurrentTimeGetter{},
	}

	shjournalnagger(os.Stdout, os.Stdin, journalConfigFile, journalCommander, elapsedTimeChecker)
}
