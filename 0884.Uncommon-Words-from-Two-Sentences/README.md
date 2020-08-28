# [884. Uncommon Words from Two Sentences](https://leetcode.com/problems/uncommon-words-from-two-sentences/)


## 题目

We are given two sentences `A` and `B`. (A *sentence* is a string of space separated words. Each *word* consists only of lowercase letters.)

A word is *uncommon* if it appears exactly once in one of the sentences, and does not appear in the other sentence.

Return a list of all uncommon words.

You may return the list in any order.

**Example 1:**

    Input: A = "this apple is sweet", B = "this apple is sour"
    Output: ["sweet","sour"]

**Example 2:**

    Input: A = "apple apple", B = "banana"
    Output: ["banana"]

**Note:**

1. `0 <= A.length <= 200`
2. `0 <= B.length <= 200`
3. `A` and `B` both contain only spaces and lowercase letters.


## 题目大意

给定两个句子 A 和 B 。（句子是一串由空格分隔的单词。每个单词仅由小写字母组成。）

如果一个单词在其中一个句子中只出现一次，在另一个句子中却没有出现，那么这个单词就是不常见的。返回所有不常用单词的列表。您可以按任何顺序返回列表。


## 解题思路

- 找出 2 个句子中不同的单词，将它们俩都打印出来。简单题，先将 2 个句子的单词都拆开放入 map 中进行词频统计，不同的两个单词的词频肯定都为 1，输出它们即可。
