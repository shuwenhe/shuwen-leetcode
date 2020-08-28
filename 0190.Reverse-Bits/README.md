# [190. Reverse Bits](https://leetcode.com/problems/reverse-bits/)


## 题目

Reverse bits of a given 32 bits unsigned integer.

**Example 1:**

    Input: 00000010100101000001111010011100
    Output: 00111001011110000010100101000000
    Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596, so return 964176192 which its binary representation is 00111001011110000010100101000000.

**Example 2:**

    Input: 11111111111111111111111111111101
    Output: 10111111111111111111111111111111
    Explanation: The input binary string 11111111111111111111111111111101 represents the unsigned integer 4294967293, so return 3221225471 which its binary representation is 10101111110010110010011101101001.

**Note:**

- Note that in some languages such as Java, there is no unsigned integer type. In this case, both input and output will be given as signed integer type and should not affect your implementation, as the internal binary representation of the integer is the same whether it is signed or unsigned.
- In Java, the compiler represents the signed integers using [2's complement notation](https://en.wikipedia.org/wiki/Two%27s_complement). Therefore, in **Example 2** above the input represents the signed integer `-3` and the output represents the signed integer `-1073741825`.

## 题目大意

颠倒给定的 32 位无符号整数的二进制位。提示：

- 请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
- 在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在上面的 示例 2 中，输入表示有符号整数 -3，输出表示有符号整数 -1073741825。

## 解题思路

- 简单题，要求反转 32 位的二进制位。
- 把 num 往右移动，不断的消灭右边最低位的 1，将这个 1 给 res，res 不断的左移即可实现反转二进制位的目的。
