package xsort

//SelectedSort 选择排序
//时间复杂度O(n)^2
//空间复杂度O(1), 没有额外开辟空间
func SelectedSort(data []int) []int {
	if data == nil || len(data) < 2 {
		return data
	}

	//获取切片长度
	l := len(data)
	//外层负责循环
	for i := 0; i < l-1; i++ {
		//最小值下标
		minIndex := i
		for j := i + 1; j < l; j++ {
			//最小下标值 > 当前值， min = 当前值下标
			if data[minIndex] > data[j] {
				minIndex = j
			}
		}

		//最小值下标 != 不等于i下标值，替换
		if minIndex != i {
			swap(data, i, minIndex)
		}
	}
	return data
}
