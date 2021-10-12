package xbitoperators

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	test := []int{1,1,1,1,2,3,3,4,5,7,2,4,5}
	res := SingleNumber(test)
	t.Log(res)
}