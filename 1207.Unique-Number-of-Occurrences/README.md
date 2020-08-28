# [1207. Unique Number of Occurrences](https://leetcode.com/problems/unique-number-of-occurrences/)


## 题目

Given an array of integers `arr`, write a function that returns `true` if and only if the number of occurrences of each value in the array is unique.

**Example 1:**

    Input: arr = [1,2,2,1,1,3]
    Output: true
    Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.

**Example 2:**

    Input: arr = [1,2]
    Output: false

**Example 3:**

    Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
    Output: true

**Constraints:**

- `1 <= arr.length <= 1000`
- `-1000 <= arr[i] <= 1000`



## 题目大意

给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。

提示：

- 1 <= arr.length <= 1000
- -1000 <= arr[i] <= 1000

## 解题思路


- 给出一个数组，先统计每个数字出现的频次，判断在这个数组中是否存在相同的频次。
- 简单题，先统计数组中每个数字的频次，然后用一个 map 判断频次是否重复。
