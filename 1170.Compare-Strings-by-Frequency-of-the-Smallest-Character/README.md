# [1170. Compare Strings by Frequency of the Smallest Character](https://leetcode.com/problems/compare-strings-by-frequency-of-the-smallest-character/)

## 题目

Let's define a function `f(s)` over a non-empty string `s`, which calculates the frequency of the smallest character in `s`. For example, if `s = "dcce"` then `f(s) = 2` because the smallest character is `"c"` and its frequency is 2.

Now, given string arrays `queries` and `words`, return an integer array `answer`, where each `answer[i]` is the number of words such that `f(queries[i])` < `f(W)`, where `W` is a word in `words`.

**Example 1:**

    Input: queries = ["cbd"], words = ["zaaaz"]
    Output: [1]
    Explanation: On the first query we have f("cbd") = 1, f("zaaaz") = 3 so f("cbd") < f("zaaaz").

**Example 2:**

    Input: queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
    Output: [1,2]
    Explanation: On the first query only f("bbb") < f("aaaa"). On the second query both f("aaa") and f("aaaa") are both > f("cc").

**Constraints:**

- `1 <= queries.length <= 2000`
- `1 <= words.length <= 2000`
- `1 <= queries[i].length, words[i].length <= 10`
- `queries[i][j]`, `words[i][j]` are English lowercase letters.


## 题目大意


我们来定义一个函数 f(s)，其中传入参数 s 是一个非空字符串；该函数的功能是统计 s  中（按字典序比较）最小字母的出现频次。

例如，若 s = "dcce"，那么 f(s) = 2，因为最小的字母是 "c"，它出现了 2 次。

现在，给你两个字符串数组待查表 queries 和词汇表 words，请你返回一个整数数组 answer 作为答案，其中每个 answer[i] 是满足 f(queries[i]) < f(W) 的词的数目，W 是词汇表 words 中的词。

提示：

- 1 <= queries.length <= 2000
- 1 <= words.length <= 2000
- 1 <= queries[i].length, words[i].length <= 10
- queries[i][j], words[i][j] 都是小写英文字母




## 解题思路

- 给出 2 个数组，`queries` 和 `words`，针对每一个 `queries[i]` 统计在 `words[j]` 中满足 `f(queries[i]) < f(words[j])` 条件的 `words[j]` 的个数。`f(string)` 的定义是 `string` 中字典序最小的字母的频次。
- 先依照题意，构造出 `f()` 函数，算出每个 `words[j]` 的 `f()` 值，然后排序。再依次计算 `queries[i]` 的 `f()` 值。针对每个 `f()` 值，在 `words[j]` 的 `f()` 值中二分搜索，查找比它大的值的下标 `k`，`n-k` 即是比 `queries[i]` 的 `f()` 值大的元素个数。依次输出到结果数组中即可。
