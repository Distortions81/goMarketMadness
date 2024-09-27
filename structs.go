/***************
* STOCK MARKET *
****************/
/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

const (
	TREND_NONE = iota
	TREND_UP
	TREND_DOWN
	TREND_MAX
)

var trendSymbol [TREND_MAX]string = [TREND_MAX]string{
	"→",
	"↑",
	"↓",
}

type playerData struct {
	Name   string
	Number int
	Money  int

	Stocks []playerStockData
	Loans  []loanData
}

type playerStockData struct {
	Name string
	StockID,
	Shares int
}

type loanData struct {
	StartWeek      int
	Interest       float64
	PaymentHistory []int
	Amount         int
}

type stockData struct {
	Name       string
	Price      float64
	Volatility float64
	Bankrupt   bool

	LastVolatility float64
	LastPrice      float64

	PriceHistory      []float64
	VolatilityHistory []float64

	Trend int
}
