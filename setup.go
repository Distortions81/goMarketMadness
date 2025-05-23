package main

import (
	"fmt"
	"math/rand"
)

func (game *gameData) setup() {

	if promptForBool(game, false, "Customize game settings?") {
		//Init settings if needed
		if len(game.Settings) == 0 {
			game.Settings = defSettings
		}

		//Prompt for each setting
		for _, item := range game.Settings {
			if item.Hide {
				continue
			}
			input := promptForString(game, game.getSettingString(item.ID), 0, 64, false, "%v: [%v] ", item.Name, item.DefSetting)
			game.putSettingString(item.ID, input)
		}
	} else {
		//Use defaults
		game.Settings = defSettings
	}

	//Prompt to create players
	oldPlayers := game.Players
	numPlayers := len(game.Players)

	if game.Players == nil {
		game.promptNumPlayers()
	} else {
		if !promptForBool(game, false, "Play with same %v players?", numPlayers) {
			game.promptNumPlayers()
		}
	}
	game.createPlayerList(game.NumPlayers)

	//Create players
	for p, player := range game.Players {
		if player == nil {
			game.Players[p] = &playerData{Number: p + 1}
			//Transfer old player name, if exists
			if oldPlayers != nil && oldPlayers[p] != nil {
				game.Players[p].Name = oldPlayers[p].Name
			}
		}
	}

	//Init players
	for p := range game.Players {
		game.Players[p].Loans = []loanData{}
		if game.Players[p].Name == "" {
			//Prompt for name
			pName := fmt.Sprintf("Player-%v", p+1)
			game.Players[p].Name = promptForString(game, pName, 0, game.getSettingInt(SET_MAXNAMELEN), true, "Name for player #%v:", p+1)
		}
		//Give starting money
		game.Players[p].Balance = game.getSettingFloat(SET_STARTMONEY)
	}

	//Prompt for game length
	game.NumWeeks = promptForInteger(game, true, game.getSettingInt(SET_DEFAULT_WEEKS), 4, game.getSettingInt(SET_MAXWEEKS), "How many weeks?")

	//Init stocks
	game.Stocks = defaultStocks
	game.StockChoices = []choiceData{}
	for s := range game.Stocks {
		//Init stock choice list
		game.StockChoices = append(game.StockChoices, choiceData{Name: game.Stocks[s].Name})

		startPrice := rand.Float64()*10 + 2
		game.Stocks[s].TrendPrice = randBool()
		game.Stocks[s].TrendVolatility = randBool()
		game.Stocks[s].setPrice(startPrice)
		game.Stocks[s].LastPrice = game.Stocks[s].Price
		game.Stocks[s].Volatility = rand.Float64() * game.getSettingFloat(SET_MAXSIG)
	}

	//Init APR
	game.APR = game.genLogRand(
		game.getSettingFloat(SET_MAXAPR) -
			game.getSettingFloat(SET_MINAPR) +
			game.getSettingFloat(SET_MINAPR))
	game.APR = roundToCent(game.APR)
	game.TrendAPR = randBool()

}

func (game *gameData) promptNumPlayers() {
	game.NumPlayers = promptForInteger(game, true, 1, 1, game.getSettingInt(SET_MAXPLAYERS), "How many players?")
}

func (game *gameData) createPlayerList(numPlayers int) {
	game.Players = make([]*playerData, numPlayers)
}
