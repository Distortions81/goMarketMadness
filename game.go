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
	"strings"
)

func (game *gameData) playGame() {

	//Prompt to create players
	numPlayers := len(game.players)
	if game.players == nil {
		numPlayers = promptForInteger("How many players?", 1, 1, maxPlayers)
		game.players = make([]*playerData, numPlayers)
	} else {
		result := promptForString("Play again with same players?", 1, 3, false)
		if strings.EqualFold(result, "n") || strings.EqualFold(result, "no") {
			numPlayers = promptForInteger("How many players?", 1, 1, maxPlayers)
			game.players = make([]*playerData, numPlayers)
		}
	}

	//Create players
	for p, player := range game.players {
		if player == nil {
			game.players[p] = &playerData{}
		}
	}

	//Init players, give starting money
	for p, player := range game.players {
		player.Number = p + 1
		if player.Name == "" {
			prompt := fmt.Sprintf("Name for player #%v", p+1)
			player.Name = promptForString(prompt, 2, maxPlayerNameLen, true)
		}
		player.Money = startingMoney
	}

	//Prompt for game length
	game.numWeeks = promptForInteger("How many weeks?", 12, minWeeks, maxWeeks)

	//Init stocks
	game.stocks = defaultStocks
	for s := range game.stocks {
		startPrice := RND()*10 + 2
		game.stocks[s].setPrice(startPrice)
		game.stocks[s].Volatility = RND() * (maxStartVolatility)
	}

	//Init APR
	game.APR = genLogRand(maxAPR-minAPR) + minAPR

	//Game loop
	game.tickStocks()
	for week := range game.numWeeks {
		game.week = week
		fmt.Printf("\n\n*** The %v week has begun! ***\n", numberNames[week])
		for p, player := range game.players {
			game.showStockPrices()
			fmt.Printf("\nPlayer #%v: (%v), it is your turn!\n", p+1, player.Name)
			fmt.Printf("Cash: $%0.2f", player.Money)
			if len(player.Loans) > 0 {
				fmt.Printf("Loans: ")
				for l, loan := range player.Loans {
					if loan.Complete {
						continue
					}
					player.loanCharges()
					fmt.Printf("Loan #%v: Loan Amount: %0.2f, Remaining: %0.2f, APR: %0.2f%%", l+1, loan.Starting, loan.Principal, loan.APR)
				}
			}
			fmt.Println("")
			promptForChoice(game, player, mainChoiceMenu)
		}
		game.tickStocks()
	}

	input := promptForString("Game over.\nStart a new game? (Y/n)", 1, 3, false)
	if strings.EqualFold(input, "y") || strings.EqualFold(input, "yes") {
		game.playGame()
	}
}
