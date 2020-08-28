# [976. Largest Perimeter Triangle](https://leetcode.com/problems/largest-perimeter-triangle/)

## 题目

Given an array A of positive lengths, return the largest perimeter of a triangle with non-zero area, formed from 3 of these lengths.

If it is impossible to form any triangle of non-zero area, return 0.


Example 1:

```c
Input: [2,1,2]
Output: 5
```

Example 2:

```c
Input: [1,2,1]
Output: 0
```

Example 3:

```c
Input: [3,2,3,4]
Output: 10
```

Example 4:

```c
Input: [3,6,2,3]
Output: 8
```

Note:

- 3 <= A.length <= 10000
- 1 <= A[i] <= 10^6

## 题目大意

找到可以组成三角形三条边的长度，要求输出三条边之和最长的，即三角形周长最长。

## 解题思路

这道题也是排序题，先讲所有的长度进行排序，从大边开始往前找，找到第一个任意两边之和大于第三边(满足能构成三角形的条件)的下标，然后输出这 3 条边之和即可，如果没有找到输出 0 。