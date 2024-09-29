/***************
* STOCK MARKET *
****************/
/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func anyKey() {
	reader := bufio.NewReader(os.Stdin)
	b, err := reader.ReadByte()
	if err != nil {
		log.Fatal("Unable to read a string.")
	}
	if b != '\n' {
		fmt.Println()
	}
	CallSound(50, 3000, 5)
}

func readLine() string {
	var buffer []byte
	var bLen int

	reader := bufio.NewReader(os.Stdin)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			log.Fatal("Unable to read a string.")
		}

		if b == '\b' || b == '\x7f' {
			if bLen > 0 {
				bLen--
				buffer = buffer[:bLen]
				fmt.Print("\b \b")
				continue
			} else {
				CallSound(100, 110, 1)
			}
		} else if b == '\n' {
			return string(buffer)
		} else {
			buffer = append(buffer, b)
			bLen++
			fmt.Print(string(b))
		}
	}
}

func promptForString(defaultAnswer string, min, max int, confirm bool, format string, args ...interface{}) string {
	fmt.Printf("Default answer '%v'\n", defaultAnswer)
	fmt.Printf(format+" ", args...)

	line := readLine()
	fmt.Println()
	lLen := len(line)
	if lLen < min {
		fmt.Printf("You must supply at least %v characters.\n", min)
		return promptForString(defaultAnswer, min, max, confirm, format, args...)
	} else if lLen > max {
		fmt.Printf("That is too long, must be less than %v characters.\n", max)
		return promptForString(defaultAnswer, min, max, confirm, format, args...)
	}

	if confirm {
		if promptForBool(true, "Confirm") {
			return line
		} else {
			promptForString(defaultAnswer, min, max, confirm, format, args...)
		}
	}

	if line == "" {
		return defaultAnswer
	}
	return line
}

func promptForBool(defaultYes bool, format string, args ...interface{}) bool {
	question := ""
	if defaultYes {
		question = " (Y/n):"
	} else {
		question = " (y/n):"
	}
	result := promptForString("", 0, 3, false, format+question, args...)

	if !defaultYes && result == "" {
		fmt.Println("That isn't a valid option, type yes/no or y/n.")
		return promptForBool(defaultYes, format, args...)
	}
	return (defaultYes && result == "") || strings.EqualFold(result, "y") || strings.EqualFold(result, "yes")
}

func promptForInteger(defaultVal, min, max int, prompt string) int {
	fmt.Printf("%v (%v-%v): (%v) ", prompt, min, max, defaultVal)

	line := readLine()
	fmt.Println()
	if line == "" {
		return defaultVal
	}
	value, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		fmt.Println("That isn't a number.")
		return promptForInteger(defaultVal, min, max, prompt)
	}
	if int(value) < min || int(value) > max {
		fmt.Printf("Must be a value between %v and %v.\n", min, max)
		return promptForInteger(defaultVal, min, max, prompt)
	}

	return int(value)
}

func promptForMoney(prompt string, defaultVal, min, max float64) float64 {
	fmt.Printf("%v ($%0.2f-$%0.2f): ($%0.2f) ", prompt, min, max, defaultVal)

	line := readLine()
	fmt.Println()
	if line == "" {
		return defaultVal
	}
	line = strings.TrimSpace(line)
	line = NumOnly(line)
	value, err := strconv.ParseFloat(line, 64)
	if err != nil {
		fmt.Println("That isn't a number.")
		return promptForMoney(prompt, defaultVal, min, max)
	}
	value = roundToCent(value)
	min = roundToCent(min)
	max = roundToCent(max)

	if value < min || value > max {
		fmt.Printf("Must be a value between $%0.2f and $%0.2f.\n", min, max)
		return promptForMoney(prompt, defaultVal, min, max)
	}

	return value
}

func promptForChoice(game *gameData, player *playerData, options []choiceData) int {
	fmt.Println("")
	for i, item := range options {
		fmt.Printf("%v) %v\n", i+1, item.Name)
	}

	num := promptForInteger(1, 1, len(options), "Choice")
	if num < len(options) {
		choice := options[num-1]
		if len(choice.Submenu) > 0 {
			promptForChoice(game, player, choice.Submenu)
			promptForChoice(game, player, options)
		} else if choice.ChoiceFunc != nil {
			choice.ChoiceFunc(game, player)
		}
		return num
	} else {
		fmt.Println("That isn't a valid choice!")
		promptForChoice(game, player, options)
	}

	return 0
}
