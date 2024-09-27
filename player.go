package main

import "fmt"

func (player *playerData) credit(income float64) {
	player.Money = roundToCent(player.Money + income)
}

func (player *playerData) debit(income float64) {
	newAmount := roundToCent(player.Money - income)
	if newAmount <= 0 {
		player.Bankrupt = true
		player.Money = 0
	}
}

func (player *playerData) loanCharges() {
	for l, loan := range player.Loans {
		if loan.Complete {
			continue
		}
		interest := calcInterest(float64(loan.Principal), loan.APR)
		if interest <= 0 || loan.Principal <= 0 {
			player.Loans[l].Complete = true
			player.Loans[l].Principal = 0
			continue
		}

		player.debit(interest)
		fmt.Printf("\nLoan #%v: Charged interest: $%0.2f\n", l+1, interest)
	}
}
