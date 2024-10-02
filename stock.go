/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"math"
	"math/rand"
)

func (stock *stockData) tickStock(game *gameData) {
	if stock.Bankrupt {
		return
	}

	stock.LastPrice = stock.Price
	stock.PriceHistory = append(stock.PriceHistory, stock.LastPrice)

	stock.tickVolatility(game)
	changePercent := 2 * stock.Volatility * rand.Float64()
	change := 1 + (changePercent / 100)

	if rand.Float64() <= game.getSettingFloat(SET_STOCK_TREND) {
		stock.TrendPrice = !stock.TrendPrice
	}

	if stock.TrendPrice {
		stock.setPrice(stock.LastPrice * change)
	} else {
		stock.setPrice(stock.LastPrice * (1 / change))
	}

	if stock.Price < 0.01 {
		stock.Price = 0
		stock.Bankrupt = true
	}

	if stock.LastPrice > stock.Price {
		stock.PriceArrow = TREND_DOWN
	} else if stock.LastPrice < stock.Price {
		stock.PriceArrow = TREND_UP
	} else {
		stock.PriceArrow = TREND_NONE
	}
}

func (stock *stockData) tickVolatility(game *gameData) {
	stock.LastVolatility = stock.Volatility
	stock.VolatilityHistory = append(stock.VolatilityHistory, stock.LastVolatility)

	changePercent := 2 * game.getSettingFloat(SET_SIGSIG) * rand.Float64()

	change := 1 + (changePercent / 100)

	if rand.Float64() <= game.getSettingFloat(SET_VOL_TREND) {
		stock.TrendVolatility = !stock.TrendVolatility
	}
	if stock.TrendVolatility {
		stock.Volatility = (stock.LastVolatility * change)
	} else {
		stock.Volatility = (stock.LastVolatility * (1 / change))
	}

	stock.Volatility = math.Max(stock.Volatility, game.getSettingFloat(SET_MINSIG))
	stock.Volatility = math.Min(stock.Volatility, game.getSettingFloat(SET_MAXSIG))
}

func (stock *stockData) setPrice(price float64) {
	stock.Price = roundToCent(price)
}

func (player *playerData) creditStock(game *gameData, stockNum, numShares int) {
	for s, stock := range player.Stocks {
		if stock.StockID == stockNum {
			player.Stocks[s].Shares += numShares
			return
		}
	}

	newStock := playerStockData{Name: game.Stocks[stockNum].Name, StockID: stockNum, Shares: numShares}
	player.Stocks = append(player.Stocks, newStock)
}

func (player *playerData) debitStock(stockNum, numShares int) bool {
	for s := range player.Stocks {
		if s == stockNum {
			if player.Stocks[s].Shares <= numShares {
				player.Stocks[s].Shares -= numShares
				return true
			}
			return false
		}
	}
	return false
}
