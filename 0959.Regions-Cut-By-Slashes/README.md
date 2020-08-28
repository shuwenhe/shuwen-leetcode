# [959. Regions Cut By Slashes](https://leetcode.com/problems/regions-cut-by-slashes/)


## 题目

In a N x N `grid` composed of 1 x 1 squares, each 1 x 1 square consists of a `/`, `\`, or blank space. These characters divide the square into contiguous regions.

(Note that backslash characters are escaped, so a `\` is represented as `"\\"`.)

Return the number of regions.

**Example 1:**

    Input:
    [
      " /",
      "/ "
    ]
    Output: 2
    Explanation: The 2x2 grid is as follows:

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/1.png)

**Example 2:**

    Input:
    [
      " /",
      "  "
    ]
    Output: 1
    Explanation: The 2x2 grid is as follows:

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/2.png)

**Example 3:**

    Input:
    [
      "\\/",
      "/\\"
    ]
    Output: 4
    Explanation: (Recall that because \ characters are escaped, "\\/" refers to \/, and "/\\" refers to /\.)
    The 2x2 grid is as follows:

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/3.png)

**Example 4:**

    Input:
    [
      "/\\",
      "\\/"
    ]
    Output: 5
    Explanation: (Recall that because \ characters are escaped, "/\\" refers to /\, and "\\/" refers to \/.)
    The 2x2 grid is as follows:

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/4.png)

**Example 5:**

    Input:
    [
      "//",
      "/ "
    ]
    Output: 3
    Explanation: The 2x2 grid is as follows:

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/5.png)

**Note:**

1. `1 <= grid.length == grid[0].length <= 30`
2. `grid[i][j]` is either `'/'`, `'\'`, or `' '`.


## 题目大意

在由 1 x 1 方格组成的 N x N 网格 grid 中，每个 1 x 1 方块由 /、\ 或空格构成。这些字符会将方块划分为一些共边的区域。(请注意，反斜杠字符是转义的，因此 \ 用 "\\" 表示)返回区域的数目。


提示：

- 1 <= grid.length == grid[0].length <= 30
- grid[i][j] 是 '/'、'\'、或 ' '。

## 解题思路


- 给出一个字符串，代表的是 `N x N` 正方形中切分的情况，有 2 种切分的情况 `'\'` 和 `'/'` ，即从左上往右下切和从右上往左下切。问按照给出的切分方法，能把 `N x N` 正方形切成几部分？
- 这一题解题思路是并查集。先将每个 `1*1` 的正方形切分成下图的样子。分成 4 小块。然后按照题目给的切分图来合并各个小块。

![](https://img.halfrost.com/Leetcode/leetcode_959.png)

- 遇到 `'\\'`，就把第 0 块和第 1 块 `union()` 起来，第 2 块和第 3 块 `union()` 起来；遇到 `'/'`，就把第 0 块和第 3 块 `union()` 起来，第 2 块和第 1 块 `union()` 起来；遇到 `' '`，就把第 0 块和第 1 块 `union()` 起来，第 2 块和第 1 块 `union()` 起来，第 2 块和第 3 块 `union()` 起来，即 4 块都 `union()` 起来；最后还需要记得上一行和下一行还需要 `union()`，即本行的第 2 块和下一行的第 0 块 `union()` 起来；左边一列和右边一列也需要 `union()`。即本列的第 1 块和右边一列的第 3 块 `union()` 起来。最后计算出集合总个数就是最终答案了。
