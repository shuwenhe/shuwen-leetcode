# [1123. Lowest Common Ancestor of Deepest Leaves](https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/)


## 题目

Given a rooted binary tree, return the lowest common ancestor of its deepest leaves.

Recall that:

- The node of a binary tree is a *leaf* if and only if it has no children
- The *depth* of the root of the tree is 0, and if the depth of a node is `d`, the depth of each of its children is `d+1`.
- The *lowest common ancestor* of a set `S` of nodes is the node `A` with the largest depth such that every node in S is in the subtree with root `A`.

**Example 1:**

    Input: root = [1,2,3]
    Output: [1,2,3]
    Explanation: 
    The deepest leaves are the nodes with values 2 and 3.
    The lowest common ancestor of these leaves is the node with value 1.
    The answer returned is a TreeNode object (not an array) with serialization "[1,2,3]".

**Example 2:**

    Input: root = [1,2,3,4]
    Output: [4]

**Example 3:**

    Input: root = [1,2,3,4,5]
    Output: [2,4,5]

**Constraints:**

- The given tree will have between 1 and 1000 nodes.
- Each node of the tree will have a distinct value between 1 and 1000.


## 题目大意


给你一个有根节点的二叉树，找到它最深的叶节点的最近公共祖先。

回想一下：

- 叶节点 是二叉树中没有子节点的节点
- 树的根节点的 深度 为 0，如果某一节点的深度为 d，那它的子节点的深度就是 d+1
- 如果我们假定 A 是一组节点 S 的 最近公共祖先，<font color="#c7254e" face="Menlo, Monaco, Consolas, Courier New, monospace">S</font> 中的每个节点都在以 A 为根节点的子树中，且 A 的深度达到此条件下可能的最大值。
 

提示：

- 给你的树中将有 1 到 1000 个节点。
- 树中每个节点的值都在 1 到 1000 之间。


## 解题思路


- 给出一颗树，找出最深的叶子节点的最近公共祖先 LCA。
- 这一题思路比较直接。先遍历找到最深的叶子节点，如果左右子树的最深的叶子节点深度相同，那么当前节点就是它们的最近公共祖先。如果左右子树的最深的深度不等，那么需要继续递归往下找符合题意的 LCA。如果最深的叶子节点没有兄弟，那么公共父节点就是叶子本身，否则返回它的 LCA。
- 有几个特殊的测试用例，见测试文件。特殊的点就是最深的叶子节点没有兄弟节点的情况。
