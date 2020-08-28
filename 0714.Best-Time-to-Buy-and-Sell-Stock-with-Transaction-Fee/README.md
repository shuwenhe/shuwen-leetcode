# [714. Best Time to Buy and Sell Stock with Transaction Fee](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/)


## 题目

Your are given an array of integers `prices`, for which the `i`-th element is the price of a given stock on day `i`; and a non-negative integer `fee` representing a transaction fee.

You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction. You may not buy more than 1 share of a stock at a time (ie. you must sell the stock share before you buy again.)

Return the maximum profit you can make.

**Example 1:**

    Input: prices = [1, 3, 2, 8, 4, 9], fee = 2
    Output: 8
    Explanation: The maximum profit can be achieved by:
    Buying at prices[0] = 1
    Selling at prices[3] = 8
    Buying at prices[4] = 4
    Selling at prices[5] = 9
    The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.

**Note:**

- `0 < prices.length <= 50000`.
- `0 < prices[i] < 50000`.
- `0 <= fee < 50000`.


## 题目大意

给定一个整数数组 prices，其中第 i 个元素代表了第 i 天的股票价格 ；非负整数 fee 代表了交易股票的手续费用。你可以无限次地完成交易，但是你每次交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。要求返回获得利润的最大值。



## 解题思路

- 给定一个数组，表示一支股票在每一天的价格。设计一个交易算法，在这些天进行自动交易，要求：每一天只能进行一次操作；在买完股票后，必须卖了股票，才能再次买入；每次卖了股票以后，需要缴纳一部分的手续费。问如何交易，能让利润最大？
- 这一题是第 121 题、第 122 题、第 309 题的变种题。
- 这一题的解题思路是 DP，需要维护买和卖的两种状态。`buy[i]` 代表第 `i` 天买入的最大收益，`sell[i]` 代表第 `i` 天卖出的最大收益，状态转移方程是 `buy[i] = max(buy[i-1], sell[i-1]-prices[i])`，`sell[i] = max(sell[i-1], buy[i-1]+prices[i]-fee)`。
