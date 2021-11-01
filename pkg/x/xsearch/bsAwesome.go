package xsearch

func LessIndex(arr []int) int {
	l := len(arr)
	if arr == nil || l == 0 {
		return -1
	}

	if l == 1 || arr[0] < arr[1] {
		return 0
	}

	if arr[l-1] < arr[l-2] {
		return l - 1
	}
	left, right, mid := 1, l-2, 0
	for left < right {
		mid = left + ((right - left) >> 1)
		if arr[mid] > arr[mid-1] {
			right = mid - 1
		} else if arr[mid] > arr[mid+1] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
