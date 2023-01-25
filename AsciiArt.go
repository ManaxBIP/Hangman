package Hangman

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func AsciiArt(file string) {
	attempts := 10
	var StockUserChoise []string
	var Redondant bool
	data, err := os.Open("/Hangman/words.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	data2, err2 := os.Open(file)
	if err2 != nil {
		log.Panicf("failed reading data from file: %s", err2)
	}
	var str []string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	fileScanner2 := bufio.NewScanner(data2)
	fileScanner2.Split(bufio.ScanLines)
	var linesAsciiArt []string
	for fileScanner2.Scan() {
		linesAsciiArt = append(linesAsciiArt, fileScanner2.Text())
	}
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	random := str[y1.Intn(len(str))]
	n := len(random)/2 - 1
	var randomSplitted []string
	var ToShow []string
	RandomRune := []rune(random)
	for i := range RandomRune {
		RandomRune[i] = RandomRune[i] - 32
	}
	var RandomUpper string
	for i := range RandomRune {
		RandomUpper += string(RandomRune[i])
	}
	randomSplitted = strings.Split(RandomUpper, "")
	for i := 0; i < n; i++ {
		randomIndex := rand.Intn(len(RandomRune))
		pick := RandomRune[randomIndex]
		if RandomRune[randomIndex] != 0 {
			ToShow = append(ToShow, string(pick))
		} else {
			i--
		}
		for k := range RandomRune {
			if RandomRune[k] == pick {
				RandomRune[k] = 0
			}
		}
	}
	res := make([]string, len(randomSplitted))
	for i := 0; i < len(randomSplitted); i++ {
		res[i] = "_"
	}
	for y := 0; y <= len(ToShow)-1; y++ {
		count := 0
		for _, i := range randomSplitted {
			if ToShow[y] == i {
				res[count] = i
			}
			count++
		}
	}
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
	print("Good Luck, you have 10 attempts.\n")
	var AsciiRep int
	StartArt := 307
	h := 0
	r := ""
	for _, i := range res {
		if i == "_" {
			h++
		}
	}
	for y := 120; y < 123; y++ {
		for i := h; i >= 0; i-- {
			r = linesAsciiArt[y]
			fmt.Print(r)
		}
		fmt.Print("\n")
	}
	for i := 0; i < 2; i++ {
		print("\n")
	}
	countFinish := 0
	for x := attempts; x > 0; x-- {
		AsciiRep = 1
		StartArt = 307
		Redondant = false
		if attempts < 1 {
			break
		}
		countFinish = 0
		for elm := range res {
			if res[elm] != "_" {
				countFinish++
			}
		}
		if countFinish == len(res) {
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
				Attempts: attempts,
				Word:     randomSplitted,
				Result:   res,
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
			for i := range randomSplitted {
				if UserChoice == randomSplitted[i] {
					res[i] = UserChoice
					x++
					found = true
				}
			}
			if found == true {
				for _, i := range res {
					if i == string('_') {
						for j := 120; j <= 123; j++ {
							fmt.Println(linesAsciiArt[j])
						}
					}
					for k := 'A'; k <= 'Z'; k++ {
						if i == string(k) && i == "A" {
							for x := 298; x <= 305; x++ {
								fmt.Println(linesAsciiArt[x])
							}
						} else if i == string(k) && i == "E" {
							for x := 334; x <= 340; x++ {
								fmt.Println(linesAsciiArt[x])
							}
						} else if i == string(k) && i == "I" {
							for x := 370; x <= 376; x++ {
								fmt.Println(linesAsciiArt[x])
							}
						} else if i == string(k) && i == "B" {
							for x := 307; x <= 313; x++ {
								fmt.Println(linesAsciiArt[x])
							}
						} else if i == string(k) && i == "C" {
							for x := 316; x <= 322; x++ {
								fmt.Println(linesAsciiArt[x])
							}
						} else if i == string(k) && i == "D" {
							for x := 325; x <= 331; x++ {
								fmt.Println(linesAsciiArt[x])
							}
						} else if i < string('I') && i != string('C') && i != string('B') && i != string('A') && i != string('E') && i != string('D') {
							if i == string(k) {
								AsciiRep = int(k)
								StartArt += (10 * (AsciiRep - 65)) - 16
								for x := StartArt; x <= StartArt+6; x++ {
									fmt.Println(linesAsciiArt[x])
								}
								StartArt = 307
							}
						} else {
							if i == string(k) {
								AsciiRep = int(k)
								StartArt += (9 * (AsciiRep - 65)) - 9
								for x := StartArt; x <= StartArt+6; x++ {
									fmt.Println(linesAsciiArt[x])
								}
								StartArt = 307
							}
						}
					}
					//print(" ")
				}
				for j := 0; j < 2; j++ {
					print("\n")
				}
			}
			if found == false {
				attempts--
				print("Not present in the word, ", attempts, " attempts remaining\n")
				if attempts == 9 {
					for i := 0; i < 8; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts == 8 {
					for i := 8; i < 16; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts == 7 {
					for i := 16; i < 24; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts == 6 {
					for i := 24; i < 32; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts == 5 {
					for i := 32; i < 40; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts == 4 {
					for i := 40; i < 48; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts == 3 {
					for i := 48; i < 56; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts == 2 {
					for i := 56; i < 64; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts == 1 {
					for i := 64; i < 72; i++ {
						fmt.Println(lines[i])
					}
				}
				if attempts < 1 {
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
				for _, i := range randomSplitted {
					StrRandomSlitted += i
				}
				if UserChoice == StrRandomSlitted {
					for i := range randomSplitted {
						res[i] = randomSplitted[i]
					}
					for _, i := range res {
						for k := 'A'; k <= 'Z'; k++ {
							if i == string(k) && i == "A" {
								for x := 298; x <= 305; x++ {
									fmt.Println(linesAsciiArt[x])
								}
							} else if i == string(k) && i == "E" {
								for x := 334; x <= 340; x++ {
									fmt.Println(linesAsciiArt[x])
								}
							} else if i == string(k) && i == "I" {
								for x := 370; x <= 376; x++ {
									fmt.Println(linesAsciiArt[x])
								}
							} else if i == string(k) && i == "B" {
								for x := 307; x <= 313; x++ {
									fmt.Println(linesAsciiArt[x])
								}
							} else if i == string(k) && i == "C" {
								for x := 316; x <= 322; x++ {
									fmt.Println(linesAsciiArt[x])
								}
							} else if i == string(k) && i == "D" {
								for x := 325; x <= 331; x++ {
									fmt.Println(linesAsciiArt[x])
								}
							} else if i < string('I') && i != string('C') && i != string('B') && i != string('A') && i != string('E') && i != string('D') {
								if i == string(k) {
									AsciiRep = int(k)
									StartArt += (10 * (AsciiRep - 65)) - 16
									for x := StartArt; x <= StartArt+6; x++ {
										fmt.Println(linesAsciiArt[x])
									}
									StartArt = 307
								}
							} else {
								if i == string(k) {
									AsciiRep = int(k)
									StartArt += (9 * (AsciiRep - 65)) - 9
									for x := StartArt; x <= StartArt+6; x++ {
										fmt.Println(linesAsciiArt[x])
									}
									StartArt = 307
								}
							}
						}
						//print(" ")
					}
					for j := 0; j < 2; j++ {
						print("\n")
					}
				} else {
					attempts -= 2
					if attempts < 0 {
						attempts = 0
					}
					print("Not present in the word, ", attempts, " attempts remaining\n")
					if attempts == 9 {
						for i := 0; i < 8; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts == 8 {
						for i := 8; i < 16; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts == 7 {
						for i := 16; i < 24; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts == 6 {
						for i := 24; i < 32; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts == 5 {
						for i := 32; i < 40; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts == 4 {
						for i := 40; i < 48; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts == 3 {
						for i := 48; i < 56; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts == 2 {
						for i := 56; i < 64; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts == 1 {
						for i := 64; i < 72; i++ {
							fmt.Println(lines[i])
						}
					}
					if attempts < 1 {
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
	if countFinish != len(res) && attempts < 1 {
		fmt.Println("You lose ! The result was ", RandomUpper, ".")
	}
}
