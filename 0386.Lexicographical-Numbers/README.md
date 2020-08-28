# [386. Lexicographical Numbers](https://leetcode.com/problems/lexicographical-numbers/)


## 题目

Given an integer n, return 1 - n in lexicographical order.

For example, given 13, return: [1,10,11,12,13,2,3,4,5,6,7,8,9].

Please optimize your algorithm to use less time and space. The input size may be as large as 5,000,000.


## 题目大意

给定一个整数 n, 返回从 1 到 n 的字典顺序。例如，给定 n =13，返回 [1,10,11,12,13,2,3,4,5,6,7,8,9] 。

请尽可能的优化算法的时间复杂度和空间复杂度。 输入的数据 n 小于等于 5,000,000。



## 解题思路


- 给出一个数字 n ，要求按照字典序对 1-n 这 n 个数排序。
- DFS 暴力求解即可。
