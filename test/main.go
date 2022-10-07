package main

import (
	"Hangman"
	"os"
)

func main() {
	args := os.Args[1:]
	Hangman.Game(args[0])
}
