/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import (
	"fmt"
	"sync"
)

var (
	sOut []string

	sInBuf,
	sLine string
	sRune  []rune
	sDirty bool
)

func main() {

	newGame := &gameData{Settings: defSettings}
	go newGame.playGame()

	startEbiten()
}

var (
	outputLock,
	inputLock sync.Mutex
)

func printf(format string, args ...interface{}) {
	outputLock.Lock()
	defer outputLock.Unlock()

	buf := fmt.Sprintf(format, args...)
	sOut = append(sOut, buf)
}

func printfln(format string, args ...interface{}) {
	outputLock.Lock()
	defer outputLock.Unlock()

	buf := fmt.Sprintf(format, args...)
	sOut = append(sOut, buf+"\n")
}

func println(output string) {
	outputLock.Lock()
	defer outputLock.Unlock()

	sOut = append(sOut, output+"\n")
}
