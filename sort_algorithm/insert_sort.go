package sort_algorithm

//插入排序
func InsertSort(values []int) []int {
	if len(values) <= 1 {
		return values
	}
	for i := 1; i < len(values); i++ {
		var value = values[i]
		var j = i - 1
		//寻找插入位置
		for ; j >= 0; j-- {
			if values[j] > value {
				values[j+1] = values[j] //数据移动
			} else {
				break
			}
		}
		values[j+1] = value //插入数据
	}
	return values
}
