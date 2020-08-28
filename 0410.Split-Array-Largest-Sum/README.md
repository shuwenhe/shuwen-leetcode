# [410. Split Array Largest Sum](https://leetcode.com/problems/split-array-largest-sum/)


## 题目

Given an array which consists of non-negative integers and an integer m, you can split the array into m non-empty continuous subarrays. Write an algorithm to minimize the largest sum among these m subarrays.

**Note:**If n is the length of array, assume the following constraints are satisfied:

- 1 ≤ n ≤ 1000
- 1 ≤ m ≤ min(50, n)

**Examples:**

    Input:
    nums = [7,2,5,10,8]
    m = 2
    
    Output:
    18
    
    Explanation:
    There are four ways to split nums into two subarrays.
    The best way is to split it into [7,2,5] and [10,8],
    where the largest sum among the two subarrays is only 18.


## 题目大意


给定一个非负整数数组和一个整数 m，你需要将这个数组分成 m 个非空的连续子数组。设计一个算法使得这 m 个子数组各自和的最大值最小。

注意:
数组长度 n 满足以下条件:

- 1 ≤ n ≤ 1000
- 1 ≤ m ≤ min(50, n)



## 解题思路

- 给出一个数组和分割的个数 M。要求把数组分成 M 个子数组，输出子数组和的最大值。
- 这一题可以用动态规划 DP 解答，也可以用二分搜索来解答。这一题是二分搜索里面的 max-min 最大最小值问题。题目可以转化为在 `M` 次划分中，求一个 `x`，使得 `x` 满足：对任意的`S(i)`，都满足 `S(i) ≤ x`。这个条件保证了 `x` 是所有 `S(i)` 中的最大值。要求的是满足该条件的最小的 `x`。`x` 的搜索范围在 `[max, sum]` 中。逐步二分逼近 low 值，直到找到能满足条件的 low 的最小值，即为最终答案。
