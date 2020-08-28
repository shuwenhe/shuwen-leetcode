# [49. Group Anagrams](https://leetcode.com/problems/group-anagrams/)

## 题目

Given an array of strings, group anagrams together.


Example:

```c
Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
```

Note:

- All inputs will be in lowercase.
- The order of your output does not matter.

## 题目大意

给出一个字符串数组，要求对字符串数组里面有 Anagrams 关系的字符串进行分组。Anagrams 关系是指两个字符串的字符完全相同，顺序不同，两者是由排列组合组成。

## 解题思路

这道题可以将每个字符串都排序，排序完成以后，相同 Anagrams 的字符串必然排序结果一样。把排序以后的字符串当做 key 存入到 map 中。遍历数组以后，就能得到一个 map，key 是排序以后的字符串，value 对应的是这个排序字符串以后的 Anagrams 字符串集合。最后再将这些 value 对应的字符串数组输出即可。
