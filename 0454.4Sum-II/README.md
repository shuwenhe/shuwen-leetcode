# [454. 4Sum II](https://leetcode.com/problems/4sum-ii/)

## 题目

Given four lists A, B, C, D of integer values, compute how many tuples (i, j, k, l) there are such that A[i] + B[j] + C[k] + D[l] is zero.

To make problem a bit easier, all A, B, C, D have same length of N where 0 ≤ N ≤ 500. All integers are in the range of -228 to 228 - 1 and the result is guaranteed to be at most 231 - 1.

Example 1:

```c
Input:
A = [ 1, 2]
B = [-2,-1]
C = [-1, 2]
D = [ 0, 2]

Output:
2

Explanation:
The two tuples are:
1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0
```


## 题目大意

给出 4 个数组，计算这些数组中存在几对 i，j，k，l 可以使得 A[i] + B[j] + C[k] + D[l] = 0 。

## 解题思路

这道题的数据量不大，0 ≤ N ≤ 500，但是如果使用暴力解法，四层循环，会超时。这道题的思路和第 1 题思路也类似，先可以将 2 个数组中的组合都存入 map 中。之后将剩下的 2 个数组进行 for 循环，找出和为 0 的组合。这样时间复杂度是 O(n^2)。当然也可以讲剩下的 2 个数组的组合也存入 map 中，不过最后在 2 个 map 中查找结果也是 O(n^2) 的时间复杂度。



