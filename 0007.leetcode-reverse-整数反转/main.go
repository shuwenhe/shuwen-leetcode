// 给出一个32位的有符号整数，你需要将这个整数中每位上的数字进行反转。
// 输入: 123
// 输出: 321

package main

import (
	"fmt"
)

func main() {
	res := reverse(123)
	fmt.Println("res = ", res)
}

func reverse(x int) int {
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		x = x / 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}
