package sort

func swap(data []int, aIdx, bIdx int) {
	tmp := data[aIdx]
	data[aIdx] = data[bIdx]
	data[bIdx] = tmp
}

func Sort(data []int) {
	sortIdx(data, 0, len(data) - 1)
}

func sortIdx(data []int, startIdx, endIdx int) {
	if startIdx < endIdx   {
		pivot := partition(data, startIdx, endIdx)
		sortIdx(data, startIdx, pivot-1)
		sortIdx(data, pivot+1, endIdx)
	}
}

func partition(data []int, left, right int) int {
	pivot := data[right]

	pointer := left
	for j := left; j < right; j++ {
		if data[j] <= pivot {
			swap(data, pointer, j)
			pointer++
		}
	}

	//Swap pointer & pivot
	swap(data, pointer, right)

	return pointer
}
