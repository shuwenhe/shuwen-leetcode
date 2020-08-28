# [437. Path Sum III](https://leetcode.com/problems/path-sum-iii/)


## 题目

You are given a binary tree in which each node contains an integer value.

Find the number of paths that sum to a given value.

The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).

The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.

**Example:**

    root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
    
          10
         /  \
        5   -3
       / \    \
      3   2   11
     / \   \
    3  -2   1
    
    Return 3. The paths that sum to 8 are:
    
    1.  5 -> 3
    2.  5 -> 2 -> 1
    3. -3 -> 11


## 题目大意

给定一个二叉树，它的每个结点都存放着一个整数值。找出路径和等于给定数值的路径总数。路径不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。二叉树不超过1000个节点，且节点数值范围是 [-1000000,1000000] 的整数。


## 解题思路


- 这一题是第 112 题 Path Sum 和第 113 题 Path Sum II 的加强版，这一题要求求出任意一条路径的和为 sum，起点不一定是根节点，可以是任意节点开始。
- 注意这一题可能出现负数的情况，节点和为 sum，并不一定是最终情况，有可能下面还有正数节点和负数节点相加正好为 0，那么这也是一种情况。一定要遍历到底。
- 一个点是否为 sum 的起点，有 3 种情况，第一种情况路径包含该 root 节点，如果包含该结点，就在它的左子树和右子树中寻找和为 `sum-root.Val` 的情况。第二种情况路径不包含该 root 节点，那么就需要在它的左子树和右子树中分别寻找和为 sum 的结点。

