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

	game.setup()

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

func (game *gameData) setup() {

	choice := promptForBool(false, "Customize game settings?")

	if len(game.settings) == 0 {
		game.settings = defSettings
	}

	if choice {
		for _, item := range game.settings {
			if item.hide {
				continue
			}
			input := promptForString(game.gGetString(item.id), 0, 64, false, "%v: (%v):", item.name, item.defSetting)
			game.gPutString(item.id, input)
		}
	} else {
		game.settings = defSettings
	}

	//Prompt to create players
	numPlayers := len(game.players)
	if game.players == nil {
		numPlayers = promptForInteger(true, 1, 1, game.gGetInt(SET_MAXPLAYERS), "How many players?")
		game.players = make([]*playerData, numPlayers)
	} else {
		if !promptForBool(false, "Play again with same %v players?", numPlayers) {
			numPlayers = promptForInteger(true, 1, 1, game.gGetInt(SET_MAXPLAYERS), "How many players?")
			game.players = make([]*playerData, numPlayers)
		}
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
	for p, _ := range game.players {
		game.players[p].Loans = []loanData{}
		if game.players[p].Name == "" {
			pName := fmt.Sprintf("Player-%v", p+1)
			game.players[p].Name = promptForString(pName, 0, game.gGetInt(SET_MAXNAMELEN), true, "Name for player #%v:", p+1)
		}
		game.players[p].Balance = game.gGetFloat(SET_STARTMONEY)
	}

	//Prompt for game length
	game.numWeeks = promptForInteger(true, 24, 4, game.gGetInt(SET_MAXWEEKS), "How many weeks?")

	//Init stocks
	game.stocks = defaultStocks
	game.stockChoices = []choiceData{}
	for s := range game.stocks {
		game.stockChoices = append(game.stockChoices, choiceData{Name: game.stocks[s].Name})
		startPrice := RND()*10 + 2
		game.stocks[s].setPrice(startPrice)
		game.stocks[s].Volatility = RND() * game.gGetFloat(SET_MAXSIG)
	}

	//Init APR
	game.APR = genLogRand(game,
		game.gGetFloat(SET_MAXAPR)-
			game.gGetFloat(SET_MINAPR)+
			game.gGetFloat(SET_MINAPR))
}
