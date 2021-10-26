package xsearch

//BinarySearch 在一个有序数组中，找某个数是否存在
func BinarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		m := low + (high-low)>>1
		if nums[m] == target {
			return m
		} else if nums[m] > target {
			high = m - 1
		} else {
			low = m + 1
		}
	}
	return -1
}

//在一个有序数组中，找>=某个数最左侧的位置

//局部最小值问题

//?对数器
