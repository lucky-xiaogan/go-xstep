package xsort

import "testing"

func TestInsertSort(t *testing.T) {
	test := []int{4, 2, 5, 6, 1, 2, 34, 9, 20, 56, 7, 8}
	res := InsertSort(test)
	t.Log(res)
}
