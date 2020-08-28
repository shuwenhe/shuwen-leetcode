# [331. Verify Preorder Serialization of a Binary Tree](https://leetcode.com/problems/verify-preorder-serialization-of-a-binary-tree/)

## 题目

One way to serialize a binary tree is to use pre-order traversal. When we encounter a non-null node, we record the node's value. If it is a null node, we record using a sentinel value such as #.

```c
     _9_
    /   \
   3     2
  / \   / \
 4   1  #  6
/ \ / \   / \
# # # #   # #
```

For example, the above binary tree can be serialized to the string "9,3,4,#,#,1,#,#,2,#,6,#,#", where # represents a null node.

Given a string of comma separated values, verify whether it is a correct preorder traversal serialization of a binary tree. Find an algorithm without reconstructing the tree.

Each comma separated value in the string must be either an integer or a character '#' representing null pointer.

You may assume that the input format is always valid, for example it could never contain two consecutive commas such as "1,,3".

Example 1:

```c
Input: "9,3,4,#,#,1,#,#,2,#,6,#,#"
Output: true
```

Example 2:

```c
Input: "1,#"
Output: false
```
Example 3:

```c
Input: "9,#,#,1"
Output: false
```

## 题目大意

给定一串以逗号分隔的序列，验证它是否是正确的二叉树的前序序列化。编写一个在不重构树的条件下的可行算法。

## 解题思路

这道题有些人用栈，有些用栈的深度求解。换个视角。如果叶子结点是 null，那么所有非 null 的结点(除了 root 结点)必然有 2 个出度，1 个入度(2 个孩子和 1 个父亲，孩子可能为空，但是这一题用 "#" 代替了，所以肯定有 2 个孩子)；所有的 null 结点只有 0 个出度，1 个入度(0 个孩子和 1 个父亲)。

我们开始构建这颗树，在构建过程中，我们记录出度和度之间的差异 `diff = outdegree - indegree`。当下一个节点到来时，我们将 diff 减 1，因为这个节点提供了一个度。如果这个节点不为 null，我们将 diff 增加 2，因为它提供两个出度。如果序列化是正确的，则 diff 应该永远不会为负，并且 diff 在完成时将为零。最后判断一下 diff 是不是为 0 即可判断它是否是正确的二叉树的前序序列化。