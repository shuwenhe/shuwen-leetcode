# [74. Search a 2D Matrix](https://leetcode.com/problems/search-a-2d-matrix/)


## 题目

Write an efficient algorithm that searches for a value in an *m* x *n* matrix. This matrix has the following properties:

- Integers in each row are sorted from left to right.
- The first integer of each row is greater than the last integer of the previous row.

**Example 1:**

    Input:
    matrix = [
      [1,   3,  5,  7],
      [10, 11, 16, 20],
      [23, 30, 34, 50]
    ]
    target = 3
    Output: true

**Example 2:**

    Input:
    matrix = [
      [1,   3,  5,  7],
      [10, 11, 16, 20],
      [23, 30, 34, 50]
    ]
    target = 13
    Output: false


## 题目大意

编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

- 每行中的整数从左到右按升序排列。
- 每行的第一个整数大于前一行的最后一个整数。


## 解题思路


- 给出一个二维矩阵，矩阵的特点是随着矩阵的下标增大而增大。要求设计一个算法能在这个矩阵中高效的找到一个数，如果找到就输出 true，找不到就输出 false。
- 虽然是一个二维矩阵，但是由于它特殊的有序性，所以完全可以按照下标把它看成一个一维矩阵，只不过需要行列坐标转换。最后利用二分搜索直接搜索即可。
