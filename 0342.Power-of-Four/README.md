# [342. Power of Four](https://leetcode.com/problems/power-of-four/)


## 题目

Given an integer (signed 32 bits), write a function to check whether it is a power of 4.

**Example 1:**

    Input: 16
    Output: true

**Example 2:**

    Input: 5
    Output: false

**Follow up**: Could you solve it without loops/recursion?

## 题目大意

给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。


## 解题思路

- 判断一个数是不是 4 的 n 次方。
- 这一题最简单的思路是循环，可以通过。但是题目要求不循环就要判断，这就需要用到数论的知识了。
- 证明 `(4^n - 1) % 3 == 0`，(1) `4^n - 1 = (2^n + 1) * (2^n - 1)`(2) 在任何连续的 3  个数中 `(2^n-1)`，`(2^n)`，`(2^n+1)`，一定有一个数是 3 的倍数。`(2^n)` 肯定不是 3 的倍数，那么 `(2^n-1)` 或者 `(2^n+1)` 中一定有一个是 3 的倍数。所以 `4^n-1` 一定是 3 的倍数。
