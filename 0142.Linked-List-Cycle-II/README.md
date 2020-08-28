# [142. Linked List Cycle II](https://leetcode.com/problems/linked-list-cycle-ii/)

## 题目

Given a linked list, return the node where the cycle begins. If there is no cycle, return null.

To represent a cycle in the given linked list, we use an integer pos which represents the position (0-indexed) in the linked list where tail connects to. If pos is -1, then there is no cycle in the linked list.

Note: Do not modify the linked list.

Example 1:

```c
Input: head = [3,2,0,-4], pos = 1
Output: tail connects to node index 1
Explanation: There is a cycle in the linked list, where tail connects to the second node.
```

Example 2:

```c
Input: head = [1,2], pos = 0
Output: tail connects to node index 0
Explanation: There is a cycle in the linked list, where tail connects to the first node.
```

Example 3:

```c
Input: head = [1], pos = -1
Output: no cycle
Explanation: There is no cycle in the linked list.
```


## 题目大意

判断链表是否有环，不能使用额外的空间。如果有环，输出环的起点指针，如果没有环，则输出空。

## 解题思路

这道题是第 141 题的加强版。在判断是否有环的基础上，还需要输出环的第一个点。

分析一下判断环的原理。fast 指针一次都 2 步，slow 指针一次走 1 步。令链表 head 到环的一个点需要 x1 步，从环的第一个点到相遇点需要 x2 步，从环中相遇点回到环的第一个点需要 x3 步。那么环的总长度是 x2 + x3 步。

fast 和 slow 会相遇，说明他们走的时间是相同的，可以知道他们走的路程有以下的关系：

```c
fast 的 t = (x1 + x2 + x3 + x2) / 2
slow 的 t = (x1 + x2) / 1

x1 + x2 + x3 + x2 = 2 * (x1 + x2)

所以 x1 = x3
```

所以 2 个指针相遇以后，如果 slow 继续往前走，fast 指针回到起点 head，两者都每次走一步，那么必定会在环的起点相遇，相遇以后输出这个点即是结果。

