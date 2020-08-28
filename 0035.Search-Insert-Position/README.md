# [35. Search Insert Position](https://leetcode.com/problems/search-insert-position/)


## 题目

Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

**Example 1:**

    Input: [1,3,5,6], 5
    Output: 2

**Example 2:**

    Input: [1,3,5,6], 2
    Output: 1

**Example 3:**

    Input: [1,3,5,6], 7
    Output: 4

**Example 4:**

    Input: [1,3,5,6], 0
    Output: 0


## 题目大意

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

## 解题思路

- 给出一个已经从小到大排序后的数组，要求在数组中找到插入 target 元素的位置。
- 这一题是经典的二分搜索的变种题，在有序数组中找到最后一个比 target 小的元素。
