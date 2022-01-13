package main

import (
	"testing"
)

func TestGcd(t *testing.T) {
	result := gcd(25, 20)
	resultExpt := 5

	if result != 5 {
		t.Errorf("gcd Test failed. Expected result is %d But you got %d", resultExpt, result)
	} else {
		t.Logf("gcd Test passed")
	}
}

func TestCountItem(t *testing.T) {
	apples := 25
	cakes := 20
	divisor := gcd(apples, cakes)
	eachApples, eachCakes := countItem(apples, cakes, divisor)

	eachApplesExpt, eachCakesExpt := 5, 4

	if eachApples != 5 && eachCakes != 4 {
		t.Errorf("countItem Test failed. Expected result {apples: %d, cakes: %d}. But you got {apples: %d, cakes: =%d}",
			eachApplesExpt, eachCakesExpt, eachApples, eachCakes)
	} else {
		t.Logf("countItem Test passed")
	}
}
