# [217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/)

## 题目

Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.


Example 1:

```c
Input: [1,2,3,1]
Output: true
```
Example 2:

```c
Input: [1,2,3,4]
Output: false
```

Example 3:

```c
Input: [1,1,1,3,3,4,3,2,4,2]
Output: true
```

## 题目大意

这是一道简单题，如果数组里面有重复数字就输出 true，否则输出 flase。

## 解题思路

用 map 判断即可。
