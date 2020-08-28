# [653. Two Sum IV - Input is a BST](https://leetcode.com/problems/two-sum-iv-input-is-a-bst/)

## 题目

Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.

**Example 1:**

    Input: 
        5
       / \
      3   6
     / \   \
    2   4   7
    
    Target = 9
    
    Output: True

**Example 2:**

    Input: 
        5
       / \
      3   6
     / \   \
    2   4   7
    
    Target = 28
    
    Output: False


## 题目大意

给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。

## 解题思路


- 在树中判断是否存在 2 个数的和是 sum。
- 这一题是 two sum 问题的变形题，只不过题目背景是在 BST 上处理的。处理思路大体一致，用 map 记录已经访问过的节点值。边遍历树边查看 map 里面是否有 sum 的另外一半。
