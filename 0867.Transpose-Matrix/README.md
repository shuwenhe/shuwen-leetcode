# [867. Transpose Matrix](https://leetcode.com/problems/transpose-matrix/)


## 题目

Given a matrix `A`, return the transpose of `A`.

The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.

**Example 1:**

    Input: [[1,2,3],[4,5,6],[7,8,9]]
    Output: [[1,4,7],[2,5,8],[3,6,9]]

**Example 2:**

    Input: [[1,2,3],[4,5,6]]
    Output: [[1,4],[2,5],[3,6]]

**Note:**

1. `1 <= A.length <= 1000`
2. `1 <= A[0].length <= 1000`


## 题目大意

给定一个矩阵 A， 返回 A 的转置矩阵。矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。


## 解题思路


- 给出一个矩阵，顺时针旋转 90°
- 解题思路很简单，直接模拟即可。
