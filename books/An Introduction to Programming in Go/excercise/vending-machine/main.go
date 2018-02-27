package main

import (
	"fmt"
)

var coinValues = map[string]int{
	"T":  10,
	"F":  5,
	"TW": 2,
	"O":  1,
}

var drickPrices = map[string]int{
	"SD": 18, // soft drink
	"CF": 12, // canned coffee
	"DW": 7,  // drinking water
}

type Wallet struct {
	total int
}

func (_wallet *Wallet) addCoin(coin string) int {
	if val, err := coinValues[coin]; err {
		_wallet.total += val
	}
	return _wallet.total
}

func (_wallet *Wallet) getTotal() int {
	return _wallet.total
}

func main() {
	_wallet := Wallet{0}

	_wallet.addCoin("T")
	_wallet.addCoin("TW")

	fmt.Println("Currently inserted money:", _wallet.getTotal())

}
