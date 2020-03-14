package main

import "fmt"

func renderMenu(journals Journals) string {
	menuStr := ""
	for i, journal := range journals.Journals {
		menuStr += fmt.Sprintf("[%d] : %s\n", i, journal.Name)
	}

	menuStr += "[q] : Exit\n\n"
	menuStr += "Select an Option: "
	return menuStr
}
