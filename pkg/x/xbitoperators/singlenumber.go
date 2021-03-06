package xbitoperators

/*
SingleNumber 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

& 与运算 两个位都是 1 时，结果才为 1，否则为 0，如
  1 0 0 1 1
& 1 1 0 0 1
------------------------------
 1 0 0 0 1

| 或运算 两个位都是 0 时，结果才为 0，否则为 1，如
  1 0 0 1 1
| 1 1 0 0 1
------------------------------
  1 1 0 1 1

^ 异或运算，两个位相同则为 0，不同则为 1，如
  1 0 0 1 1
^ 1 1 0 0 1
-----------------------------
  0 1 0 1 0
1） 0 ^ N = N, N ^ N = 0
2) 满足交换率和结合率: a ^ b = b ^ a, (a ^ b) ^ c = a ^ (b ^ c)
3) N 个数异或的值与异或的顺序无关
4）交换两个值 a= 甲, b= 乙
	(1)a = a ^ b, (2)b = a ^ b, (3)a = a ^ b
	推理方式(代入法)：
	第一次执行：a = 甲 ^ 乙 , b = 乙
	第二次执行：b = a ^ b = 甲 ^ 乙 ^ 乙 = 甲 ^ 0 = 甲
	第三次执行：a = 甲 ^ 乙 ^ 甲 = 0 ^ 乙 = 乙
注意：a，b属于两块不同的内存空间

~ 取反运算，0 则变为 1，1 则变为 0，如
~ 1 0 0 1 1
-----------------------------
  0 1 1 0 0
*/

func SingleNumber(data []int) int {
	res := 0
	for _, v := range data {
		res ^= v
	}
	return res
}

func PrintOddTimesNum2(data []int) (int, int) {
	eor, onlyOne := 0, 0
	for _, v := range data {
		eor ^= v
	}

	//eor = a ^ b
	//eor != 0
	//eor必然有一个位置为1
	onlyOne = eor & (^eor + 1) //提取出最右的1
	for _, v := range data {
		if onlyOne&v == 0 {
			onlyOne ^= v
		}
	}
	return onlyOne, eor ^ onlyOne
}
