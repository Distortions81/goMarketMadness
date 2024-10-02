/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"fmt"
	"math"
)

func accBalance(data cData) bool {
	printfLn("Available balance: $%0.2f", data.player.Balance)
	return false
}

func printLoans(data cData) bool {
	count := 0
	for l, loan := range data.player.Loans {
		if loan.Complete || loan.Principal <= 0 {
			continue
		}
		count++
		loan.printLoan(l)
	}
	if count == 0 {
		printLn("You do not have any loans!")
	}
	return false
}

func payLoan(data cData) bool {
	numLoans := data.player.getLoanCount()
	if numLoans == 0 {
		printLn("You don't have any loans.")
		return false
	}

	choice := 1
	if numLoans > 1 {
		for l, loan := range data.player.Loans {
			if loan.Principal <= 0 || loan.Complete {
				continue
			}
			loan.printLoan(l)
		}
		choice = promptForInteger(data.game, false, 1, 1, numLoans, "Which loan do you want to make a payment on?")
	}
	loan := data.player.Loans[choice-1]
	if loan.Principal <= 0 || loan.Complete {
		printLn("That loan is already paid off.")
		return false
	}
	loan.printLoan(choice)
	amount := promptForMoney(data.game, "How much do you want to pay?", loan.Principal, math.Min(10, loan.Principal), math.Min(loan.Principal, data.player.Balance))
	data.player.debit(amount)
	loan.makeLoanPayment(amount)
	printfLn("Made payment of $%0.2f", amount)
	loan.PaymentHistory = append(loan.PaymentHistory, amount)
	return false
}

func takeLoan(data cData) bool {
	printfLn("Current APR %0.2f%%", data.game.APR)

	numLoans := data.player.getLoanCount()

	maxLoan := data.player.maxLoanAmount(data.game)
	if maxLoan < 1.00 || numLoans > data.game.getSettingInt(SET_MAXLOANNUM) {
		printLn("Sorry, you have too many loans!")
		return false
	}

	maxLoan = roundToDollar(maxLoan)
	printfLn("Maximum loan $%0.2f", maxLoan)
	loanAmount := promptForMoney(data.game, "Borrow how much?", maxLoan, 1.00, maxLoan)

	remainingWeeks := data.game.NumWeeks - data.game.Week - 1
	if remainingWeeks <= 1 {
		printLn("Not enough time left!")
		return false
	}
	prompt := fmt.Sprintf("Loan term in weeks: 1-%v", remainingWeeks)
	loanTerm := promptForInteger(data.game, true, remainingWeeks, 1, remainingWeeks, prompt)

	newLoan := loanData{Starting: loanAmount, Principal: loanAmount, APR: data.game.APR, StartWeek: data.game.Week, TermWeeks: loanTerm}
	totalInterest := newLoan.calcTotalInterest()
	payments := newLoan.calcLoanPayment()

	if promptForBool(data.game, false, "Loan terms: Total interest: $%0.2f over %v weeks.\nWeekly payments: $%0.2f\nAccept", totalInterest, loanTerm, payments) {
		data.player.Loans = append(data.player.Loans, newLoan)
		data.player.credit(loanAmount)
	}
	return false
}

func (player *playerData) maxLoanAmount(game *gameData) float64 {
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
	printfLn("Loan #%v: Loan Amount: $%0.2f\nPrincipal: $%0.2f, APR: %0.2f%%", num, loan.Starting, loan.Principal, loan.APR)
}
