# [528. Random Pick with Weight](https://leetcode.com/problems/random-pick-with-weight/)


## 题目

Given an array `w` of positive integers, where `w[i]` describes the weight of index `i`, write a function `pickIndex` which randomly picks an index in proportion to its weight.

Note:

1. `1 <= w.length <= 10000`
2. `1 <= w[i] <= 10^5`
3. `pickIndex` will be called at most `10000` times.

**Example 1:**

    Input: 
    ["Solution","pickIndex"]
    [[[1]],[]]
    Output: [null,0]

**Example 2:**

    Input: 
    ["Solution","pickIndex","pickIndex","pickIndex","pickIndex","pickIndex"]
    [[[1,3]],[],[],[],[],[]]
    Output: [null,0,1,1,1,0]

**Explanation of Input Syntax:**

The input is two lists: the subroutines called and their arguments. `Solution`'s constructor has one argument, the array `w`. `pickIndex` has no arguments. Arguments are always wrapped with a list, even if there aren't any.


## 题目大意

给定一个正整数数组 w ，其中 w[i] 代表位置 i 的权重，请写一个函数 pickIndex ，它可以随机地获取位置 i，选取位置 i 的概率与 w[i] 成正比。

说明:

1. 1 <= w.length <= 10000
2. 1 <= w[i] <= 10^5
3. pickIndex 将被调用不超过 10000 次


输入语法说明：

输入是两个列表：调用成员函数名和调用的参数。Solution 的构造函数有一个参数，即数组 w。pickIndex 没有参数。输入参数是一个列表，即使参数为空，也会输入一个 [] 空列表。



## 解题思路

- 给出一个数组，每个元素值代表该下标的权重值，`pickIndex()` 随机取一个位置 i，这个位置出现的概率和该元素值成正比。
- 由于涉及到了权重的问题，这一题可以先考虑用前缀和处理权重。在 `[0,prefixSum)` 区间内随机选一个整数 `x`，下标 `i` 是满足 `x< prefixSum[i]` 条件的最小下标，求这个下标 `i` 即是最终解。二分搜索查找下标 `i` 。对于某些下标 `i`，所有满足 `prefixSum[i] - w[i] ≤ v < prefixSum[i]` 的整数 `v` 都映射到这个下标。因此，所有的下标都与下标权重成比例。
- 时间复杂度：预处理的时间复杂度是 O(n)，`pickIndex()` 的时间复杂度是 O(log n)。空间复杂度 O(n)。
