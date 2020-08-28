# [713. Subarray Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k/)

## 题目

Your are given an array of positive integers nums.

Count and print the number of (contiguous) subarrays where the product of all the elements in the subarray is less than k.

Example 1:

```c
Input: nums = [10, 5, 2, 6], k = 100
Output: 8
Explanation: The 8 subarrays that have product less than 100 are: [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6].
Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.
```


Note:

- 0 < nums.length <= 50000.
- 0 < nums[i] < 1000.
- 0 <= k < 10^6.

## 题目大意

给出一个数组，要求在输出符合条件的窗口数，条件是，窗口中所有数字乘积小于 K 。

## 解题思路

这道题也是滑动窗口的题目，在窗口滑动的过程中不断累乘，直到乘积大于 k，大于 k 的时候就缩小左窗口。有一种情况还需要单独处理一下，即类似 [100] 这种情况。这种情况窗口内乘积等于 k，不小于 k，左边窗口等于右窗口，这个时候需要左窗口和右窗口同时右移。


