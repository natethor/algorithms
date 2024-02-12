package bubbleSort

import (
	"testing"
	"reflect"
)

func TestBubbleSort(t *testing.T) {
	input := []int{3, 9, 20, 12, 4, 15, 7}
	expected := []int{3, 4, 7, 9, 12, 15, 20}

	BubbleSort(&input)

	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Bubble Sort failed. Expected %v, got %v", expected, input)
	}
}
