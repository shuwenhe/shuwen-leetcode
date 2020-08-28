# [637. Average of Levels in Binary Tree](https://leetcode.com/problems/average-of-levels-in-binary-tree/)

## 题目

Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.

Example 1:

```c
Input:
    3
   / \
  9  20
    /  \
   15   7
Output: [3, 14.5, 11]
Explanation:
The average value of nodes on level 0 is 3,  on level 1 is 14.5, and on level 2 is 11. Hence return [3, 14.5, 11].
```

Note:  

The range of node's value is in the range of 32-bit signed integer.
 

## 题目大意

按层序从上到下遍历一颗树，计算每一层的平均值。


## 解题思路

- 用一个队列即可实现。
- 第 102 题和第 107 题都是按层序遍历的。


