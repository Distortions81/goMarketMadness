package main

import (
	"math"
	"math/rand"
)

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
		printfLn("APR %v%0.2f%% to %0.2f%%", trendSymbol[2], game.LastAPR-game.APR, game.APR)
	} else if game.APR > game.LastAPR {
		printfLn("APR %v%0.2f%% to %0.2f%%", trendSymbol[1], game.APR-game.LastAPR, game.APR)
	}

	game.APRHistory = append(game.APRHistory, game.APR)
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
