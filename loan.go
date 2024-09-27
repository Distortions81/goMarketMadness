package main

import (
	"fmt"
	"math"
	"math/rand"
)

func checkBalance(game *gameData, player *playerData) {
	fmt.Printf("Available balance: $%0.2f\n", player.Balance)
}

func payLoan(game *gameData, player *playerData) {
	numLoans := player.getCount()
	if numLoans == 0 {
		fmt.Println("You don't have any loans.")
		return
	}

	choice := 1
	if numLoans > 1 {
		for l, loan := range player.Loans {
			if loan.Complete {
				continue
			}
			printLoan(l, loan)
		}
		choice = promptForInteger(1, 1, numLoans, "What loan do you want to make a payment on?")
	}
	loan := player.Loans[choice-1]
	if loan.Principal <= 0 || loan.Complete {
		fmt.Println("That loan is already paid off.")
		return
	}
	printLoan(choice, loan)
	amount := promptForMoney("How much do you want to pay?", loan.Principal, math.Min(10, loan.Principal), math.Min(loan.Principal, player.Balance))
	player.debit(amount)
	loan.makePayment(amount)
	fmt.Printf("Made payment of $%0.2f\n", amount)
	loan.PaymentHistory = append(loan.PaymentHistory, amount)
}

func printLoan(num int, loan loanData) {
	fmt.Printf("Loan #%v: Loan Amount: $%0.2f, Principal: $%0.2f, APR: %0.2f%%\n", num, loan.Starting, loan.Principal, loan.APR)
}

func takeLoan(game *gameData, player *playerData) {
	fmt.Printf("Current APR %0.2f%%\n", game.APR)

	numLoans := player.getCount()

	maxLoan := calcMaxLoan(game, player)
	if maxLoan < 1.00 || numLoans > maxLoanCount {
		fmt.Print("Sorry, the bank refuses to loan you any money.")
		return
	}

	maxLoan = roundToDollar(maxLoan)
	fmt.Printf("Maximum loan the bank will offer $%0.2f\n", maxLoan)
	loanAmount := promptForMoney("How much do you want to borrow?", maxLoan, 1.00, maxLoan)

	remainingWeeks := game.numWeeks - game.week - 1
	if remainingWeeks <= 0 {
		fmt.Println("There isn't enough time left in the game for a loan!")
		return
	}
	prompt := fmt.Sprintf("Loan term in weeks: 1-%v", remainingWeeks)
	loanTerm := promptForInteger(remainingWeeks, 1, remainingWeeks, prompt)

	newLoan := loanData{Starting: loanAmount, Principal: loanAmount, APR: game.APR, StartWeek: game.week, TermWeeks: loanTerm}
	totalInterest := calcTotalInterest(newLoan)
	payments := calcPayment(newLoan)

	if promptForBool(false, "Loan terms: Total interest: $%0.2f over %v weeks. Weekly payments: $%0.2f\nAccept", totalInterest, loanTerm, payments) {
		player.Loans = append(player.Loans, newLoan)
		player.credit(loanAmount)
	}
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

		payment := calcPayment(loan)

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

		//handle bankrupt

		player.Loans[l].PaymentHistory = append(loan.PaymentHistory, payment)

		principalPayment := payment - interestForWeek

		player.Loans[l].makePayment(principalPayment)

		fmt.Printf("Loan #%v: Payment: $%0.2f, Principal Reduction: $%0.2f, Interest Charged: $%0.2f\n", l+1, payment, principalPayment, interestForWeek)

		if player.Loans[l].Principal <= 0.01 {
			player.Loans[l].Principal = 0
			player.Loans[l].Complete = true
			fmt.Printf("Loan #%v is fully paid off.\n", l+1)
		}
	}
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

func processLoans(player *playerData) int {

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
	changePercent := 2 * volatilityAPR * RND()
	change := 1 + (changePercent / 100)
	if rand.Float64() > 0.5 {
		game.APR = (game.APR * change)
	} else {
		game.APR = (game.APR * (1 / change))
	}

	game.APR = math.Max(game.APR, minAPR)
	game.APR = math.Min(game.APR, maxAPR)

	//fmt.Printf("New interest rate: %0.2f%%\n", game.APR)
}

func (player *playerData) getCount() int {
	count := 0
	for _, loan := range player.Loans {
		if loan.Complete {
			continue
		}
		count++
	}

	return count
}

func (loan *loanData) makePayment(amount float64) {
	loan.Principal -= amount
	loan.PaymentHistory = append(loan.PaymentHistory, amount)
	if loan.Principal <= 0 {
		loan.Principal = 0
		loan.Complete = true
	}
}
