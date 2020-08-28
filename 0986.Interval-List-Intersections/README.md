# [986. Interval List Intersections](https://leetcode.com/problems/interval-list-intersections/)

## 题目

Given two lists of closed intervals, each list of intervals is pairwise disjoint and in sorted order.

Return the intersection of these two interval lists.

(Formally, a closed interval [a, b] (with a <= b) denotes the set of real numbers x with a <= x <= b.  The intersection of two closed intervals is a set of real numbers that is either empty, or can be represented as a closed interval.  For example, the intersection of [1, 3] and [2, 4] is [2, 3].)




Example 1:

![](https://assets.leetcode.com/uploads/2019/01/30/interval1.png)

```c
Input: A = [[0,2],[5,10],[13,23],[24,25]], B = [[1,5],[8,12],[15,24],[25,26]]
Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
Reminder: The inputs and the desired output are lists of Interval objects, and not arrays or lists.
```

Note:  

- 0 <= A.length < 1000
- 0 <= B.length < 1000
- 0 <= A[i].start, A[i].end, B[i].start, B[i].end < 10^9

NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

## 题目大意

这道题考察的是滑动窗口的问题。

给出 2 个数组 A 和数组 B。要求求出这 2 个数组的交集数组。题意见图。

## 解题思路

交集的左边界应该为，`start := max(A[i].Start, B[j].Start)`，右边界为，`end := min(A[i].End, B[j].End)`，如果 `start <= end`，那么这个就是一个满足条件的交集，放入最终数组中。如果 `A[i].End <= B[j].End`，代表 B 数组范围比 A 数组大，A 的游标右移。如果 `A[i].End > B[j].End`，代表 A 数组范围比 B 数组大，B 的游标右移。
