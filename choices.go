/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

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
