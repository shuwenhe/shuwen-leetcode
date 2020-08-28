# [992. Subarrays with K Different Integers](https://leetcode.com/problems/subarrays-with-k-different-integers/)

## 题目

Given an array A of positive integers, call a (contiguous, not necessarily distinct) subarray of A good if the number of different integers in that subarray is exactly K.

(For example, [1,2,3,1,2] has 3 different integers: 1, 2, and 3.)

Return the number of good subarrays of A.


Example 1:

```c
Input: A = [1,2,1,2,3], K = 2
Output: 7
Explanation: Subarrays formed with exactly 2 different integers: [1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].
```

Example 2:

```c
Input: A = [1,2,1,3,4], K = 3
Output: 3
Explanation: Subarrays formed with exactly 3 different integers: [1,2,1,3], [2,1,3], [1,3,4].
```

Note:  

- 1 <= A.length <= 20000
- 1 <= A[i] <= A.length
- 1 <= K <= A.length

## 题目大意

这道题考察的是滑动窗口的问题。

给出一个数组 和 K，K 代表窗口能能包含的不同数字的个数。K = 2 代表窗口内只能有 2 个不同的数字。求数组中满足条件 K 的窗口个数。

## 解题思路

如果只是单纯的滑动窗口去做，会错过一些解。比如在例子 1 中，滑动窗口可以得到 [1,2], [1,2,1], [1,2,1,2], [2,1,2], [1,2], [2,3], 会少 [2,1] 这个解，原因是右边窗口滑动到最右边了，左边窗口在缩小的过程中，右边窗口不会再跟着动了。有同学可能会说，每次左边窗口移动的时候，右边窗口都再次从左边窗口的位置开始重新滑动。这样做确实可以，但是这样做完会发现超时。因为中间包含大量的重复计算。

这道题就需要第 3 个指针。原有滑动窗口的 2 个指针，右窗口保留这个窗口里面最长的子数组，正好有 K 个元素，左窗口右移的逻辑不变。再多用一个指针用来标识正好有 K - 1 个元素的位置。那么正好有 K 个不同元素的解就等于 ans = atMostK(A, K) - atMostK(A, K - 1)。最多有 K 个元素减去最多有 K - 1 个元素得到的窗口中正好有 K 个元素的解。

以例子 1 为例，先求最多有 K 个元素的窗口个数。

```c
[1]     
[1,2], [2]     
[1,2,1], [2,1], [1]  
[1,2,1,2], [2,1,2], [1,2], [2]  
[2,3], [3]  
```

每当窗口滑动到把 K 消耗为 0 的时候，res = right - left + 1 。为什么要这么计算，right - left + 1 代表的含义是，终点为 right，至多为 K 个元素的窗口有多少个。[left,right], [left + 1,right], [left + 2,right] …… [right,right]。这样算出来的解是包含这道题最终求得的解的，还多出了一部分解。多出来的部分减掉即可，即减掉最多为 K - 1 个元素的解。

最多为 K - 1 个元素的解如下：

```c
[1]
[2]
[1]
[2]
[3]
```

两者相减以后得到的结果就是最终结果：

```c   
[1,2]    
[1,2,1], [2,1]  
[1,2,1,2], [2,1,2], [1,2]  
[2,3]  
```


