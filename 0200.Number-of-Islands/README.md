# [200. Number of Islands](https://leetcode.com/problems/number-of-islands/)


## 题目

Given a 2d grid map of `'1'`s (land) and `'0'`s (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

**Example 1:**

    Input:
    11110
    11010
    11000
    00000
    
    Output: 1

**Example 2:**

    Input:
    11000
    11000
    00100
    00011
    
    Output: 3

## 题目大意

给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。


## 解题思路

- 要求找出地图中的孤岛。孤岛的含义是四周被海水包围的岛。
- 这一题可以按照第 79 题的思路进行搜索，只要找到为 "1" 的岛以后，从这里开始搜索这周连通的陆地，也都标识上访问过。每次遇到新的 "1" 且没有访问过，就相当于遇到了新的岛屿了。

