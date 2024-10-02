/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"math/rand"
	"os"
	"os/signal"
	"regexp"
	"syscall"
)

const (
	TREND_NONE = iota
	TREND_UP
	TREND_DOWN
	TREND_MAX
)

var trendSymbol [TREND_MAX]string = [TREND_MAX]string{
	"",
	"^",
	"v",
}

func handleExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(1)
	}()
}

func NumOnly(str string) string {
	alphafilter, _ := regexp.Compile("[^0-9.]+")
	str = alphafilter.ReplaceAllString(str, "")
	return str
}

func randBool() bool {
	return rand.Float64() <= 0.5
}

func getTrend(a, b float64) string {
	if a > b {
		return trendSymbol[1]
	} else if b > a {
		return trendSymbol[2]
	} else {
		return trendSymbol[0]
	}
}

func (game *gameData) promptNumPlayers() {
	game.NumPlayers = promptForInteger(true, 1, 1, game.getSettingInt(SET_MAXPLAYERS), "How many players?")
}

func (game *gameData) createPlayerList(numPlayers int) {
	game.Players = make([]*playerData, numPlayers)
}

func (game *gameData) showGameStats() {
	printfLn("Game over!\n\nSynopsis:")
	if game.APRHistory[0] < game.APR {
		printfLn("APR: %v$%0.2f: $%0.2f", trendSymbol[1], game.APR-game.APRHistory[0], game.APR)
	} else if game.APR < game.APRHistory[0] {
		printfLn("APR: %v$%0.2f: $%0.2f", trendSymbol[2], game.APRHistory[0]-game.APR, game.APR)
	} else {
		printfLn("APR: %v$%0.2f", trendSymbol[0], game.APR)
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
}
