# [172. Factorial Trailing Zeroes](https://leetcode.com/problems/factorial-trailing-zeroes/)


## 题目

Given an integer n, return the number of trailing zeroes in n!.

**Example 1:**

    Input: 3
    Output: 0
    Explanation: 3! = 6, no trailing zero.

**Example 2:**

    Input: 5
    Output: 1
    Explanation: 5! = 120, one trailing zero.

**Note:** Your solution should be in logarithmic time complexity.


## 题目大意


给定一个整数 n，返回 n! 结果尾数中零的数量。说明: 你算法的时间复杂度应为 O(log n) 。




## 解题思路

- 给出一个数 n，要求 n！末尾 0 的个数。
- 这是一道数学题。计算 N 的阶乘有多少个后缀 0，即计算 N! 里有多少个 10，也是计算 N! 里有多少个 2 和 5（分解质因数），最后结果即 2 的个数和 5 的个数取较小值。每两个数字就会多一个质因数 2，而每五个数字才多一个质因数 5。每 5 个数字就会多一个质因数 5。0~4 的阶乘里没有质因数 5，5~9 的阶乘里有 1 个质因数 5，10~14 的阶乘里有 2 个质因数 5，依此类推。所以 0 的个数即为 `min(阶乘中 5 的个数和 2 的个数)`。
- N! 有多少个后缀 0，即 N! 有多少个质因数 5。N! 有多少个质因数 5，即 N 可以划分成多少组 5个数字一组，加上划分成多少组 25 个数字一组，加上划分多少组成 125 个数字一组，等等。即 `res = N/5 + N/(5^2) + N/(5^3) + ... = ((N / 5) / 5) / 5 /...` 。最终算法复杂度为 O(logN)。
