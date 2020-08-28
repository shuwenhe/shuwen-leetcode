# [844. Backspace String Compare](https://leetcode.com/problems/backspace-string-compare/)

## 题目

Given two strings S and T, return if they are equal when both are typed into empty text editors. # means a backspace character.


Example 1:

```c
Input: S = "ab#c", T = "ad#c"
Output: true
Explanation: Both S and T become "ac".
```

Example 2:

```c
Input: S = "ab##", T = "c#d#"
Output: true
Explanation: Both S and T become "".
```

Example 3:

```c
Input: S = "a##c", T = "#a#c"
Output: true
Explanation: Both S and T become "c".
```

Example 4:

```c
Input: S = "a#c", T = "b"
Output: false
Explanation: S becomes "c" while T becomes "b".
```


Note:

- 1 <= S.length <= 200
- 1 <= T.length <= 200
- S and T only contain lowercase letters and '#' characters.


Follow up:

- Can you solve it in O(N) time and O(1) space?

## 题目大意


给 2 个字符串，如果遇到 # 号字符，就回退一个字符。问最终的 2 个字符串是否完全一致。

## 解题思路

这一题可以用栈的思想来模拟，遇到 # 字符就回退一个字符。不是 # 号就入栈一个字符。比较最终 2 个字符串即可。











