package main

func CallClear() {
	outputLock.Lock()
	defer outputLock.Unlock()

	sOut = []string{}
}
