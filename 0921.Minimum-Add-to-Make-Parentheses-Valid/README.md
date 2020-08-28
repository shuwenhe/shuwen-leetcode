# [921. Minimum Add to Make Parentheses Valid](https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/)

## 题目

Given a string S of '(' and ')' parentheses, we add the minimum number of parentheses ( '(' or ')', and in any positions ) so that the resulting parentheses string is valid.

Formally, a parentheses string is valid if and only if:

- It is the empty string, or
- It can be written as AB (A concatenated with B), where A and B are valid strings, or
- It can be written as (A), where A is a valid string.
Given a parentheses string, return the minimum number of parentheses we must add to make the resulting string valid.

 

Example 1:

```c
Input: "())"
Output: 1

```

Example 2:

```c
Input: "((("
Output: 3
```

Example 3:

```c
Input: "()"
Output: 0
```

Example 4:

```c
Input: "()))(("
Output: 4
```

Note:

1. S.length <= 1000
2. S only consists of '(' and ')' characters.

## 题目大意

给一个括号的字符串，如果能在这个括号字符串中的任意位置添加括号，问能使得这串字符串都能完美匹配的最少添加数是多少。

## 解题思路

这题也是栈的题目，利用栈进行括号匹配。最后栈里剩下几个括号，就是最少需要添加的数目。