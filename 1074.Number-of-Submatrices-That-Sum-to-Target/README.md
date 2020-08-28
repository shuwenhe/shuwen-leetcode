# [1074. Number of Submatrices That Sum to Target](https://leetcode.com/problems/number-of-submatrices-that-sum-to-target/)


## 题目

Given a `matrix`, and a `target`, return the number of non-empty submatrices that sum to target.

A submatrix `x1, y1, x2, y2` is the set of all cells `matrix[y]` with `x1 <= x <= x2` and `y1 <= y <= y2`.

Two submatrices `(x1, y1, x2, y2)` and `(x1', y1', x2', y2')` are different if they have some coordinate that is different: for example, if `x1 != x1'`.

**Example 1:**

    Input: matrix = [[0,1,0],[1,1,1],[0,1,0]], target = 0
    Output: 4
    Explanation: The four 1x1 submatrices that only contain 0.

**Example 2:**

    Input: matrix = [[1,-1],[-1,1]], target = 0
    Output: 5
    Explanation: The two 1x2 submatrices, plus the two 2x1 submatrices, plus the 2x2 submatrix.

**Note:**

1. `1 <= matrix.length <= 300`
2. `1 <= matrix[0].length <= 300`
3. `-1000 <= matrix[i] <= 1000`
4. `-10^8 <= target <= 10^8`


## 题目大意

给出矩阵 matrix 和目标值 target，返回元素总和等于目标值的非空子矩阵的数量。

子矩阵 x1, y1, x2, y2 是满足 x1 <= x <= x2 且 y1 <= y <= y2 的所有单元 matrix[x][y] 的集合。

如果 (x1, y1, x2, y2) 和 (x1', y1', x2', y2') 两个子矩阵中部分坐标不同（如：x1 != x1'），那么这两个子矩阵也不同。


提示：

1. 1 <= matrix.length <= 300
2. 1 <= matrix[0].length <= 300
3. -1000 <= matrix[i] <= 1000
4. -10^8 <= target <= 10^8




## 解题思路

- 给出一个矩阵，要求在这个矩阵中找出子矩阵的和等于 target 的矩阵个数。
- 这一题读完题感觉是滑动窗口的二维版本。如果把它拍扁，在一维数组中，求连续的子数组和为 target，这样就很好做。如果这题不降维，纯暴力解是 O(n^6)。如何优化降低时间复杂度呢？
- 联想到第 1 题 Two Sum 问题，可以把 2 个数求和的问题优化到 O(n)。这里也用类似的思想，用一个 map 来保存行方向上曾经出现过的累加和，相减就可以得到本行的和。这里可能读者会有疑惑，为什么不能每一行都单独保存呢？为什么一定要用累加和相减的方式来获取每一行的和呢？因为这一题要求子矩阵所有解，如果只单独保存每一行的和，只能求得小的子矩阵，子矩阵和子矩阵组成的大矩阵的情况会漏掉(当然再循环一遍，把子矩阵累加起来也可以，但是这样就多了一层循环了)，例如子矩阵是 1*4 的，但是 2 个这样的子矩阵摞在一起形成 2 * 4 也能满足条件，如果不用累加和的办法，只单独存每一行的和，最终还要有组合的步骤。经过这样的优化，可以从 O(n^6) 优化到 O(n^4)，能 AC 这道题，但是时间复杂度太高了。如何优化？
- 首先，子矩阵需要上下左右 4 个边界，4 个变量控制循环就需要 O(n^4)，行和列的区间累加还需要 O(n^2)。行和列的区间累加可以通过 preSum 来解决。例如 `sum[i,j] = sum[j] - sum[i - 1]`，其中 sum[k] 中存的是从 0 到 K 的累加和： ![](https://img.halfrost.com/Leetcode/leetcode_1074.gif)

    那么一个区间内的累加和可以由这个区间的右边界减去区间左边界左边的那个累加和得到(由于是闭区间，所需要取左边界左边的和)。经过这样的处理，列方向的维度就被我们拍扁了。

- 再来看看行方向的和，现在每一列的和都可以通过区间相减的方法得到。那么这道题就变成了第 1 题 Two Sum 的问题了。Two Sum 问题只需要 O(n) 的时间复杂度求解，这一题由于是二维的，所以两个列的边界还需要循环，所以最终优化下来的时间复杂度是 O(n^3)。计算 presum 可以直接用原数组，所以空间复杂度只有一个 O(n) 的字典。
- 类似思路的题目有第 560 题，第 304 题。
