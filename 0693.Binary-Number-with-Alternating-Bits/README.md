# [693. Binary Number with Alternating Bits](https://leetcode.com/problems/binary-number-with-alternating-bits/)

## 题目

Given a positive integer, check whether it has alternating bits: namely, if two adjacent bits will always have different values.

**Example 1:**

    Input: 5
    Output: True
    Explanation:
    The binary representation of 5 is: 101

**Example 2:**

    Input: 7
    Output: False
    Explanation:
    The binary representation of 7 is: 111.

**Example 3:**

    Input: 11
    Output: False
    Explanation:
    The binary representation of 11 is: 1011.

**Example 4:**

    Input: 10
    Output: True
    Explanation:
    The binary representation of 10 is: 1010.


## 题目大意

给定一个正整数，检查他是否为交替位二进制数：换句话说，就是他的二进制数相邻的两个位数永不相等。

## 解题思路


- 判断一个数的二进制位相邻两个数是不相等的，即 `0101` 交叉间隔的，如果是，输出 true。这一题有多种做法，最简单的方法就是直接模拟。比较巧妙的方法是通过位运算，合理构造特殊数据进行位运算到达目的。`010101` 构造出 `101010` 两者相互 `&` 位运算以后就为 0，因为都“插空”了。
