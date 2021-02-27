package sort_algorithm

//快速排序
func QuickSort(values []int) []int {
	n := len(values) - 1
	QuickSortC(values, 0, n)
	return values
}

func QuickSortC(values []int, p, r int) {
	if p >= r {
		return
	}
	q := Partition(values, p, r)
	QuickSortC(values, p, q-1)
	QuickSortC(values, q+1, r)
}

func Partition(values []int, start, end int) int {
	pivot := values[start]
	left, right := start, end
	for left != right {
		for left < right && pivot < values[right] {
			right--
		}
		for left < right && pivot >= values[left] {
			left++
		}
		if left < right {
			values[left], values[right] = values[right], values[left]
		}
	}
	values[start], values[left] = values[left], values[start]
	return left
}
