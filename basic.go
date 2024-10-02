package main

func CallClear() {
	consoleOutLock.Lock()
	defer consoleOutLock.Unlock()

	consoleOut = []string{}
}

func CallBGColor(i int) {
	if i < 0 || i > 15 {
		return
	}

	colorBG = tiColor[i]
}
