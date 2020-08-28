# [763. Partition Labels](https://leetcode.com/problems/partition-labels/)

## 题目

A string S of lowercase letters is given. We want to partition this string into as many parts as possible so that each letter appears in at most one part, and return a list of integers representing the size of these parts.



Example 1:


```c
Input: S = "ababcbacadefegdehijhklij"
Output: [9,7,8]
Explanation:
The partition is "ababcbaca", "defegde", "hijhklij".
This is a partition so that each letter appears in at most one part.
A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.
```

Note:

- S will have length in range [1, 500].
- S will consist of lowercase letters ('a' to 'z') only.


## 题目大意

这道题考察的是滑动窗口的问题。

给出一个字符串，要求输出满足条件窗口的长度，条件是在这个窗口内，字母中出现在这一个窗口内，不出现在其他窗口内。

## 解题思路

这一题有 2 种思路，第一种思路是先记录下每个字母的出现次数，然后对滑动窗口中的每个字母判断次数是否用尽为 0，如果这个窗口内的所有字母次数都为 0，这个窗口就是符合条件的窗口。时间复杂度为 O(n^2)

另外一种思路是记录下每个字符最后一次出现的下标，这样就不用记录次数。在每个滑动窗口中，依次判断每个字母最后一次出现的位置，如果在一个下标内，所有字母的最后一次出现的位置都包含进来了，那么这个下标就是这个满足条件的窗口大小。时间复杂度为 O(n^2)
