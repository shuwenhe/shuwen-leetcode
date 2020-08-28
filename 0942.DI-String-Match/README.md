# [942. DI String Match](https://leetcode.com/problems/di-string-match/)


## 题目

Given a string `S` that **only** contains "I" (increase) or "D" (decrease), let `N = S.length`.

Return **any** permutation `A` of `[0, 1, ..., N]` such that for all `i = 0, ..., N-1`:

- If `S[i] == "I"`, then `A[i] < A[i+1]`
- If `S[i] == "D"`, then `A[i] > A[i+1]`

**Example 1:**

    Input: "IDID"
    Output: [0,4,1,3,2]

**Example 2:**

    Input: "III"
    Output: [0,1,2,3]

**Example 3:**

    Input: "DDI"
    Output: [3,2,0,1]

**Note:**

1. `1 <= S.length <= 10000`
2. `S` only contains characters `"I"` or `"D"`.


## 题目大意

给定只含 "I"（增大）或 "D"（减小）的字符串 S ，令 N = S.length。返回 [0, 1, ..., N] 的任意排列 A 使得对于所有 i = 0, ..., N-1，都有：

- 如果 S[i] == "I"，那么 A[i] < A[i+1]
- 如果 S[i] == "D"，那么 A[i] > A[i+1]



## 解题思路


- 给出一个字符串，字符串中只有字符 `"I"` 和字符 `"D"`。字符 `"I"` 代表 `A[i] < A[i+1]`，字符 `"D"` 代表 `A[i] > A[i+1]` ，要求找到满足条件的任意组合。
- 这一题也是水题，取出字符串长度即是最大数的数值，然后按照题意一次排出最终数组即可。

