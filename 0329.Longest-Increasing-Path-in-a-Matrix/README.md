# [329. Longest Increasing Path in a Matrix](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/)


## 题目

Given an integer matrix, find the length of the longest increasing path.

From each cell, you can either move to four directions: left, right, up or down. You may NOT move diagonally or move outside of the boundary (i.e. wrap-around is not allowed).

**Example 1:**

    Input: nums = 
    [
      [9,9,4],
      [6,6,8],
      [2,1,1]
    ] 
    Output: 4 
    Explanation: The longest increasing path is [1, 2, 6, 9].

**Example 2:**

    Input: nums = 
    [
      [3,4,5],
      [3,2,6],
      [2,2,1]
    ] 
    Output: 4 
    Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.


## 题目大意

给定一个整数矩阵，找出最长递增路径的长度。对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。


## 解题思路


- 给出一个矩阵，要求在这个矩阵中找到一个最长递增的路径。路径有上下左右 4 个方向。
- 这一题解题思路很明显，用 DFS 即可。在提交完第一版以后会发现 TLE，因为题目给出了一个非常大的矩阵，搜索次数太多。所以需要用到记忆化，把曾经搜索过的最大长度缓存起来，增加了记忆化以后再次提交 AC。
