# [904. Fruit Into Baskets](https://leetcode.com/problems/fruit-into-baskets/)

## 题目

In a row of trees, the i-th tree produces fruit with type tree[i].

You start at any tree of your choice, then repeatedly perform the following steps:

1. Add one piece of fruit from this tree to your baskets.  If you cannot, stop.
2. Move to the next tree to the right of the current tree.  If there is no tree to the right, stop.

Note that you do not have any choice after the initial choice of starting tree: you must perform step 1, then step 2, then back to step 1, then step 2, and so on until you stop.

You have two baskets, and each basket can carry any quantity of fruit, but you want each basket to only carry one type of fruit each.

What is the total amount of fruit you can collect with this procedure?


Example 1:

```c
Input: [1,2,1]
Output: 3
Explanation: We can collect [1,2,1].
```

Example 2:

```c
Input: [0,1,2,2]
Output: 3
Explanation: We can collect [1,2,2].
If we started at the first tree, we would only collect [0, 1].
```

Example 3:

```c
Input: [1,2,3,2,2]
Output: 4
Explanation: We can collect [2,3,2,2].
If we started at the first tree, we would only collect [1, 2].
```

Example 4:

```c
Input: [3,3,3,1,2,1,1,2,3,3,4]
Output: 5
Explanation: We can collect [1,2,1,1,2].
If we started at the first tree or the eighth tree, we would only collect 4 fruits.
```

Note:

- 1 <= tree.length <= 40000
- 0 <= tree[i] < tree.length

## 题目大意

这道题考察的是滑动窗口的问题。

给出一个数组，数组里面的数字代表每个果树上水果的种类，1 代表一号水果，不同数字代表的水果不同。现在有 2 个篮子，每个篮子只能装一个种类的水果，这就意味着只能选 2 个不同的数字。摘水果只能从左往右摘，直到右边没有水果可以摘就停下。问可以连续摘水果的最长区间段的长度。


## 解题思路

简化一下题意，给出一段数字，要求找出包含 2 个不同数字的最大区间段长度。这个区间段内只能包含这 2 个不同数字，可以重复，但是不能包含其他数字。

用典型的滑动窗口的处理方法处理即可。


