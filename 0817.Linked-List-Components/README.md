# [817. Linked List Components](https://leetcode.com/problems/linked-list-components/)

## 题目

We are given head, the head node of a linked list containing unique integer values.

We are also given the list G, a subset of the values in the linked list.

Return the number of connected components in G, where two values are connected if they appear consecutively in the linked list.

Example 1:

```c
Input: 
head: 0->1->2->3
G = [0, 1, 3]
Output: 2
Explanation: 
0 and 1 are connected, so [0, 1] and [3] are the two connected components.
```

Example 2:

```c
Input: 
head: 0->1->2->3->4
G = [0, 3, 1, 4]
Output: 2
Explanation: 
0 and 1 are connected, 3 and 4 are connected, so [0, 1] and [3, 4] are the two connected components.
```

Note:

- If N is the length of the linked list given by head, 1 <= N <= 10000.
- The value of each node in the linked list will be in the range [0, N - 1].
- 1 <= G.length <= 10000.
- G is a subset of all values in the linked list.



## 题目大意

这道题题目的意思描述的不是很明白，我提交了几次 WA 以后才悟懂题意。

这道题的意思是，在 G 中能组成多少组子链表，这些子链表的要求是能在原链表中是有序的。

## 解题思路

这个问题再抽象一下就成为这样：在原链表中去掉 G 中不存在的数，会被切断成几段链表。例如，将原链表中 G 中存在的数标为 0，不存在的数标为 1 。原链表标识为 0-0-0-1-0-1-1-0-0-1-0-1，那么这样原链表被断成了 4 段。只要在链表中找 0-1 组合就可以认为是一段，因为这里必定会有一段生成。

考虑末尾的情况，0-1，1-0，0-0，1-1，这 4 种情况的特征都是，末尾一位只要是 0，都会新产生一段。所以链表末尾再单独判断一次，是 0 就再加一。





