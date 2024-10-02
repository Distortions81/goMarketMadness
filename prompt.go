/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"strconv"
	"strings"
)

func promptForString(game *gameData, defaultAnswer string, min, max int, confirm bool, format string, args ...interface{}) string {
	game.showCursor = true
	defer func() { game.showCursor = false }()

	printfLn(format, args...)

	line := <-newInput
	if line == "" {
		line = defaultAnswer
	}
	lLen := len(line)
	if lLen < min {
		printfLn("At least %v characters.", min)
		return promptForString(game, defaultAnswer, min, max, confirm, format, args...)
	} else if lLen > max {
		printfLn("Only %v characters max.", max)
		return promptForString(game, defaultAnswer, min, max, confirm, format, args...)
	}

	if confirm {
		if promptForBool(game, true, "Confirm: [%v]", line) {
			if line == "" {
				return defaultAnswer
			}
			return line
		} else {
			promptForString(game, defaultAnswer, min, max, confirm, format, args...)
		}
	}

	return line
}

func promptForBool(game *gameData, defaultYes bool, format string, args ...interface{}) bool {
	game.showCursor = true
	defer func() { game.showCursor = false }()

	question := ""
	if defaultYes {
		question = " [Y/n]"
	} else {
		question = " [y/N]"
	}
	result := promptForString(game, "", 0, 3, false, format+question, args...)

	if result == "" {
		return defaultYes
	} else if strings.EqualFold(result, "n") || strings.EqualFold(result, "no") {
		return false
	} else if strings.EqualFold(result, "y") || strings.EqualFold(result, "yes") {
		return true
	} else {
		printLn("Yes or no y/n?")
		return promptForBool(game, defaultYes, format, args...)
	}
}

func promptForInteger(game *gameData, useDefault bool, defaultVal, min, max int, prompt string) int {
	game.showCursor = true
	defer func() { game.showCursor = false }()

	if useDefault {
		printfLn("%v %v-%v [%v]", prompt, min, max, defaultVal)
	} else {
		printfLn("%v %v-%v", prompt, min, max)
	}

	line := <-newInput
	if useDefault && line == "" {
		return defaultVal
	}
	value, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		printLn("Must be a number.")
		return promptForInteger(game, useDefault, defaultVal, min, max, prompt)
	}
	if int(value) < min || int(value) > max {
		printfLn("Must be %v to %v.", min, max)
		return promptForInteger(game, useDefault, defaultVal, min, max, prompt)
	}

	return int(value)
}

func promptForMoney(game *gameData, prompt string, defaultVal, min, max float64) float64 {
	game.showCursor = true
	defer func() { game.showCursor = false }()

	printfLn("%v $%0.2f-%0.2f: $%0.2f", prompt, min, max, defaultVal)

	line := <-newInput
	if line == "" {
		return defaultVal
	}
	line = strings.TrimSpace(line)
	line = NumOnly(line)
	value, err := strconv.ParseFloat(line, 64)
	if err != nil {
		printLn("That isn't a number.")
		return promptForMoney(game, prompt, defaultVal, min, max)
	}
	value = roundToCent(value)
	min = roundToCent(min)
	max = roundToCent(max)

	if value < min || value > max {
		printfLn("Must be $%0.2f-%0.2f.", min, max)
		return promptForMoney(game, prompt, defaultVal, min, max)
	}

	return value
}

func promptForChoice(game *gameData, player *playerData, options []choiceData) int {
	game.showCursor = true
	defer func() { game.showCursor = false }()

	printLn("")

	player.lastMenu = options
	for i, item := range options {
		printfLn("%v) %v", i+1, item.Name)
	}

	num := promptForInteger(game, true, 1, 1, len(options), "Choice")
	if num <= len(options) {
		choice := options[num-1]
		if len(choice.Submenu) > 0 {
			promptForChoice(game, player, choice.Submenu)
			promptForChoice(game, player, options)
		} else if choice.ChoiceFunc != nil {
			if !choice.ChoiceFunc(cData{game: game, player: player}) {
				promptForChoice(game, player, player.lastMenu)
			}
		}
		return num
	} else {
		printLn("That isn't a valid choice!")
		promptForChoice(game, player, options)
	}

	return 0
}
