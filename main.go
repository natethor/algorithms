package main

import (
	"fmt"
	"github.com/natethor/algorithms/binarySearch"
	"github.com/natethor/algorithms/bubbleSort"
	"github.com/natethor/algorithms/quickSort"
	"math/rand"
	"sort"
	"time"
)

func generateRandomArray(size int, min int, max int) []int {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = rng.Intn(max-min+1) + min
	}

	return array
}

func main() {
	fmt.Println("\nChoose an algorithm:")
	fmt.Println("1. Binary Search")
	fmt.Println("2. Bubble Sort")
	fmt.Println("3. Quick Sort")

	var choice int
	fmt.Print("Enter your choice (1, 2, or 3): ")
	if _, err := fmt.Scan(&choice); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	size := 20
	var randomArray []int
	for {
		randomArray = generateRandomArray(size, 1, 100)
		if !sort.IntsAreSorted(randomArray) {
			break
		}
	}

	fmt.Println("randomArray: ", randomArray)
	switch choice {
	case 1: // Binary Search (must be a sorted array)
		if !sort.IntsAreSorted(randomArray) {
			sort.Ints(randomArray)
		}
		fmt.Println("sortedArray: ", randomArray)
		target := randomArray[size-4]
		fmt.Println("target value: ", target)
		result := binarySearch.BinarySearch(&randomArray, target)
		fmt.Println("result index: ", result)

	case 2: // Bubble Sort
		bubbleSort.BubbleSort(&randomArray)
		fmt.Println("sortedArray: ", randomArray)

	case 3: // Quick Sort
		quickSort.QuickSort(&randomArray)
		fmt.Println("sortedArray: ", randomArray)

	default:
		fmt.Println("Invalid choice")
	}
}
