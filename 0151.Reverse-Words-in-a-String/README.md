# [151. Reverse Words in a String](https://leetcode.com/problems/reverse-words-in-a-string/)



## 题目

Given an input string, reverse the string word by word.

**Example 1:**

    Input: "the sky is blue"
    Output: "blue is sky the"

**Example 2:**

    Input: "  hello world!  "
    Output: "world! hello"
    Explanation: Your reversed string should not contain leading or trailing spaces.

**Example 3:**

    Input: "a good   example"
    Output: "example good a"
    Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

**Note:**

- A word is defined as a sequence of non-space characters.
- Input string may contain leading or trailing spaces. However, your reversed string should not contain leading or trailing spaces.
- You need to reduce multiple spaces between two words to a single space in the reversed string.

**Follow up:**

For C programmers, try to solve it *in-place* in *O*(1) extra space.


## 题目大意

给定一个字符串，逐个翻转字符串中的每个单词。

说明：

- 无空格字符构成一个单词。
- 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
- 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。
 

进阶：

- 请选用 C 语言的用户尝试使用 O(1) 额外空间复杂度的原地解法。


## 解题思路


- 给出一个中间有空格分隔的字符串，要求把这个字符串按照单词的维度前后翻转。
- 依照题意，先把字符串按照空格分隔成每个小单词，然后把单词前后翻转，最后再把每个单词中间添加空格。
