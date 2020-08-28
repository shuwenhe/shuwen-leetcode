# [1019. Next Greater Node In Linked List](https://leetcode.com/problems/next-greater-node-in-linked-list/)

## 题目

We are given a linked list with head as the first node.  Let's number the nodes in the list: node\_1, node\_2, node\_3, ... etc.

Each node may have a next larger value: for node_i, next\_larger(node\_i) is the node\_j.val such that j > i, node\_j.val > node\_i.val, and j is the smallest possible choice.  If such a j does not exist, the next larger value is 0.

Return an array of integers answer, where answer[i] = next\_larger(node\_{i+1}).

Note that in the example inputs (not outputs) below, arrays such as [2,1,5] represent the serialization of a linked list with a head node value of 2, second node value of 1, and third node value of 5.

 

Example 1:

```c
Input: [2,1,5]
Output: [5,5,0]
```

Example 2:

```c
Input: [2,7,4,3,5]
Output: [7,0,5,5,0]
```

Example 3:

```c
Input: [1,7,5,1,9,2,5,1]
Output: [7,9,9,9,0,5,0,0]
```

Note:

- 1 <= node.val <= 10^9 for each node in the linked list.
- The given list has length in the range [0, 10000].


## 题目大意

给出一个链表，要求找出每个结点后面比该结点值大的第一个结点，如果找不到这个结点，则输出 0 。


## 解题思路

这一题和第 739 题、第 496 题、第 503 题类似。也有 2 种解题方法。先把链表中的数字存到数组中，整道题的思路就和第 739 题完全一致了。普通做法就是 2 层循环。优化的做法就是用单调栈，维护一个单调递减的栈即可。


