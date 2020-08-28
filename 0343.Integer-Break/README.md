# [343. Integer Break](https://leetcode.com/problems/integer-break/)


## 题目

Given a positive integer n, break it into the sum of **at least** two positive integers and maximize the product of those integers. Return the maximum product you can get.

**Example 1:**

    Input: 2
    Output: 1
    Explanation: 2 = 1 + 1, 1 × 1 = 1.

**Example 2:**

    Input: 10
    Output: 36
    Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.

**Note**: You may assume that n is not less than 2 and not larger than 58.


## 题目大意

给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。


## 解题思路

- 这一题是 DP 的题目，将一个数字分成多个数字之和，至少分为 2 个数字之和，求解分解出来的数字乘积最大是多少。
- 这一题的动态转移方程是 `dp[i] = max(dp[i], j * (i - j), j * dp[i-j])` ，一个数分解成 `j` 和 `i - j` 两个数字，或者分解成 `j` 和 `更多的分解数`，`更多的分解数`即是 `dp[i-j]`，由于 `dp[i-j]` 下标小于 `i` ，所以 `dp[i-j]` 在计算 `dp[i]` 的时候一定计算出来了。

