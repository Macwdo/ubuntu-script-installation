package main

import "testing"

func TestIsOddSuccess(t *testing.T) {
	number_as_2 := 2

	result := IsOdd(number_as_2)

	if !result {
		t.Errorf("Result for 2 is not odd")
	}

}
