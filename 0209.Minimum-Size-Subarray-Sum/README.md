# [209. Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/)

## 题目

Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

Example 1:

```c
Input: s = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: the subarray [4,3] has the minimal length under the problem constraint.
```

Follow up:       
  
If you have figured out the O(n) solution, try coding another solution of which the time complexity is O(n log n). 

## 题目大意

给定一个整型数组和一个数字 s，找到数组中最短的一个连续子数组，使得连续子数组的数字之和 sum>=s，返回最短的连续子数组的返回值。

## 解题思路

这一题的解题思路是用滑动窗口。在滑动窗口 [i,j]之间不断往后移动，如果总和小于 s，就扩大右边界 j，不断加入右边的值，直到 sum > s，之和再缩小 i 的左边界，不断缩小直到 sum < s，这时候右边界又可以往右移动。以此类推。

