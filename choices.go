/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"math"
	"sort"
)

var mainChoiceMenu []choiceData = []choiceData{
	{Name: "End turn", ChoiceFunc: endTurn},
	{Name: "Stocks", Submenu: stockChoices},
	{Name: "Banking", Submenu: bankChoices},
	{Name: "Leaderboard", ChoiceFunc: leaderboard},
	{Name: "Leave the game", ChoiceFunc: leaveTable},
}

var bankChoices []choiceData = []choiceData{
	{Name: "Diplay loans", ChoiceFunc: displayAllLoans},
	{Name: "Take out a loan", ChoiceFunc: takeLoan},
	{Name: "Make a payment on a loan", ChoiceFunc: payLoan},
	{Name: "See account balance", ChoiceFunc: checkBalance},
	{Name: "Go back"},
}

var stockChoices []choiceData = []choiceData{
	{Name: "Display shares", ChoiceFunc: displayShares},
	{Name: "Buy shares", ChoiceFunc: buyShares},
	{Name: "Sell shares", ChoiceFunc: sellShares},
	{Name: "Go back"},
}

type cData struct {
	player *playerData
	game   *gameData
}

type choiceData struct {
	Name,
	Desc string

	DefaultSetting any

	ChoiceFunc func(data cData) bool
	Submenu    []choiceData
	GoBack     bool
}

type leaderData struct {
	Name     string
	StockVal float64
	BankVal  float64
	Debts    float64
	NetWorth float64
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
		tmp.NetWorth = netWorth
		leaderBoard = append(leaderBoard, tmp)
	}

	sort.Slice(leaderBoard, func(i, j int) bool {
		return leaderBoard[i].NetWorth > leaderBoard[j].NetWorth
	})

	println("\nLeaderboard:")
	for v, victim := range leaderBoard {
		printfln("#%v -- %v: Stocks: %v,\nBank: %v, Debts: %v, Net: %v",
			v+1, victim.Name, victim.StockVal, victim.BankVal, victim.Debts, victim.NetWorth)
	}

	if data.game.Week == data.game.NumWeeks+1 {
		printfln("\n%v won the game!", leaderBoard[0].Name)
	}
	return false
}

func leaveTable(data cData) bool {
	printfln("Player #%v: %v\nLeft the game.", data.player.Number, data.player.Name)
	data.player.Gone = true
	return true
}

func endTurn(data cData) bool {
	printfln("Player #%v: %v\nHas ended their turn.", data.player.Number, data.player.Name)
	return true
}

func buyShares(data cData) bool {
	printfln("\nBuy which stock?")

	//Print stock list
	maxLen := 0
	for _, stock := range data.game.Stocks {
		maxLen = maxInt(maxLen, len(stock.Name))
	}
	for s, stock := range data.game.Stocks {
		printfln("%v) %*v $%0.2f", s+1, maxLen, stock.Name, stock.Price)
	}

	choice := promptForInteger(false, 1, 1, len(data.game.Stocks), "Buy which stock?")
	choice -= 1
	maxAfford := math.Floor(data.player.Balance / data.game.Stocks[choice].Price)
	maxAfford = floorToCent(maxAfford)
	if maxAfford < 1 {
		printfln("You can't afford any.")
		return false
	}

	maxBuy := math.Min(data.game.getSettingFloat(SET_MAXSHARES), maxAfford)
	suggest := math.Min(10, maxBuy)

	numShares := promptForInteger(true, int(suggest), 1, int(maxBuy), "How many shares?")
	dollarValue := roundToCent(data.game.Stocks[choice].Price * float64(numShares))
	checkBalance(data)
	if promptForBool(false, "Buy %v shares of\n%v for $%0.2f?", numShares, data.game.Stocks[choice].Name, dollarValue) {
		data.player.debit(dollarValue)
		printfln("Debit: $%0.2f, New balance: $%0.2f", dollarValue, data.player.Balance)
		data.player.creditStock(data.game, choice, numShares)
	}
	return false
}

func sellShares(data cData) bool {
	printfln("\nSell which stock?")

	//Print stock list
	maxLen := 0
	for _, stock := range data.player.Stocks {
		if stock.Shares <= 0 {
			continue
		}
		maxLen = maxInt(maxLen, len(stock.Name))
	}
	for s, stock := range data.player.Stocks {
		if stock.Shares <= 0 {
			continue
		}
		dollarValue := roundToCent(data.game.Stocks[stock.StockID].Price * float64(stock.Shares))
		printfln("%v) %*v %v shares $%0.2f", s+1,
			maxLen, stock.Name, stock.Shares, dollarValue)
	}

	choice := promptForInteger(false, 1, 1, len(data.game.Stocks), "Sell which stock?")
	choice -= 1
	stock := data.player.Stocks[choice]
	suggest := min(10, float64(stock.Shares))
	numShares := promptForInteger(true, int(suggest), 1, int(stock.Shares), "How many shares?")
	dollarValue := roundToCent(data.game.Stocks[stock.StockID].Price * float64(numShares))
	if promptForBool(false, "Sell %v shares of %v for $%0.2f?", numShares, stock.Name, dollarValue) {
		data.player.credit(dollarValue)
		printfln("Credit: $%0.2f\nNew balance: $%0.2f", dollarValue, data.player.Balance)
		data.player.debitStock(stock.StockID, numShares)
	}
	return false
}

func displayShares(data cData) bool {

	count := 0
	println("")
	for _, stock := range data.player.Stocks {
		if stock.Shares <= 0 {
			continue
		}
		printfln("Stock %v, %v shares.\nCurrent value: $%0.2f",
			stock.Name, stock.Shares, float64(stock.Shares)*data.game.Stocks[stock.StockID].Price)
		count++
	}

	if count == 0 {
		println("You have no stocks.")
	}
	return false
}
