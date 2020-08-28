package leetcode

import (
	"strings"
)

// 解法一，DFS 深搜
func letterCasePermutation(S string) []string {
	if len(S) == 0 {
		return []string{}
	}
	res, pos, c := []string{}, []int{}, []int{}
	SS := strings.ToLower(S)
	for i := 0; i < len(SS); i++ {
		if isLowerLetter(SS[i]) {
			pos = append(pos, i)
		}
	}
	for i := 0; i <= len(pos); i++ {
		findLetterCasePermutation(SS, pos, i, 0, c, &res)
	}
	return res
}

func isLowerLetter(v byte) bool {
	if v >= 'a' && v <= 'z' {
		return true
	}
	return false
}

func findLetterCasePermutation(s string, pos []int, target, index int, c []int, res *[]string) {
	if len(c) == target {
		b := []byte(s)
		for _, v := range c {
			b[pos[v]] -= 'a' - 'A'
		}
		*res = append(*res, string(b))
		return
	}
	for i := index; i < len(pos)-(target-len(c))+1; i++ {
		c = append(c, i)
		findLetterCasePermutation(s, pos, target, i+1, c, res)
		c = c[:len(c)-1]
	}
}

// 解法二，先讲第一个字母变大写，然后依次把后面的字母变大写。最终的解数组中答案是翻倍增长的
// 第一步：
// [mqe] -> [mqe, Mqe]
// 第二步：
// [mqe, Mqe] -> [mqe Mqe mQe MQe]
// 第二步：
// [mqe Mqe mQe MQe] -> [mqe Mqe mQe MQe mqE MqE mQE MQE]

func letterCasePermutation1(S string) []string {
	res := make([]string, 0, 1<<uint(len(S)))
	S = strings.ToLower(S)
	for k, v := range S {
		if isLetter784(byte(v)) {
			switch len(res) {
			case 0:
				res = append(res, S, toUpper(S, k))
			default:
				for _, s := range res {
					res = append(res, toUpper(s, k))
				}
			}
		}
	}
	if len(res) == 0 {
		res = append(res, S)
	}
	return res
}

func isLetter784(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func toUpper(s string, i int) string {
	b := []byte(s)
	b[i] -= 'a' - 'A'
	return string(b)
}
