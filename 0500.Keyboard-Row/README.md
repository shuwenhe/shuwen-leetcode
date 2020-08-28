# [500. Keyboard Row](https://leetcode.com/problems/keyboard-row/)


## 题目

Given a List of words, return the words that can be typed using letters of **alphabet** on only one row's of American keyboard like the image below.

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/keyboard.png)

**Example:**

    Input: ["Hello", "Alaska", "Dad", "Peace"]
    Output: ["Alaska", "Dad"]

**Note:**

1. You may use one character in the keyboard more than once.
2. You may assume the input string will only contain letters of alphabet.


## 题目大意

给定一个单词列表，只返回可以使用在键盘同一行的字母打印出来的单词。键盘如上图所示。

## 解题思路

- 给出一个字符串数组，要求依次判断数组中的每个字符串是否都位于键盘上的同一个行，如果是就输出。这也是一道水题。
