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
	"math"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"syscall"
)

func setupTerm() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

func handleExit() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fixTerm()
		os.Exit(1)
	}()
}

func fixTerm() {
	exec.Command("stty", "-F", "/dev/tty", "sane").Run()
	fmt.Println("\nGame will now close.")
}

func showChange(stock stockData) string {
	buf := fmt.Sprintf("%v: $%0.2f", stock.Name, stock.Price)
	if stock.Trend == TREND_UP || stock.Trend == TREND_DOWN {
		buf = buf + fmt.Sprintf(" %v $%0.2f", trendSymbol[stock.Trend], math.Abs(stock.Price-stock.LastPrice))
	} else {
		buf = buf + fmt.Sprintf(" %v", trendSymbol[stock.Trend])
	}
	return buf
}

func (game *gameData) showStockPrices() {
	fmt.Print("Stock prices: ")
	for s, stock := range game.stocks {
		if s > 0 {
			fmt.Print(" -- ")
		}
		fmt.Printf(showChange(stock))
	}
}

func (game *gameData) tickStocks() {
	for s := range game.stocks {
		game.stocks[s].tickStock()
	}
}

func NumOnly(str string) string {
	alphafilter, _ := regexp.Compile("[^0-9.]+")
	str = alphafilter.ReplaceAllString(str, "")
	return str
}
