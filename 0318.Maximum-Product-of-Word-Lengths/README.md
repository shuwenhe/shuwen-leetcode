# [318. Maximum Product of Word Lengths](https://leetcode.com/problems/maximum-product-of-word-lengths/)


## 题目

Given a string array `words`, find the maximum value of `length(word[i]) * length(word[j])` where the two words do not share common letters. You may assume that each word will contain only lower case letters. If no such two words exist, return 0.

**Example 1:**

    Input: ["abcw","baz","foo","bar","xtfn","abcdef"]
    Output: 16 
    Explanation: The two words can be "abcw", "xtfn".

**Example 2:**

    Input: ["a","ab","abc","d","cd","bcd","abcd"]
    Output: 4 
    Explanation: The two words can be "ab", "cd".

**Example 3:**

    Input: ["a","aa","aaa","aaaa"]
    Output: 0 
    Explanation: No such pair of words.



## 题目大意

给定一个字符串数组 words，找到 length(word[i]) * length(word[j]) 的最大值，并且这两个单词不含有公共字母。你可以认为每个单词只包含小写字母。如果不存在这样的两个单词，返回 0。


## 解题思路

- 在字符串数组中找到 2 个没有公共字符的字符串，并且这两个字符串的长度乘积要是最大的，求这个最大的乘积。
- 这里需要利用位运算 `&` 运算的性质，如果 `X & Y = 0`，说明 X 和 Y 完全不相同。那么我们将字符串都编码成二进制数，进行 `&` 运算即可分出没有公共字符的字符串，最后动态维护长度乘积最大值即可。将字符串编码成二进制数的规则比较简单，每个字符相对于 'a' 的距离，根据这个距离将 1 左移多少位。

```c
    a 1->1  
    b 2->10  
    c 4->100  
    ab 3->11  
    ac 5->101  
    abc 7->111  
    az 33554433->10000000000000000000000001  
```
