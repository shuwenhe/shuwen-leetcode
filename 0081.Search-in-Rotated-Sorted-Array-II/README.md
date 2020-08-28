# [81. Search in Rotated Sorted Array II](https://leetcode.com/problems/search-in-rotated-sorted-array-ii/)


## 题目

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., `[0,0,1,2,2,5,6]` might become `[2,5,6,0,0,1,2]`).

You are given a target value to search. If found in the array return `true`, otherwise return `false`.

**Example 1:**

    Input: nums = [2,5,6,0,0,1,2], target = 0
    Output: true

**Example 2:**

    Input: nums = [2,5,6,0,0,1,2], target = 3
    Output: false

**Follow up:**

- This is a follow up problem to [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/description/), where `nums` may contain duplicates.
- Would this affect the run-time complexity? How and why?


## 题目大意

假设按照升序排序的数组在预先未知的某个点上进行了旋转。( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。

编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。

进阶:

- 这是搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
- 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？


## 解题思路


- 给出一个数组，数组中本来是从小到大排列的，并且数组中有重复数字。但是现在把后面随机一段有序的放到数组前面，这样形成了前后两端有序的子序列。在这样的一个数组里面查找一个数，设计一个 O(log n) 的算法。如果找到就输出 `true`，如果没有找到，就输出 `false` 。
- 这一题是第 33 题的加强版，实现代码完全一样，只不过输出变了。这一题输出 `true` 和 `false` 了。具体思路见第 33 题。
