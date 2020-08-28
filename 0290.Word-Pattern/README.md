# [290. Word Pattern](https://leetcode.com/problems/word-pattern/)

## 题目

Given a pattern and a string str, find if str follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.

Example 1:

```c
Input: pattern = "abba", str = "dog cat cat dog"
Output: true
```

Example 2:

```c
Input:pattern = "abba", str = "dog cat cat fish"
Output: false
```

Example 3:

```c
Input: pattern = "aaaa", str = "dog cat cat dog"
Output: false
```

Example 4:

```c
Input: pattern = "abba", str = "dog dog dog dog"
Output: false
```

Note:

You may assume pattern contains only lowercase letters, and str contains lowercase letters separated by a single space.


## 题目大意

给定一个模式串，判断字符串是否和给定的模式串，是一样的模式。

## 解题思路

这道题用 2 个 map 即可。1 个 map 记录模式与字符串的匹配关系，另外一个 map 记录字符串和模式的匹配关系。为什么需要记录双向的关系呢？因为 Example 4 中，a 对应了 dog，这个时候 b 如果再对应 dog 是错误的，所以这里需要从 dog 查询它是否已经和某个模式匹配过了。所以需要双向的关系。



