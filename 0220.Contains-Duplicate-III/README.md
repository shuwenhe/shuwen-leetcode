# [220. Contains Duplicate III](https://leetcode.com/problems/contains-duplicate-iii/)

## 题目

Given an array of integers, find out whether there are two distinct indices i and j in the array such that the **absolute** difference between **nums[i]** and **nums[j]** is at most t and the **absolute** difference between i and j is at most k.

**Example 1:**

    Input: nums = [1,2,3,1], k = 3, t = 0
    Output: true

**Example 2:**

    Input: nums = [1,0,1,1], k = 1, t = 2
    Output: true

**Example 3:**

    Input: nums = [1,5,9,1,5,9], k = 2, t = 3
    Output: false




## 题目大意

给出一个数组 num，再给 K 和 t。问在 num 中能否找到一组 i 和 j，使得 num[i] 和 num[j] 的绝对差值最大为 t，并且 i 和 j 之前的绝对差值最大为 k。

## 解题思路


- 给出一个数组，要求在数组里面找到 2 个索引，`i` 和 `j`，使得 `| nums[i] - nums[j] | ≤ t` ，并且 `| i - j | ≤  k` 。
- 这是一道滑动窗口的题目。第一想法就是用 `i` 和 `j` 两个指针，针对每个 `i` ，都从 `i + 1` 往后扫完整个数组，判断每个 `i` 和 `j` ，判断是否满足题意。`j` 在循环的过程中注意判断剪枝条件 `| i - j | ≤  k`。这个做法的时间复杂度是 O(n^2)。这个做法慢的原因在于滑动窗口的左边界和右边界在滑动过程中不是联动滑动的。
- 于是考虑，如果数组是有序的呢？把数组按照元素值从小到大进行排序，如果元素值相等，就按照 index 从小到大进行排序。在这样有序的数组中找满足题意的 `i` 和 `j`，滑动窗口左边界和右边界就是联动的了。窗口的右边界滑到与左边界元素值的差值 ≤ t 的地方，满足了这个条件再判断 `| i - j | ≤  k`，如果右边界与左边界元素值的差值 > t 了，说明该把左边界往右移动了(能这样移动的原因就是因为我们将数组元素大小排序了，右移是增大元素的方向)。移动左边界的时候需要注意左边界不能超过右边界。这样滑动窗口一次滑过整个排序后的数组，就可以判断是否存在满足题意的 `i` 和 `j` 。这个做法的时间主要花在排序上了，时间复杂度是 O(n log n)。
