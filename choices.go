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
	{Name: "End turn"},
	{Name: "Buy/Sell stock"},
	{Name: "Banking"},
	{Name: "Leave the table"},
}

type choiceData struct {
	Name       string
	ChoiceFunc func(player playerData)
	Enabled    bool
}

func endTurn(player playerData) {
	fmt.Printf("Player #%v: (%v) has ended their turn.", player.Number, player.Name)
	return
}
