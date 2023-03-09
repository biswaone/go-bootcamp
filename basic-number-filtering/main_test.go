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
	got_1 := IsEven(input_1)
	got_2 := IsEven(input_2)

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
	got_1 := IsOdd(input_1)
	got_2 := IsOdd(input_2)

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
		got := IsPrime(tc.n)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted: %v but got: %v", want, got)
		}
	}
}

func TestOddPrime(t *testing.T) {
	//given
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{3, 5, 7}

	//got
	got := OddPrime(input)

	//then
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}

}

func TestOddMultipleOf_3_GreaterThan_10(t *testing.T) {
	//given
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
	want := []int{15, 21}

	//got
	got := OddMultipleOf_3_GreaterThan_10(input)

	//then
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}

}

func TestFilerAll(t *testing.T) {
	//given
	IsGreaterThan5 := func(number int) bool { return number > 5 }
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int{9, 15}

	//got
	got := filterAll(input, IsOdd, IsGreaterThan5, IsMultipleOf3)

	//then
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}

}

func TestFilterAll(t *testing.T) {
	//given
	IsGreaterThan5 := func(number int) bool { return number > 5 }
	IsLessThan15 := func(number int) bool { return number < 15 }

	testCases := []struct {
		input     []int
		functions []func(int) bool
		want      []int
	}{
		{
			input:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			functions: []func(int) bool{IsOdd, IsGreaterThan5, IsMultipleOf3},
			want:      []int{9, 15},
		},
		{
			input:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			functions: []func(int) bool{IsEven, IsLessThan15, IsMultipleOf3},
			want:      []int{6, 12},
		},
	}

	for _, tc := range testCases {
		got := filterAll(tc.input, tc.functions...)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted: %v but got: %v", want, got)
		}
	}
}

func TestFilterAny(t *testing.T) {
	//given
	IsGreaterThan15 := func(number int) bool { return number > 15 }
	IsLessThan6 := func(number int) bool { return number < 6 }
	IsMultipleOf5 := func(number int) bool { return number%5 == 0 }

	testCases := []struct {
		input     []int
		functions []func(int) bool
		want      []int
	}{
		{
			input:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			functions: []func(int) bool{IsPrime, IsGreaterThan15, IsMultipleOf5},
			want:      []int{2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20},
		},
		{
			input:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			functions: []func(int) bool{IsLessThan6, IsMultipleOf3},
			want:      []int{1, 2, 3, 4, 5, 6, 9, 12, 15, 18},
		},
	}

	for _, tc := range testCases {
		got := filterAny(tc.input, tc.functions...)
		want := tc.want
		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted: %v but got: %v", want, got)
		}
	}
}
