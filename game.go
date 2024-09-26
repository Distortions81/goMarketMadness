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

func playGame() {
	numPlayers = promptForInteger("How many players?", 1, 1, maxPlayers)

	//Init players, give starting money
	players = make([]playerData, numPlayers)
	for p := 0; p < numPlayers; p++ {
		players[p].Number = p + 1
		prompt := fmt.Sprintf("Name for player #%v", p+1)
		players[p].Name = promptForString(prompt, 2, maxPlayerNameLen, true)
		players[p].Money = startingMoney
	}

	//Prompt for game length
	numWeeks = promptForInteger("How many weeks?", 12, minWeeks, maxWeeks)

	//Init stocks
	fmt.Println("\nStocks init:")
	for s := range stockList {
		startPrice := RND()*10 + 2
		stockList[s].Price = startPrice
		stockList[s].Volatility = RND() * (maxStartVolatility)
	}

	//Game loop
	for week := range numWeeks {
		fmt.Printf("\n*** The %v week has begun! ***\n", numberNames[week])
		for p, player := range players {
			fmt.Printf("Player #%v: (%v), it is your turn:\n", p+1, player.Name)
			promptForChoice(player, mainChoiceMenu)
		}
	}

	input := promptForString("Game over.\nStart a new game? (Y/n)", 1, 3, false)
	if strings.EqualFold(input, "y") || strings.EqualFold(input, "yes") {
		playGame()
	}
}
