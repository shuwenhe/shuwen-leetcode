# [66. Plus One](https://leetcode.com/problems/plus-one/)


## 题目

Given a **non-empty** array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

**Example 1:**

    Input: [1,2,3]
    Output: [1,2,4]
    Explanation: The array represents the integer 123.

**Example 2:**

    Input: [4,3,2,1]
    Output: [4,3,2,2]
    Explanation: The array represents the integer 4321.


## 题目大意


给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。你可以假设除了整数 0 之外，这个整数不会以零开头。



## 解题思路

- 给出一个数组，代表一个十进制数，数组的 0 下标是十进制数的高位。要求计算这个十进制数加一以后的结果。
- 简单的模拟题。从数组尾部开始往前扫，逐位进位即可。最高位如果还有进位需要在数组里面第 0 位再插入一个 1 。
