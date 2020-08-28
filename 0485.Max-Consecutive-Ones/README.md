# [485. Max Consecutive Ones](https://leetcode.com/problems/max-consecutive-ones/)


## 题目

Given a binary array, find the maximum number of consecutive 1s in this array.

**Example 1**:

```
Input: [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
    The maximum number of consecutive 1s is 3.
```

**Note**:

- The input array will only contain `0` and `1`.
- The length of input array is a positive integer and will not exceed 10,000


## 题目大意

给定一个二进制数组， 计算其中最大连续1的个数。

注意：

- 输入的数组只包含 0 和 1。
- 输入数组的长度是正整数，且不超过 10,000。


## 解题思路

- 给定一个二进制数组， 计算其中最大连续1的个数。
- 简单题。扫一遍数组，累计 1 的个数，动态维护最大的计数，最终输出即可。

## 代码

```go
func findMaxConsecutiveOnes(nums []int) int {
	maxCount, currentCount := 0, 0
	for _, v := range nums {
		if v == 1 {
			currentCount++
		} else {
			currentCount = 0
		}
		if currentCount > maxCount {
			maxCount = currentCount
		}
	}
	return maxCount
}
```