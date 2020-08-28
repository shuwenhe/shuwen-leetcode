# [766. Toeplitz Matrix](https://leetcode.com/problems/toeplitz-matrix/)


## 题目

A matrix is *Toeplitz* if every diagonal from top-left to bottom-right has the same element.

Now given an `M x N` matrix, return `True` if and only if the matrix is *Toeplitz*.

**Example 1:**

    Input:
    matrix = [
      [1,2,3,4],
      [5,1,2,3],
      [9,5,1,2]
    ]
    Output: True
    Explanation:
    In the above grid, the diagonals are:
    "[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]".
    In each diagonal all elements are the same, so the answer is True.

**Example 2:**

    Input:
    matrix = [
      [1,2],
      [2,2]
    ]
    Output: False
    Explanation:
    The diagonal "[1, 2]" has different elements.

**Note:**

1. `matrix` will be a 2D array of integers.
2. `matrix` will have a number of rows and columns in range `[1, 20]`.
3. `matrix[i][j]` will be integers in range `[0, 99]`.

**Follow up:**

1. What if the matrix is stored on disk, and the memory is limited such that you can only load at most one row of the matrix into the memory at once?
2. What if the matrix is so large that you can only load up a partial row into the memory at once?


## 题目大意

如果一个矩阵的每一方向由左上到右下的对角线上具有相同元素，那么这个矩阵是托普利茨矩阵。给定一个 M x N 的矩阵，当且仅当它是托普利茨矩阵时返回 True。



## 解题思路


- 给出一个矩阵，要求判断矩阵所有对角斜线上的数字是否都是一个数字。
- 水题，直接循环判断即可。

