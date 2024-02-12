package binarySearch

import "fmt"

func BinarySearch(input *[]int, target int) int {
	low := 0
	high := len((*input)) - 1

	for low <= high {
		mid := low + (high-low)/2
		fmt.Println("mid index: ", mid)

		if (*input)[mid] == target {
			return mid
		} else if (*input)[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
