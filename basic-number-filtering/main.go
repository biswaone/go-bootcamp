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
