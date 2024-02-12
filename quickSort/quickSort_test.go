package quickSort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	input := []int{3, 9, 20, 12, 4, 15, 7}
	expected := []int{3, 4, 7, 9, 12, 15, 20}

	QuickSort(&input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Quick Sort failed. Expected %v, got %v", expected, input)
	}
}
