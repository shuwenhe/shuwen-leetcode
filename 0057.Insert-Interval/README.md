# [57. Insert Interval](https://leetcode.com/problems/insert-interval/)

## 题目

Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

Example 1:  

```
Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
```

Example 2:  

```
Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
```

## 题目大意

这一题是第 56 题的加强版。给出多个没有重叠的区间，然后再给一个区间，要求把如果有重叠的区间进行合并。

## 解题思路

可以分 3 段处理，先添加原来的区间，即在给的 newInterval 之前的区间。然后添加 newInterval ，注意这里可能需要合并多个区间。最后把原来剩下的部分添加到最终结果中即可。