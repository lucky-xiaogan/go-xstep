package xsearch

import "testing"

// NearestRightIndex 在arr上，找满足>=value的最左位置
func TestNearestRightIndex(t *testing.T) {
	i := NearestRightIndex([]int{1, 2, 3, 3, 3, 3, 4, 5, 6, 7, 7, 7, 8}, 3)
	t.Log(i)
}
