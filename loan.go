package main

func (loan loanData) totalPaid() float64 {
	amount := 0.0
	for _, payment := range loan.PaymentHistory {
		amount += payment
	}
	return roundToCent(amount)
}
