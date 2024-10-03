package main

func CallClear() {
	consoleOutLock.Lock()
	defer consoleOutLock.Unlock()

	setScreenDirty(true)
	consoleOut = []string{}
}

func CallBGColor(i int) {
	if i < 0 || i > 15 {
		return
	}
	setScreenDirty(true)
	colorBG = tiColor[i]
}

func EnterKey(game *gameData, input string) {
	if input == "" {
		input = "Press enter to continue."
	}
	printfLn(input)

	game.showCursor = true
	defer func() { game.showCursor = false }()
	<-newInput
	unprintln()
}
