# [371. Sum of Two Integers](https://leetcode.com/problems/sum-of-two-integers/)


## 题目

Calculate the sum of two integers a and b, but you are **not allowed** to use the operator `+` and `-`.

**Example 1:**

    Input: a = 1, b = 2
    Output: 3

**Example 2:**

    Input: a = -2, b = 3
    Output: 1


## 题目大意

不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。

## 解题思路

- 要求不用加法和减法运算符计算 `a+b`。这一题需要用到 `^` 和 `&` 运算符的性质，两个数 ^ 可以实现两个数不带进位的二进制加法。这里需要实现加法，肯定需要进位。所以如何找到进位是本题的关键。
- 在二进制中，只有 1 和 1 加在一起才会进位，0 和 0，0 和 1，1 和 0，这三种情况都不会进位，规律就是 `a & b` 为 0 的时候就不用进位，为 1 的时候代表需要进位。进位是往前进一位，所以还需要左移操作，所以加上的进位为 `(a&b)<<1`。
