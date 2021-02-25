package sort_algorithm

//冒泡排序
func BubbleSort(values []int) []int {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			//前后交换
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	return values
}
