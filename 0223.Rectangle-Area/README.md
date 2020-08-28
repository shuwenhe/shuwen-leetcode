# [223. Rectangle Area](https://leetcode.com/problems/rectangle-area/)


## 题目

Find the total area covered by two **rectilinear** rectangles in a **2D** plane.

Each rectangle is defined by its bottom left corner and top right corner as shown in the figure.

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rectangle_area.png)

**Example:**

    Input: A = -3, B = 0, C = 3, D = 4, E = 0, F = -1, G = 9, H = 2
    Output: 45

**Note:**

Assume that the total area is never beyond the maximum possible value of **int**.



## 题目大意

在二维平面上计算出两个由直线构成的矩形重叠后形成的总面积。每个矩形由其左下顶点和右上顶点坐标表示，如图所示。说明: 假设矩形面积不会超出 int 的范围。

## 解题思路


- 给出两个矩形的坐标，求这两个矩形在坐标轴上覆盖的总面积。
- 几何题，由于只有 2 个矩形，所以按照题意做即可。先分别求两个矩形的面积，加起来再减去两个矩形重叠的面积。
