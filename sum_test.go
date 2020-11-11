package main

import "testing"

func TestSumOnePlusOne(t *testing.T) {
	firstNum := 1
	secondNum := 1
	expectedResult := 2
	result := SumNumber(firstNum, secondNum)

	if result != expectedResult {
		t.Fatalf("Expected result : %d, got : %d", expectedResult, result)
	}
}
