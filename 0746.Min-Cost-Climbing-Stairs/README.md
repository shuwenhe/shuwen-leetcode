# [746. Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/)


## 题目

On a staircase, the `i`-th step has some non-negative cost `cost[i]` assigned (0 indexed).

Once you pay the cost, you can either climb one or two steps. You need to find minimum cost to reach the top of the floor, and you can either start from the step with index 0, or the step with index 1.

**Example 1:**

    Input: cost = [10, 15, 20]
    Output: 15
    Explanation: Cheapest is start on cost[1], pay that cost and go to the top.

**Example 2:**

    Input: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
    Output: 6
    Explanation: Cheapest is start on cost[0], and only step on 1s, skipping cost[3].

**Note:**

1. `cost` will have a length in the range `[2, 1000]`.
2. Every `cost[i]` will be an integer in the range `[0, 999]`.


## 题目大意

数组的每个索引做为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost\[i\] (索引从 0 开始)。每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。


## 解题思路


- 这一题算是第 70 题的加强版。依旧是爬楼梯的问题，解题思路也是 DP。在爬楼梯的基础上增加了一个新的条件，每层楼梯都有一个 cost 花费，问上到最终楼层，花费最小值是多少。
- `dp[i]` 代表上到第 n 层的最小花费，状态转移方程是 `dp[i] = cost[i] + min(dp[i-2], dp[i-1])`，最终第 n 层的最小花费是 `min(dp[n-2], dp[n-1])` 。
- 由于每层的花费只和前两层有关系，所以每次 DP 迭代的时候只需要 2 个临时变量即可。可以用这种方式来优化辅助空间。

