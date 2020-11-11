package main

import "fmt"

func SumNumber(firstNum int, secondNum int) int {
	result := firstNum + secondNum
	fmt.Printf("Adding number : %d and %d = %d \n", firstNum, secondNum, result)
	return result
}

func main() {
	SumNumber(1, 1)
}
