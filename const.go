/***************
* STOCK MARKET *
****************/
/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"strconv"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
)

const (
	SET_MAXPLAYERS = iota
	SET_MAXWEEKS
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
	SET_END
)

var defSettings = []settingsData{
	{name: "max players", id: SET_MAXPLAYERS, defSetting: 25},
	{name: "max weeks", id: SET_MAXWEEKS, defSetting: 52},
	{name: "max name length", id: SET_MAXNAMELEN, defSetting: 64},
	{name: "RNG logarithm ratio", id: SET_RANDLOG, defSetting: 100},
	{name: "max loan number", id: SET_MAXLOANNUM, defSetting: 10},
	{name: "max shares", id: SET_MAXSHARES, defSetting: 10000},
	{name: "starting money", id: SET_STARTMONEY, defSetting: 5000},
	{name: "min volatility", id: SET_MINSIG, defSetting: 3},
	{name: "max volatility", id: SET_MAXSIG, defSetting: 10},
	{name: "volatility volatility", id: SET_SIGSIG, defSetting: 10},
	{name: "apr volatility", id: SET_SIGAPR, defSetting: 10},
	{name: "max loan amount", id: SET_MAXLOAN, defSetting: 1000000},
	{name: "min loan amount", id: SET_MINLOAN, defSetting: 1000},
	{name: "max apr", id: SET_MAXAPR, defSetting: 19},
	{name: "min apr", id: SET_MINAPR, defSetting: 2.5},
}

// Copy defaults to setting
func init() {
	for s := range defSettings {
		defSettings[s].setting = defSettings[s].defSetting
	}
}

func (game *gameData) gGetInt(id int) int {
	for _, item := range game.settings {
		if item.id == id {
			val := item.setting

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

func (game *gameData) gGetFloat(id int) float64 {
	for _, item := range game.settings {
		if item.id == id {
			val := item.setting

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

func (game *gameData) gGetString(id int) string {
	for _, item := range game.settings {
		if item.id == id {
			switch v := game.settings[id].setting.(type) {
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

func (game *gameData) gPutString(id int, val string) {
	for _, item := range game.settings {
		if item.id == id {
			valType := item.setting

			switch valType.(type) {
			case int:
				newVal, _ := strconv.ParseInt(val, 10, 64)
				game.settings[id].setting = newVal
			case int64:
				newVal, _ := strconv.ParseInt(val, 10, 64)
				game.settings[id].setting = newVal
			case string:
				game.settings[id].setting = val
			case float64:
				newVal, _ := strconv.ParseFloat(val, 64)
				game.settings[id].setting = newVal
			case float32:
				newVal, _ := strconv.ParseFloat(val, 64)
				game.settings[id].setting = newVal
			}
			return
		}
	}
}

var (
	defaultStocks []stockData = []stockData{
		{Name: "US STEEL"}, {Name: "PAN AM"}, {Name: "FORD"}, {Name: "SANYO"}, {Name: "XEROX"}, {Name: "AT&T"},
	}
	games []gameData

	sr beep.SampleRate
)

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

func init() {
	sr = beep.SampleRate(44000)
	speaker.Init(sr, 4800)
}
