package main

import (
	"Hangman"
	"flag"
	"os"
)

func main() {
	var saveLoaded bool
	var AsciiArt bool
	args := os.Args[1:]
	flag.BoolVar(&saveLoaded, "startWith", false, "start with save")
	flag.BoolVar(&AsciiArt, "letterFile", false, "start with Ascii Art")
	flag.Parse()
	if saveLoaded {
		Hangman.Load(args[1])

	} else if AsciiArt {
		Hangman.AsciiArt(args[1])
	} else if len(args) != 0 {
		Hangman.Game(args[0])
	}
}
