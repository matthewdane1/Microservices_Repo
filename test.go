package main

import (
	"fmt"
	"sort"
)

// func findMinimum(numbers ...int) int {
// 	smallest := numbers[0]
// 	for _, x := range numbers {
// 		if x < smallest {
// 			smallest = x
// 		}
// 	}
// 	return smallest
// }

// func main() {
// 	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 987436}
// 	smallest := findMinimum(nums...)
// 	fmt.Println(smallest)
// }

// can also use sort

func SmallestIntegerFinder(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 987436}
	smallest := SmallestIntegerFinder(nums)
	fmt.Println(smallest)
}


// did it work