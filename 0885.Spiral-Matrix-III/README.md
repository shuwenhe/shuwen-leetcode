# [885. Spiral Matrix III](https://leetcode.com/problems/spiral-matrix-iii/)


## 题目

On a 2 dimensional grid with `R` rows and `C` columns, we start at `(r0, c0)` facing east.

Here, the north-west corner of the grid is at the first row and column, and the south-east corner of the grid is at the last row and column.

Now, we walk in a clockwise spiral shape to visit every position in this grid.

Whenever we would move outside the boundary of the grid, we continue our walk outside the grid (but may return to the grid boundary later.)

Eventually, we reach all `R * C` spaces of the grid.

Return a list of coordinates representing the positions of the grid in the order they were visited.

**Example 1:**

    Input: R = 1, C = 4, r0 = 0, c0 = 0
    Output: [[0,0],[0,1],[0,2],[0,3]]

![](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/08/24/example_1.png)

**Example 2:**

    Input: R = 5, C = 6, r0 = 1, c0 = 4
    Output: [[1,4],[1,5],[2,5],[2,4],[2,3],[1,3],[0,3],[0,4],[0,5],[3,5],[3,4],
    [3,3],[3,2],[2,2],[1,2],[0,2],[4,5],[4,4],[4,3],[4,2],[4,1],[3,1],[2,1],[1,1],
    [0,1],[4,0],[3,0],[2,0],[1,0],[0,0]]

![](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/08/24/example_2.png)

**Note:**

1. `1 <= R <= 100`
2. `1 <= C <= 100`
3. `0 <= r0 < R`
4. `0 <= c0 < C`


## 题目大意

在 R 行 C 列的矩阵上，我们从 (r0, c0) 面朝东面开始。这里，网格的西北角位于第一行第一列，网格的东南角位于最后一行最后一列。现在，我们以顺时针按螺旋状行走，访问此网格中的每个位置。每当我们移动到网格的边界之外时，我们会继续在网格之外行走（但稍后可能会返回到网格边界）。最终，我们到过网格的所有 R * C 个空间。

要求输出按照访问顺序返回表示网格位置的坐标列表。


## 解题思路


- 给出一个二维数组的行 `R`，列 `C`，以及这个数组中的起始点 `(r0,c0)`。从这个起始点开始出发，螺旋的访问数组中各个点，输出途径经过的每个坐标。注意每个螺旋的步长在变长，第一个螺旋是 1 步，第二个螺旋是 1 步，第三个螺旋是 2 步，第四个螺旋是 2 步……即 1，1，2，2，3，3，4，4，5……这样的步长。
- 这一题是第 59 题的加强版。除了有螺旋以外，还加入了步长的限制。步长其实是有规律的，第 0 次移动的步长是 `0/2+1`，第 1 次移动的步长是 `1/2+1`，第 n 次移动的步长是 `n/2+1`。其他的做法和第 59 题一致。

