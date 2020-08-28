# [101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/)

## 题目


Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:


```c
    1
   / \
  2   2
 / \ / \
3  4 4  3
```

But the following [1,2,2,null,3,null,3] is not:

```c
    1
   / \
  2   2
   \   \
   3    3
```

Note:   

Bonus points if you could solve it both recursively and iteratively.

## 题目大意

这一题要求判断 2 颗树是否是左右对称的。


## 解题思路

- 这道题是几道题的综合题。将根节点的左字数反转二叉树，然后再和根节点的右节点进行比较，是否完全相等。
- 反转二叉树是第 226 题。判断 2 颗树是否完全相等是第 100 题。


