package main

import "testing"

func TestWeather25andBangkokFewCloud(t *testing.T) {
	display := weatherCelcius(25, "Bangkok few cloud")
	output := `
 _  _ 
 _||_ 
|_  _| c
Bangkok few cloud
`
	if display != output {
		t.Error(
			"expected", output,
			"got", display)
	}
}

func TestWeather34andTakSunny(t *testing.T) {
	display := weatherCelcius(34, "Tak Sunny")
	output := `
 _    
 _||_|
 _|  | c
Tak Sunny
`
	if display != output {
		t.Error(
			"expected", output,
			"got", display)
	}
}
