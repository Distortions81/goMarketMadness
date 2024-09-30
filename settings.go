/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"strconv"
)

var (
	defaultStocks []stockData = []stockData{
		{Name: "US STEEL"}, {Name: "PAN AM"}, {Name: "FORD"}, {Name: "SANYO"}, {Name: "XEROX"}, {Name: "AT&T"},
	}
)

const (
	//Re-arrange or delete will break saves
	SET_MAXPLAYERS = iota
	SET_MAXWEEKS
	SET_DEFAULT_WEEKS
	SET_MAXNAMELEN
	SET_RANDLOG
	SET_MAXLOANNUM
	SET_MAXSHARES
	SET_STARTMONEY
	SET_MINSIG
	SET_MAXSIG
	SET_SIGSIG
	SET_SIGAPR
	SET_MAXLOAN
	SET_MINLOAN
	SET_MAXAPR
	SET_MINAPR
	SET_APR_TREND
	SET_VOL_TREND
	SET_STOCK_TREND
	SET_END
)

var defSettings = []settingsData{
	{Name: "Max game players", ID: SET_MAXPLAYERS, DefSetting: 25, Hide: true},
	{Name: "Max game weeks", ID: SET_MAXWEEKS, DefSetting: 52, Hide: true},
	{Name: "Default game weeks", ID: SET_DEFAULT_WEEKS, DefSetting: 24},
	{Name: "Max player name length", ID: SET_MAXNAMELEN, DefSetting: 64, Hide: true},
	{Name: "RNG logarithm ratio", ID: SET_RANDLOG, DefSetting: 100, Hide: true},
	{Name: "Max player loan count", ID: SET_MAXLOANNUM, DefSetting: 10},
	{Name: "Max buy shares", ID: SET_MAXSHARES, DefSetting: 10000},
	{Name: "Player starting money", ID: SET_STARTMONEY, DefSetting: 5000},
	{Name: "Min stock volatility", ID: SET_MINSIG, DefSetting: 1},
	{Name: "Max stock volatility", ID: SET_MAXSIG, DefSetting: 3},
	{Name: "Volatility volatility", ID: SET_SIGSIG, DefSetting: 5},
	{Name: "Apr volatility", ID: SET_SIGAPR, DefSetting: 2},
	{Name: "Max single loan amount", ID: SET_MAXLOAN, DefSetting: 1000000},
	{Name: "Min single loan amount", ID: SET_MINLOAN, DefSetting: 1000},
	{Name: "Max loan APR", ID: SET_MAXAPR, DefSetting: 19},
	{Name: "Min loan APR", ID: SET_MINAPR, DefSetting: 2.5},
	{Name: "APR trend change chance 0.01-1.0", ID: SET_APR_TREND, DefSetting: 0.2},
	{Name: "Volatility trend change chance 0.01-1.0", ID: SET_VOL_TREND, DefSetting: 0.2},
	{Name: "Stock trend change chance 0.01-1.0", ID: SET_STOCK_TREND, DefSetting: 0.2},
}

// Copy defaults to setting
func init() {
	for s := range defSettings {
		defSettings[s].Setting = defSettings[s].DefSetting
	}
}

func (game *gameData) getSettingInt(id int) int {
	for _, item := range game.Settings {
		if item.ID == id {
			val := item.Setting

			switch v := val.(type) {
			case int:
				return v
			case int64:
				return int(v)
			case string:
				vint, _ := strconv.ParseInt(v, 10, 64)
				return int(vint)
			case float64:
				return int(v)
			case float32:
				return int(v)
			}
		}
	}

	return -1
}

func (game *gameData) getSettingFloat(id int) float64 {
	for _, item := range game.Settings {
		if item.ID == id {
			val := item.Setting

			switch v := val.(type) {
			case int:
				return float64(v)
			case int64:
				return float64(v)
			case string:
				vint, _ := strconv.ParseFloat(v, 64)
				return vint
			case float64:
				return v
			case float32:
				return float64(v)
			}
		}
	}

	return -1
}

func (game *gameData) getSettingString(id int) string {
	for _, item := range game.Settings {
		if item.ID == id {
			switch v := game.Settings[id].Setting.(type) {
			case int:
				return strconv.FormatInt(int64(v), 10)
			case int64:
				return strconv.FormatInt(int64(v), 10)
			case string:
				return v
			case float64:
				return strconv.FormatFloat(v, 'f', -1, 64)
			case float32:
				return strconv.FormatFloat(float64(v), 'f', -1, 64)
			}
		}
	}

	return "Error"
}

func (game *gameData) putSettingString(id int, val string) {
	for _, item := range game.Settings {
		if item.ID == id {
			valType := item.Setting

			switch valType.(type) {
			case int:
				newVal, _ := strconv.ParseInt(val, 10, 64)
				game.Settings[id].Setting = newVal
			case int64:
				newVal, _ := strconv.ParseInt(val, 10, 64)
				game.Settings[id].Setting = newVal
			case string:
				game.Settings[id].Setting = val
			case float64:
				newVal, _ := strconv.ParseFloat(val, 64)
				game.Settings[id].Setting = newVal
			case float32:
				newVal, _ := strconv.ParseFloat(val, 64)
				game.Settings[id].Setting = newVal
			}
			return
		}
	}
}
