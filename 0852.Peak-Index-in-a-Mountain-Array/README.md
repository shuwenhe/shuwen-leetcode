# [852. Peak Index in a Mountain Array](https://leetcode.com/problems/peak-index-in-a-mountain-array/)


## 题目

Let's call an array `A` a *mountain* if the following properties hold:

- `A.length >= 3`
- There exists some `0 < i < A.length - 1` such that `A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]`

Given an array that is definitely a mountain, return any `i` such that `A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]`.

**Example 1:**

    Input: [0,1,0]
    Output: 1

**Example 2:**

    Input: [0,2,1,0]
    Output: 1

**Note:**

1. `3 <= A.length <= 10000`
2. `0 <= A[i] <= 10^6`
3. A is a mountain, as defined above.


## 题目大意

我们把符合下列属性的数组 A 称作山脉：

- A.length >= 3
- 存在 0 < i < A.length - 1 使得A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
给定一个确定为山脉的数组，返回任何满足 A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1] 的 i 的值。

提示：

- 3 <= A.length <= 10000
- 0 <= A[i] <= 10^6
- A 是如上定义的山脉
 
 

## 解题思路

- 给出一个数组，数组里面存在有且仅有一个“山峰”，(山峰的定义是，下标 `i` 比 `i-1`、`i+1` 位置上的元素都要大)，找到这个“山峰”，并输出其中一个山峰的下标。
- 这一题直接用二分搜索即可，数组中的元素算基本有序。判断是否为山峰的条件为比较左右两个数，如果当前的数比左右两个数都大，即找到了山峰。其他的情况都在山坡上。这一题有两种写法，第一种写法是标准的二分写法，第二种写法是变形的二分写法。
