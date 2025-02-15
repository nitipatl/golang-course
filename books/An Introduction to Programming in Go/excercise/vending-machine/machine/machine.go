package machine

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

type Machine struct {
	total          int
	selectedDrinks string
}

func (_machine *Machine) addCoin(coin string) int {
	if val, ok := coinValues[coin]; ok {
		_machine.total += val
	}
	return _machine.total
}

func (_machine *Machine) getTotal() int {
	return _machine.total
}

func (_machine *Machine) selectDrink(action string) string {
	if val, ok := drickPrices[action]; ok {
		if val <= _machine.total {
			_machine.total -= val
			_machine.selectedDrinks += action + " "
		}
	}
	return _machine.selectedDrinks
}

func (_machine *Machine) selectReturn() (totalReturn string) {
	var keys []string
	for k := range coinValues {
		keys = append(keys, k)
	}
	for _, k := range keys {
		if coinValues[k] <= _machine.total {
			coins := _machine.total / coinValues[k]
			for i := 0; i < coins; i++ {
				totalReturn += k + " "
			}
			_machine.total %= coinValues[k]
		}
	}
	return
}
