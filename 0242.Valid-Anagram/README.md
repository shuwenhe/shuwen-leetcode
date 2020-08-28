# [242. Valid Anagram](https://leetcode.com/problems/valid-anagram/)

## 题目

Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

```c
Input: s = "anagram", t = "nagaram"
Output: true
```

Example 2:

```c
Input: s = "rat", t = "car"
Output: false
```

Note:  

  
You may assume the string contains only lowercase alphabets.
  
  
    
Follow up:  


What if the inputs contain unicode characters? How would you adapt your solution to such case?

## 题目大意

给出 2 个字符串 s 和 t，如果 t 中的字母在 s 中都存在，输出 true，否则输出 false。

## 解题思路

这道题可以用打表的方式做。先把 s 中的每个字母都存在一个 26 个容量的数组里面，每个下标依次对应 26 个字母。s 中每个字母都对应表中一个字母，每出现一次就加 1。然后再扫字符串 t，每出现一个字母就在表里面减一。如果都出现了，最终表里面的值肯定都是 0 。最终判断表里面的值是否都是 0 即可，有非 0 的数都输出 false 。