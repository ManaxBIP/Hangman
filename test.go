package Hangman

import "fmt"

func Test(word string, AvancePlayer string, UserInput string) string {
	var res string
	var InputIn bool
	var indice int
	var restemp []string
	var nbrUserIn int
	for i := 0; i < len(AvancePlayer); i++ {
		restemp = append(restemp, string(AvancePlayer[i]))
	}
	fmt.Println(restemp)
	for i := 0; i < len(word); i++ {
		if string(word[i]) == UserInput {
			InputIn = true
			indice = i
			nbrUserIn++
		}
	}
	if InputIn && nbrUserIn == 1 {
		if indice == 0 {
			if string(AvancePlayer[0]) == "_" {
				restemp[0] = UserInput
			}
		} else {
			if string(AvancePlayer[indice*2]) == "_" {
				restemp[indice*2] = UserInput
			}
		}
	}
	if InputIn && nbrUserIn > 1 {
		for i := 0; i < len(word); i++ {
			if string(word[i]) == UserInput {
				restemp[i*2] = UserInput
			}
		}
	}
	for i := 0; i < len(restemp); i++ {
		res += restemp[i]
	}
	fmt.Println(res)
	return res
}
