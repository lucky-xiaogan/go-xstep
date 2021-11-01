package xsearch

import "testing"

func TestLessIndex(t *testing.T) {
	list := []int{10, 80, 1, 2, 3, 4, 5, 6, 7, 8}
	index := LessIndex(list)
	t.Log(list[index])
}
