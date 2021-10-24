package xsort

func InsertSort(data []int) []int {
	if data == nil || len(data) < 2 {
		return data
	}

	l := len(data)
	//0~0有序
	//0~1 有序
	// ...
	for i := 1; i < l; i++ { //0~i 做到有序
		temp := data[i]
		j := i
		for j > 0 && data[j] < data[j-1] {
			swap(data, j, j-1)
			j--
		}
		data[j] = temp
	}
	return data
}
