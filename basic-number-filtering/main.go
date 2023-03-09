package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

}

func IsOdd(number int) bool {
	return number%2 != 0
}

func IsEven(number int) bool {
	return number%2 == 0
}

func EvenNumbers(numberArray []int) []int {
	var evenNumberArray []int
	for _, num := range numberArray {
		if IsEven(num) {
			evenNumberArray = append(evenNumberArray, num)
		}
	}
	return evenNumberArray

}

func IsMultipleOf3(number int) bool {
	return number%3 == 0
}

func IsPrime(number int) bool {
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

func OddNumbers(numberArray []int) []int {
	var OddNumberArray []int
	for _, num := range numberArray {
		if IsOdd(num) {
			OddNumberArray = append(OddNumberArray, num)
		}
	}
	return OddNumberArray
}

func OddPrime(numberArray []int) []int {
	var OddPrimeNumbers []int
	for _, num := range numberArray {
		if IsOdd(num) && IsPrime(num) {
			OddPrimeNumbers = append(OddPrimeNumbers, num)
		}
	}
	return OddPrimeNumbers
}

func OddMultipleOf_3_GreaterThan_10(numberArray []int) []int {
	var numbers []int
	for _, num := range numberArray {
		if IsOdd(num) && IsMultipleOf3(num) && num > 10 {
			numbers = append(numbers, num)
		}
	}
	return numbers
}

func All[T any](v T, functions ...func(T) bool) bool {
	for _, f := range functions {
		if !f(v) {
			return false
		}
	}
	return true
}

func Any[T any](v T, functions ...func(T) bool) bool {
	for _, f := range functions {
		if f(v) {
			return true
		}
	}
	return false
}

func filterAll(numberArray []int, conditions ...func(int) bool) []int {
	var result []int
	for _, num := range numberArray {
		if All(num, conditions...) {
			result = append(result, num)
		}
	}
	return result
}

func filterAny(numberArray []int, conditions ...func(int) bool) []int {
	var result []int
	for _, num := range numberArray {
		if Any(num, conditions...) {
			result = append(result, num)
		}
	}
	return result
}
