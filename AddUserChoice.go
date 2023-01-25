package Hangman

func AddUserChoice(StockUserChoice string, UserChoice string) string {
	res := StockUserChoice
	res += UserChoice
	return res
}
