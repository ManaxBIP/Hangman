package Hangman

func CheckUserInput(StockUser string, UserInput string) bool {
	var res bool
	for i := range StockUser {
		if UserInput == string(StockUser[i]) {
			res = true
		}
	}
	return res
}
