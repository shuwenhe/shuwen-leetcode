# [1208. Get Equal Substrings Within Budget](https://leetcode.com/problems/get-equal-substrings-within-budget/)


## 题目

You are given two strings `s` and `t` of the same length. You want to change `s` to `t`. Changing the `i`-th character of `s` to `i`-th character of `t` costs `|s[i] - t[i]|` that is, the absolute difference between the ASCII values of the characters.

You are also given an integer `maxCost`.

Return the maximum length of a substring of `s` that can be changed to be the same as the corresponding substring of `t`with a cost less than or equal to `maxCost`.

If there is no substring from `s` that can be changed to its corresponding substring from `t`, return `0`.

**Example 1:**

    Input: s = "abcd", t = "bcdf", maxCost = 3
    Output: 3
    Explanation: "abc" of s can change to "bcd". That costs 3, so the maximum length is 3.

**Example 2:**

    Input: s = "abcd", t = "cdef", maxCost = 3
    Output: 1
    Explanation: Each character in s costs 2 to change to charactor in t, so the maximum length is 1.

**Example 3:**

    Input: s = "abcd", t = "acde", maxCost = 0
    Output: 1
    Explanation: You can't make any change, so the maximum length is 1.

**Constraints:**

- `1 <= s.length, t.length <= 10^5`
- `0 <= maxCost <= 10^6`
- `s` and `t` only contain lower case English letters.

## 题目大意

给你两个长度相同的字符串，s 和 t。将 s 中的第 i 个字符变到 t 中的第 i 个字符需要 |s[i] - t[i]| 的开销（开销可能为 0），也就是两个字符的 ASCII 码值的差的绝对值。

用于变更字符串的最大预算是 maxCost。在转化字符串时，总开销应当小于等于该预算，这也意味着字符串的转化可能是不完全的。如果你可以将 s 的子字符串转化为它在 t 中对应的子字符串，则返回可以转化的最大长度。如果 s 中没有子字符串可以转化成 t 中对应的子字符串，则返回 0。

提示：

- 1 <= s.length, t.length <= 10^5
- 0 <= maxCost <= 10^6
- s 和 t 都只含小写英文字母。

## 解题思路

- 给出 2 个字符串 `s` 和 `t` 和一个“预算”，要求把“预算”尽可能的花完，`s` 中最多连续有几个字母能变成 `t` 中的字母。“预算”的定义是：|s[i] - t[i]| 。
- 这一题是滑动窗口的题目，滑动窗口右边界每移动一格，就减少一定的预算，直到预算不能减少，再移动滑动窗口的左边界，这个时候注意要把预算还原回去。当整个窗口把字符 `s` 或 `t` 都滑动完了的时候，取出滑动过程中窗口的最大值即为结果。
