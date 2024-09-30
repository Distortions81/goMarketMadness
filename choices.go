/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"fmt"
	"math"
)

var mainChoiceMenu []choiceData = []choiceData{
	{Name: "End turn", ChoiceFunc: endTurn},
	{Name: "Stocks", Submenu: stockChoices},
	{Name: "Banking", Submenu: bankChoices},
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

	ChoiceFunc func(data cData)
	Submenu    []choiceData
	Enabled    bool
}

func leaveTable(data cData) {
	fmt.Printf("Player #%v: (%v) has left the game.\n", data.player.Number, data.player.Name)
	data.player.Gone = true
}

func endTurn(data cData) {
	fmt.Printf("Player #%v: (%v) has ended their turn.\n", data.player.Number, data.player.Name)
}

func buyShares(data cData) {
	fmt.Printf("\nBuy shares of which stock?\n")

	//Print stock list
	maxLen := 0
	for _, stock := range data.game.Stocks {
		maxLen = maxInt(maxLen, len(stock.Name))
	}
	for s, stock := range data.game.Stocks {
		fmt.Printf("%v) %*v -- $%0.2f\n", s+1, maxLen, stock.Name, stock.Price)
	}

	choice := promptForInteger(false, 1, 1, len(data.game.Stocks), "Buy which stock?")
	choice -= 1
	maxAfford := math.Floor(data.player.Balance / data.game.Stocks[choice].Price)
	maxAfford = floorToCent(maxAfford)
	if maxAfford < 1 {
		fmt.Printf("You can't afford to buy any shares.")
		return
	}

	maxBuy := math.Min(data.game.getSettingFloat(SET_MAXSHARES), maxAfford)
	suggest := math.Min(10, maxBuy)

	numShares := promptForInteger(true, int(suggest), 1, int(maxBuy), "How many shares?")
	dollarValue := roundToCent(data.game.Stocks[choice].Price * float64(numShares))
	checkBalance(data)
	if promptForBool(false, "Buy %v shares of %v for $%0.2f?", numShares, data.game.Stocks[choice].Name, dollarValue) {
		data.player.debit(dollarValue)
		fmt.Printf("Debit: $%0.2f, New balance: $%0.2f\n", dollarValue, data.player.Balance)
		data.player.creditStock(data.game, choice, numShares)
	}
}

func sellShares(data cData) {
	fmt.Printf("\nSell shares of which stock?\n")

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
		fmt.Printf("%v) %*v -- (%v shares) $%0.2f\n", s+1,
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
		fmt.Printf("Credit: $%0.2f, New balance: $%0.2f\n", dollarValue, data.player.Balance)
		data.player.debitStock(stock.StockID, numShares)
	}
}

func displayShares(data cData) {

	count := 0
	fmt.Println()
	for _, stock := range data.player.Stocks {
		if stock.Shares <= 0 {
			continue
		}
		fmt.Printf("Stock %v, %v shares. Current value: $%0.2f",
			stock.Name, stock.Shares, float64(stock.Shares)*data.game.Stocks[stock.StockID].Price)
		count++
	}

	if count == 0 {
		fmt.Println("You don't have any stock shares at the moment.")
	}
	fmt.Println()
}
