package main

import (
	"fmt"
	"sync"
)

var (
	screenMagnify = 4

	fontScale             = 1
	fontSizeX, fontSizeY  = 8 * fontScale, 8 * fontScale
	termWidth, termHeight = 32, 24
	scrollBack            = 10000

	screenOverscan = 0.1640

	baseX = (fontSizeX / fontScale) * termWidth
	baseY = (fontSizeY + 1/fontScale) * (termHeight - 1)

	xMargin = int(float64(baseX*fontScale) * screenOverscan)
	yMargin = int(float64(baseY*fontScale) * screenOverscan)

	screenWidth  = int(baseX+xMargin) * fontScale
	screenHeight = int(baseY+yMargin) * fontScale

	colorBG = tiColor[C_CYAN]
	colorFG = tiColor[C_BLACK]

	cursorChar = 127

	consoleOutLock sync.Mutex
	consoleOut     []string

	cInputRune []rune
	consoleIn  string
	newInput   chan string = make(chan string)
)

func printf(format string, args ...interface{}) {
	setScreenDirty(true)

	consoleOutLock.Lock()
	defer consoleOutLock.Unlock()

	buf := fmt.Sprintf(format, args...)
	consoleOut = append(consoleOut, buf)
	end := len(consoleOut)

	consoleOut = consoleOut[max(0, end-scrollBack):end]
	scroll = 0
}

func printfLn(format string, args ...interface{}) {
	setScreenDirty(true)

	consoleOutLock.Lock()
	defer consoleOutLock.Unlock()

	buf := fmt.Sprintf(format, args...)
	consoleOut = append(consoleOut, buf+"\n")
	end := len(consoleOut)

	consoleOut = consoleOut[max(0, end-scrollBack):end]
	scroll = 0
}

func printLn(output string) {
	setScreenDirty(true)

	consoleOutLock.Lock()
	defer consoleOutLock.Unlock()

	consoleOut = append(consoleOut, output+"\n")
	end := len(consoleOut)

	consoleOut = consoleOut[max(0, end-scrollBack):end]
	scroll = 0
}

func unprintln() {
	setScreenDirty(true)

	consoleOutLock.Lock()
	defer consoleOutLock.Unlock()

	end := len(consoleOut)
	if end <= 0 {
		return
	}
	consoleOut = consoleOut[max(0, (end-1)-scrollBack) : end-1]
}
