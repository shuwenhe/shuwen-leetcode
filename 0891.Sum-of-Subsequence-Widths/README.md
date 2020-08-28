# [891. Sum of Subsequence Widths](https://leetcode.com/problems/sum-of-subsequence-widths/)

## 题目

Given an array of integers A, consider all non-empty subsequences of A.

For any sequence S, let the width of S be the difference between the maximum and minimum element of S.

Return the sum of the widths of all subsequences of A. 

As the answer may be very large, return the answer modulo 10^9 + 7.

 

Example 1:

```c
Input: [2,1,3]
Output: 6
Explanation:
Subsequences are [1], [2], [3], [2,1], [2,3], [1,3], [2,1,3].
The corresponding widths are 0, 0, 0, 1, 1, 2, 2.
The sum of these widths is 6.
```

Note:

- 1 <= A.length <= 20000
- 1 <= A[i] <= 20000


## 题目大意

给定一个整数数组 A ，考虑 A 的所有非空子序列。对于任意序列 S ，设 S 的宽度是 S 的最大元素和最小元素的差。返回 A 的所有子序列的宽度之和。由于答案可能非常大，请返回答案模 10^9+7。


## 解题思路

- 理解题意以后，可以发现，数组内元素的顺序并不影响最终求得的所有子序列的宽度之和。 
 
		[2,1,3]:[1],[2],[3],[2,1],[2,3],[1,3],[2,1,3]
		[1,2,3]:[1],[2],[3],[1,2],[2,3],[1,3],[1,2,3]
	针对每个 A[i] 而言，A[i] 对最终结果的贡献是在子序列的左右两边的时候才有贡献，当 A[i] 位于区间中间的时候，不影响最终结果。先对 A[i] 进行排序，排序以后，有 i 个数 <= A[i]，有 n - i - 1 个数 >= A[i]。所以 A[i] 会在 2^i 个子序列的右边界出现，2^(n-i-1) 个左边界出现。那么 A[i] 对最终结果的贡献是 A[i] * 2^i - A[i] * 2^(n-i-1) 。举个例子，[1,4,5,7]，A[2] = 5，那么 5 作为右边界的子序列有 2^2 = 4 个，即 [5],[1,5],[4,5],[1,4,5]，5 作为左边界的子序列有 2^(4-2-1) = 2 个，即 [5],[5,7]。A[2] = 5 对最终结果的影响是 5 * 2^2 - 5 * 2^(4-2-1) = 10 。
- 题目要求所有子序列的宽度之和，也就是求每个区间最大值减去最小值的总和。那么 `Ans = SUM{ A[i]*2^i - A[n-i-1] * 2^(n-i-1) }`，其中 `0 <= i < n`。需要注意的是 2^i 可能非常大，所以在计算中就需要去 mod 了，而不是最后计算完了再 mod。注意取模的结合律：`(a * b) % c = (a % c) * (b % c) % c`。