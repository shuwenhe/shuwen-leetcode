# [212. Word Search II](https://leetcode.com/problems/word-search-ii/)


## 题目

Given a 2D board and a list of words from the dictionary, find all words in the board.

Each word must be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

**Example:**

    Input: 
    board = [
      ['o','a','a','n'],
      ['e','t','a','e'],
      ['i','h','k','r'],
      ['i','f','l','v']
    ]
    words = ["oath","pea","eat","rain"]
    
    Output: ["eat","oath"]

**Note:**

1. All inputs are consist of lowercase letters `a-z`.
2. The values of `words` are distinct.

## 题目大意

给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。


## 解题思路

- 这一题是第 79 题的加强版，在第 79 题的基础上增加了一个 word 数组，要求找出所有出现在地图中的单词。思路还是可以按照第 79 题 DFS 搜索，不过时间复杂度特别高！
- 想想更优的解法。
