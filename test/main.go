package main

import (
	"Hangman"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 0 {
		Hangman.Game(args[0])
	}
}
