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

func Game() {
	data, err := os.Open("words.txt")
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

	fmt.Println(random)

	n := len(random)/2 - 1
	print(n)
	print("\n")
	var randomSplitted []string
	var ToShow []string
	randomSplitted = strings.Split(random, "")
	print(randomSplitted[0])
	print("\n")
	for i := 0; i < n; i++ {
		RandomRune := []rune(random)
		randomIndex := rand.Intn(len(RandomRune))
		pick := RandomRune[randomIndex]
		ToShow = append(ToShow, string(pick))
	}
	fmt.Println(ToShow)
	res := make([]string, len(randomSplitted))
	for i := 0; i < len(randomSplitted); i++ {
		res[i] = "_"
	}
	fmt.Println(res)
	for y := 0; y <= len(ToShow); y++ {
		count := 0
		for _, i := range randomSplitted {
			if ToShow[y] == i {
				res[count] = i
			}
			count++
		}
	}

	fmt.Println(res)

}
