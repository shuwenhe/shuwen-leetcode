# [645. Set Mismatch](https://leetcode.com/problems/set-mismatch/)


## 题目

The set `S` originally contains numbers from 1 to `n`. But unfortunately, due to the data error, one of the numbers in the set got duplicated to **another** number in the set, which results in repetition of one number and loss of another number.

Given an array `nums` representing the data status of this set after the error. Your task is to firstly find the number occurs twice and then find the number that is missing. Return them in the form of an array.

**Example 1:**

    Input: nums = [1,2,2,4]
    Output: [2,3]

**Note:**

1. The given array size will in the range [2, 10000].
2. The given array's numbers won't have any order.


## 题目大意


集合 S 包含从1到 n 的整数。不幸的是，因为数据错误，导致集合里面某一个元素复制了成了集合里面的另外一个元素的值，导致集合丢失了一个整数并且有一个元素重复。给定一个数组 nums 代表了集合 S 发生错误后的结果。你的任务是首先寻找到重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。

注意:  

- 给定数组的长度范围是 [2, 10000]。
- 给定的数组是无序的。


## 解题思路


- 给出一个数组，数组里面装的是 1-n 的数字，由于错误导致有一个数字变成了另外一个数字，要求找出重复的一个数字和正确的数字。这一题是简单题，根据下标比对就可以找到哪个数字重复了，哪个数字缺少了。
