# [226. Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/)

## 题目

Invert a binary tree.

Example:

Input:

```c
     4
   /   \
  2     7
 / \   / \
1   3 6   9
```

Output:

```c
     4
   /   \
  7     2
 / \   / \
9   6 3   1
```

Trivia:   

This problem was inspired by this original tweet by Max Howell:

>Google: 90% of our engineers use the software you wrote (Homebrew), but you can’t invert a binary tree on a whiteboard so f*** off.
 

## 题目大意

"经典"的反转二叉树的问题。


## 解题思路

还是用递归来解决，先递归调用反转根节点的左孩子，然后递归调用反转根节点的右孩子，然后左右交换根节点的左孩子和右孩子。


