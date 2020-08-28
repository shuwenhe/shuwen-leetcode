# [1221. Split a String in Balanced Strings](https://leetcode.com/problems/split-a-string-in-balanced-strings/)


## 题目

Balanced strings are those who have equal quantity of 'L' and 'R' characters.

Given a balanced string `s` split it in the maximum amount of balanced strings.

Return the maximum amount of splitted balanced strings.

**Example 1**:

    Input: s = "RLRRLLRLRL"
    Output: 4
    Explanation: s can be split into "RL", "RRLL", "RL", "RL", each substring contains same number of 'L' and 'R'.

**Example 2**:

    Input: s = "RLLLLRRRLR"
    Output: 3
    Explanation: s can be split into "RL", "LLLRRR", "LR", each substring contains same number of 'L' and 'R'.

**Example 3**:

    Input: s = "LLLLRRRR"
    Output: 1
    Explanation: s can be split into "LLLLRRRR".

**Constraints**:

- `1 <= s.length <= 1000`
- `s[i] = 'L' or 'R'`

## 题目大意


在一个「平衡字符串」中，'L' 和 'R' 字符的数量是相同的。给出一个平衡字符串 s，请你将它分割成尽可能多的平衡字符串。返回可以通过分割得到的平衡字符串的最大数量。

提示：

- 1 <= s.length <= 1000
- s[i] = 'L' 或 'R'


## 解题思路

- 给出一个字符串，要求把这个字符串切成一些子串，这些子串中 R 和 L 的字符数是相等的。问能切成多少个满足条件的子串。
- 这道题是简单题，按照题意模拟即可。从左往右扫，遇到 `R` 就加一，遇到 `L` 就减一，当计数是 `0` 的时候就是平衡的时候，就切割。

## 代码

```go

package leetcode

func balancedStringSplit(s string) int {
	count, res := 0, 0
	for _, r := range s {
		if r == 'R' {
			count++
		} else {
			count--
		}
		if count == 0 {
			res++
		}
	}
	return res
}

```