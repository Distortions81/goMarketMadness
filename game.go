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
	"math/rand"

	"github.com/Distortions81/goCardinal"
)

func (game *gameData) playGame() {

	game.setup()

	//Game loop
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

	game.showGameStats()

	if promptForBool(false, "\nPlay again?") {
		game.playGame()
	}
}

func (game *gameData) setup() {

	if promptForBool(false, "Customize game settings?") {
		//Init settings if needed
		if len(game.settings) == 0 {
			game.settings = defSettings
		}

		//Prompt for each setting
		for _, item := range game.settings {
			if item.hide {
				continue
			}
			input := promptForString(game.getSettingString(item.id), 0, 64, false, "%v: (%v):", item.name, item.defSetting)
			game.putSettingString(item.id, input)
		}
	} else {
		//Use defaults
		game.settings = defSettings
	}

	//Prompt to create players
	numPlayers := len(game.players)
	if game.players == nil {
		game.promptNumPlayers()
		game.createPlayerList(game.numPlayers)
	} else {
		if !promptForBool(false, "Play again with same %v players?", numPlayers) {
			game.promptNumPlayers()
			game.players = make([]*playerData, numPlayers)
		}
	}
	oldPlayers := game.players

	//Create players
	for p, player := range game.players {
		if player == nil {
			game.players[p] = &playerData{}
			//Transfer old player name, if exists
			if oldPlayers[p] != nil {
				game.players[p].Name = oldPlayers[p].Name
			}
		}
	}

	//Init players
	for p := range game.players {
		game.players[p].Loans = []loanData{}
		if game.players[p].Name == "" {
			//Prompt for name
			pName := fmt.Sprintf("Player-%v", p+1)
			game.players[p].Name = promptForString(pName, 0, game.getSettingInt(SET_MAXNAMELEN), true, "Name for player #%v:", p+1)
		}
		//Give starting money
		game.players[p].Balance = game.getSettingFloat(SET_STARTMONEY)
	}

	//Prompt for game length
	game.numWeeks = promptForInteger(true, 52, 4, game.getSettingInt(SET_MAXWEEKS), "How many weeks?")

	//Init stocks
	game.stocks = defaultStocks
	game.stockChoices = []choiceData{}
	for s := range game.stocks {
		//Init stock choice list
		game.stockChoices = append(game.stockChoices, choiceData{Name: game.stocks[s].Name})

		startPrice := rand.Float64()*10 + 2
		game.stocks[s].trendPrice = randBool()
		game.stocks[s].trendVolatility = randBool()
		game.stocks[s].setPrice(startPrice)
		game.stocks[s].Volatility = rand.Float64() * game.getSettingFloat(SET_MAXSIG)
	}

	//Init APR
	game.apr = genLogRand(game,
		game.getSettingFloat(SET_MAXAPR)-
			game.getSettingFloat(SET_MINAPR)+
			game.getSettingFloat(SET_MINAPR))
	game.apr = roundToCent(game.apr)
	game.trendAPR = randBool()
}
