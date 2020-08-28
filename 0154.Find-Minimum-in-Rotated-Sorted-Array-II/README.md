# [154. Find Minimum in Rotated Sorted Array II](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/)


## 题目

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., `[0,1,2,4,5,6,7]` might become `[4,5,6,7,0,1,2]`).

Find the minimum element.

The array may contain duplicates.

**Example 1:**

    Input: [1,3,5]
    Output: 1

**Example 2:**

    Input: [2,2,2,0,1]
    Output: 0

**Note:**

- This is a follow up problem to [Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/).
- Would allow duplicates affect the run-time complexity? How and why?


## 题目大意

假设按照升序排序的数组在预先未知的某个点上进行了旋转。( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。请找出其中最小的元素。

注意数组中可能存在重复的元素。

## 解题思路


- 给出一个原本从小到大排序过的数组，注意数组中有重复的元素。但是在某一个分割点上，把数组切分后的两部分对调位置，数值偏大的放到了数组的前部。求这个数组中最小的元素。
- 这一题是第 153 题的加强版，增加了重复元素的条件。但是实际做法还是没有变，还是用二分搜索，只不过在相等元素上多增加一个判断即可。时间复杂度 O(log n)。
