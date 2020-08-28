# [633. Sum of Square Numbers](https://leetcode.com/problems/sum-of-square-numbers/)


## 题目

Given a non-negative integer `c`, your task is to decide whether there're two integers `a` and `b` such that a^2 + b^2 = c.

**Example 1:**

    Input: 5
    Output: True
    Explanation: 1 * 1 + 2 * 2 = 5

**Example 2:**

    Input: 3
    Output: False


## 题目大意

给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a^2 + b^2 = c。


## 解题思路

- 给出一个数，要求判断这个数能否由由 2 个完全平方数组成。能则输出 true，不能则输出 false。
- 可以用二分搜索来解答这道题。判断题意，依次计算 `low * low + high * high`  和 c 是否相等。从 [1, sqrt(n)] 区间内进行二分，若能找到则返回 true，找不到就返回 false 。
