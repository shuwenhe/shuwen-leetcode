# [496. Next Greater Element I](https://leetcode.com/problems/next-greater-element-i/)

## 题目

You are given two arrays (without duplicates) nums1 and nums2 where nums1’s elements are subset of nums2. Find all the next greater numbers for nums1's elements in the corresponding places of nums2.

The Next Greater Number of a number x in nums1 is the first greater number to its right in nums2. If it does not exist, output -1 for this number.


Example 1:

```c
Input: nums1 = [4,1,2], nums2 = [1,3,4,2].
Output: [-1,3,-1]
Explanation:
    For number 4 in the first array, you cannot find the next greater number for it in the second array, so output -1.
    For number 1 in the first array, the next greater number for it in the second array is 3.
    For number 2 in the first array, there is no next greater number for it in the second array, so output -1.
```

Example 2:

```c
Input: nums1 = [2,4], nums2 = [1,2,3,4].
Output: [3,-1]
Explanation:
    For number 2 in the first array, the next greater number for it in the second array is 3.
    For number 4 in the first array, there is no next greater number for it in the second array, so output -1.
```

Note:  

- All elements in nums1 and nums2 are unique.  
- The length of both nums1 and nums2 would not exceed 1000.


## 题目大意

这道题也是简单题。题目给出 2 个数组 A 和 B，针对 A 中的每个数组中的元素，要求在 B 数组中找出比 A 数组中元素大的数，B 中元素之间的顺序保持不变。如果找到了就输出这个值，如果找不到就输出 -1。


## 解题思路

简单题，依题意做即可。