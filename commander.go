package main

// Commander sends commands
type Commander interface {
	command(int)
}
