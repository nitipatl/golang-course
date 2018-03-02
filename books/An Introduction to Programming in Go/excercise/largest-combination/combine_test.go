package combine

import (
	"testing"
)

func TestInput50and2and1and9Return95021(t *testing.T) {

	if combinedNumber(50, 2, 1, 9) != "95021" {
		t.Error("combinedNumber(50, 2, 1, 9) != 95021 as expected.")
	}
}

func TestInput51and50and56Return565150(t *testing.T) {
	if combinedNumber(51, 50, 56) != "565150" {
		t.Error("combinedNumber(51, 50, 56) != 565150 as expected.")
	}
}

func TestInput5and50and56Return56550(t *testing.T) {
	if combinedNumber(5, 50, 56) != "56550" {
		t.Error("combinedNumber(5, 50, 56) != 56550 as expected.")
	}
}

func TestInput420and42and423Return42423420(t *testing.T) {
	if combinedNumber(420, 42, 423) != "42423420" {
		t.Error("combinedNumber(420, 42, 423) != 42423420 as expected.")
	}
}

func TestInput420and421and423Return423421420(t *testing.T) {
	if combinedNumber(420, 421, 423) != "423421420" {
		t.Error("combinedNumber(420, 421, 423) != 423421420 as expected.")
	}
}
