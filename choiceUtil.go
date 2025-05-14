package main

import (
	"sort"
)

func leaveTable(data cData) bool {
	printfLn("Player #%v: %v\nLeft the game.", data.player.Number, data.player.Name)
	data.player.Gone = true
	return true
}

func endTurn(data cData) bool {
	printfLn("Player #%v: %v\nHas ended their turn.\n", data.player.Number, data.player.Name)
	return true
}

func leaderboard(data cData) bool {
	var leaderBoard []leaderData
	for _, player := range data.game.Players {
		tmp := leaderData{Name: player.Name}

		stockVal := 0.0
		for _, stock := range player.Stocks {
			stockVal += roundToCent(data.game.Stocks[stock.StockID].Price * float64(stock.Shares))
		}
		tmp.StockVal = stockVal

		debts := 0.0
		for _, loan := range player.Loans {
			debts += loan.Principal
		}
		tmp.Debts = debts
		tmp.BankVal = player.Balance

		netWorth := stockVal + player.Balance - debts
		tmp.NetWorth = roundToCent(netWorth)
		leaderBoard = append(leaderBoard, tmp)
	}

	sort.Slice(leaderBoard, func(i, j int) bool {
		return leaderBoard[i].NetWorth > leaderBoard[j].NetWorth
	})

	printLn("\nLeaderboard:")
	for v, victim := range leaderBoard {
		if v > 0 {
			EnterKey(data.game, "")
		}
		printfLn("#%v %v\nStocks: %v, Bank: %v\n Debts: %v, Net: %v",
			v+1, victim.Name, victim.StockVal, victim.BankVal, victim.Debts, victim.NetWorth)
	}

	if data.game.Week == data.game.NumWeeks+1 {
		printfLn("\n%v won the game!", leaderBoard[0].Name)
	}
	return false
}
