package Hangman

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Wordchosee(file string) string {
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
	//print("Good Luck, you have 10 attempts.\n")

	return RandomUpper

}
