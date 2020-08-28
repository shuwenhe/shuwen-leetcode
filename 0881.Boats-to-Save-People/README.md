# [881. Boats to Save People](https://leetcode.com/problems/boats-to-save-people/)

## 题目

The i-th person has weight people[i], and each boat can carry a maximum weight of limit.

Each boat carries at most 2 people at the same time, provided the sum of the weight of those people is at most limit.

Return the minimum number of boats to carry every given person.  (It is guaranteed each person can be carried by a boat.)


Example 1:

```c
Input: people = [1,2], limit = 3
Output: 1
Explanation: 1 boat (1, 2)
```


Example 2:

```c
Input: people = [3,2,2,1], limit = 3
Output: 3
Explanation: 3 boats (1, 2), (2) and (3)
```


Example 3:

```c
Input: people = [3,5,3,4], limit = 5
Output: 4
Explanation: 4 boats (3), (3), (4), (5)
```

Note:

- 1 <= people.length <= 50000
- 1 <= people[i] <= limit <= 30000


## 题目大意

给出人的重量数组，和一个船最大载重量 limit。一个船最多装 2 个人。要求输出装下所有人，最小需要多少艘船。

## 解题思路

先对人的重量进行排序，然后用 2 个指针分别指向一前一后，一起计算这两个指针指向的重量之和，如果小于 limit，左指针往右移动，并且右指针往左移动。如果大于等于 limit，右指针往左移动。每次指针移动，需要船的个数都要 ++。


