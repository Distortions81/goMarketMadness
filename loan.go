package main

import (
	"math"
)

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

		printfLn("Loan #%v: Payment: $%0.2f\nPrincipal: $%0.2f, Interest: $%0.2f", l+1, payment, player.Loans[l].Principal, interestForWeek)

		if player.Loans[l].Principal <= 0.01 {
			player.Loans[l].Principal = 0
			player.Loans[l].Complete = true
			printfLn("Loan #%v is paid off.", l+1)
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
