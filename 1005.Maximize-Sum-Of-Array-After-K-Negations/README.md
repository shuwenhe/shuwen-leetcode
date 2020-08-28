# [1005. Maximize Sum Of Array After K Negations](https://leetcode.com/problems/maximize-sum-of-array-after-k-negations/)

## 题目

Given an array A of integers, we must modify the array in the following way: we choose an i and replace A[i] with -A[i], and we repeat this process K times in total.  (We may choose the same index i multiple times.)

Return the largest possible sum of the array after modifying it in this way.


Example 1:

```c
Input: A = [4,2,3], K = 1
Output: 5
Explanation: Choose indices (1,) and A becomes [4,-2,3].
```

Example 2:

```c
Input: A = [3,-1,0,2], K = 3
Output: 6
Explanation: Choose indices (1, 2, 2) and A becomes [3,1,0,2].
```

Example 3:

```c
Input: A = [2,-3,-1,5,-4], K = 2
Output: 13
Explanation: Choose indices (1, 4) and A becomes [2,3,-1,5,4].
```

Note:

- 1 <= A.length <= 10000
- 1 <= K <= 10000
- -100 <= A[i] <= 100

## 题目大意

将数组中的元素变成它的相反数，这种操作执行 K 次之后，求出数组中所有元素的总和最大。

## 解题思路

这一题可以用最小堆来做，构建最小堆，每次将最小的元素变成它的相反数。然后最小堆调整，再将新的最小元素变成它的相反数。执行 K 次以后求数组中所有的值之和就是最大值。

这道题也可以用排序来实现。排序一次，从最小值开始往后扫，依次将最小值变为相反数。这里需要注意一点，负数都改变成正数以后，接着不是再改变这些变成正数的负数，而是接着改变顺序的正数。因为这些正数是比较小的正数。负数越小，变成正数以后值越大。正数越小，变成负数以后对总和影响最小。具体实现见代码。
