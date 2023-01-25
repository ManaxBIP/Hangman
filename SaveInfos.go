package Hangman

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func Saveinfos(username string, score int) {
	fileSave, err := os.OpenFile("Scoreboard.txt", os.O_CREATE|os.O_APPEND, 0677)
	if err != nil {
		log.Fatal(err)
	}
	fileSave.WriteString(username + "\n")
	fileSave.WriteString(strconv.Itoa(score) + "\n")
	defer fileSave.Close()
	fmt.Println("Infos Saved in Scoreboard.txt.")
}
