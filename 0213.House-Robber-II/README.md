# [213. House Robber II](https://leetcode.com/problems/house-robber-ii/)


## 题目

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed. All houses at this place are **arranged in a circle.** That means the first house is the neighbor of the last one. Meanwhile, adjacent houses have security system connected and **it will automatically contact the police if two adjacent houses were broken into on the same night**.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight **without alerting the police**.

**Example 1:**

    Input: [2,3,2]
    Output: 3
    Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
                 because they are adjacent houses.

**Example 2:**

    Input: [1,2,3,1]
    Output: 4
    Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
                 Total amount you can rob = 1 + 3 = 4.

## 题目大意

你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都**围成一圈**，这意味着第一个房屋和最后一个房屋是紧挨着的。同时，相邻的房屋装有相互连通的防盗系统，**如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警**。

给定一个代表每个房屋存放金额的非负整数数组，计算你**在不触动警报装置的情况下**，能够偷窃到的最高金额。


## 解题思路

- 这一题是第 198 题的加强版。不过这次是在一个环形的街道中，即最后一个元素和第一个元素是邻居，在不触碰警报的情况下，问能够窃取的财产的最大值是多少？
- 解题思路和第 198 完全一致，只需要增加额外的一个转换。由于首尾是相邻的，所以在取了第一个房子以后就不能取第 n 个房子，那么就在 [0,n - 1] 的区间内找出总价值最多的解，然后再 [1,n] 的区间内找出总价值最多的解，两者取最大值即可。

