# [46. Permutations](https://leetcode.com/problems/permutations/)


## 题目

Given a collection of **distinct** integers, return all possible permutations.

**Example:**


    Input: [1,2,3]
    Output:
    [
      [1,2,3],
      [1,3,2],
      [2,1,3],
      [2,3,1],
      [3,1,2],
      [3,2,1]
    ]


## 题目大意

给定一个没有重复数字的序列，返回其所有可能的全排列。


## 解题思路

- 求出一个数组的排列组合中的所有排列，用 DFS 深搜即可。
