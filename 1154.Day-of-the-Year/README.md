# [1154. Day of the Year](https://leetcode.com/problems/day-of-the-year/)


## 题目

Given a string `date` representing a [Gregorian calendar](https://en.wikipedia.org/wiki/Gregorian_calendar) date formatted as `YYYY-MM-DD`, return the day number of the year.

**Example 1:**

    Input: date = "2019-01-09"
    Output: 9
    Explanation: Given date is the 9th day of the year in 2019.

**Example 2:**

    Input: date = "2019-02-10"
    Output: 41

**Example 3:**

    Input: date = "2003-03-01"
    Output: 60

**Example 4:**

    Input: date = "2004-03-01"
    Output: 61

**Constraints:**

- `date.length == 10`
- `date[4] == date[7] == '-'`, and all other `date[i]`'s are digits
- `date` represents a calendar date between Jan 1st, 1900 and Dec 31, 2019.

## 题目大意


实现一个 MajorityChecker 的类，它应该具有下述几个 API：

- MajorityChecker(int[] arr) 会用给定的数组 arr 来构造一个 MajorityChecker 的实例。
- int query(int left, int right, int threshold) 有这么几个参数：  
	- 0 <= left <= right < arr.length 表示数组 arr 的子数组的长度。
	- 2 * threshold > right - left + 1，也就是说阈值 threshold 始终比子序列长度的一半还要大。

每次查询 query(...) 会返回在 arr[left], arr[left+1], ..., arr[right] 中至少出现阈值次数 threshold 的元素，如果不存在这样的元素，就返回 -1。

提示：

- 1 <= arr.length <= 20000
- 1 <= arr[i] <= 20000
- 对于每次查询，0 <= left <= right < len(arr)
- 对于每次查询，2 * threshold > right - left + 1
- 查询次数最多为 10000





## 解题思路

- 给出一个时间字符串，求出这一天是这一年当中的第几天。
- 简单题。依照题意处理即可。
