package main

import (
	"fmt"
	"math"
)

func (loan loanData) totalPaid() float64 {
	amount := 0.0
	for _, payment := range loan.PaymentHistory {
		amount += payment
	}
	return roundToCent(amount)
}

func takeLoan(game *gameData, player *playerData) {
	fmt.Printf("\nCurrent APR %0.2f%%\n", game.APR)

	maxLoan := calcMaxLoan(game, player)
	if maxLoan < 1.00 {
		fmt.Print("Sorry, the bank refuses to loan you any money.")
		return
	}

	maxLoan = roundToCent(maxLoan)
	loanAmount := promptForMoney("How much do you want to borrow?", maxLoan, 1.00, maxLoan)

	newLoan := loanData{Starting: loanAmount, Principal: loanAmount, APR: game.APR, StartWeek: game.week}
	player.Loans = append(player.Loans, newLoan)
}

func calcMaxLoan(game *gameData, player *playerData) float64 {
	stockAssets := 0.0
	debt := 0.0

	for _, stock := range player.Stocks {
		value := game.stocks[stock.StockID].Price
		value = roundToCent(value)
		stockAssets += (value * float64(stock.Shares))
	}

	for _, loan := range player.Loans {
		debt += loan.Principal
	}

	stockWeight := 0.60
	totalAssets := player.Money + (stockAssets * stockWeight)
	maxLoanAmount := totalAssets - debt
	maxLoanAmount = math.Max(maxLoanAmount, 0.0)

	return maxLoanAmount
}
