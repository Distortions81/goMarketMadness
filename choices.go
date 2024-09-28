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
	{Name: "Buy/Sell stock", Submenu: buySellChoices},
	{Name: "Banking", Submenu: mainBankChoices},
	{Name: "Leave the table"},
}

var mainBankChoices []choiceData = []choiceData{
	{Name: "Take out a loan", ChoiceFunc: takeLoan},
	{Name: "Make a payment on a loan", ChoiceFunc: payLoan},
	{Name: "See balance", ChoiceFunc: checkBalance},
}

var buySellChoices []choiceData = []choiceData{
	{Name: "Buy shares", ChoiceFunc: buyShares},
	{Name: "Sell shares", ChoiceFunc: sellShares},
	{Name: "Go back"},
}

type choiceData struct {
	Name       string
	ChoiceFunc func(game *gameData, player *playerData)
	Submenu    []choiceData
	Enabled    bool
}

func endTurn(game *gameData, player *playerData) {
	fmt.Printf("Player #%v: (%v) has ended their turn.\n", player.Number, player.Name)
}

func buyShares(game *gameData, player *playerData) {
	fmt.Printf("Buy shares of which stock?\n")

	maxLen := 0
	for _, stock := range game.stocks {
		maxLen = max(maxLen, len(stock.Name))
	}
	for s, stock := range game.stocks {
		fmt.Printf("#%v %*v -- $%0.2f\n", s+1, maxLen, stock.Name, stock.Price)
	}

	choice := promptForInteger(1, 1, len(game.stocks), "Buy which stock?")
	maxAfford := math.Floor(player.Balance / game.stocks[choice].Price)
	maxAfford = floorToCent(maxAfford)
	if maxAfford < 1 {
		fmt.Printf("You can't afford to buy any shares.")
		return
	}

	maxBuy := math.Min(maxShares, maxAfford)
	suggest := math.Min(10, maxBuy)

	numShares := promptForInteger(int(suggest), 1, int(maxBuy), "How many shares?")
	dollarValue := roundToCent(game.stocks[choice].Price * float64(numShares))
	if promptForBool(false, "Buy %v shares of %v for $%0.2f?", numShares, game.stocks[choice].Name, dollarValue) {
		player.debit(dollarValue)
		checkBalance(game, player)
		//add stock to player
	}
}

func sellShares(game *gameData, player *playerData) {
}
