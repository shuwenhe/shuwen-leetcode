# [309. Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/)


## 题目

Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (ie, buy one and sell one share of the stock multiple times) with the following restrictions:

- You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
- After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)

**Example:**

    Input: [1,2,3,0,2]
    Output: 3 
    Explanation: transactions = [buy, sell, cooldown, buy, sell]

## 题目大意

给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

- 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
- 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。



## 解题思路

- 给定一个数组，表示一支股票在每一天的价格。设计一个交易算法，在这些天进行自动交易，要求：每一天只能进行一次操作；在买完股票后，必须卖了股票，才能再次买入；每次卖了股票以后，在下一天是不能购买的。问如何交易，能让利润最大？
- 这一题是第 121 题和第 122 题的变种题。
- 每天都有 3 种操作，`buy`，`sell`，`cooldown`。`sell` 之后的一天一定是 `cooldown`，但是 `cooldown` 可以出现在任意一天。例如：`buy，cooldown，cooldown，sell，cooldown，cooldown`。`buy[i]` 代表第 `i` 天通过 `buy` 或者 `cooldown` 结束此天能获得的最大收益。例如：`buy, sell, buy` 或者 `buy, cooldown, cooldown`。`sell[i]` 代表第 `i` 天通过 `sell` 或者 `cooldown` 结束此天能获得的最大收益。例如：`buy, sell, buy, sell` 或者 `buy, sell, cooldown, cooldown`。`price[i-1]` 代表第 `i` 天的股票价格(由于 price 是从 0 开始的)。
- 第 i 天如果是 sell，那么这天能获得的最大收益是 `buy[i - 1] + price[i - 1]`，因为只有 buy 了才能 sell。如果这一天是 cooldown，那么这天能获得的最大收益还是 sell[i - 1]。所以 sell[i] 的状态转移方程 `sell[i] = max(buy[i - 1] + price[i - 1], sell[i - 1])`。`sell[0] = 0` 代表第一天就卖了，由于第一天不持有股票，所以 sell[0] = 0。`sell[1] = max(sell[0], buy[0]+prices[1])` 代表第一天卖了，和第一天不卖，第二天卖做对比，钱多的保存至 sell[1]。
- 第 i 天如果是 buy，那么这天能获得的最大收益是 `sell[i - 2] - price[i - 1]`，因为 i - 1 天是 cooldown。如果这一天是 cooldown，那么这天能获得的最大收益还是 buy[i - 1]。所以 buy[i] 的状态转移方程 `buy[i] = max(sell[i - 2] - price[i - 1], buy[i - 1])`。`buy[0] = -prices[0]` 代表第一天就买入，所以金钱变成了负的。`buy[1] = max(buy[0], -prices[1])` 代表第一天不买入，第二天再买入。

