package Hangman

import (
	"math/rand"
	"strings"
)

func PrintLetter(word string) string {
	var result string
	n := len(word)/2 - 1
	var randomSplitted []string
	var ToShow []string
	RandomRune := []rune(word)
	var RandomUpper string
	for i := range RandomRune {
		RandomUpper += string(RandomRune[i])
		//print(string(RandomUpper[i]))
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
				break
			}
			count++
		}
	}
	for i := 0; i < len(res); i++ {
		result += res[i]
		result += " "
	}
	print(result)
	return result
}
