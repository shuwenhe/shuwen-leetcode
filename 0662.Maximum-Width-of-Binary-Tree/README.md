# [662. Maximum Width of Binary Tree](https://leetcode.com/problems/maximum-width-of-binary-tree/)


## 题目

Given a binary tree, write a function to get the maximum width of the given tree. The width of a tree is the maximum width among all levels. The binary tree has the same structure as a **full binary tree**, but some nodes are null.

The width of one level is defined as the length between the end-nodes (the leftmost and right most non-null nodes in the level, where the `null` nodes between the end-nodes are also counted into the length calculation.

**Example 1:**

    Input: 
    
               1
             /   \
            3     2
           / \     \  
          5   3     9 
    
    Output: 4
    Explanation: The maximum width existing in the third level with the length 4 (5,3,null,9).

**Example 2:**

    Input: 
    
              1
             /  
            3    
           / \       
          5   3     
    
    Output: 2
    Explanation: The maximum width existing in the third level with the length 2 (5,3).

**Example 3:**

    Input: 
    
              1
             / \
            3   2 
           /        
          5      
    
    Output: 2
    Explanation: The maximum width existing in the second level with the length 2 (3,2).

**Example 4:**

    Input: 
    
              1
             / \
            3   2
           /     \  
          5       9 
         /         \
        6           7
    Output: 8
    Explanation:The maximum width existing in the fourth level with the length 8 (6,null,null,null,null,null,null,7).

**Note:** Answer will in the range of 32-bit signed integer.


## 题目大意

给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。这个二叉树与满二叉树（full binary tree）结构相同，但一些节点为空。

每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。

注意: 答案在32位有符号整数的表示范围内。



## 解题思路


- 给出一个二叉树，求这棵树最宽的部分。
- 这一题可以用 BFS 也可以用 DFS，但是用 BFS 比较方便。按照层序遍历，依次算出每层最左边不为 `null` 的节点和最右边不为 `null` 的节点。这两个节点之间都是算宽度的。最终输出最大的宽度即可。此题的关键在于如何有效的找到每一层的左右边界。
- 这一题可能有人会想着先补全满二叉树，然后每层分别找左右边界。这种方法提交以后会卡在 `104 / 108` 这组测试用例上，这组测试用例会使得最后某几层填充出现的满二叉树节点特别多，最终导致 `Memory Limit Exceeded` 了。
- 由于此题要找每层的左右边界，实际上每个节点的 `Val` 值是我们不关心的，那么可以把这个值用来标号，标记成该节点在每层中的序号。父亲节点在上一层中的序号是 x，那么它的左孩子在下一层满二叉树中的序号是 `2*x`，它的右孩子在下一层满二叉树中的序号是 `2*x + 1`。将所有节点都标上号，用 BFS 层序遍历每一层，每一层都找到左右边界，相减拿到宽度，动态维护最大宽度，就是本题的最终答案。