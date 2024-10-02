package main

import (
	"fmt"
	"math"
)

func (stock stockData) showChange() string {
	buf := fmt.Sprintf("%v:", stock.Name)
	if stock.PriceArrow == TREND_UP || stock.PriceArrow == TREND_DOWN {
		buf = buf + fmt.Sprintf(" %v$%0.2f to", trendSymbol[stock.PriceArrow], math.Abs(stock.Price-stock.LastPrice))
	}
	buf = buf + fmt.Sprintf(" $%0.2f", stock.Price)

	return buf
}

func (game *gameData) showStockPrices() {
	printfLn("Stock prices: ")
	for _, stock := range game.Stocks {
		printfLn(stock.showChange())
	}
}

func (game *gameData) tickStocks() {
	for s := range game.Stocks {
		game.Stocks[s].tickStock(game)
	}
}
