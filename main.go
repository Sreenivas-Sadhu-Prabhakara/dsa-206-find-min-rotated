package main

import "fmt"

// FindMin returns the smallest value of a sorted array of distinct values
// rotated an unknown number of times. Compare the midpoint against the
// RIGHT end: larger means the drop is to the right, otherwise at or left.
func FindMin(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > nums[hi] {
			lo = mid + 1 // the cliff (max->min) lies right of mid
		} else {
			hi = mid // mid might itself be the minimum
		}
	}
	return nums[lo]
}

func main() {
	fmt.Println(FindMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(FindMin([]int{11, 13, 15, 17}))
}
