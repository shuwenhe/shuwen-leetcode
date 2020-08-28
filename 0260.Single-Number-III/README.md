# [260. Single Number III](https://leetcode.com/problems/single-number-iii/)


## 题目

Given an array of numbers `nums`, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once.

**Example:**

    Input:  [1,2,1,3,2,5]
    Output: [3,5]

**Note**:

1. The order of the result is not important. So in the above example, `[5, 3]` is also correct.
2. Your algorithm should run in linear runtime complexity. Could you implement it using only constant space complexity?


## 题目大意

给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。

注意：  

- 结果输出的顺序并不重要，对于上面的例子，[5, 3] 也是正确答案。
- 要求算法时间复杂度是线性的，并且不使用额外的辅助空间。



## 解题思路

- 这一题是第 136 题的加强版。第 136 题里面只有一个数出现一次，其他数都出现 2 次。而这一次有 2 个数字出现一次，其他数出现 2 次。
- 解题思路还是利用异或，把出现 2 次的数先消除。最后我们要找的 2 个数肯定也是不同的，所以最后 2 个数对一个数进行异或，答案肯定是不同的。那么我们找哪个数为参照物呢？可以随便取，不如就取 lsb 最低位为 1 的数吧
- 于是整个数组会被分为 2 部分，异或 lsb 为 0 的和异或 lsb 为 1 的，在这 2 部分中，用异或操作把出现 2 次的数都消除，那么剩下的 2 个数分别就在这 2 部分中。
