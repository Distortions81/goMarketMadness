package main

import (
	"fmt"
	"math"
)

func (stock stockData) showTrend() string {
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
		printfLn(stock.showTrend())
	}
}

func (game *gameData) tickStocks() {
	for s := range game.Stocks {
		game.Stocks[s].tickStock(game)
	}
}

func displayShares(data cData) bool {

	count := 0
	printLn("")
	for _, stock := range data.player.Stocks {
		if stock.Shares <= 0 {
			continue
		}
		printfLn("Stock %v, %v shares.\nCurrent value: $%0.2f",
			stock.Name, stock.Shares, float64(stock.Shares)*data.game.Stocks[stock.StockID].Price)
		count++
	}

	if count == 0 {
		printLn("You have no stocks.")
	}
	return false
}

func buyShares(data cData) bool {
	printfLn("\nBuy which stock?")

	//Print stock list
	maxLen := 0
	for _, stock := range data.game.Stocks {
		maxLen = max(maxLen, len(stock.Name))
	}
	for s, stock := range data.game.Stocks {
		printfLn("%v) %*v $%0.2f", s+1, maxLen, stock.Name, stock.Price)
	}

	choice := promptForInteger(false, 1, 1, len(data.game.Stocks), "Buy which stock?")
	choice -= 1
	maxAfford := math.Floor(data.player.Balance / data.game.Stocks[choice].Price)
	maxAfford = floorToCent(maxAfford)
	if maxAfford < 1 {
		printfLn("You can't afford any.")
		return false
	}

	maxBuy := math.Min(data.game.getSettingFloat(SET_MAXSHARES), maxAfford)
	suggest := math.Min(10, maxBuy)

	numShares := promptForInteger(true, int(suggest), 1, int(maxBuy), "How many shares?")
	dollarValue := roundToCent(data.game.Stocks[choice].Price * float64(numShares))
	accBalance(data)
	if promptForBool(false, "Buy %v shares of\n%v for $%0.2f?", numShares, data.game.Stocks[choice].Name, dollarValue) {
		data.player.debit(dollarValue)
		printfLn("Debit: $%0.2f, New balance: $%0.2f", dollarValue, data.player.Balance)
		data.player.creditStock(data.game, choice, numShares)
	}
	return false
}

func sellShares(data cData) bool {
	printfLn("\nSell which stock?")

	//Print stock list
	maxLen := 0
	for _, stock := range data.player.Stocks {
		if stock.Shares <= 0 {
			continue
		}
		maxLen = max(maxLen, len(stock.Name))
	}
	for s, stock := range data.player.Stocks {
		if stock.Shares <= 0 {
			continue
		}
		dollarValue := roundToCent(data.game.Stocks[stock.StockID].Price * float64(stock.Shares))
		printfLn("%v) %*v %v shares $%0.2f", s+1,
			maxLen, stock.Name, stock.Shares, dollarValue)
	}

	choice := promptForInteger(false, 1, 1, len(data.game.Stocks), "Sell which stock?")
	choice -= 1
	stock := data.player.Stocks[choice]
	suggest := min(10, float64(stock.Shares))
	numShares := promptForInteger(true, int(suggest), 1, int(stock.Shares), "How many shares?")
	dollarValue := roundToCent(data.game.Stocks[stock.StockID].Price * float64(numShares))
	if promptForBool(false, "Sell %v shares of %v for $%0.2f?", numShares, stock.Name, dollarValue) {
		data.player.credit(dollarValue)
		printfLn("Credit: $%0.2f\nNew balance: $%0.2f", dollarValue, data.player.Balance)
		data.player.debitStock(stock.StockID, numShares)
	}
	return false
}
