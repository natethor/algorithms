package quickSort

func QuickSort(input *[]int) {
	sort(input, 0, len(*input)-1)
}

func sort(input *[]int, low int, high int) {
	if low >= high {
		return
	}

	pivot_index := partition(input, low, high)
	sort(input, low, pivot_index-1)
	sort(input, pivot_index+1, high)
}

func partition(input *[]int, low int, high int) int {
	pivot := (*input)[high]
	idx := low - 1

	for i := low; i < high; i++ {
		if (*input)[i] <= pivot {
			idx++
			// do the swapping
			temp := (*input)[i]
			(*input)[i] = (*input)[idx]
			(*input)[idx] = temp
		}
	}

	// swap the pivot with the index
	idx++
	temp := (*input)[idx]
	(*input)[idx] = (*input)[high]
	(*input)[high] = temp
	return idx
}
