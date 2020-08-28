# [92. Reverse Linked List II](https://leetcode.com/problems/reverse-linked-list-ii/)

## 题目

Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

```
Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
```


## 题目大意

给定 2 个链表中结点的位置 m, n，反转这个两个位置区间内的所有结点。

## 解题思路

由于有可能整个链表都被反转，所以构造一个新的头结点指向当前的头。之后的处理方法是：找到第一个需要反转的结点的前一个结点 p，从这个结点开始，依次把后面的结点用“头插”法，插入到 p 结点的后面。循环次数用 n-m 来控制。

这一题结点可以原地变化，更改各个结点的 next 指针就可以。不需要游标 p 指针。因为每次逆序以后，原有结点的相对位置就发生了变化，相当于游标指针已经移动了，所以不需要再有游标 p = p.Next 的操作了。