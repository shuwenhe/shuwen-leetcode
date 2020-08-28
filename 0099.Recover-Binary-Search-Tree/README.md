# [99. Recover Binary Search Tree](https://leetcode.com/problems/recover-binary-search-tree/)


## 题目

Two elements of a binary search tree (BST) are swapped by mistake.

Recover the tree without changing its structure.

**Example 1:**

    Input: [1,3,null,null,2]
    
       1
      /
     3
      \
       2
    
    Output: [3,1,null,null,2]
    
       3
      /
     1
      \
       2

**Example 2:**

    Input: [3,1,4,null,null,2]
    
      3
     / \
    1   4
       /
      2
    
    Output: [2,1,4,null,null,3]
    
      2
     / \
    1   4
       /
      3

**Follow up:**

- A solution using O(*n*) space is pretty straight forward.
- Could you devise a constant space solution?

## 题目大意

二叉搜索树中的两个节点被错误地交换。请在不改变其结构的情况下，恢复这棵树。


## 解题思路

- 在二叉搜索树中，有 2 个结点的值出错了，要求修复这两个结点。
- 这一题按照先根遍历 1 次就可以找到这两个出问题的结点，因为先访问根节点，然后左孩子，右孩子。用先根遍历二叉搜索树的时候，根结点比左子树都要大，根结点比右子树都要小。所以`左子树比根结点大的话，就是出现了乱序`；`根节点比右子树大的话，就是出现了乱序`。遍历过程中在左子树中如果出现了前一次遍历的结点的值大于此次根节点的值，这就出现了出错结点了，记录下来。继续遍历直到找到第二个这样的结点。最后交换这两个结点的时候，只是交换他们的值就可以了，而不是交换这两个结点相应的指针指向。
