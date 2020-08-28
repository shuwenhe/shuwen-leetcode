# [508. Most Frequent Subtree Sum](https://leetcode.com/problems/most-frequent-subtree-sum/)


## 题目

Given the root of a tree, you are asked to find the most frequent subtree sum. The subtree sum of a node is defined as the sum of all the node values formed by the subtree rooted at that node (including the node itself). So what is the most frequent subtree sum value? If there is a tie, return all the values with the highest frequency in any order.

**Examples 1**Input:

      5
     /  \
    2   -3

return [2, -3, 4], since all the values happen only once, return all of them in any order.

**Examples 2**Input:

      5
     /  \
    2   -5

return [2], since 2 happens twice, however -5 only occur once.

**Note:** You may assume the sum of values in any subtree is in the range of 32-bit signed integer.


## 题目大意

给出二叉树的根，找出出现次数最多的子树元素和。一个结点的子树元素和定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。然后求出出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的元素（不限顺序）。提示： 假设任意子树元素和均可以用 32 位有符号整数表示。

## 解题思路


- 给出一个树，要求求出每个节点以自己为根节点的子树的所有节点值的和，最后按照这些和出现的频次，输出频次最多的和，如果频次出现次数最多的对应多个和，则全部输出。
- 递归找出每个节点的累加和，用 map 记录频次，最后把频次最多的输出即可。
