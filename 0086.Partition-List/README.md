# [86. Partition List](https://leetcode.com/problems/partition-list/)

## 题目

Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.

You should preserve the original relative order of the nodes in each of the two partitions.

Example:

```
Input: head = 1->4->3->2->5->2, x = 3
Output: 1->2->2->4->3->5
```


## 题目大意

给定一个数 x，比 x 大或等于的数字都要排列在比 x 小的数字后面，并且相对位置不能发生变化。由于相对位置不能发生变化，所以不能用类似冒泡排序的思想。

## 解题思路

这道题最简单的做法是构造双向链表，不过时间复杂度是 O(n^2)。

(以下描述定义，大于等于 x 的都属于比 x 大)

更优的方法是新构造 2 个链表，一个链表专门存储比 x 小的结点，另一个专门存储比 x 大的结点。在原链表头部开始扫描一边，依次把这两类点归类到 2 个新建链表中，有点入栈的意思。由于是从头开始扫描的原链表，所以原链表中的原有顺序会依旧被保存下来。最后 2 个新链表里面会存储好各自的结果，把这两个链表，比 x 小的链表拼接到 比 x 大的链表的前面，就能得到最后的答案了。


