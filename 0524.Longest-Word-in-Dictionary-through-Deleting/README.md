# [524. Longest Word in Dictionary through Deleting](https://leetcode.com/problems/longest-word-in-dictionary-through-deleting/)

## 题目

Given a string and a string dictionary, find the longest string in the dictionary that can be formed by deleting some characters of the given string. If there are more than one possible results, return the longest word with the smallest lexicographical order. If there is no possible result, return the empty string.


Example 1:

```c
Input:
s = "abpcplea", d = ["ale","apple","monkey","plea"]

Output: 
"apple"
```


Example 2:

```c
Input:
s = "abpcplea", d = ["a","b","c"]

Output: 
"a"
```

Note:  

- All the strings in the input will only contain lower-case letters.
- The size of the dictionary won't exceed 1,000.
- The length of all the strings in the input won't exceed 1,000.


## 题目大意


给出一个初始串，再给定一个字符串数组，要求在字符串数组中找到能在初始串中通过删除字符得到的最长的串，如果最长的串有多组解，要求输出字典序最小的那组解。

## 解题思路


这道题就单纯的用 O(n^2) 暴力循环即可，注意最终解的要求，如果都是最长的串，要求输出字典序最小的那个串，只要利用字符串比较得到字典序最小的串即可。

