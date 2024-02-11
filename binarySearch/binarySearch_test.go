package binarySearch

import "testing"

func TestBinarySearch(t *testing.T) {
	input := [13]int{1, 3, 4, 7, 11, 12, 15, 21, 25, 31, 38, 42, 59}

	expected := 11 // Index where 42 is located in the array
	actual := binarySearch(input[:], 42)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	} else {
		t.Logf("Search successful. Element found at index %d", actual)
	}
}
