# [897. Increasing Order Search Tree](https://leetcode.com/problems/increasing-order-search-tree/)


## 题目

Given a binary search tree, rearrange the tree in **in-order** so that the leftmost node in the tree is now the root of the tree, and every node has no left child and only 1 right child.

    Example 1:
    Input: [5,3,6,2,4,null,8,1,null,null,null,7,9]
    
           5
          / \
        3    6
       / \    \
      2   4    8
     /        / \ 
    1        7   9
    
    Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
    
     1
      \
       2
        \
         3
          \
           4
            \
             5
              \
               6
                \
                 7
                  \
                   8
                    \
                     9

**Note:**

1. The number of nodes in the given tree will be between 1 and 100.
2. Each node will have a unique integer value from 0 to 1000.


## 题目大意

给定一个树，按中序遍历重新排列树，使树中最左边的结点现在是树的根，并且每个结点没有左子结点，只有一个右子结点。


提示：

- 给定树中的结点数介于 1 和 100 之间。
- 每个结点都有一个从 0 到 1000 范围内的唯一整数值。


## 解题思路

- 给出一颗树，要求把树的所有孩子都排列到右子树上。
- 按照题意，可以先中根遍历原树，然后按照中根遍历的顺序，把所有节点都放在右子树上。见解法二。
- 上一种解法会重新构造一颗新树，有没有办法可以直接更改原有的树呢？节约存储空间。虽然平时软件开发过程中不建议更改原有的值，但是算法题中追求空间和时间的最优，可以考虑一下。**树可以看做是有多个孩子的链表**。这一题可以看成是链表的类似反转的操作。这一点想通以后，就好做了。先找到左子树中最左边的节点，这个节点是新树的根节点。然后依次往回遍历，不断的记录下上一次遍历的最后节点 tail，边遍历，边把后续节点串起来。最终“反转”完成以后，就得到了题目要求的样子了。代码实现见解法一。
