package Hangman

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Game(file string) {
	attempts := 10
	data, err := os.Open(file)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	var str []string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	random := str[y1.Intn(len(str))]

	n := len(random)/2 - 1
	var randomSplitted []string
	var ToShow []string
	randomSplitted = strings.Split(random, "")
	RandomRune := []rune(random)
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
	print("Good Luck, you have 10 attempts.\n")
	for _, i := range res {
		print(i)
		print(" ")
	}
	for i := 0; i < 2; i++ {
		print("\n")
	}
	countFinish := 0
	for x := attempts; x > 0; x-- {
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
		if len(UserChoice) == 1 && UserChoice > string(rune(96)) && UserChoice < string(rune(123)) {
			for i := range randomSplitted {
				if UserChoice == randomSplitted[i] {
					res[i] = UserChoice
					x++
					found = true
					for _, i := range res {
						print(i)
						print(" ")
					}
					for j := 0; j < 2; j++ {
						print("\n")
					}
				}
			}
			if found == false {
				print("Not present in the word, ", x-1, " attempts remaining\n")
				for i := 0; i < 10; i++ {
					print("\n")
				}
			}
		} else {
			x++
		}
	}
	if countFinish != len(res) {
		print("You lose ! The result was ", random, ".")
	}
}
