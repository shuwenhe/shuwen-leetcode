# [405. Convert a Number to Hexadecimal](https://leetcode.com/problems/convert-a-number-to-hexadecimal/)


## 题目

Given an integer, write an algorithm to convert it to hexadecimal. For negative integer, [two’s complement](https://en.wikipedia.org/wiki/Two%27s_complement) method is used.

**Note:**

1. All letters in hexadecimal (`a-f`) must be in lowercase.
2. The hexadecimal string must not contain extra leading `0`s. If the number is zero, it is represented by a single zero character `'0'`; otherwise, the first character in the hexadecimal string will not be the zero character.
3. The given number is guaranteed to fit within the range of a 32-bit signed integer.
4. You **must not use any method provided by the library** which converts/formats the number to hex directly.

**Example 1:**

    Input:
    26
    
    Output:
    "1a"

**Example 2:**

    Input:
    -1
    
    Output:
    "ffffffff"


## 题目大意

给定一个整数，编写一个算法将这个数转换为十六进制数。对于负整数，我们通常使用[补码运算](https://baike.baidu.com/item/%E8%A1%A5%E7%A0%81/6854613?fr=aladdin)方法。

注意:

1. 十六进制中所有字母(a-f)都必须是小写。
2. 十六进制字符串中不能包含多余的前导零。如果要转化的数为 0，那么以单个字符 '0' 来表示；对于其他情况，十六进制字符串中的第一个字符将不会是 0 字符。 
3. 给定的数确保在 32 位有符号整数范围内。
4. 不能使用任何由库提供的将数字直接转换或格式化为十六进制的方法。



## 解题思路

- 这一题是水题，将十进制数转换成十六进制的数。需要额外注意 0 和负数的情况。

