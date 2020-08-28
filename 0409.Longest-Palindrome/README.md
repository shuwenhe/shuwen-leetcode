# [409. Longest Palindrome](https://leetcode.com/problems/longest-palindrome/)


## 题目

Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.

This is case sensitive, for example `"Aa"` is not considered a palindrome here.

**Note:**Assume the length of given string will not exceed 1,010.

**Example:**

    Input:
    "abccccdd"
    
    Output:
    7
    
    Explanation:
    One longest palindrome that can be built is "dccaccd", whose length is 7.


## 题目大意

给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。注意:假设字符串的长度不会超过 1010。


## 解题思路


- 给出一个字符串，要求用这个字符串里面的字符组成一个回文串，问回文串最长可以组合成多长的？
- 这也是一题水题，先统计每个字符的频次，然后每个字符能取 2 个的取 2 个，不足 2 个的并且当前构造中的回文串是偶数的情况下(即每 2 个都配对了)，可以取 1 个。最后组合出来的就是最长回文串。
