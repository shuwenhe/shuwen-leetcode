# [328. Odd Even Linked List](https://leetcode.com/problems/odd-even-linked-list/)

## 题目

Given a singly linked list, group all odd nodes together followed by the even nodes. Please note here we are talking about the node number and not the value in the nodes.

You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.

Example 1:

```c
Input: 1->2->3->4->5->NULL
Output: 1->3->5->2->4->NULL
```

Example 2:

```c
Input: 2->1->3->5->6->4->7->NULL
Output: 2->3->6->7->1->5->4->NULL
```

Note:

- The relative order inside both the even and odd groups should remain as it was in the input.
- The first node is considered odd, the second node even and so on ...

## 题目大意

这道题和第 86 题非常类型。第 86 题是把排在某个点前面的小值放在一个链表中，排在某个点后端的大值放在另外一个链表中，最后 2 个链表首尾拼接一下就是答案。

## 解题思路

这道题思路也是一样的，分别把奇数和偶数都放在 2 个链表中，最后首尾拼接就是答案。