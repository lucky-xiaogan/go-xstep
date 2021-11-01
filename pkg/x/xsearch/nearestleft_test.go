package xsearch

import "testing"

func TestNearestIndex(t *testing.T)  {
	i := NearestLeftIndex([]int{1,2,3,3,3,3,4,5,6,7,7,7,8}, 3)
	t.Log(i)
}
