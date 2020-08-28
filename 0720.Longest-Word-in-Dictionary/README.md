# [720. Longest Word in Dictionary](https://leetcode.com/problems/longest-word-in-dictionary/)


## 题目

Given a list of strings `words` representing an English Dictionary, find the longest word in `words` that can be built one character at a time by other words in `words`. If there is more than one possible answer, return the longest word with the smallest lexicographical order.

If there is no answer, return the empty string.

**Example 1:**

    Input: 
    words = ["w","wo","wor","worl", "world"]
    Output: "world"
    Explanation: 
    The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".

**Example 2:**

    Input: 
    words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
    Output: "apple"
    Explanation: 
    Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is lexicographically smaller than "apply".

**Note:**

- All the strings in the input will only contain lowercase letters.
- The length of `words` will be in the range `[1, 1000]`.
- The length of `words[i]` will be in the range `[1, 30]`.


## 题目大意

给出一个字符串数组 words 组成的一本英语词典。从中找出最长的一个单词，该单词是由 words 词典中其他单词逐步添加一个字母组成。若其中有多个可行的答案，则返回答案中字典序最小的单词。若无答案，则返回空字符串。



## 解题思路


- 给出一个字符串数组，要求找到长度最长的，并且可以由字符串数组里面其他字符串拼接一个字符组成的字符串。如果存在多个这样的最长的字符串，则输出字典序较小的那个字符串，如果找不到这样的字符串，输出空字符串。
- 这道题解题思路是先排序，排序完成以后就是字典序从小到大了。之后再用 map 辅助记录即可。
