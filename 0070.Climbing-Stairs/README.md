# [70. Climbing Stairs](https://leetcode.com/problems/climbing-stairs/)


## 题目

You are climbing a stair case. It takes *n* steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

**Note:** Given *n* will be a positive integer.

**Example 1:**

    Input: 2
    Output: 2
    Explanation: There are two ways to climb to the top.
    1. 1 step + 1 step
    2. 2 steps

**Example 2:**

    Input: 3
    Output: 3
    Explanation: There are three ways to climb to the top.
    1. 1 step + 1 step + 1 step
    2. 1 step + 2 steps
    3. 2 steps + 1 step


## 题目大意

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？注意：给定 n 是一个正整数


## 解题思路

- 简单的 DP，经典的爬楼梯问题。一个楼梯可以由 `n-1` 和 `n-2` 的楼梯爬上来。
- 这一题求解的值就是斐波那契数列。
