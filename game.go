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

	"github.com/Distortions81/goCardinal"
)

func (game *gameData) playGame() {

	//Prompt to create players
	numPlayers := len(game.players)
	if game.players == nil {
		numPlayers = promptForInteger(1, 1, maxPlayers, "How many players?")
		game.players = make([]*playerData, numPlayers)
	} else {
		if !promptForBool(false, "Play again with same %v players?", numPlayers) {
			numPlayers = promptForInteger(1, 1, maxPlayers, "How many players?")
			game.players = make([]*playerData, numPlayers)
		}
	}

	choice := promptForBool(false, "Change game settings?")
	if choice {
		promptForChoice(game, nil, configChoices)
	} else {
		//Use defaults, or use previous settings.
	}
	oldPlayers := game.players

	//Create players
	for p, player := range game.players {
		if player == nil {
			game.players[p] = &playerData{}
			if oldPlayers[p] != nil {
				game.players[p].Name = oldPlayers[p].Name
			}
		}
	}

	//Init players, give starting money
	for p, player := range game.players {
		player.Number = p + 1
		player.Loans = []loanData{}
		if player.Name == "" {
			pName := fmt.Sprintf("Player #%v", player.Number)
			player.Name = promptForString(pName, 2, maxNameLen, true, "Name for player #%v:", p+1)
		}
		player.Balance = startingMoney
	}

	//Prompt for game length
	game.numWeeks = promptForInteger(12, minWeeks, maxWeeks, "How many weeks?")

	//Init stocks
	game.stocks = defaultStocks
	game.stockChoices = []choiceData{}
	for s := range game.stocks {
		game.stockChoices = append(game.stockChoices, choiceData{Name: game.stocks[s].Name})
		startPrice := RND()*10 + 2
		game.stocks[s].setPrice(startPrice)
		game.stocks[s].Volatility = RND() * (maxVolatility)
	}

	//Init APR
	game.APR = genLogRand(maxAPR-minAPR) + minAPR

	//Game loop
	game.tickStocks()
	for week := range game.numWeeks {
		game.week = week + 1
		fmt.Printf("\n*** The %v week has begun! ***\n", goCardinal.NumberToOrdinal(int64(game.week)))
		game.tickAPR()

		for p, player := range game.players {
			game.showStockPrices()
			fmt.Printf("\nPlayer #%v: (%v), it is your turn!\n", p+1, player.Name)
			processLoans(game.players[p])
			fmt.Printf("Bank balance: $%0.2f\n", player.Balance)
			promptForChoice(game, player, mainChoiceMenu)
		}
		game.tickStocks()
	}

	if promptForBool(false, "Game over!\nPlay again?") {
		game.playGame()
	}
}
