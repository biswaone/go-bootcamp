package main

import (
	"reflect"
	"testing"
)

func TestIsEven(t *testing.T) {
	//given
	input_1 := 1
	input_2 := 2
	want_1 := false
	want_2 := true

	//when
	got_1 := isEven(input_1)
	got_2 := isEven(input_2)

	//then
	if !reflect.DeepEqual(got_1, want_1) {
		t.Errorf("wanted: %v but got: %v", want_1, got_1)
	}
	if !reflect.DeepEqual(got_2, want_2) {
		t.Errorf("wanted: %v but got: %v", want_2, got_2)
	}
}

func TestEvenNumbers(t *testing.T) {
	// given
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 4, 6, 8, 10}

	// when
	output := EvenNumbers(input)

	//then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}

func TestIsOdd(t *testing.T) {
	//given
	input_1 := 1
	input_2 := 2
	want_1 := true
	want_2 := false

	//when
	got_1 := isOdd(input_1)
	got_2 := isOdd(input_2)

	//then
	if !reflect.DeepEqual(got_1, want_1) {
		t.Errorf("wanted: %v but got: %v", want_1, got_1)
	}
	if !reflect.DeepEqual(got_2, want_2) {
		t.Errorf("wanted: %v but got: %v", want_2, got_2)
	}

}

func TestOddNumbers(t *testing.T) {
	// given
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{1, 3, 5, 7, 9}

	// when
	got := OddNumbers(input)

	//then
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
}

func TestIsPrime(t *testing.T) {
	//given
	testCases := []struct {
		n    int
		want bool
	}{
		{2, true},
		{3, true},
		{5, true},
		{7, true},
		{11, true},
		{4, false},
		{6, false},
		{8, false},
		{9, false},
		{10, false},
	}

	for _, tc := range testCases {
		got := isPrime(tc.n)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted: %v but got: %v", want, got)
		}
	}
}
