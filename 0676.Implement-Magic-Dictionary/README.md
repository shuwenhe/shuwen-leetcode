# [676. Implement Magic Dictionary](https://leetcode.com/problems/implement-magic-dictionary/)


## 题目

Implement a magic directory with `buildDict`, and `search` methods.

For the method `buildDict`, you'll be given a list of non-repetitive words to build a dictionary.

For the method `search`, you'll be given a word, and judge whether if you modify **exactly** one character into **another**character in this word, the modified word is in the dictionary you just built.

**Example 1:**

    Input: buildDict(["hello", "leetcode"]), Output: Null
    Input: search("hello"), Output: False
    Input: search("hhllo"), Output: True
    Input: search("hell"), Output: False
    Input: search("leetcoded"), Output: False

**Note:**

1. You may assume that all the inputs are consist of lowercase letters `a-z`.
2. For contest purpose, the test data is rather small by now. You could think about highly efficient algorithm after the contest.
3. Please remember to **RESET** your class variables declared in class MagicDictionary, as static/class variables are **persisted across multiple test cases**. Please see [here](https://leetcode.com/faq/#different-output) for more details.


## 题目大意

实现一个带有 buildDict, 以及 search 方法的魔法字典。对于 buildDict 方法，你将被给定一串不重复的单词来构建一个字典。对于 search 方法，你将被给定一个单词，并且判定能否只将这个单词中一个字母换成另一个字母，使得所形成的新单词存在于你构建的字典中。



## 解题思路


- 实现 `MagicDictionary` 的数据结构，这个数据结构内会存储一个字符串数组，当执行 `Search` 操作的时候要求判断传进来的字符串能否只改变一个字符(不能增加字符也不能删除字符)就能变成 `MagicDictionary` 中存储的字符串，如果可以，就输出 `true`，如果不能，就输出 `false`。
- 这题的解题思路比较简单，用 Map 判断即可。
