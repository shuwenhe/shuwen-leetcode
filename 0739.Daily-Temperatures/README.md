# [739. Daily Temperatures](https://leetcode.com/problems/daily-temperatures/)

## 题目


Given a list of daily temperatures T, return a list such that, for each day in the input, tells you how many days you would have to wait until a warmer temperature. If there is no future day for which this is possible, put 0 instead.

For example, given the list of temperatures T = [73, 74, 75, 71, 69, 72, 76, 73], your output should be [1, 1, 4, 2, 1, 1, 0, 0].

Note: The length of temperatures will be in the range [1, 30000]. Each temperature will be an integer in the range [30, 100].


## 题目大意

给出一个温度数组，要求输出比当天温度高的在未来的哪一天，输出未来第几天的天数。例如比 73 度高的在未来第 1 天出现，比 75 度高的在未来第 4 天出现。

## 解题思路

这道题根据题意正常处理就可以了。2 层循环。另外一种做法是单调栈，维护一个单调递减的单调栈即可。


