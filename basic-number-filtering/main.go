package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

}

func isEven(number int) bool {
	return number%2 == 0
}

func EvenNumbers(numberArray []int) []int {
	var evenNumberArray []int
	for _, num := range numberArray {
		if isEven(num) {
			evenNumberArray = append(evenNumberArray, num)
		}
	}
	return evenNumberArray

}

func isOdd(number int) bool {
	return number%2 != 0
}

func OddNumbers(numberArray []int) []int {
	var OddNumberArray []int
	for _, num := range numberArray {
		if isOdd(num) {
			OddNumberArray = append(OddNumberArray, num)
		}
	}
	return OddNumberArray
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
