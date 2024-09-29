/***************
* STOCK MARKET *
****************/
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
	{Name: "Leave the game"},
}

var bankChoices []choiceData = []choiceData{
	{Name: "Diplay loans", ChoiceFunc: displayLoans},
	{Name: "Take out a loan", ChoiceFunc: takeLoan},
	{Name: "Make a payment on a loan", ChoiceFunc: payLoan},
	{Name: "See account balance", ChoiceFunc: checkBalance},
}

var stockChoices []choiceData = []choiceData{
	{Name: "Display shares", ChoiceFunc: displayShares},
	{Name: "Buy shares", ChoiceFunc: buyShares},
	{Name: "Sell shares", ChoiceFunc: sellShares},
	{Name: "Go back"},
}

type choiceData struct {
	Name,
	Desc string

	DefaultSetting any

	ChoiceFunc func(game *gameData, player *playerData)
	Submenu    []choiceData
	Enabled    bool
}

func endTurn(game *gameData, player *playerData) {
	fmt.Printf("Player #%v: (%v) has ended their turn.\n", player.Number, player.Name)
}

func buyShares(game *gameData, player *playerData) {
	fmt.Printf("\nBuy shares of which stock?\n")

	//Pretty-print list
	maxLen := 0
	for _, stock := range game.stocks {
		maxLen = max(maxLen, len(stock.Name))
	}
	for s, stock := range game.stocks {
		fmt.Printf("%v) %*v -- $%0.2f\n", s+1, maxLen, stock.Name, stock.Price)
	}

	choice := promptForInteger(1, 1, len(game.stocks), "Buy which stock?")
	maxAfford := math.Floor(player.Balance / game.stocks[choice].Price)
	maxAfford = floorToCent(maxAfford)
	if maxAfford < 1 {
		fmt.Printf("You can't afford to buy any shares.")
		return
	}

	maxBuy := math.Min(game.gGetFloat(SET_MAXSHARES), maxAfford)
	suggest := math.Min(10, maxBuy)

	numShares := promptForInteger(int(suggest), 1, int(maxBuy), "How many shares?")
	dollarValue := roundToCent(game.stocks[choice].Price * float64(numShares))
	checkBalance(game, player)
	if promptForBool(false, "Buy %v shares of %v for $%0.2f?", numShares, game.stocks[choice].Name, dollarValue) {
		player.debit(dollarValue)
		fmt.Printf("Debit: $%0.2f, New balance: $%0.2f\n", dollarValue, player.Balance)
		player.creditStock(game, choice, numShares)
	}
}

func sellShares(game *gameData, player *playerData) {
}

func displayShares(game *gameData, player *playerData) {

	count := 0
	fmt.Println()
	for _, stock := range player.Stocks {
		if stock.Shares > 0 {
			fmt.Printf("Stock %v, %v shares. Current value: $%0.2f",
				stock.Name, stock.Shares, float64(stock.Shares)*game.stocks[stock.StockID].Price)
			count++
		}
	}

	if count == 0 {
		fmt.Println("You don't have any stock shares at the moment.")
	}
	fmt.Println()
}
