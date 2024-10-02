/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"strconv"
	"strings"
	"time"
)

func readLine() string {

	for {
		time.Sleep(time.Millisecond * 10)

		inputLock.Lock()
		if !sDirty {
			inputLock.Unlock()
			continue
		}

		output := sLine
		sDirty = false
		inputLock.Unlock()
		return output
	}
}

func promptForString(defaultAnswer string, min, max int, confirm bool, format string, args ...interface{}) string {
	printfln(format+" ", args...)

	line := readLine()
	if line == "" {
		line = defaultAnswer
	}
	lLen := len(line)
	if lLen < min {
		printfln("At least %v characters.", min)
		return promptForString(defaultAnswer, min, max, confirm, format, args...)
	} else if lLen > max {
		printfln("Only %v characters max.", max)
		return promptForString(defaultAnswer, min, max, confirm, format, args...)
	}

	if confirm {
		if promptForBool(true, "Confirm: [%v]", line) {
			if line == "" {
				return defaultAnswer
			}
			return line
		} else {
			promptForString(defaultAnswer, min, max, confirm, format, args...)
		}
	}

	return line
}

func promptForBool(defaultYes bool, format string, args ...interface{}) bool {
	question := ""
	if defaultYes {
		question = " [Y/n] "
	} else {
		question = " [y/N] "
	}
	result := promptForString("", 0, 3, false, format+question, args...)

	if result == "" {
		return defaultYes
	} else if strings.EqualFold(result, "n") || strings.EqualFold(result, "no") {
		return false
	} else if strings.EqualFold(result, "y") || strings.EqualFold(result, "yes") {
		return true
	} else {
		println("Yes or no y/n?")
		return promptForBool(defaultYes, format, args...)
	}
}

func promptForInteger(useDefault bool, defaultVal, min, max int, prompt string) int {

	if useDefault {
		printfln("%v %v-%v [%v] ", prompt, min, max, defaultVal)
	} else {
		printfln("%v %v-%v ", prompt, min, max)
	}

	line := readLine()
	if useDefault && line == "" {
		return defaultVal
	}
	value, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		println("Must be a number.")
		return promptForInteger(useDefault, defaultVal, min, max, prompt)
	}
	if int(value) < min || int(value) > max {
		printfln("Must be %v to %v.", min, max)
		return promptForInteger(useDefault, defaultVal, min, max, prompt)
	}

	return int(value)
}

func promptForMoney(prompt string, defaultVal, min, max float64) float64 {

	printfln("%v $%0.2f-%0.2f: $%0.2f ", prompt, min, max, defaultVal)

	line := readLine()
	if line == "" {
		return defaultVal
	}
	line = strings.TrimSpace(line)
	line = NumOnly(line)
	value, err := strconv.ParseFloat(line, 64)
	if err != nil {
		println("That isn't a number.")
		return promptForMoney(prompt, defaultVal, min, max)
	}
	value = roundToCent(value)
	min = roundToCent(min)
	max = roundToCent(max)

	if value < min || value > max {
		printfln("Must be $%0.2f-%0.2f.", min, max)
		return promptForMoney(prompt, defaultVal, min, max)
	}

	return value
}

func promptForChoice(game *gameData, player *playerData, options []choiceData) int {
	println("")

	player.lastMenu = options
	for i, item := range options {
		printfln("%v) %v", i+1, item.Name)
	}

	num := promptForInteger(true, 1, 1, len(options), "Choice")
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
		println("That isn't a valid choice!")
		promptForChoice(game, player, options)
	}

	return 0
}
