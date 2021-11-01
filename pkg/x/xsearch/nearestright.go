package xsearch

// NearestRightIndex 在arr上，找满足>=value的最左位置
func NearestRightIndex(arr []int, val int) int {
	//index 记录最左的对号
	l, r, index := 0, len(arr)-1, -1
	for l <= r { // 至少一个数的时候
		mid := l + ((r - l) >> 1)
		if arr[mid] <= val {
			index = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return index
}
