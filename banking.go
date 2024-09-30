/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"fmt"
	"math"
	"math/rand"
)

func checkBalance(data cData) {
	fmt.Printf("Available balance: $%0.2f\n", data.player.Balance)
}

func payLoan(data cData) {
	numLoans := data.player.getLoanCount()
	if numLoans == 0 {
		fmt.Println("You don't have any loans.")
		return
	}

	choice := 1
	if numLoans > 1 {
		for l, loan := range data.player.Loans {
			if loan.Principal <= 0 || loan.Complete {
				continue
			}
			loan.printLoan(l)
		}
		choice = promptForInteger(false, 1, 1, numLoans, "Which loan do you want to make a payment on?")
	}
	loan := data.player.Loans[choice-1]
	if loan.Principal <= 0 || loan.Complete {
		fmt.Println("That loan is already paid off.")
		return
	}
	loan.printLoan(choice)
	amount := promptForMoney("How much do you want to pay?", loan.Principal, math.Min(10, loan.Principal), math.Min(loan.Principal, data.player.Balance))
	data.player.debit(amount)
	loan.makeLoanPayment(amount)
	fmt.Printf("Made payment of $%0.2f\n", amount)
	loan.PaymentHistory = append(loan.PaymentHistory, amount)
}

func displayAllLoans(data cData) {
	count := 0
	for l, loan := range data.player.Loans {
		if loan.Complete || loan.Principal <= 0 {
			continue
		}
		count++
		loan.printLoan(l)
	}
	if count == 0 {
		fmt.Println("You do not have any loans!")
	}
}

func takeLoan(data cData) {
	fmt.Printf("Current APR %0.2f%%\n", data.game.APR)

	numLoans := data.player.getLoanCount()

	maxLoan := data.player.calcMaxLoan(data.game)
	if maxLoan < 1.00 || numLoans > data.game.getSettingInt(SET_MAXLOANNUM) {
		fmt.Print("Sorry, you have too many loans, the bank refuses.")
		return
	}

	maxLoan = roundToDollar(maxLoan)
	fmt.Printf("Maximum loan the bank will offer $%0.2f\n", maxLoan)
	loanAmount := promptForMoney("How much do you want to borrow?", maxLoan, 1.00, maxLoan)

	remainingWeeks := data.game.NumWeeks - data.game.Week - 1
	if remainingWeeks <= 1 {
		fmt.Println("There isn't enough time left in the game for a loan!")
		return
	}
	prompt := fmt.Sprintf("Loan term in weeks: 1-%v", remainingWeeks)
	loanTerm := promptForInteger(true, remainingWeeks, 1, remainingWeeks, prompt)

	newLoan := loanData{Starting: loanAmount, Principal: loanAmount, APR: data.game.APR, StartWeek: data.game.Week, TermWeeks: loanTerm}
	totalInterest := newLoan.calcTotalInterest()
	payments := newLoan.calcLoanPayment()

	if promptForBool(false, "Loan terms: Total interest: $%0.2f over %v weeks. Weekly payments: $%0.2f\nAccept", totalInterest, loanTerm, payments) {
		data.player.Loans = append(data.player.Loans, newLoan)
		data.player.credit(loanAmount)
	}
}

func (player *playerData) calcMaxLoan(game *gameData) float64 {
	stockAssets := 0.0
	debt := 0.0

	//Calculate stock values
	for _, stock := range player.Stocks {
		value := game.Stocks[stock.StockID].Price
		value = roundToCent(value)
		stockAssets += (value * float64(stock.Shares))
	}

	//Add up debt
	for _, loan := range player.Loans {
		debt += loan.Principal
	}

	stockWeight := 0.60
	totalAssets := player.Balance + (stockAssets * stockWeight)
	maxLoanAmount := totalAssets - debt
	maxLoanAmount = math.Max(maxLoanAmount, 0.0)

	return maxLoanAmount
}

func (player *playerData) loanCharges() {
	for l, loan := range player.Loans {
		if loan.Complete {
			continue
		}

		payment := loan.calcLoanPayment()

		weeklyInterestRate := loan.APR / 100 / 52
		interestForWeek := loan.Principal * weeklyInterestRate

		if loan.Principal <= payment {
			payment = loan.Principal + interestForWeek
		}

		if loan.Principal <= 0 || payment <= 0 {
			player.Loans[l].Complete = true
			player.Loans[l].Principal = 0
			continue
		}

		player.debit(payment)

		//TODO handle bankrupt here

		player.Loans[l].PaymentHistory = append(loan.PaymentHistory, payment)

		principalPayment := payment - interestForWeek

		player.Loans[l].makeLoanPayment(principalPayment)

		fmt.Printf("Loan #%v: Payment: $%0.2f, Principal: $%0.2f, Interest Charged: $%0.2f\n", l+1, payment, player.Loans[l].Principal, interestForWeek)

		if player.Loans[l].Principal <= 0.01 {
			player.Loans[l].Principal = 0
			player.Loans[l].Complete = true
			fmt.Printf("Loan #%v is now paid off.\n", l+1)
		}
	}
}

func (loan loanData) calcLoanPayment() float64 {
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

func (loan loanData) calcTotalInterest() float64 {
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

func (player *playerData) processLoans() int {

	player.loanCharges()

	total := 0
	for l := range player.Loans {
		if player.Loans[l].Complete {
			continue
		}
		total++
	}

	return total
}

func (game *gameData) tickAPR() {
	game.LastAPR = game.APR
	changePercent := 2 * game.getSettingFloat(SET_SIGAPR) * rand.Float64()
	change := 1 + (changePercent / 100)

	if rand.Float64() <= game.getSettingFloat(SET_APR_TREND) {
		game.TrendAPR = !game.TrendAPR
	}
	if game.TrendAPR {
		game.APR = (game.APR * change)
	} else {
		game.APR = (game.APR * (1 / change))
	}

	game.APR = math.Max(game.APR, game.getSettingFloat(SET_MINAPR))
	game.APR = math.Min(game.APR, game.getSettingFloat(SET_MAXAPR))
	game.APR = roundToCent(game.APR)

	if game.LastAPR > game.APR {
		fmt.Printf("APR %v%0.2f%% to %0.2f%%\n", trendSymbol[2], game.LastAPR-game.APR, game.APR)
	} else if game.APR > game.LastAPR {
		fmt.Printf("APR %v%0.2f%% to %0.2f%%\n", trendSymbol[1], game.APR-game.LastAPR, game.APR)
	}

	game.APRHistory = append(game.APRHistory, game.APR)
}

func (player *playerData) getLoanCount() int {
	count := 0
	for _, loan := range player.Loans {
		if loan.Complete {
			continue
		}
		count++
	}

	return count
}

func (loan *loanData) makeLoanPayment(amount float64) {
	loan.Principal -= amount
	loan.PaymentHistory = append(loan.PaymentHistory, amount)
	if loan.Principal <= 0 {
		loan.Principal = 0
		loan.Complete = true
	}
}

func (player *playerData) credit(income float64) {
	player.Balance = roundToCent(player.Balance + income)
}

func (player *playerData) debit(charge float64) bool {
	newAmount := roundToCent(player.Balance - charge)

	if newAmount <= 0 {
		player.Bankrupt = true
		return false
	}
	player.Balance = newAmount
	return true
}

func (loan loanData) printLoan(num int) {
	fmt.Printf("Loan #%v: Loan Amount: $%0.2f, Principal: $%0.2f, APR: %0.2f%%\n", num, loan.Starting, loan.Principal, loan.APR)
}
