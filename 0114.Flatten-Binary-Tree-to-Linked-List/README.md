# [114. Flatten Binary Tree to Linked List](https://leetcode.com/problems/flatten-binary-tree-to-linked-list/)


## 题目

Given a binary tree, flatten it to a linked list in-place.

For example, given the following tree:

    		1
       / \
      2   5
     / \   \
    3   4   6

The flattened tree should look like:

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

## 题目大意

给定一个二叉树，原地将它展开为链表。

## 解题思路

- 要求把二叉树“打平”，按照先根遍历的顺序，把树的结点都放在右结点中。
- 按照递归和非递归思路实现即可。
- 递归的思路可以这么想：倒序遍历一颗树，即是先遍历右孩子，然后遍历左孩子，最后再遍历根节点。

        		1
           / \
          2   5
         / \   \
        3   4   6
        -----------        
        pre = 5
        cur = 4
        
            1
           / 
          2   
         / \   
        3   4
             \
              5
               \
                6
        -----------        
        pre = 4
        cur = 3
        
            1
           / 
          2   
         /   
        3 
         \
          4
           \
            5
             \
              6
        -----------        
        cur = 2
        pre = 3
        
            1
           / 
          2   
           \
            3 
             \
              4
               \
                5
                 \
                  6
        -----------        
        cur = 1
        pre = 2
        
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

- 可以先仿造先根遍历的代码，写出这个倒序遍历的逻辑：

        public void flatten(TreeNode root) {
            if (root == null)
                return;
            flatten(root.right);
            flatten(root.left);
        }

- 实现了倒序遍历的逻辑以后，再进行结点之间的拼接：

        private TreeNode prev = null;
        
        public void flatten(TreeNode root) {
            if (root == null)
                return;
            flatten(root.right);
            flatten(root.left);
            root.right = prev;
            root.left = null;
            prev = root;
        }
