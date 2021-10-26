package xsearch

import "testing"

func TestBinarySearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	index := BinarySearch(list, 3)
	t.Log(index)
}
