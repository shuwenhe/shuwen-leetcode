# [628. Maximum Product of Three Numbers](https://leetcode.com/problems/maximum-product-of-three-numbers/)


## 题目

Given an integer array, find three numbers whose product is maximum and output the maximum product.

**Example 1:**

    Input: [1,2,3]
    Output: 6

**Example 2:**

    Input: [1,2,3,4]
    Output: 24

**Note:**

1. The length of the given array will be in range [3,10^4] and all elements are in the range [-1000, 1000].
2. Multiplication of any three numbers in the input won't exceed the range of 32-bit signed integer.


## 题目大意

给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。




## 解题思路


- 给出一个数组，要求求出这个数组中任意挑 3 个数能组成的乘积最大的值。
- 题目的 test case 数据量比较大，如果用排序的话，时间复杂度高，可以直接考虑模拟，挑出 3 个数组成乘积最大值，必然是一个正数和二个负数，或者三个正数。那么选出最大的三个数和最小的二个数，对比一下就可以求出最大值了。时间复杂度 O(n)

