# [338. Counting Bits](https://leetcode.com/problems/counting-bits/)


## 题目

Given a non negative integer number **num**. For every numbers **i** in the range **0 ≤ i ≤ num** calculate the number of 1's in their binary representation and return them as an array.

**Example 1:**

    Input: 2
    Output: [0,1,1]

**Example 2:**

    Input: 5
    Output: [0,1,1,2,1,2]

**Follow up:**

- It is very easy to come up with a solution with run time **O(n*sizeof(integer))**. But can you do it in linear time **O(n)** /possibly in a single pass?
- Space complexity should be **O(n)**.
- Can you do it like a boss? Do it without using any builtin function like **\_\_builtin\_popcount** in c++ or in any other language.

## 题目大意


给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

## 解题思路

- 给出一个数，要求计算出 0 ≤ i ≤ num 中每个数的二进制位 1 的个数。
- 这一题就是利用二进制位运算的经典题。

        X&1==1or==0，可以用 X&1 判断奇偶性，X&1>0 即奇数。
        X = X & (X-1) 清零最低位的1
        X & -X => 得到最低位的1 
        X&~X=>0
