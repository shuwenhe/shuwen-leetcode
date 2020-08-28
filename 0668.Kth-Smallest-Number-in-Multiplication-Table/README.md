# [668. Kth Smallest Number in Multiplication Table](https://leetcode.com/problems/kth-smallest-number-in-multiplication-table/)


## 题目

Nearly every one have used the [Multiplication Table](https://en.wikipedia.org/wiki/Multiplication_table). But could you find out the `k-th` smallest number quickly from the multiplication table?

Given the height `m` and the length `n` of a `m * n` Multiplication Table, and a positive integer `k`, you need to return the `k-th` smallest number in this table.

**Example 1:**

    Input: m = 3, n = 3, k = 5
    Output: 
    Explanation: 
    The Multiplication Table:
    1	2	3
    2	4	6
    3	6	9
    
    The 5-th smallest number is 3 (1, 2, 2, 3, 3).

**Example 2:**

    Input: m = 2, n = 3, k = 6
    Output: 
    Explanation: 
    The Multiplication Table:
    1	2	3
    2	4	6
    
    The 6-th smallest number is 6 (1, 2, 2, 3, 4, 6).

**Note:**

1. The `m` and `n` will be in the range [1, 30000].
2. The `k` will be in the range [1, m * n]


## 题目大意

几乎每一个人都用乘法表。但是你能在乘法表中快速找到第 k 小的数字吗？给定高度 m 、宽度 n 的一张 m * n 的乘法表，以及正整数 k，你需要返回表中第 k 小的数字。


注意：

- m 和 n 的范围在 [1, 30000] 之间。
- k 的范围在 [1, m * n] 之间。

## 解题思路

- 给出 3 个数字，m，n，k。m  和 n 分别代表乘法口诀表的行和列。要求在这个乘法口诀表中找第 k 小的数字。
- 这一题是第 378 题变种题。利用二分搜索，在 `[1,m*n]` 的区间内搜索第 `k` 小的数。每次二分统计 `≤ mid` 数字的个数。由于是在两数乘法构成的矩阵中计数，知道乘数，被乘数也就知道了，所以计数只需要一层循环。整体代码和第 378 题完全一致，只是计数的部分不同罢了。可以对比第 378 题一起练习。
