/***************
* STOCK MARKET *
****************/
/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

type gameData struct {
	players        []*playerData
	apr, lastAPR   float64
	stocks         []stockData
	stockChoices   []choiceData
	settings       []settingsData
	trendAPR       bool
	week, numWeeks int
}

type settingsData struct {
	hide bool
	name string

	id int

	defSetting,
	setting any
}

type playerData struct {
	Name     string
	Number   int
	Balance  float64
	Bankrupt bool

	Stocks []playerStockData
	Loans  []loanData
}

type playerStockData struct {
	Name string
	StockID,
	Shares int
}

type loanData struct {
	Starting,
	Principal,
	APR float64

	StartWeek      int
	TermWeeks      int
	PaymentHistory []float64
	Complete       bool
}

type stockData struct {
	Name       string
	Price      float64
	PriceArrow int
	Volatility float64
	Bankrupt   bool

	LastVolatility float64
	LastPrice      float64

	PriceHistory      []float64
	VolatilityHistory []float64

	trendPrice,
	trendVolatility bool
}
