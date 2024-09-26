/***************
* STOCK MARKET *
****************/
/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

import "fmt"

func main() {
	defer fixTerm()
	handleExit()
	setupTerm()

	CallClear()
	fmt.Println("Press any key to continue")
	anyKey()

	playGame()
}
