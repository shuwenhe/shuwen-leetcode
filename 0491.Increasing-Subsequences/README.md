# [491. Increasing Subsequences](https://leetcode.com/problems/increasing-subsequences/)


## 题目

Given an integer array, your task is to find all the different possible increasing subsequences of the given array, and the length of an increasing subsequence should be at least 2.

**Example:**

    Input: [4, 6, 7, 7]
    Output: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]

**Note:**

1. The length of the given array will not exceed 15.
2. The range of integer in the given array is [-100,100].
3. The given array may contain duplicates, and two equal integers should also be considered as a special case of increasing sequence.



## 题目大意


给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是 2。

说明:

1. 给定数组的长度不会超过15。
2. 数组中的整数范围是 [-100,100]。
3. 给定数组中可能包含重复数字，相等的数字应该被视为递增的一种情况。




## 解题思路


- 给出一个数组，要求找出这个数组中所有长度大于 2 的非递减子序列。子序列顺序和原数组元素下标必须是顺序的，不能是逆序的。
- 这一题和第 78 题和第 90 题是类似的题目。第 78 题和第 90 题是求所有子序列，这一题在这两题的基础上增加了非递减和长度大于 2 的条件。需要注意的两点是，原数组中元素可能会重复，最终结果输出的时候需要去重。最终结果输出的去重用 map 处理，数组中重复元素用 DFS 遍历搜索。在每次 DFS 中，用 map 记录遍历过的元素，保证本轮 DFS 中不出现重复的元素，递归到下一层还可以选择值相同，但是下标不同的另外一个元素。外层循环也要加一个  map，这个 map 是过滤每组解因为重复元素导致的重复解，经过过滤以后，起点不同了，最终的解也会不同。
- 这一题和第 78 题，第 90 题类似，可以一起解答和复习。
