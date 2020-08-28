# [930. Binary Subarrays With Sum](https://leetcode.com/problems/binary-subarrays-with-sum/)

## 题目

In an array A of 0s and 1s, how many non-empty subarrays have sum S?



Example 1:

```c
Input: A = [1,0,1,0,1], S = 2
Output: 4
Explanation: 
The 4 subarrays are bolded below:
[1,0,1,0,1]
[1,0,1,0,1]
[1,0,1,0,1]
[1,0,1,0,1]
```


Note:

- A.length <= 30000
- 0 <= S <= A.length
- A[i] is either 0 or 1.


## 题目大意

给定一个数组，数组里面的元素只有 0 和 1 两种。问这个数组有多少个和为 S 的子数组。

## 解题思路

这道题也是滑动窗口的题目。不断的加入右边的值，直到总和等于 S。[i,j] 区间内的和可以等于 [0,j] 的和减去 [0,i-1] 的和。在 freq 中不断的记下能使得和为 sum 的组合方法数，例如 freq[1] = 2 ，代表和为 1 有两种组合方法，(可能是 1 和 1，0 或者 0，1，这道题只管组合总数，没要求输出具体的组合对)。这道题的做法就是不断的累加，如果遇到比 S 多的情况，多出来的值就在 freq 中查表，看多出来的值可能是由几种情况构成的。一旦和与 S 相等以后，之后比 S 多出来的情况会越来越多(因为在不断累积，总和只会越来越大)，不断的查 freq 表就可以了。
