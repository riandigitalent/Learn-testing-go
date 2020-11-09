package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if Calculatate(2) != 4 {
		t.Error("expected 2 +2 to equal 4")
	}

}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-2, 0},
		{99, 101},
		{0, 2},
		{-7, -5},
	}

	for _, test := range tests {
		if output := Calculatate(test.input); output != test.expected {
			t.Error("expected ", test.input, " +2 to equal ", test.expected, " , not :", output)
		}
	}
}
