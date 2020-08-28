# [435. Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/)


## 题目

Given a collection of intervals, find the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

**Note:**

1. You may assume the interval's end point is always bigger than its start point.
2. Intervals like [1,2] and [2,3] have borders "touching" but they don't overlap each other.

**Example 1:**

    Input: [ [1,2], [2,3], [3,4], [1,3] ]
    
    Output: 1
    
    Explanation: [1,3] can be removed and the rest of intervals are non-overlapping.

**Example 2:**

    Input: [ [1,2], [1,2], [1,2] ]
    
    Output: 2
    
    Explanation: You need to remove two [1,2] to make the rest of intervals non-overlapping.

**Example 3:**

    Input: [ [1,2], [2,3] ]
    
    Output: 0
    
    Explanation: You don't need to remove any of the intervals since they're already non-overlapping.

**NOTE:** input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.


## 题目大意

给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。

注意:

1. 可以认为区间的终点总是大于它的起点。
2. 区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。



## 解题思路


- 给定一组区间，问最少删除多少个区间，可以让这些区间之间互相不重叠。注意，给定区间的起始点永远小于终止点。[1,2] 和 [2,3] 不叫重叠。
- 这一题可以反过来考虑，给定一组区间，问最多保留多少区间，可以让这些区间之间相互不重叠。先排序，判断区间是否重叠。
- 这一题一种做法是利用动态规划，模仿最长上升子序列的思想，来解题。
- 这道题另外一种做法是按照区间的结尾进行排序，每次选择结尾最早的，且和前一个区间不重叠的区间。选取结尾最早的，就可以给后面留出更大的空间，供后面的区间选择。这样可以保留更多的区间。这种做法是贪心算法的思想。

