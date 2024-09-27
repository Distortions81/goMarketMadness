/***************
* STOCK MARKET *
****************/
/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import "fmt"

var mainChoiceMenu []choiceData = []choiceData{
	{Name: "End turn", ChoiceFunc: endTurn},
	{Name: "Buy/Sell stock"},
	{Name: "Banking", Submenu: mainBankChoices},
	{Name: "Leave the table"},
}

var mainBankChoices []choiceData = []choiceData{
	{Name: "Take out a loan", ChoiceFunc: takeLoan},
	{Name: "Make a payment on a loan"},
	{Name: "See balance"},
}

type choiceData struct {
	Name       string
	ChoiceFunc func(game *gameData, player *playerData)
	Submenu    []choiceData
	Enabled    bool
}

func endTurn(game *gameData, player *playerData) {
	fmt.Printf("\nPlayer #%v: (%v) has ended their turn.\n", player.Number, player.Name)
}
