# [268. Missing Number](https://leetcode.com/problems/missing-number/)


## 题目

Given an array containing n distinct numbers taken from `0, 1, 2, ..., n`, find the one that is missing from the array.

**Example 1:**

    Input: [3,0,1]
    Output: 2

**Example 2:**

    Input: [9,6,4,2,3,5,7,0,1]
    Output: 8

**Note**:Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?


## 题目大意

给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。算法应该具有线性时间复杂度。你能否仅使用额外常数空间来实现?



## 解题思路


- 要求找出 `0, 1, 2, ..., n` 中缺失的那个数。还是利用异或的性质，`X^X = 0`。这里我们需要构造一个 X，用数组下标就可以了。数字下标是从 `[0，n-1]`，数字是 `[0，n]`，依次把数组里面的数组进行异或，把结果和 n 再异或一次，中和掉出现的数字，剩下的那个数字就是之前没有出现过的，缺失的数字。
