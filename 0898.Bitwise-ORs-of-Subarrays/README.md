# [898. Bitwise ORs of Subarrays](https://leetcode.com/problems/bitwise-ors-of-subarrays/)


## 题目

We have an array `A` of non-negative integers.

For every (contiguous) subarray `B = [A[i], A[i+1], ..., A[j]]` (with `i <= j`), we take the bitwise OR of all the elements in `B`, obtaining a result `A[i] | A[i+1] | ... | A[j]`.

Return the number of possible results. (Results that occur more than once are only counted once in the final answer.)

**Example 1:**

    Input: [0]
    Output: 1
    Explanation: 
    There is only one possible result: 0.

**Example 2:**

    Input: [1,1,2]
    Output: 3
    Explanation: 
    The possible subarrays are [1], [1], [2], [1, 1], [1, 2], [1, 1, 2].
    These yield the results 1, 1, 2, 1, 3, 3.
    There are 3 unique values, so the answer is 3.

**Example 3:**

    Input: [1,2,4]
    Output: 6
    Explanation: 
    The possible results are 1, 2, 3, 4, 6, and 7.

**Note:**

1. `1 <= A.length <= 50000`
2. `0 <= A[i] <= 10^9`


## 题目大意

我们有一个非负整数数组 A。对于每个（连续的）子数组 B = [A[i], A[i+1], ..., A[j]] （ i <= j），我们对 B 中的每个元素进行按位或操作，获得结果 A[i] | A[i+1] | ... | A[j]。返回可能结果的数量。（多次出现的结果在最终答案中仅计算一次。）



## 解题思路

- 给出一个数组，要求求出这个数组所有的子数组中，每个集合内所有数字取 `|` 运算以后，不同结果的种类数。
- 这道题可以这样考虑，第一步，先考虑所有的子数组如何得到，以 `[001, 011, 100, 110, 101]` 为例，所有的子数组集合如下：

```c
    [001]  
    [001 011] [011]  
    [001 011 100] [011 100] [100]  
    [001 011 100 110] [011 100 110] [100 110] [110]  
    [001 011 100 110 101] [011 100 110 101] [100 110 101] [110 101] [101]  
```

可以发现，从左往右遍历原数组，每次新来的一个元素，依次加入到之前已经生成过的集合中，再以自己为单独集合。这样就可以生成原数组的所有子集。

- 第二步，将每一行的子集内的所有元素都进行 `|` 运算，得到：

```c
    001  
    011 011  
    111 111 100  
    111 111 110 110  
    111 111 111 111 101  
```

- 第三步，去重：

```c
    001  
    011  
    111 100  
    111 110  
    111 101  
```

由于二进制位不超过 32 位，所以这里每一行最多不会超过 32 个数。所以最终时间复杂度不会超过 O(32 N)，即 O(K * N)。最后将这每一行的数字都放入最终的 map 中去重即可。
