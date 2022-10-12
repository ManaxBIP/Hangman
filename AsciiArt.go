package Hangman

import (
	"fmt"
	"io/ioutil"
	"log"
)

func AsciiArt(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Println(data)
}
