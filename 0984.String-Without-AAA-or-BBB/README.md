# [984. String Without AAA or BBB](https://leetcode.com/problems/string-without-aaa-or-bbb/)


## 题目

Given two integers `A` and `B`, return **any** string `S` such that:

- `S` has length `A + B` and contains exactly `A` `'a'` letters, and exactly `B` `'b'`letters;
- The substring `'aaa'` does not occur in `S`;
- The substring `'bbb'` does not occur in `S`.

**Example 1:**

    Input: A = 1, B = 2
    Output: "abb"
    Explanation: "abb", "bab" and "bba" are all correct answers.

**Example 2:**

    Input: A = 4, B = 1
    Output: "aabaa"

**Note:**

1. `0 <= A <= 100`
2. `0 <= B <= 100`
3. It is guaranteed such an `S` exists for the given `A` and `B`.


## 题目大意

给定两个整数 A 和 B，返回任意字符串 S，要求满足：

- S 的长度为 A + B，且正好包含 A 个 'a' 字母与 B 个 'b' 字母；
- 子串 'aaa' 没有出现在 S 中；
- 子串 'bbb' 没有出现在 S 中。


提示：

- 0 <= A <= 100
- 0 <= B <= 100
- 对于给定的 A 和 B，保证存在满足要求的 S。


## 解题思路


- 给出 A 和 B 的个数，要求组合出字符串，不能出现 3 个连续的 A 和 3 个连续的 B。这题由于只可能有 4 种情况，暴力枚举就可以了。假设 B 的个数比 A 多(如果 A 多，就交换一下 A 和 B)，最终可能的情况只可能是这 4 种情况中的一种： `ba`，`bbabb`，`bbabbabb`，`bbabbabbabbabababa`。
