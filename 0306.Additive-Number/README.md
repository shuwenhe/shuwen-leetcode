# [306. Additive Number](https://leetcode.com/problems/additive-number/)


## 题目

Additive number is a string whose digits can form additive sequence.

A valid additive sequence should contain **at least** three numbers. Except for the first two numbers, each subsequent number in the sequence must be the sum of the preceding two.

Given a string containing only digits `'0'-'9'`, write a function to determine if it's an additive number.

**Note:** Numbers in the additive sequence **cannot** have leading zeros, so sequence `1, 2, 03` or `1, 02, 3` is invalid.

**Example 1:**

    Input: "112358"
    Output: true 
    Explanation: The digits can form an additive sequence: 1, 1, 2, 3, 5, 8. 
                 1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8

**Example 2:**

    Input: "199100199"
    Output: true 
    Explanation: The additive sequence is: 1, 99, 100, 199. 
                 1 + 99 = 100, 99 + 100 = 199

**Follow up:**How would you handle overflow for very large input integers?


## 题目大意

累加数是一个字符串，组成它的数字可以形成累加序列。一个有效的累加序列必须至少包含 3 个数。除了最开始的两个数以外，字符串中的其他数都等于它之前两个数相加的和。给定一个只包含数字 '0'-'9' 的字符串，编写一个算法来判断给定输入是否是累加数。说明: 累加序列里的数不会以 0 开头，所以不会出现 1, 2, 03 或者 1, 02, 3 的情况。


## 解题思路

- 在给出的字符串中判断该字符串是否为斐波那契数列形式的字符串。
- 由于每次判断需要累加 2 个数字，所以在 DFS 遍历的过程中需要维护 2 个数的边界，`firstEnd` 和 `secondEnd`，两个数加起来的和数的起始位置是 `secondEnd + 1`。每次在移动 `firstEnd` 和 `secondEnd` 的时候，需要判断 `strings.HasPrefix(num[secondEnd + 1:], strconv.Itoa(x1 + x2))`，即后面的字符串中是否以和为开头。
- 如果第一个数字起始数字出现了 0 ，或者第二个数字起始数字出现了 0，都算非法异常情况，都应该直接返回 false。

