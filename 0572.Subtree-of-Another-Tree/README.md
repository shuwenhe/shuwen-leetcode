# [572. Subtree of Another Tree](https://leetcode.com/problems/subtree-of-another-tree/)


## 题目

Given two non-empty binary trees **s** and **t**, check whether tree **t** has exactly the same structure and node values with a subtree of **s**. A subtree of **s** is a tree consists of a node in **s** and all of this node's descendants. The tree **s** could also be considered as a subtree of itself.

**Example 1:** Given tree s:

		 3
        / \
       4   5
      / \
     1   2

Given tree t:

       4 
      / \
     1   2

Return **true**, because t has the same structure and node values with a subtree of s.

**Example 2:**Given tree s:

         3
        / \
       4   5
      / \
     1   2
        /
       0

Given tree t:

       4
      / \
     1   2

Return **false**.


## 题目大意

给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。


## 解题思路


- 给出 2 棵树 `s` 和 `t`，要求判断 `t` 是否是 `s` 的子树🌲。
- 这一题比较简单，针对 3 种情况依次递归判断，第一种情况 `s` 和 `t` 是完全一样的两棵树，第二种情况 `t` 是 `s` 左子树中的子树，第三种情况 `t` 是 `s` 右子树中的子树。第一种情况判断两棵数是否完全一致是第 100 题的原题。

