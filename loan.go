package main

import (
	"fmt"
	"math"
	"strings"
)

func takeLoan(game *gameData, player *playerData) {
	fmt.Printf("\nCurrent APR %0.2f%%\n", game.APR)

	maxLoan := calcMaxLoan(game, player)
	if maxLoan < 1.00 {
		fmt.Print("Sorry, the bank refuses to loan you any money.")
		return
	}

	maxLoan = roundToCent(maxLoan)
	loanAmount := promptForMoney("How much do you want to borrow?", maxLoan, 1.00, maxLoan)

	remainingWeeks := game.numWeeks - game.week - 2
	if remainingWeeks <= 0 {
		fmt.Println("There isn't enough time left in the game for a loan!")
		return
	}
	prompt := fmt.Sprintf("Loan term in weeks: 1-%v", remainingWeeks)
	loanTerm := promptForInteger(prompt, remainingWeeks, 1, remainingWeeks)

	newLoan := loanData{Starting: loanAmount, Principal: loanAmount, APR: game.APR, StartWeek: game.week, TermWeeks: loanTerm}
	totalInterest := calcTotalInterest(newLoan)
	payments := calcPayment(newLoan) + calcInterest(newLoan)

	confirmLoan := fmt.Sprintf("Loan terms: Total interest: $%0.2f over %v weeks. Weekly payments: $%0.2f\nAccept (y/n)?", totalInterest, loanTerm, payments)
	response := promptForString(confirmLoan, 1, 3, false)
	if strings.EqualFold(response, "n") || strings.EqualFold(response, "no") {
		return
	}

	player.Loans = append(player.Loans, newLoan)

	player.credit(loanAmount)
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

func (player *playerData) loanCharges() {
	for l, loan := range player.Loans {
		if loan.Complete {
			continue
		}
		payment := calcPayment(loan)
		interest := calcInterest(loan)
		if loan.Principal <= 0 || payment <= 0 {
			player.Loans[l].Complete = true
			player.Loans[l].Principal = 0
			continue
		}

		player.debit(payment + interest)
		player.Loans[l].PaymentHistory = append(loan.PaymentHistory, payment)
		player.Loans[l].Principal -= payment
		fmt.Printf("\nLoan #%v: Payment: $%0.2f\n", l+1, payment)
	}
}

func calcInterest(loan loanData) float64 {
	weeklyInterestRate := loan.APR / 100 / 52
	interest := loan.Principal * weeklyInterestRate * float64(loan.TermWeeks)
	return roundToCent(interest)
}

func calcPayment(loan loanData) float64 {
	weeklyInterestRate := loan.APR / 100 / 52
	if weeklyInterestRate == 0 {
		return roundToCent(loan.Starting / float64(loan.TermWeeks))
	}

	numerator := weeklyInterestRate * loan.Starting
	denominator := 1 - (1 / math.Pow(1+weeklyInterestRate, float64(loan.TermWeeks)))
	weeklyPayment := numerator / denominator

	weeklyPayment = math.Min(weeklyPayment, loan.Principal)
	return roundToCent(weeklyPayment)
}

func calcTotalInterest(loan loanData) float64 {
	weeklyInterestRate := loan.APR / 100 / 52
	if weeklyInterestRate == 0 {
		return 0
	}

	numerator := weeklyInterestRate * loan.Starting
	denominator := 1 - (1 / math.Pow(1+weeklyInterestRate, float64(loan.TermWeeks)))
	weeklyPayment := numerator / denominator

	totalInterest := 0.0
	remainingPrincipal := loan.Starting

	// Loop through the weeks and accumulate total interest paid
	for week := 1; week <= loan.TermWeeks; week++ {
		weeklyInterest := remainingPrincipal * weeklyInterestRate
		totalInterest += weeklyInterest
		principalPayment := weeklyPayment - weeklyInterest
		remainingPrincipal -= principalPayment
	}

	return roundToCent(totalInterest)
}

func popLoanByIndex(loans []loanData, index int) ([]loanData, error) {
	if index < 0 || index >= len(loans) {
		return loans, fmt.Errorf("index out of range")
	}

	// Remove the loan by slicing out the element at the given index
	return append(loans[:index], loans[index+1:]...), nil
}

func processLoans(player *playerData) int {

	total := 0
	for l := range player.Loans {
		player.loanCharges()
		if player.Loans[l].Complete {
			continue
		}
		total++
	}

	return total
}
