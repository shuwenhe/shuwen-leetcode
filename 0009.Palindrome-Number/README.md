# [9. Palindrome Number](https://leetcode.com/problems/palindrome-number/)


## 题目

Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

**Example 1**:

```
Input: 121
Output: true
```

**Example 2**:

```
Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
```

**Example 3**:

```
Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
```

**Follow up**:

Coud you solve it without converting the integer to a string?

## 题目大意

判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

## 解题思路

- 判断一个整数是不是回文数。
- 简单题。注意会有负数的情况，负数，个位数，10 都不是回文数。其他的整数再按照回文的规则判断。

## 代码

```go

package leetcode

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

```