# [918. Maximum Sum Circular Subarray](https://leetcode.com/problems/maximum-sum-circular-subarray/)


## 题目

Given a **circular array** **C** of integers represented by `A`, find the maximum possible sum of a non-empty subarray of **C**.

Here, a *circular array* means the end of the array connects to the beginning of the array. (Formally, `C[i] = A[i]` when `0 <= i < A.length`, and `C[i+A.length] = C[i]` when `i >= 0`.)

Also, a subarray may only include each element of the fixed buffer `A` at most once. (Formally, for a subarray `C[i], C[i+1], ..., C[j]`, there does not exist `i <= k1, k2 <= j` with `k1 % A.length = k2 % A.length`.)

**Example 1:**

    Input: [1,-2,3,-2]
    Output: 3
    Explanation: Subarray [3] has maximum sum 3

**Example 2:**

    Input: [5,-3,5]
    Output: 10
    Explanation: Subarray [5,5] has maximum sum 5 + 5 = 10

**Example 3:**

    Input: [3,-1,2,-1]
    Output: 4
    Explanation: Subarray [2,-1,3] has maximum sum 2 + (-1) + 3 = 4

**Example 4:**

    Input: [3,-2,2,-3]
    Output: 3
    Explanation: Subarray [3] and [3,-2,2] both have maximum sum 3

**Example 5:**

    Input: [-2,-3,-1]
    Output: -1
    Explanation: Subarray [-1] has maximum sum -1

**Note:**

1. `-30000 <= A[i] <= 30000`
2. `1 <= A.length <= 30000`


## 题目大意

给定一个由整数数组 A 表示的环形数组 C，求 C 的非空子数组的最大可能和。在此处，环形数组意味着数组的末端将会与开头相连呈环状。（形式上，当0 <= i < A.length 时 C[i] = A[i]，而当 i >= 0 时 C[i+A.length] = C[i]）

此外，子数组最多只能包含固定缓冲区 A 中的每个元素一次。（形式上，对于子数组 C[i], C[i+1], ..., C[j]，不存在 i <= k1, k2 <= j 其中 k1 % A.length = k2 % A.length）

提示：

- -30000 <= A[i] <= 30000
- 1 <= A.length <= 30000


## 解题思路


- 给出一个环形数组，要求出这个环形数组中的连续子数组的最大和。
- 拿到这题最先想到的思路是把这个数组再拼接一个，在这两个数组中查找连续子数组的最大和。这种做法是错误的，例如在 `[5,-3,5]` 这个数组中会得出 `7` 的结果，但是实际结果是 `10` 。那么这题怎么做呢？仔细分析可以得到，环形数组的最大连续子段和有两种情况，第一种情况是这个连续子段就出现在数组中， 不存在循环衔接的情况。针对这种情况就比较简单，用 `kadane` 算法(也是动态规划的思想)，`O(n)` 的时间复杂度就可以求出结果。第二种情况是这个连续的子段出现在跨数组的情况，即会出现首尾相连的情况。要想找到这样一个连续子段，可以反向考虑。想找到跨段的连续子段，那么这个数组剩下的这一段就是不跨段的连续子段。想要跨段的子段和最大，那么剩下的这段连续子段和最小。如果能找到这个数组的每个元素取相反数组成的数组中的最大连续子段和，那么反过来就能找到原数组的连续子段和最小。举个例子：`[1，2，-3，-4，5]` ，取它的每个元素的相反数 `[-1，-2，3，4，-5]`，构造的数组中最大连续子段和是 `3 + 4 = 7`，由于取了相反数，所以可以得到原数组中最小连续子段和是 `-7` 。所以跨段的最大连续子段和就是剩下的那段 `[1,2,5]`。
- 还有一些边界的情况，例如，`[1，2，-2，-3，5，5，-4，6]` 和 `[1，2，-2，-3，5，5，-4，8]`，所以还需要比较一下情况一和情况二的值，它们两者最大值才是最终环形数组的连续子数组的最大和。
