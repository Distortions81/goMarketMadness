/***************
* STOCK MARKET *
****************/
/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"math"
	"math/rand"
)

func (stock *stockData) tickStock() {
	if stock.Bankrupt {
		return
	}

	stock.LastPrice = stock.Price
	stock.PriceHistory = append(stock.PriceHistory, stock.LastPrice)

	stock.tickVolatility()
	changePercent := 2 * stock.Volatility * RND()
	change := 1 + (changePercent / 100)

	if rand.Float64() > 0.5 {
		stock.Price = (stock.LastPrice * change)
	} else {
		stock.Price = (stock.LastPrice * (1 / change))
	}

	if stock.Price < 0.01 {
		stock.Price = 0
		stock.Bankrupt = true
	}
}

func (stock *stockData) tickVolatility() {
	stock.LastVolatility = stock.Volatility
	stock.VolatilityHistory = append(stock.VolatilityHistory, stock.LastVolatility)

	changePercent := 2 * volatilityVolatility * RND()

	change := 1 + (changePercent / 100)
	if rand.Float64() > 0.5 {
		stock.Volatility = (stock.LastVolatility * change)
	} else {
		stock.Volatility = (stock.LastVolatility * (1 / change))
	}

	stock.Volatility = math.Max(stock.Volatility, minVolatility)
	stock.Volatility = math.Min(stock.Volatility, maxVolatility)
}
