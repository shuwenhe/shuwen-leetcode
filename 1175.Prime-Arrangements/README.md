# [1175. Prime Arrangements](https://leetcode.com/problems/prime-arrangements/)


## 题目

Return the number of permutations of 1 to `n` so that prime numbers are at prime indices (1-indexed.)

*(Recall that an integer is prime if and only if it is greater than 1, and cannot be written as a product of two positive integers both smaller than it.)*

Since the answer may be large, return the answer **modulo `10^9 + 7`**.

**Example 1:**

    Input: n = 5
    Output: 12
    Explanation: For example [1,2,5,4,3] is a valid permutation, but [5,2,3,4,1] is not because the prime number 5 is at index 1.

**Example 2:**

    Input: n = 100
    Output: 682289015

**Constraints:**

- `1 <= n <= 100`


## 题目大意


请你帮忙给从 1 到 n 的数设计排列方案，使得所有的「质数」都应该被放在「质数索引」（索引从 1 开始）上；你需要返回可能的方案总数。让我们一起来回顾一下「质数」：质数一定是大于 1 的，并且不能用两个小于它的正整数的乘积来表示。由于答案可能会很大，所以请你返回答案 模 mod 10^9 + 7 之后的结果即可。

提示：

- 1 <= n <= 100

## 解题思路

- 给出一个数 n，要求在 1-n 这 n 个数中，素数在素数索引下标位置上的全排列个数。
- 由于这一题的 `n` 小于 100，所以可以用打表法。先把小于 100 个素数都打表打出来。然后对小于 n 的素数进行全排列，即 n！，然后再对剩下来的非素数进行全排列，即 (n-c)！。两个的乘积即为最终答案。
