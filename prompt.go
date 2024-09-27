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

func promptForString(prompt string, min, max int, confirm bool) string {
	fmt.Printf("\n%v: ", prompt)
	line := readLine()
	lLen := len(line)
	if lLen < min {
		fmt.Printf("\nYou must supply at least %v characters.", min)
		return promptForString(prompt, min, max, confirm)
	} else if lLen > max {
		fmt.Printf("\nThat is too long, must be less than %v characters.", max)
		return promptForString(prompt, min, max, confirm)
	}

	if confirm {
		conPrompt := fmt.Sprintf("'%v' Confirm? (Y/N) (yes)", line)
		input := promptForString(conPrompt, 0, 3, false)
		if input == "" || strings.EqualFold(input, "y") || strings.EqualFold(input, "yes") {
			return line
		} else {
			promptForString(prompt, min, max, confirm)
		}
	}
	return line
}

func promptForInteger(prompt string, defaultVal, min, max int) int {
	fmt.Printf("\n%v (%v-%v): (%v) ", prompt, min, max, defaultVal)

	line := readLine()
	if line == "" {
		return defaultVal
	}
	value, err := strconv.ParseInt(line, 10, 64)
	if err != nil {
		fmt.Println("\nThat isn't a number.")
		return promptForInteger(prompt, defaultVal, min, max)
	}
	if int(value) < min || int(value) > max {
		fmt.Printf("\nMust be a value between %v and %v.", min, max)
		return promptForInteger(prompt, defaultVal, min, max)
	}

	return int(value)
}

func promptForChoice(player playerData, options []choiceData) {
	for i, item := range options {
		fmt.Printf("%v) %v\n", i+1, item.Name)
	}

	num := promptForInteger("Choice", 1, 1, len(options))
	if num < len(options) {
		choice := options[num-1]
		if len(choice.Submenu) > 0 {
			promptForChoice(player, choice.Submenu)
		} else if choice.ChoiceFunc != nil {
			choice.ChoiceFunc(player)
		}
	} else {
		fmt.Println("That isn't a valid choice!")
		promptForChoice(player, options)
	}
}
