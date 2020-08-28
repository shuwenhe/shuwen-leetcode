# [1160. Find Words That Can Be Formed by Characters](https://leetcode.com/problems/find-words-that-can-be-formed-by-characters/)


## 题目

You are given an array of strings `words` and a string `chars`.

A string is *good* if it can be formed by characters from `chars` (each character can only be used once).

Return the sum of lengths of all good strings in `words`.

**Example 1:**

    Input: words = ["cat","bt","hat","tree"], chars = "atach"
    Output: 6
    Explanation: 
    The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 = 6.

**Example 2:**

    Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"
    Output: 10
    Explanation: 
    The strings that can be formed are "hello" and "world" so the answer is 5 + 5 = 10.

**Note:**

1. `1 <= words.length <= 1000`
2. `1 <= words[i].length, chars.length <= 100`
3. All strings contain lowercase English letters only.


## 题目大意


给你一份『词汇表』（字符串数组） words 和一张『字母表』（字符串） chars。假如你可以用 chars 中的『字母』（字符）拼写出 words 中的某个『单词』（字符串），那么我们就认为你掌握了这个单词。注意：每次拼写时，chars 中的每个字母都只能用一次。返回词汇表 words 中你掌握的所有单词的 长度之和。

提示：

1. 1 <= words.length <= 1000
2. 1 <= words[i].length, chars.length <= 100
3. 所有字符串中都仅包含小写英文字母



## 解题思路

- 给出一个字符串数组 `words` 和一个字符串 `chars`，要求输出 `chars` 中能构成 `words` 字符串的字符数总数。
- 简单题。先分别统计 `words` 和 `chars` 里面字符的频次。然后针对 `words` 中每个 `word` 判断能够能由 `chars` 构成，如果能构成，最终结果加上这个 `word` 的长度。
