# [878. Nth Magical Number](https://leetcode.com/problems/nth-magical-number/)


## 题目

A positive integer is *magical* if it is divisible by either A or B.

Return the N-th magical number. Since the answer may be very large, **return it modulo** `10^9 + 7`.

**Example 1:**

    Input: N = 1, A = 2, B = 3
    Output: 2

**Example 2:**

    Input: N = 4, A = 2, B = 3
    Output: 6

**Example 3:**

    Input: N = 5, A = 2, B = 4
    Output: 10

**Example 4:**

    Input: N = 3, A = 6, B = 4
    Output: 8

**Note:**

1. `1 <= N <= 10^9`
2. `2 <= A <= 40000`
3. `2 <= B <= 40000`


## 题目大意


如果正整数可以被 A 或 B 整除，那么它是神奇的。返回第 N 个神奇数字。由于答案可能非常大，返回它模 10^9 + 7 的结果。


提示：

1. 1 <= N <= 10^9
2. 2 <= A <= 40000
3. 2 <= B <= 40000


## 解题思路


- 给出 3 个数字，a，b，n。要求输出可以整除 a 或者整除 b 的第 n 个数。
- 这一题是第 1201 题的缩水版，代码和解题思路也基本不变，这一题的二分搜索的区间是 `[min(A, B)，N * min(A, B)] = [2, 10 ^ 14]`。其他代码和第 1201 题一致，思路见第 1201 题。
