# [1054. Distant Barcodes](https://leetcode.com/problems/distant-barcodes/)


## 题目

In a warehouse, there is a row of barcodes, where the `i`-th barcode is `barcodes[i]`.

Rearrange the barcodes so that no two adjacent barcodes are equal. You may return any answer, and it is guaranteed an answer exists.

**Example 1:**

    Input: [1,1,1,2,2,2]
    Output: [2,1,2,1,2,1]

**Example 2:**

    Input: [1,1,1,1,2,2,3,3]
    Output: [1,3,1,3,2,1,2,1]

**Note:**

1. `1 <= barcodes.length <= 10000`
2. `1 <= barcodes[i] <= 10000`


## 题目大意

在一个仓库里，有一排条形码，其中第 i 个条形码为 barcodes[i]。请你重新排列这些条形码，使其中两个相邻的条形码 不能 相等。 你可以返回任何满足该要求的答案，此题保证存在答案。



## 解题思路


- 这一题和第 767 题原理是完全一样的。第 767 题是 Google 的面试题。
- 解题思路比较简单，先按照每个数字的频次从高到低进行排序，注意会有频次相同的数字。排序以后，分别从第 0 号位和中间的位置开始往后取数，取完以后即为最终解。
