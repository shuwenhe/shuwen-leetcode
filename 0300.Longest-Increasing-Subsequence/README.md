# [300. Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/)


## 题目

Given an unsorted array of integers, find the length of longest increasing subsequence.

**Example:**

    Input: [10,9,2,5,3,7,101,18]
    Output: 4 
    Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

**Note:**

- There may be more than one LIS combination, it is only necessary for you to return the length.
- Your algorithm should run in O(n^2) complexity.

**Follow up:** Could you improve it to O(n log n) time complexity?

## 题目大意

给定一个无序的整数数组，找到其中最长上升子序列的长度。


## 解题思路

- 给定一个整数序列，求其中的最长上升子序列的长度。这一题就是经典的最长上升子序列的问题。
- `dp[i]` 代表为第 i 个数字为结尾的最长上升子序列的长度。换种表述，dp[i] 代表 [0,i] 范围内，选择数字 nums[i] 可以获得的最长上升子序列的长度。状态转移方程为 `dp[i] = max( 1 + dp[j]) ，其中 j < i && nums[j] > nums[i]`，取所有满足条件的最大值。时间复杂度 O(n^2)
- 这道题还有一种更快的解法。考虑这样一个问题，我们是否能用一个数组，记录上升子序列的最末尾的一个数字呢？如果这个数字越小，那么这个子序列往后面添加数字的几率就越大，那么就越可能成为最长的上升子序列。举个例子：nums = [4,5,6,3]，它的所有的上升子序列为

```
    len = 1   :      [4], [5], [6], [3]   => tails[0] = 3
    len = 2   :      [4, 5], [5, 6]       => tails[1] = 5
    len = 3   :      [4, 5, 6]            => tails[2] = 6
```
- 其中 `tails[i]` 中存储的是所有长度为 i + 1 的上升子序列中末尾最小的值。也很容易证明 `tails` 数组里面的值一定是递增的(因为我们用末尾的数字描述最长递增子序列)。既然 tails 是有序的，我们就可以用二分查找的方法去更新这个 tail 数组里面的值。更新策略如下：(1). 如果 x 比所有的 tails 元素都要大，那么就直接放在末尾，并且 tails 数组长度加一；(2). 如果 `tails[i-1] < x <= tails[i]`，则更新 tails[i]，因为 x 更小，更能获得最长上升子序列。最终 tails 数组的长度即为最长的上升子序列。这种做法的时间复杂度 O(n log n)。

