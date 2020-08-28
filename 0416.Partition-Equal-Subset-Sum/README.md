# [416. Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/)


## 题目

Given a **non-empty** array containing **only positive integers**, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

**Note:**

1. Each of the array element will not exceed 100.
2. The array size will not exceed 200.

**Example 1:**

    Input: [1, 5, 11, 5]
    
    Output: true
    
    Explanation: The array can be partitioned as [1, 5, 5] and [11].

**Example 2:**

    Input: [1, 2, 3, 5]
    
    Output: false
    
    Explanation: The array cannot be partitioned into equal sum subsets.


## 题目大意

给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

注意:

1. 每个数组中的元素不会超过 100
2. 数组的大小不会超过 200



## 解题思路


- 给定一个非空的数组，其中所有的数字都是正整数。问是否可以将这个数组的元素分为两部分，使得每部分的数字和相等。
- 这一题是典型的完全背包的题型。在 n 个物品中选出一定物品，完全填满 sum/2 的背包。
- `F(n,C)` 代表将 n 个物品填满容量为 C 的背包，状态转移方程为 `F(i,C) = F(i - 1,C) || F(i - 1, C - w[i])`。当 i - 1 个物品就可以填满 C ，这种情况满足题意。同时如果 i - 1 个物品不能填满背包，加上第 i 个物品以后恰好可以填满这个背包，也可以满足题意。时间复杂度 `O( n * sum/2 ) = O( n * sum)`。
