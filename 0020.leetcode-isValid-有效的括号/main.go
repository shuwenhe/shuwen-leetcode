// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。
// 输入: "()[]{}"
// 输出: true

package main

import (
	"fmt"
)

func main() {
	b := isValid("[]")
	fmt.Println("b = ", b)
}

func isValid(s string) bool {
	if len(s) == 0 { // 空字符串直接返回true
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if v == '[' || v == '(' || v == '{' {
			stack = append(stack, v)
		} else if v == ']' && len(stack) > 0 && stack[len(stack)-1] == '[' ||
			((v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((v == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
