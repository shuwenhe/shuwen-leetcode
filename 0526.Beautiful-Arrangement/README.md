# [526. Beautiful Arrangement](https://leetcode.com/problems/beautiful-arrangement/)


## 题目

Suppose you have **N** integers from 1 to N. We define a beautiful arrangement as an array that is constructed by these **N** numbers successfully if one of the following is true for the ith position (1 <= i <= N) in this array:

1. The number at the i position is divisible by **i**.th
2. **i** is divisible by the number at the i position.th

Now given N, how many beautiful arrangements can you construct?

**Example 1:**

    Input: 2
    Output: 2
    Explanation: 
    
    The first beautiful arrangement is [1, 2]:
    
    Number at the 1st position (i=1) is 1, and 1 is divisible by i (i=1).
    
    Number at the 2nd position (i=2) is 2, and 2 is divisible by i (i=2).
    
    The second beautiful arrangement is [2, 1]:
    
    Number at the 1st position (i=1) is 2, and 2 is divisible by i (i=1).
    
    Number at the 2nd position (i=2) is 1, and i (i=2) is divisible by 1.

**Note:**

1. **N** is a positive integer and will not exceed 15.


## 题目大意

假设有从 1 到 N 的 N 个整数，如果从这 N 个数字中成功构造出一个数组，使得数组的第 i 位 (1 <= i <= N) 满足如下两个条件中的一个，我们就称这个数组为一个优美的排列。条件：

- 第 i 位的数字能被 i 整除
- i 能被第 i 位上的数字整除

现在给定一个整数 N，请问可以构造多少个优美的排列？



## 解题思路


- 这一题是第 46 题的加强版。由于这一题给出的数组里面的数字都是不重复的，所以可以当做第 46 题来做。
- 这题比第 46 题多的一个条件是，要求数字可以被它对应的下标 + 1 整除，或者下标 + 1 可以整除下标对应的这个数字。在 DFS 回溯过程中加入这个剪枝条件就可以了。
- 当前做法时间复杂度不是最优的，大概只有 33.3%
