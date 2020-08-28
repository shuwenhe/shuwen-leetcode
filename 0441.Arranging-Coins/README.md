# [441. Arranging Coins](https://leetcode.com/problems/arranging-coins/)

## 题目

You have a total of n coins that you want to form in a staircase shape, where every k-th row must have exactly k coins.

Given n, find the total number of **full** staircase rows that can be formed.

n is a non-negative integer and fits within the range of a 32-bit signed integer.

**Example 1:**

    n = 5
    
    The coins can form the following rows:
    ¤
    ¤ ¤
    ¤ ¤
    
    Because the 3rd row is incomplete, we return 2.

**Example 2:**

    n = 8
    
    The coins can form the following rows:
    ¤
    ¤ ¤
    ¤ ¤ ¤
    ¤ ¤
    
    Because the 4th row is incomplete, we return 3.


## 题目大意

你总共有 n 枚硬币，你需要将它们摆成一个阶梯形状，第 k 行就必须正好有 k 枚硬币。给定一个数字 n，找出可形成完整阶梯行的总行数。n 是一个非负整数，并且在32位有符号整型的范围内。



## 解题思路


- n 个硬币，按照递增的方式排列搭楼梯，第一层一个，第二层二个，……第 n 层需要 n 个硬币。问硬币 n 能够搭建到第几层？
- 这一题有 2 种解法，第一种解法就是解方程求出 X，`(1+x)x/2 = n`，即 `x = floor(sqrt(2*n+1/4) - 1/2)`，第二种解法是模拟。
