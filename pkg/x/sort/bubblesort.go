package sort

//BubbleSort 冒泡排序
func BubbleSort(data []int) []int {
	n := len(data)
	//外层循环
	for i := 0; i < n-1; i++ {
		//内层实现交换
		for j := 0; j < n-1-i; j++ {
			if data[j] > data[j+1] {
				//执行交换
				swap(data, j, j+1)
			}
		}
	}
	return data
}
