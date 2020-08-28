# [162. Find Peak Element](https://leetcode.com/problems/find-peak-element/)


## 题目

A peak element is an element that is greater than its neighbors.

Given an input array `nums`, where `nums[i] ≠ nums[i+1]`, find a peak element and return its index.

The array may contain multiple peaks, in that case return the index to any one of the peaks is fine.

You may imagine that `nums[-1] = nums[n] = -∞`.

**Example 1:**

    Input: nums = [1,2,3,1]
    Output: 2
    Explanation: 3 is a peak element and your function should return the index number 2.

**Example 2:**

    Input: nums = [1,2,1,3,5,6,4]
    Output: 1 or 5 
    Explanation: Your function can return either index number 1 where the peak element is 2, 
                 or index number 5 where the peak element is 6.

**Note:**

Your solution should be in logarithmic complexity.

## 题目大意

峰值元素是指其值大于左右相邻值的元素。给定一个输入数组 nums，其中 nums[i] ≠ nums[i+1]，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。你可以假设 nums[-1] = nums[n] = -∞。

说明:

- 你的解法应该是 O(logN) 时间复杂度的。


## 解题思路

- 给出一个数组，数组里面存在多个“山峰”，(山峰的定义是，下标 `i` 比 `i-1`、`i+1` 位置上的元素都要大)，找到这个“山峰”，并输出其中一个山峰的下标。
- 这一题是第 852 题的伪加强版，第 852 题中只存在一个山峰，这一题存在多个山峰。但是实际上搜索的代码是一样的，因为此题只要求随便输出一个山峰的下标即可。思路同第 852 题。
