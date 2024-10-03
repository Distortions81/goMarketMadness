/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"strings"
)

func (game *gameData) playGame() {

	countDown := 3
	for x := 01; x <= countDown; x++ {
		CallClear()
		printLn("Byte-99/4U")
		//		println("Copyright 2024 Carl Otto III")
		//		println("All rights reserved.")
		printfLn("\nLoading%v", strings.Repeat(".", x))
		//time.Sleep(time.Second)
	}
	CallClear()
	//time.Sleep(time.Millisecond * 500)
	printLn("Market Madness!")
	for x := 1; x < 8; x++ {
		CallBGColor(x)
		//time.Sleep(time.Millisecond * 200)
	}
	game.setup()

	//Game loop
	for week := range game.NumWeeks + 1 {
		game.Week = week + 1

		for p, player := range game.Players {
			if player.Gone {
				continue
			}
			game.showStockPrices()
			printfLn("\nPlayer #%v: %v\nIt is your turn!", p+1, player.Name)
			game.Players[p].processLoans()
			printfLn("Bank balance: $%0.2f", player.Balance)
			promptForChoice(game, player, mainChoiceMenu)
		}

		if game.Week == game.NumWeeks {
			printLn("** LAST WEEK!!! ***")
		} else if game.Week < game.NumWeeks {
			printfLn("*** WEEK %v of %v ***", game.Week, game.NumWeeks)
		}
		game.tickStocks()
		game.tickAPR()
	}

	game.showGameStats()
	game.playGame()
}

func (game *gameData) showGameStats() {
	printfLn("Game over!\n\nSynopsis:")
	if game.APRHistory[0] < game.APR {
		printfLn("APR: %v%0.2f%%: $%0.2f", trendSymbol[1], game.APR-game.APRHistory[0], game.APR)
	} else if game.APR < game.APRHistory[0] {
		printfLn("APR: %v%0.2f%%: $%0.2f", trendSymbol[2], game.APRHistory[0]-game.APR, game.APR)
	} else {
		printfLn("APR: %v%0.2f%%", trendSymbol[0], game.APR)
	}

	for _, stock := range game.Stocks {
		if stock.PriceHistory[0] < stock.Price {
			printfLn("%v: %v$%0.2f: $%0.2f", stock.Name, trendSymbol[1], stock.Price-stock.PriceHistory[0], stock.Price)
		} else if stock.Price < stock.PriceHistory[0] {
			printfLn("%v: %v$%0.2f: $%0.2f", stock.Name, trendSymbol[2], stock.PriceHistory[0]-stock.Price, stock.Price)
		} else {
			printfLn("%v: %v$%0.2f", stock.Name, trendSymbol[0], stock.Price)
		}
	}

	game.Week++
	leaderboard(cData{game: game, player: nil})
	EnterKey(game, "")
}
