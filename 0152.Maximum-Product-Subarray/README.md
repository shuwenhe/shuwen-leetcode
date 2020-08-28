# [152. Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/)


## 题目

Given an integer array `nums`, find the contiguous subarray within an array (containing at least one number) which has the largest product.

**Example 1:**

    Input: [2,3,-2,4]
    Output: 6
    Explanation: [2,3] has the largest product 6.

**Example 2:**

    Input: [-2,0,-1]
    Output: 0
    Explanation: The result cannot be 2, because [-2,-1] is not a subarray.


## 题目大意

给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。


## 解题思路

- 给出一个数组，要求找出这个数组中连续元素乘积最大的值。
- 这一题是 DP 的题，状态转移方程是：最大值是 `Max(f(n)) = Max( Max(f(n-1)) * n, Min(f(n-1)) * n)`；最小值是 `Min(f(n)) = Min( Max(f(n-1)) * n, Min(f(n-1)) * n)`。只要动态维护这两个值，如果最后一个数是负数，最大值就在负数 * 最小值中产生，如果最后一个数是正数，最大值就在正数 * 最大值中产生。

