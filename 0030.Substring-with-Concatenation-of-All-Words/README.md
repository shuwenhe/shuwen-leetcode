# [30. Substring with Concatenation of All Words](https://leetcode.com/problems/substring-with-concatenation-of-all-words/)

## 题目

You are given a string, s, and a list of words, words, that are all of the same length. Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.

Example 1:

```c
Input:
  s = "barfoothefoobarman",
  words = ["foo","bar"]
Output: [0,9]
Explanation: Substrings starting at index 0 and 9 are "barfoor" and "foobar" respectively.
The output order does not matter, returning [9,0] is fine too.
```

Example 2:

```c
Input:
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
Output: []
```

## 题目大意

给定一个源字符串 s，再给一个字符串数组，要求在源字符串中找到由字符串数组各种组合组成的连续串的起始下标，如果存在多个，在结果中都需要输出。

## 解题思路

这一题看似很难，但是有 2 个限定条件也导致这题不是特别难。1. 字符串数组里面的字符串长度都是一样的。2. 要求字符串数组中的字符串都要连续连在一起的，前后顺序可以是任意排列组合。

解题思路，先将字符串数组里面的所有字符串都存到 map 中，并累计出现的次数。然后从源字符串从头开始扫，每次判断字符串数组里面的字符串时候全部都用完了(计数是否为 0)，如果全部都用完了，并且长度正好是字符串数组任意排列组合的总长度，就记录下这个组合的起始下标。如果不符合，就继续考察源字符串的下一个字符，直到扫完整个源字符串。



