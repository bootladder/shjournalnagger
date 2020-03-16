package main

import "fmt"

func renderMenu(journals JournalConfigFile) string {
	menuStr := ""
	for i, journal := range journals.Journals {
		menuStr += fmt.Sprintf("[%d] : %s\n", i+1, journal.Name)
	}

	menuStr += "[e] : Edit Config\n"
	menuStr += "[q] : Exit\n\n"
	menuStr += "Select an Option: "
	return menuStr
}
