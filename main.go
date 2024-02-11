package main

import "fmt"

func main() {
	fmt.Println("\nChoose an algorithm:")
	fmt.Println("1. Binary Search")
	fmt.Println("2. Bubble Sort")

	var choice int
	fmt.Print("Enter your choice (1 or 2): ")
	if _, err := fmt.Scan(&choice); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	switch choice {
	case 1:
		fmt.Println("Binary Search")
	case 2:
		fmt.Println("Bubble Sort.")

	default:
		fmt.Println("Invalid choice.")
	}
}
