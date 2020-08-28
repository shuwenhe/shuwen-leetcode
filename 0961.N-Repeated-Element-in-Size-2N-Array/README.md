# [961. N-Repeated Element in Size 2N Array](https://leetcode.com/problems/n-repeated-element-in-size-2n-array/)


## 题目

In a array `A` of size `2N`, there are `N+1` unique elements, and exactly one of these elements is repeated N times.

Return the element repeated `N` times.

**Example 1:**

    Input: [1,2,3,3]
    Output: 3

**Example 2:**

    Input: [2,1,2,5,3,2]
    Output: 2

**Example 3:**

    Input: [5,1,5,2,5,3,5,4]
    Output: 5

**Note:**

1. `4 <= A.length <= 10000`
2. `0 <= A[i] < 10000`
3. `A.length` is even


## 题目大意

在大小为 2N 的数组 A 中有 N+1 个不同的元素，其中有一个元素重复了 N 次。返回重复了 N 次的那个元素。


## 解题思路


- 简单题。数组中有 2N 个数，有 N + 1 个数是不重复的，这之中有一个数重复了 N 次，请找出这个数。解法非常简单，把所有数都存入 map 中，如果遇到存在的 key 就返回这个数。
