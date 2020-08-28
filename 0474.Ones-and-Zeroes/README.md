# [474. Ones and Zeroes](https://leetcode.com/problems/ones-and-zeroes/)


## 题目

In the computer world, use restricted resource you have to generate maximum benefit is what we always want to pursue.

For now, suppose you are a dominator of **m** `0s` and **n** `1s` respectively. On the other hand, there is an array with strings consisting of only `0s` and `1s`.

Now your task is to find the maximum number of strings that you can form with given **m**`0s` and **n** `1s`. Each `0` and `1` can be used at most **once**.

**Note:**

1. The given numbers of `0s` and `1s` will both not exceed `100`
2. The size of given string array won't exceed `600`.

**Example 1:**

    Input: Array = {"10", "0001", "111001", "1", "0"}, m = 5, n = 3
    Output: 4
    
    Explanation: This are totally 4 strings can be formed by the using of 5 0s and 3 1s, which are “10,”0001”,”1”,”0”

**Example 2:**

    Input: Array = {"10", "0", "1"}, m = 1, n = 1
    Output: 2
    
    Explanation: You could form "10", but then you'd have nothing left. Better form "0" and "1".


## 题目大意

在计算机界中，我们总是追求用有限的资源获取最大的收益。现在，假设你分别支配着 m 个 0 和 n 个 1。另外，还有一个仅包含 0 和 1 字符串的数组。你的任务是使用给定的 m 个 0 和 n 个 1 ，找到能拼出存在于数组中的字符串的最大数量。每个 0 和 1 至多被使用一次。

注意:

1. 给定 0 和 1 的数量都不会超过 100。
2. 给定字符串数组的长度不会超过 600。



## 解题思路

- 给定一个字符串数组和 m，n，其中所有的字符都是由 0 和 1 组成的。问能否从数组中取出最多的字符串，使得这些取出的字符串中所有的 0 的个数 ≤ m，1 的个数 ≤ n。
- 这一题是典型的 01 背包的题型。只不过是一个二维的背包问题。在 n 个物品中选出一定物品，**尽量完全填满** m 维和 n 维的背包。为什么是尽量填满？因为不一定能完全填满背包。
- `dp[i][j]` 代表尽量填满容量为 `(i,j)` 的背包装下的物品总数，状态转移方程为 `dp[i][j] = max(dp[i][j], 1+dp[i-zero][j-one])`。其中 zero 代表的当前装入物品在 m 维上的体积，也即 0 的个数。one 代表的是当前装入物品在 n 维上的体积，也即 1 的个数。每添加一个物品，比较当前 (i,j) 的背包装下的物品总数和 (i-zero,j-one) 的背包装下的物品总数 + 1，比较这两者的大小，保存两者的最大值。每添加一个物品就刷新这个二维背包，直到所有物品都扫完一遍。dp[m][n] 中存储的就是最终的答案。时间复杂度 `O( n * M * N )`。
