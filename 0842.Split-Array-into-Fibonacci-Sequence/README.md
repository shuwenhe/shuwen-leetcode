# [842. Split Array into Fibonacci Sequence](https://leetcode.com/problems/split-array-into-fibonacci-sequence/)


## 题目

Given a string `S` of digits, such as `S = "123456579"`, we can split it into a *Fibonacci-like sequence* `[123, 456, 579].`

Formally, a Fibonacci-like sequence is a list `F` of non-negative integers such that:

- `0 <= F[i] <= 2^31 - 1`, (that is, each integer fits a 32-bit signed integer type);
- `F.length >= 3`;
- and `F[i] + F[i+1] = F[i+2]` for all `0 <= i < F.length - 2`.

Also, note that when splitting the string into pieces, each piece must not have extra leading zeroes, except if the piece is the number 0 itself.

Return any Fibonacci-like sequence split from `S`, or return `[]` if it cannot be done.

**Example 1:**

    Input: "123456579"
    Output: [123,456,579]

**Example 2:**

    Input: "11235813"
    Output: [1,1,2,3,5,8,13]

**Example 3:**

    Input: "112358130"
    Output: []
    Explanation: The task is impossible.

**Example 4:**

    Input: "0123"
    Output: []
    Explanation: Leading zeroes are not allowed, so "01", "2", "3" is not valid.

**Example 5:**

    Input: "1101111"
    Output: [110, 1, 111]
    Explanation: The output [11, 0, 11, 11] would also be accepted.

**Note:**

1. `1 <= S.length <= 200`
2. `S` contains only digits.


## 题目大意

给定一个数字字符串 S，比如 S = "123456579"，我们可以将它分成斐波那契式的序列 [123, 456, 579]。斐波那契式序列是一个非负整数列表 F，且满足：

- 0 <= F[i] <= 2^31 - 1，（也就是说，每个整数都符合 32 位有符号整数类型）；
- F.length >= 3；
- 对于所有的0 <= i < F.length - 2，都有 F[i] + F[i+1] = F[i+2] 成立。

另外，请注意，将字符串拆分成小块时，每个块的数字一定不要以零开头，除非这个块是数字 0 本身。返回从 S 拆分出来的所有斐波那契式的序列块，如果不能拆分则返回 []。



## 解题思路


- 这一题是第 306 题的加强版。第 306 题要求判断字符串是否满足斐波那契数列形式。这一题要求输出按照斐波那契数列形式分割之后的数字数组。
- 这一题思路和第 306 题基本一致，需要注意的是题目中的一个限制条件，`0 <= F[i] <= 2^31 - 1`，注意这个条件，笔者开始没注意，后面输出解就出现错误了，可以看笔者的测试文件用例的最后两组数据，这两组都是可以分解成斐波那契数列的，但是由于分割以后的数字都大于了 `2^31 - 1`，所以这些解都不能要！
- 这一题也要特别注意剪枝条件，没有剪枝条件，时间复杂度特别高，加上合理的剪枝条件以后，0ms 通过。

