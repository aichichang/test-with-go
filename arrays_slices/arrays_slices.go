package main

// array has a fixed size
// func Sum(numbers [5]int) int {
// 	sum := 0
//
// 	for _, number := range numbers {
// 		sum += number
// 	}
//
// 	return sum
// }

// slice does not have a fixed size
func Sum(numbers []int) int {
	sum := 0

	// range lets you iterate the array
	// range returns two values - the index and the value.
	// We are choosing to ignore the index value by using _ blank identifier
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// // println(lengthOfNumbers)
	// // make allows you to create a slice with a staring capacity of the len
	// sums := make([]int, lengthOfNumbers)
	//
	// for i, numbers := range numbersToSum {
	// 	// You can index slices like arrays with mySlice[N] to get the value out or assign it a new value with =
	// 	sums[i] = Sum(numbers)
	// }

	// slices have capacity, we can use append to add new values
	var sums []int

	for _, nums := range numbersToSum {
		sums = append(sums, Sum(nums))
	}

	return sums
}

// The tail of a collection is all items in the collection except the first one (the "head").
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, nums := range numbersToSum {

		if len(nums) < 1 {
			sums = append(sums, 0)
		} else {
			// we are saying "take from 1 to the end" with numbers[1:]
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
