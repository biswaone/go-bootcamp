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
