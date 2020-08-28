# [315. Count of Smaller Numbers After Self](https://leetcode.com/problems/count-of-smaller-numbers-after-self/)


## 题目

You are given an integer array nums and you have to return a new counts array. The counts array has the property where `counts[i]` is the number of smaller elements to the right of `nums[i]`.

**Example:**

    Input: [5,2,6,1]
    Output: [2,1,1,0] 
    Explanation:
    To the right of 5 there are 2 smaller elements (2 and 1).
    To the right of 2 there is only 1 smaller element (1).
    To the right of 6 there is 1 smaller element (1).
    To the right of 1 there is 0 smaller element.


## 题目大意


给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。

示例:

```
输入: [5,2,6,1]
输出: [2,1,1,0] 
解释:
5 的右侧有 2 个更小的元素 (2 和 1).
2 的右侧仅有 1 个更小的元素 (1).
6 的右侧有 1 个更小的元素 (1).
1 的右侧有 0 个更小的元素.
```


## 解题思路


- 给出一个数组，要求输出数组中每个元素相对于数组中的位置右边比它小的元素。
- 这一题是第 327 题的缩水版。由于需要找数组位置右边比当前位置元素小的元素，所以从数组右边开始往左边扫。构造一颗线段树，线段树里面父节点存的是子节点出现的次数和。有可能给的数据会很大，所以构造线段树的时候先离散化。还需要注意的是数组里面可能有重复元素，所以构造线段树要先去重并排序。从右往左扫的过程中，依次添加数组中的元素，添加了一次就立即 query 一次。query 的区间是 [minNum, nums[i]-1]。如果是 minNum 则输出 0，并且也要记得插入这个最小值。这一题的思路和第 327 题大体类似，详解可见第 327 题。
