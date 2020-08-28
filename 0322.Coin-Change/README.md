# [322. Coin Change](https://leetcode.com/problems/coin-change/)


## 题目

You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return `-1`.

**Example 1:**

    Input: coins = [1, 2, 5], amount = 11
    Output: 3 
    Explanation: 11 = 5 + 5 + 1

**Example 2:**

    Input: coins = [2], amount = 3
    Output: -1

**Note**:

You may assume that you have an infinite number of each kind of coin.

## 题目大意

给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。



## 解题思路

- 给出一些硬币和一个总数，问组成这个总数的硬币数最少是多少个？
- 这一题是经典的硬币问题，利用 DP 求解。不过这一题的测试用例有一个很大的值，这样开 DP 数组会比较浪费空间。例如 [1,1000000000,500000] 有这样的硬币种类，要求组成 2389412493027523 这样的总数。那么按照下面的解题方法，数组会开的很大，非常浪费空间。这个时候用 DFS 解题会节约一点空间。

