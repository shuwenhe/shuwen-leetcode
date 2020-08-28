# [1017. Convert to Base -2](https://leetcode.com/problems/convert-to-base-2/)


## 题目

Given a number `N`, return a string consisting of `"0"`s and `"1"`s that represents its value in base **`-2`** (negative two).

The returned string must have no leading zeroes, unless the string is `"0"`.

**Example 1:**

    Input: 2
    Output: "110"
    Explantion: (-2) ^ 2 + (-2) ^ 1 = 2

**Example 2:**

    Input: 3
    Output: "111"
    Explantion: (-2) ^ 2 + (-2) ^ 1 + (-2) ^ 0 = 3

**Example 3:**

    Input: 4
    Output: "100"
    Explantion: (-2) ^ 2 = 4

**Note:**

1. `0 <= N <= 10^9`


## 题目大意

给出数字 N，返回由若干 "0" 和 "1"组成的字符串，该字符串为 N 的负二进制（base -2）表示。除非字符串就是 "0"，否则返回的字符串中不能含有前导零。

提示：

- 0 <= N <= 10^9



## 解题思路

- 给出一个十进制的数，要求转换成 -2 进制的数
- 这一题仿造十进制转二进制的思路，短除法即可。

