# [836. Rectangle Overlap](https://leetcode.com/problems/rectangle-overlap/)


## 题目

A rectangle is represented as a list `[x1, y1, x2, y2]`, where `(x1, y1)` are the coordinates of its bottom-left corner, and `(x2, y2)` are the coordinates of its top-right corner.

Two rectangles overlap if the area of their intersection is positive. To be clear, two rectangles that only touch at the corner or edges do not overlap.

Given two (axis-aligned) rectangles, return whether they overlap.

**Example 1:**

    Input: rec1 = [0,0,2,2], rec2 = [1,1,3,3]
    Output: true

**Example 2:**

    Input: rec1 = [0,0,1,1], rec2 = [1,0,2,1]
    Output: false

**Notes:**

1. Both rectangles `rec1` and `rec2` are lists of 4 integers.
2. All coordinates in rectangles will be between `-10^9` and `10^9`.


## 题目大意

矩形以列表 [x1, y1, x2, y2] 的形式表示，其中 (x1, y1) 为左下角的坐标，(x2, y2) 是右上角的坐标。如果相交的面积为正，则称两矩形重叠。需要明确的是，只在角或边接触的两个矩形不构成重叠。给出两个矩形，判断它们是否重叠并返回结果。

说明：

1. 两个矩形 rec1 和 rec2 都以含有四个整数的列表的形式给出。
2. 矩形中的所有坐标都处于 -10^9 和 10^9 之间。


## 解题思路

- 给出两个矩形的坐标，判断两个矩形是否重叠。
- 几何题，按照几何方法判断即可。
