# [91. Decode Ways](https://leetcode.com/problems/decode-ways/)


## 题目

A message containing letters from `A-Z` is being encoded to numbers using the following mapping:

    'A' -> 1
    'B' -> 2
    ...
    'Z' -> 26

Given a **non-empty** string containing only digits, determine the total number of ways to decode it.

**Example 1:**

    Input: "12"
    Output: 2
    Explanation: It could be decoded as "AB" (1 2) or "L" (12).

**Example 2:**

    Input: "226"
    Output: 3
    Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

## 题目大意

一条包含字母 A-Z 的消息通过以下方式进行了编码：

```c
'A' -> 1
'B' -> 2
...
'Z' -> 26
```

给定一个只包含数字的非空字符串，请计算解码方法的总数。



## 解题思路

- 给出一个数字字符串，题目要求把数字映射成 26 个字母，映射以后问有多少种可能的翻译方法。
- 这题思路也是 DP。`dp[n]` 代表翻译长度为 n 个字符的字符串的方法总数。由于题目中的数字可能出现 0，0 不能翻译成任何字母，所以出现 0 要跳过。dp[0] 代表空字符串，只有一种翻译方法，`dp[0] = 1`。dp[1] 需要考虑原字符串是否是 0 开头的，如果是 0 开头的，`dp[1] = 0`，如果不是 0 开头的，`dp[1] = 1`。状态转移方程是 `dp[i] += dp[i-1] (当 1 ≤ s[i-1 : i] ≤ 9)；dp[i] += dp[i-2] (当 10 ≤ s[i-2 : i] ≤ 26)`。最终结果是 `dp[n]`。
