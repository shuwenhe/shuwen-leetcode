# [75. Sort Colors](https://leetcode.com/problems/sort-colors/)

## 题目

Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.

Example 1:  

```
Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```

Follow up:

- A rather straight forward solution is a two-pass algorithm using counting sort.  
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
- Could you come up with a one-pass algorithm using only constant space?


## 题目大意

抽象题意其实就是排序。这题可以用快排一次通过。

## 解题思路

题目末尾的 Follow up 提出了一个更高的要求，能否用一次循环解决问题？这题由于数字只会出现 0，1，2 这三个数字，所以用游标移动来控制顺序也是可以的。具体做法：0 是排在最前面的，所以只要添加一个 0，就需要放置 1 和 2。1 排在 2 前面，所以添加 1 的时候也需要放置 2 。至于最后的 2，只用移动游标即可。

这道题可以用计数排序，适合待排序数字很少的题目。用一个 3 个容量的数组分别计数，记录 0，1，2 出现的个数。然后再根据个数排列 0，1，2 即可。时间复杂度 O(n)，空间复杂度 O(K)。这一题 K = 3。

这道题也可以用一次三路快排。数组分为 3 部分，第一个部分都是 0，中间部分都是 1，最后部分都是 2 。

