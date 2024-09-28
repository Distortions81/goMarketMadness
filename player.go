/***************
* STOCK MARKET *
****************/
/*
 * Copyright (C) 2024 Carl Frank Otto III
 * All rights reserved.
 */

package main

func (player *playerData) credit(income float64) {
	player.Balance = roundToCent(player.Balance + income)
}

func (player *playerData) debit(charge float64) bool {
	newAmount := roundToCent(player.Balance - charge)

	if newAmount <= 0 {
		player.Bankrupt = true
		return false
	}
	player.Balance = newAmount
	return true
}
