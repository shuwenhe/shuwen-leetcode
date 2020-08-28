# [697. Degree of an Array](https://leetcode.com/problems/degree-of-an-array/)


## 题目

Given a non-empty array of non-negative integers `nums`, the **degree** of this array is defined as the maximum frequency of any one of its elements.

Your task is to find the smallest possible length of a (contiguous) subarray of `nums`, that has the same degree as `nums`.

**Example 1**:

```
Input: [1, 2, 2, 3, 1]
Output: 2
Explanation: 
The input array has a degree of 2 because both elements 1 and 2 appear twice.
Of the subarrays that have the same degree:
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
The shortest length is 2. So return 2.

```

**Example 2**:

```
Input: [1,2,2,3,1,4,2]
Output: 6
```

**Note**:

- `nums.length` will be between 1 and 50,000.
- `nums[i]` will be an integer between 0 and 49,999.


## 题目大意

给定一个非空且只包含非负数的整数数组 nums, 数组的度的定义是指数组里任一元素出现频数的最大值。你的任务是找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。

注意:

- nums.length 在 1 到 50,000 区间范围内。
- nums[i] 是一个在 0 到 49,999 范围内的整数。


## 解题思路

- 找一个与给定数组相同度的最短连续子数组，输出其长度。数组的度的定义是任一元素出现频数的最大值。
- 简单题。先统计各个元素的频次，并且动态维护最大频次和子数组的起始和终点位置。这里最短连续子数组有点“迷惑人”。这个最短子数组其实处理起来很简单。只需从前往后扫一遍，记录各个元素第一次出现的位置和最后一次出现的位置即是最短的连续子数组。然后在频次字典里面寻找和最大频次相同的所有解，有可能有多个子数组能满足题意，取出最短的输出即可。

## 代码

```go
func findShortestSubArray(nums []int) int {
	frequency, maxFreq, smallest := map[int][]int{}, 0, len(nums)
	for i, num := range nums {
		if _, found := frequency[num]; !found {
			frequency[num] = []int{1, i, i}
		} else {
			frequency[num][0]++
			frequency[num][2] = i
		}
		if maxFreq < frequency[num][0] {
			maxFreq = frequency[num][0]
		}
	}
	for _, indices := range frequency {
		if indices[0] == maxFreq {
			if smallest > indices[2]-indices[1]+1 {
				smallest = indices[2] - indices[1] + 1
			}
		}
	}
	return smallest
}
```