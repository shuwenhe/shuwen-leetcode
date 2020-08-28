# [240. Search a 2D Matrix II](https://leetcode.com/problems/search-a-2d-matrix-ii/)


## 题目

Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

- Integers in each row are sorted in ascending from left to right.
- Integers in each column are sorted in ascending from top to bottom.

**Example:**

Consider the following matrix:

    [
      [1,   4,  7, 11, 15],
      [2,   5,  8, 12, 19],
      [3,   6,  9, 16, 22],
      [10, 13, 14, 17, 24],
      [18, 21, 23, 26, 30]
    ]

Given target = `5`, return `true`.

Given target = `20`, return `false`.


## 题目大意

编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：

- 每行的元素从左到右升序排列。
- 每列的元素从上到下升序排列。



## 解题思路

- 给出一个二维矩阵，矩阵的特点是每一个行内，元素随着下标增大而增大，每一列内，元素也是随着下标增大而增大。但是相邻两行的元素并没有大小关系。例如第一行最后一个元素就比第二行第一个元素要大。要求设计一个算法能在这个矩阵中高效的找到一个数，如果找到就输出 true，找不到就输出 false。
- 这一题是第 74 题的加强版。第 74 题中的二维矩阵完全是一个有序的一维矩阵，但是这一题如果把它拍扁成一维，并不是有序的。首先每一个行或者每一列是有序的 ，那么我们可以依次在每一行或者每一列中利用二分去搜索。这样做时间复杂度为  O(n log n)。
- 还有一个模拟的解法。通过观察，我们发现了这个矩阵的一个特点，最右边一列的元素是本行中最大的元素，所以我们可以先从最右边一列开始找到第一个比 target 元素大的元素，这个元素所在的行，是我们接着要搜索的。在行中搜索是从最右边开始往左边搜索，时间复杂度是 O(n)，算上一开始在最右边一列中查找的时间复杂度是 O(m)，所以最终的时间复杂度为 O(m+n)。
