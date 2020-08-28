# [771. Jewels and Stones](https://leetcode.com/problems/jewels-and-stones/)



## 题目

You're given strings `J` representing the types of stones that are jewels, and `S` representing the stones you have. Each character in `S` is a type of stone you have. You want to know how many of the stones you have are also jewels.

The letters in `J` are guaranteed distinct, and all characters in `J` and `S` are letters. Letters are case sensitive, so `"a"` is considered a different type of stone from `"A"`.

**Example 1:**

    Input: J = "aA", S = "aAAbbbb"
    Output: 3

**Example 2:**

    Input: J = "z", S = "ZZ"
    Output: 0

**Note:**

- `S` and `J` will consist of letters and have length at most 50.
- The characters in `J` are distinct.


## 题目大意

给定字符串 J 代表石头中宝石的类型，和字符串 S 代表你拥有的石头。S 中每个字符代表了一种你拥有的石头的类型，你想知道你拥有的石头中有多少是宝石。

J 中的字母不重复，J 和 S 中的所有字符都是字母。字母区分大小写，因此 "a" 和 "A" 是不同类型的石头。



## 解题思路


- 给出 2 个字符串，要求在 S 字符串中找出在 J 字符串里面出现的字符个数。这是一道简单题。
