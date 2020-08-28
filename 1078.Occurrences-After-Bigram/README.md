# [1078. Occurrences After Bigram](https://leetcode.com/problems/occurrences-after-bigram/)


## 题目

Given words `first` and `second`, consider occurrences in some `text` of the form "`first second third`", where `second` comes immediately after `first`, and `third`comes immediately after `second`.

For each such occurrence, add "`third`" to the answer, and return the answer.

**Example 1:**

    Input: text = "alice is a good girl she is a good student", first = "a", second = "good"
    Output: ["girl","student"]

**Example 2:**

    Input: text = "we will we will rock you", first = "we", second = "will"
    Output: ["we","rock"]

**Note:**

1. `1 <= text.length <= 1000`
2. `text` consists of space separated words, where each word consists of lowercase English letters.
3. `1 <= first.length, second.length <= 10`
4. `first` and `second` consist of lowercase English letters.


## 题目大意


给出第一个词 first 和第二个词 second，考虑在某些文本 text 中可能以 "first second third" 形式出现的情况，其中 second 紧随 first 出现，third 紧随 second 出现。对于每种这样的情况，将第三个词 "third" 添加到答案中，并返回答案。




## 解题思路


- 简单题。给出一个 text，要求找出紧接在 first 和 second 后面的那个字符串，有多个就输出多个。解法很简单，先分解出 words 每个字符串，然后依次遍历进行字符串匹配。匹配到 first 和 second 以后，输出之后的那个字符串。
