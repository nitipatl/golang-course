package main

import (
	"fmt"
)

func main() {
	_machine := Machine{total: 0, selectedDrinks: ""}

	_machine.addCoin("T")
	_machine.addCoin("TW")
	_machine.addCoin("O")

	fmt.Println("Currently inserted money:", _machine.getTotal())

	fmt.Println(_machine.selectDrink("DW"))

	fmt.Println("Return:", _machine.selectReturn())

}
