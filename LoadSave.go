package Hangman

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type ElementSaved struct {
	Attempts int
	Word     []string
	Result   []string
}

func Load(save string) {
	data, err := ioutil.ReadFile(save)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	var SavedValue *ElementSaved
	json.Unmarshal(data, &SavedValue)
	Attempts := SavedValue.Attempts
	Word := SavedValue.Word
	Res := SavedValue.Result
	dataPosition, err := os.Open("hangman.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fileScanner := bufio.NewScanner(dataPosition)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	fmt.Println("Welcome Back, you have", Attempts, "attempts remaining.")
	for _, i := range Res {
		print(i)
		print(" ")
	}
	for i := 0; i < 2; i++ {
		print("\n")
	}
	var StockUserChoise []string
	var Redondant bool
	countFinish := 0
	for x := Attempts; x > 0; x-- {
		Redondant = false
		if Attempts < 1 {
			break
		}
		countFinish = 0
		for elm := range Res {
			if Res[elm] != "_" {
				countFinish++
			}
		}
		if countFinish == len(Res) {
			print("Congrats !")
			break
		}
		var UserChoice string
		found := false
		fmt.Print("Choose: ")
		fmt.Scan(&UserChoice)
		if UserChoice == "STOP" {
			type ElementSaved struct {
				Attempts int
				Word     []string
				Result   []string
			}
			ElmtSaved := ElementSaved{
				Attempts: Attempts,
				Word:     Word,
				Result:   Res,
			}
			save, err := json.Marshal(ElmtSaved)
			if err != nil {
				fmt.Println("error:", err)
			}
			fileSave, err := os.Create("save.txt")
			if err != nil {
				log.Fatal(err)
			}
			errWrite := ioutil.WriteFile("save.txt", save, 0777)
			if errWrite != nil {
				fmt.Println(errWrite)
			}
			defer fileSave.Close()
			fmt.Println("Game Saved in save.txt.")
			break
		}
		for _, i := range StockUserChoise {
			if UserChoice == i {
				Redondant = true
			}
		}
		if Redondant == false {
			StockUserChoise = append(StockUserChoise, UserChoice)
		}
		if len(UserChoice) == 1 && UserChoice > string(rune(64)) && UserChoice < string(rune(91)) && Redondant == false {
			for i := range Word {
				if UserChoice == Word[i] {
					Res[i] = UserChoice
					x++
					found = true
				}
			}
			if found == true {
				for _, i := range Res {
					print(i)
					print(" ")
				}
				for j := 0; j < 2; j++ {
					print("\n")
				}
			}
			if found == false {
				Attempts--
				print("Not present in the word, ", Attempts, " attempts remaining\n")
				if Attempts == 9 {
					for i := 0; i < 8; i++ {
						fmt.Println(lines[i])
					}
				}
				if Attempts == 8 {
					for i := 8; i < 16; i++ {
						fmt.Println(lines[i])
					}
				}
				if Attempts == 7 {
					for i := 16; i < 24; i++ {
						fmt.Println(lines[i])
					}
				}
				if Attempts == 6 {
					for i := 24; i < 32; i++ {
						fmt.Println(lines[i])
					}
				}
				if Attempts == 5 {
					for i := 32; i < 40; i++ {
						fmt.Println(lines[i])
					}
				}
				if Attempts == 4 {
					for i := 40; i < 48; i++ {
						fmt.Println(lines[i])
					}
				}
				if Attempts == 3 {
					for i := 48; i < 56; i++ {
						fmt.Println(lines[i])
					}
				}
				if Attempts == 2 {
					for i := 56; i < 64; i++ {
						fmt.Println(lines[i])
					}
				}
				if Attempts == 1 {
					for i := 64; i < 72; i++ {
						fmt.Println(lines[i])
					}
				}
				if Attempts < 1 {
					for i := 72; i < 80; i++ {
						fmt.Println(lines[i])
					}
				}
			}
		} else if len(UserChoice) > 1 && Redondant == false {
			count := 0
			for _, i := range UserChoice {
				if string(i) > string(rune(64)) && string(i) < string(rune(91)) {
					count++
				}
			}
			if count == len(UserChoice) {
				var StrRandomSlitted string
				for _, i := range Word {
					StrRandomSlitted += i
				}
				if UserChoice == StrRandomSlitted {
					for i := range Word {
						Res[i] = Word[i]
					}
					for _, i := range Res {
						print(i)
						print(" ")
					}
					for j := 0; j < 2; j++ {
						print("\n")
					}
				} else {
					Attempts -= 2
					if Attempts < 0 {
						Attempts = 0
					}
					print("Not present in the word, ", Attempts, " attempts remaining\n")
					if Attempts == 9 {
						for i := 0; i < 8; i++ {
							fmt.Println(lines[i])
						}
					}
					if Attempts == 8 {
						for i := 8; i < 16; i++ {
							fmt.Println(lines[i])
						}
					}
					if Attempts == 7 {
						for i := 16; i < 24; i++ {
							fmt.Println(lines[i])
						}
					}
					if Attempts == 6 {
						for i := 24; i < 32; i++ {
							fmt.Println(lines[i])
						}
					}
					if Attempts == 5 {
						for i := 32; i < 40; i++ {
							fmt.Println(lines[i])
						}
					}
					if Attempts == 4 {
						for i := 40; i < 48; i++ {
							fmt.Println(lines[i])
						}
					}
					if Attempts == 3 {
						for i := 48; i < 56; i++ {
							fmt.Println(lines[i])
						}
					}
					if Attempts == 2 {
						for i := 56; i < 64; i++ {
							fmt.Println(lines[i])
						}
					}
					if Attempts == 1 {
						for i := 64; i < 72; i++ {
							fmt.Println(lines[i])
						}
					}
					if Attempts < 1 {
						for i := 72; i < 80; i++ {
							fmt.Println(lines[i])
						}
					}
					x++
				}
			}
		} else if Redondant == true {
			fmt.Println("Already used ! ")
			x++
		}
	}
	if countFinish != len(Res) && Attempts < 1 {
		var WordToPrint string
		for _, i := range Word {
			WordToPrint += i
		}
		fmt.Println("You lose ! The result was ", WordToPrint, ".")
	}
}
