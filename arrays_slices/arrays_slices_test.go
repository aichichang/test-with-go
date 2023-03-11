package main

import (
	"reflect"
	"testing"
)

func TestSliceSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {

		// how we define an array
		// numbers := [5]int{1, 2, 3, 4, 5}

		// how we define a slice
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9, 3})
	want := []int{3, 12}

	// go does not let you use equality operators with slices
	// if got != want {
	// 	t.Errorf("got %v want %v", got, want)
	// }

	// so we will have to use DeepEqual to compare two values
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9, 3}, []int{})
	want := []int{2, 12, 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
