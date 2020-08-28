# [978. Longest Turbulent Subarray](https://leetcode.com/problems/longest-turbulent-subarray/)

## 题目

A subarray `A[i], A[i+1], ..., A[j]` of `A` is said to be *turbulent* if and only if:

- For `i <= k < j`, `A[k] > A[k+1]` when `k` is odd, and `A[k] < A[k+1]` when `k` is even;
- **OR**, for `i <= k < j`, `A[k] > A[k+1]` when `k` is even, and `A[k] < A[k+1]` when `k` is odd.

That is, the subarray is turbulent if the comparison sign flips between each adjacent pair of elements in the subarray.

Return the **length** of a maximum size turbulent subarray of A.

**Example 1:**

    Input: [9,4,2,10,7,8,8,1,9]
    Output: 5
    Explanation: (A[1] > A[2] < A[3] > A[4] < A[5])

**Example 2:**

    Input: [4,8,12,16]
    Output: 2

**Example 3:**

    Input: [100]
    Output: 1

**Note:**

1. `1 <= A.length <= 40000`
2. `0 <= A[i] <= 10^9`


## 题目大意


当 A 的子数组 A[i], A[i+1], ..., A[j] 满足下列条件时，我们称其为湍流子数组：

若 i <= k < j，当 k 为奇数时， A[k] > A[k+1]，且当 k 为偶数时，A[k] < A[k+1]；
或 若 i <= k < j，当 k 为偶数时，A[k] > A[k+1] ，且当 k 为奇数时， A[k] < A[k+1]。
也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。

返回 A 的最大湍流子数组的长度。

提示：

- 1 <= A.length <= 40000
- 0 <= A[i] <= 10^9



## 解题思路


- 给出一个数组，要求找出“摆动数组”的最大长度。所谓“摆动数组”的意思是，元素一大一小间隔的。
- 这一题可以用滑动窗口来解答。用一个变量记住下次出现的元素需要大于还是需要小于前一个元素。也可以用模拟的方法，用两个变量分别记录上升和下降数字的长度。一旦元素相等了，上升和下降数字长度都置为 1，其他时候按照上升和下降的关系增加队列长度即可，最后输出动态维护的最长长度。
