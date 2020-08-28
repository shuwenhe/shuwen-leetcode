# [350. Intersection of Two Arrays II](https://leetcode.com/problems/intersection-of-two-arrays-ii/)

## 题目

Given two arrays, write a function to compute their intersection.



Example 1:

```c
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
```

Example 2:

```c
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
```

Note:

- Each element in the result should appear as many times as it shows in both arrays.
- The result can be in any order.


Follow up:

- What if the given array is already sorted? How would you optimize your algorithm?
- What if nums1's size is small compared to nums2's size? Which algorithm is better?
- What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?

## 题目大意

这题是第 349 题的加强版。要求输出 2 个数组的交集元素，如果元素出现多次，要输出多次。

## 解题思路

这一题还是延续第 349 题的思路。把数组一中的数字都放进字典中，另外字典的 key 是数组中的数字，value 是这个数字出现的次数。在扫描数组二的时候，每取出一个存在的数组，把字典中的 value 减一。如果 value 是 0 代表不存在这个数字。



