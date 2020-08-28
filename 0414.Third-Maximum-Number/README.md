# [414. Third Maximum Number](https://leetcode.com/problems/third-maximum-number/)

## 题目

Given a **non-empty** array of integers, return the **third** maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).

**Example 1:**

    Input: [3, 2, 1]
    
    Output: 1
    
    Explanation: The third maximum is 1.

**Example 2:**

    Input: [1, 2]
    
    Output: 2
    
    Explanation: The third maximum does not exist, so the maximum (2) is returned instead.

**Example 3:**

    Input: [2, 2, 3, 1]
    
    Output: 1
    
    Explanation: Note that the third maximum here means the third maximum distinct number.
    Both numbers with value 2 are both considered as second maximum.


## 题目大意

给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是 O(n)。


## 解题思路

- 水题，动态维护 3 个最大值即可。注意数组中有重复数据的情况。如果只有 2 个数或者 1 个数，则返回 2 个数中的最大值即可。