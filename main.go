package main

import (
	"fmt"
	"os"
)

func main() {

	shjournalnagger(os.Stdout, os.Stdin, Journals{}, JournalCommander{})
	fmt.Println("hello world")
}
