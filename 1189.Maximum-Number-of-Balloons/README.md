# [1189. Maximum Number of Balloons](https://leetcode.com/problems/maximum-number-of-balloons/)


## 题目

Given a string `text`, you want to use the characters of `text` to form as many instances of the word **"balloon"** as possible.

You can use each character in `text` **at most once**. Return the maximum number of instances that can be formed.

**Example 1:**

![](https://assets.leetcode.com/uploads/2019/09/05/1536_ex1_upd.JPG)

    Input: text = "nlaebolko"
    Output: 1

**Example 2:**

![](https://assets.leetcode.com/uploads/2019/09/05/1536_ex2_upd.JPG)

    Input: text = "loonbalxballpoon"
    Output: 2

**Example 3:**

    Input: text = "leetcode"
    Output: 0

**Constraints:**

- `1 <= text.length <= 10^4`
- `text` consists of lower case English letters only.


## 题目大意

给你一个字符串 text，你需要使用 text 中的字母来拼凑尽可能多的单词 "balloon"（气球）。字符串 text 中的每个字母最多只能被使用一次。请你返回最多可以拼凑出多少个单词 "balloon"。

提示：

- 1 <= text.length <= 10^4
- text 全部由小写英文字母组成

## 解题思路


- 给出一个字符串，问这个字符串里面的数组能组成多少个 **balloon** 这个单词。
- 简单题，先统计 26 个字母每个字母的频次，然后取出 balloon 这 5 个字母出现频次最小的值就是结果。
