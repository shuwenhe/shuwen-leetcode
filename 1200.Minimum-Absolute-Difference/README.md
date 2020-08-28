# [1200. Minimum Absolute Difference](https://leetcode.com/problems/minimum-absolute-difference/)


## 题目

Given an array of **distinct** integers `arr`, find all pairs of elements with the minimum absolute difference of any two elements.

Return a list of pairs in ascending order(with respect to pairs), each pair `[a, b]` follows

- `a, b` are from `arr`
- `a < b`
- `b - a` equals to the minimum absolute difference of any two elements in `arr`

**Example 1:**

    Input: arr = [4,2,1,3]
    Output: [[1,2],[2,3],[3,4]]
    Explanation: The minimum absolute difference is 1. List all pairs with difference equal to 1 in ascending order.

**Example 2:**

    Input: arr = [1,3,6,10,15]
    Output: [[1,3]]

**Example 3:**

    Input: arr = [3,8,-10,23,19,-4,-14,27]
    Output: [[-14,-10],[19,23],[23,27]]

**Constraints:**

- `2 <= arr.length <= 10^5`
- `-10^6 <= arr[i] <= 10^6`


## 题目大意

给出一个数组，要求找出所有满足条件的数值对 [a,b]：`a<b` 并且 `b-a` 的差值是数组中所有两个元素差值的最小值。

## 解题思路


- 给出一个数组，要求找出所有满足条件的数值对 [a,b]：`a<b` 并且 `b-a` 的差值是数组中所有两个元素差值的最小值。
- 简单题，按照题意先排序，然后依次求出两个相邻元素的差值，求出最小的差值。最后遍历一遍数组，把所有等于最小差值的数值对都输出。
