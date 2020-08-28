# [219. Contains Duplicate II](https://leetcode.com/problems/contains-duplicate-ii/)

## 题目

Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.


Example 1:

```c
Input: nums = [1,2,3,1], k = 3
Output: true
```
Example 2:

```c
Input: nums = [1,0,1,1], k = 1
Output: true
```

Example 3:

```c
Input: nums = [1,2,3,1,2,3], k = 2
Output: false
```

## 题目大意

这是一道简单题，如果数组里面有重复数字，并且重复数字的下标差值小于等于 K 就输出 true，如果没有重复数字或者下标差值超过了 K ，则输出 flase。

## 解题思路

这道题可以维护一个只有 K 个元素的 map，每次只需要判断这个 map 里面是否存在这个元素即可。如果存在就代表重复数字的下标差值在 K 以内。map 的长度如果超过了 K 以后就删除掉 i-k 的那个元素，这样一直维护 map 里面只有 K 个元素。
