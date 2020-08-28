# [41. First Missing Positive](https://leetcode.com/problems/first-missing-positive/description/)

## 题目

Given an unsorted integer array, find the smallest missing positive integer.

Example 1:  

```
Input: [1,2,0]  
Output: 3  
```

Example 2:  

```
Input: [3,4,-1,1]  
Output: 2  
```

Example 3:  

```
Input: [7,8,9,11,12]  
Output: 1  
```

Note:  

Your algorithm should run in O(n) time and uses constant extra space.

## 题目大意

找到缺失的第一个正整数。

## 解题思路


为了减少时间复杂度，可以把 input 数组都装到 map 中，然后 i 循环从 1 开始，依次比对 map 中是否存在 i，只要不存在 i 就立即返回结果，即所求。