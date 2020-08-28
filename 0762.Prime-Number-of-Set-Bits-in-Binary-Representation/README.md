# [762. Prime Number of Set Bits in Binary Representation](https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/)


## 题目

Given two integers `L` and `R`, find the count of numbers in the range `[L, R]` (inclusive) having a prime number of set bits in their binary representation.

(Recall that the number of set bits an integer has is the number of `1`s present when written in binary. For example, `21` written in binary is `10101` which has 3 set bits. Also, 1 is not a prime.)

**Example 1:**

    Input: L = 6, R = 10
    Output: 4
    Explanation:
    6 -> 110 (2 set bits, 2 is prime)
    7 -> 111 (3 set bits, 3 is prime)
    9 -> 1001 (2 set bits , 2 is prime)
    10->1010 (2 set bits , 2 is prime)

**Example 2:**

    Input: L = 10, R = 15
    Output: 5
    Explanation:
    10 -> 1010 (2 set bits, 2 is prime)
    11 -> 1011 (3 set bits, 3 is prime)
    12 -> 1100 (2 set bits, 2 is prime)
    13 -> 1101 (3 set bits, 3 is prime)
    14 -> 1110 (3 set bits, 3 is prime)
    15 -> 1111 (4 set bits, 4 is not prime)

**Note:**

1. `L, R` will be integers `L <= R` in the range `[1, 10^6]`.
2. `R - L` will be at most 10000.


## 题目大意

给定两个整数 L 和 R ，找到闭区间 [L, R] 范围内，计算置位位数为质数的整数个数。（注意，计算置位代表二进制表示中1的个数。例如 21 的二进制表示 10101 有 3 个计算置位。还有，1 不是质数。）


注意:

- L, R 是 L <= R 且在 [1, 10^6] 中的整数。
- R - L 的最大值为 10000。



## 解题思路


- 题目给出 `[L, R]` 区间，在这个区间内的每个整数的二进制表示中 1 的个数如果是素数，那么最终结果就加一，问最终结果是多少？这一题是一个组合题，判断一个数的二进制位有多少位 1，是第 191 题。题目中限定了区间最大不超过 10^6 ，所以 1 的位数最大是 19 位，也就是说素数最大就是 19 。那么素数可以有限枚举出来。最后按照题目的意思累积结果就可以了。
