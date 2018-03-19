package main

import (
	"fmt"

	"github.com/nitipatl/vendingMachine"
)

func main() {
	_machine := vendingMachine.NewVendingMachine(0, "")

	_machine.AddCoin("T")
	_machine.AddCoin("TW")
	_machine.AddCoin("O")

	fmt.Println("Currently inserted money:", _machine.GetTotal())

	fmt.Println(_machine.SelectDrink("DW"))

	fmt.Println("Return:", _machine.SelectReturn())

}
