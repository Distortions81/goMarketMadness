package main

import "fmt"

func (loan loanData) totalPaid() float64 {
	amount := 0.0
	for _, payment := range loan.PaymentHistory {
		amount += payment
	}
	return roundToCent(amount)
}

func takeLoan(game *gameData, player *playerData) {
	fmt.Printf("Current APR %0.2f%%", game.APR)
	fmt.Printf("How much do you want to borrow?")
}
