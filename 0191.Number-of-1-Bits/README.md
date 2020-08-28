# [191. Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/)

## 题目

Write a function that takes an unsigned integer and return the number of '1' bits it has (also known as the [Hamming weight](http://en.wikipedia.org/wiki/Hamming_weight)).

**Example 1:**

    Input: 00000000000000000000000000001011
    Output: 3
    Explanation: The input binary string 00000000000000000000000000001011 has a total of three '1' bits.

**Example 2:**

    Input: 00000000000000000000000010000000
    Output: 1
    Explanation: The input binary string 00000000000000000000000010000000 has a total of one '1' bit.

**Example 3:**

    Input: 11111111111111111111111111111101
    Output: 31
    Explanation: The input binary string 11111111111111111111111111111101 has a total of thirty one '1' bits.

**Note:**

- Note that in some languages such as Java, there is no unsigned integer type. In this case, the input will be given as signed integer type and should not affect your implementation, as the internal binary representation of the integer is the same whether it is signed or unsigned.
- In Java, the compiler represents the signed integers using [2's complement notation](https://en.wikipedia.org/wiki/Two%27s_complement). Therefore, in **Example 3** above the input represents the signed integer `-3`.


## 题目大意

编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）。

## 解题思路

- 求 uint32 数的二进制位中 1 的个数。
- 这一题的解题思路就是利用二进制位操作。`X = X & ( X -1 )` 这个操作可以清除最低位的二进制位 1，利用这个操作，直至把数清零。操作了几次即为有几个二进制位 1 。
- 最简单的方法即是直接调用库函数 `bits.OnesCount(uint(num))` 。
