# [594. Longest Harmonious Subsequence](https://leetcode.com/problems/longest-harmonious-subsequence/)


## 题目

We define a harmounious array as an array where the difference between its maximum value and its minimum value is **exactly** 1.

Now, given an integer array, you need to find the length of its longest harmonious subsequence among all its possible [subsequences](https://en.wikipedia.org/wiki/Subsequence).

**Example 1:**

    Input: [1,3,2,2,5,2,3,7]
    Output: 5
    Explanation: The longest harmonious subsequence is [3,2,2,2,3].

**Note:** The length of the input array will not exceed 20,000.


## 题目大意

和谐数组是指一个数组里元素的最大值和最小值之间的差别正好是1。现在，给定一个整数数组，你需要在所有可能的子序列中找到最长的和谐子序列的长度。说明: 输入的数组长度最大不超过20,000.

## 解题思路

- 在给出的数组里面找到这样一个子数组：要求子数组中的最大值和最小值相差 1 。这一题是简单题。先统计每个数字出现的频次，然后在 map 找相差 1 的 2 个数组的频次和，动态的维护两个数的频次和就是最后要求的子数组的最大长度。
