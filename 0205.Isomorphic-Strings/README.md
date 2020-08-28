# [205. Isomorphic Strings](https://leetcode.com/problems/isomorphic-strings/)

## 题目

Given two strings s and t, determine if they are isomorphic.

Two strings are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.

Example 1:

```c
Input: s = "egg", t = "add"
Output: true
```

Example 2:

```c
Input: s = "foo", t = "bar"
Output: false
```

Example 3:

```c
Input: s = "paper", t = "title"
Output: true
```

Note:   

You may assume both s and t have the same length.




## 题目大意

这道题和第 290 题基本是一样的。第 290 题是模式匹配，这道题的题意是字符串映射，实质是一样的。

给定一个初始字符串串，判断初始字符串是否可以通过字符映射的方式，映射到目标字符串，如果可以映射，则输出 true，反之输出 false。

## 解题思路

这道题做法和第 290 题基本一致。



