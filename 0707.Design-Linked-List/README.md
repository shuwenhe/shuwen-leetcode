# [707. Design Linked List](https://leetcode.com/problems/design-linked-list/)

## 题目

Design your implementation of the linked list. You can choose to use the singly linked list or the doubly linked list. A node in a singly linked list should have two attributes: val and next. val is the value of the current node, and next is a pointer/reference to the next node. If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node in the linked list. Assume all nodes in the linked list are 0-indexed.

Implement these functions in your linked list class:

- get(index) : Get the value of the index-th node in the linked list. If the index is invalid, return -1.
- addAtHead(val) : Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
- addAtTail(val) : Append a node of value val to the last element of the linked list.
- addAtIndex(index, val) : Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
- deleteAtIndex(index) : Delete the index-th node in the linked list, if the index is valid.

Example:

```c
MyLinkedList linkedList = new MyLinkedList();
linkedList.addAtHead(1);
linkedList.addAtTail(3);
linkedList.addAtIndex(1, 2);  // linked list becomes 1->2->3
linkedList.get(1);            // returns 2
linkedList.deleteAtIndex(1);  // now the linked list is 1->3
linkedList.get(1);            // returns 3
```

Note:

- All values will be in the range of [1, 1000].
- The number of operations will be in the range of [1, 1000].
- Please do not use the built-in LinkedList library.

## 题目大意

这道题比较简单，设计一个链表，实现相关操作即可。

## 解题思路

这题有一个地方比较坑，题目中 Note 里面写的数值取值范围是 [1, 1000]，笔者把 0 当做无效值。结果 case 里面出现了 0 是有效值。case 和题意不符。
