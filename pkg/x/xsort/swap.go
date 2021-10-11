package xsort

//swap 交换函数
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
