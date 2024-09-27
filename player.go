package main

func (player *playerData) credit(income float64) {
	player.Money = roundToCent(player.Money + income)
}

func (player *playerData) debit(charge float64) bool {
	newAmount := roundToCent(player.Money - charge)

	if newAmount <= 0 {
		player.Bankrupt = true
		return false
	}
	player.Money = newAmount
	return true
}
