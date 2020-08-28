# [187. Repeated DNA Sequences](https://leetcode.com/problems/repeated-dna-sequences/)


## 题目

All DNA is composed of a series of nucleotides abbreviated as A, C, G, and T, for example: "ACGAATTCCG". When studying DNA, it is sometimes useful to identify repeated sequences within the DNA.

Write a function to find all the 10-letter-long sequences (substrings) that occur more than once in a DNA molecule.

**Example:**

    Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
    
    Output: ["AAAAACCCCC", "CCCCCAAAAA"]


## 题目大意

所有 DNA 由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：“ACGAATTCCG”。在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。编写一个函数来查找 DNA 分子中所有出现超多一次的10个字母长的序列（子串）。

## 解题思路

- 这一题不用位运算比较好做，维护一个长度为 10 的字符串，在 map 中出现次数 > 1 就输出。
- 用位运算想做这一题，需要动态的维护长度为 10 的 hashkey，先计算开头长度为 9 的 hash，在往后面扫描的过程中，如果长度超过了 10 ，就移除 hash 开头的一个字符，加入后面一个字符。具体做法是先将 ATCG 变成 00，01，10，11 的编码，那么长度为 10 ，hashkey 就需要维护在 20 位。mask = 0xFFFFF 就是 20 位的。维护了 hashkey 以后，根据这个 hashkey 进行去重和统计频次。
