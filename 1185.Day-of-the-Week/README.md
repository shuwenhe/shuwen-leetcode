# [1185. Day of the Week](https://leetcode.com/problems/day-of-the-week/)


## 题目

Given a date, return the corresponding day of the week for that date.

The input is given as three integers representing the `day`, `month` and `year` respectively.

Return the answer as one of the following values `{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}`.

**Example 1:**

    Input: day = 31, month = 8, year = 2019
    Output: "Saturday"

**Example 2:**

    Input: day = 18, month = 7, year = 1999
    Output: "Sunday"

**Example 3:**

    Input: day = 15, month = 8, year = 1993
    Output: "Sunday"

**Constraints:**

- The given dates are valid dates between the years `1971` and `2100`.


## 题目大意

给你一个日期，请你设计一个算法来判断它是对应一周中的哪一天。输入为三个整数：day、month 和 year，分别表示日、月、年。

您返回的结果必须是这几个值中的一个 {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}。

提示：

- 给出的日期一定是在 1971 到 2100 年之间的有效日期。

## 解题思路


- 给出一个日期，要求算出这一天是星期几。
- 简单题，按照常识计算即可。
