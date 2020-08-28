# [367. Valid Perfect Square](https://leetcode.com/problems/valid-perfect-square/)

## 题目

Given a positive integer num, write a function which returns True if num is a perfect square else False.

**Note:** **Do not** use any built-in library function such as `sqrt`.

**Example 1:**

    Input: 16
    Output: true

**Example 2:**

    Input: 14
    Output: false


## 题目大意

给定一个正整数 num，编写一个函数，如果 num 是一个完全平方数，则返回 True，否则返回 False。

说明：不要使用任何内置的库函数，如 sqrt。




## 解题思路


- 给出一个数，要求判断这个数是不是完全平方数。
- 可以用二分搜索来解答这道题。判断完全平方数，根据它的定义来，是否能被开根号，即找到一个数的平方是否可以等于待判断的数字。从 [1, n] 区间内进行二分，若能找到则返回 true，找不到就返回 false 。
