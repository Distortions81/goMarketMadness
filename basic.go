package main

func CallClear() {
	outputLock.Lock()
	defer outputLock.Unlock()

	sOut = []string{}
}

func CallBGColor(i int) {
	if i < 0 || i > 15 {
		return
	}

	colorBG = tiColor[i]
}
