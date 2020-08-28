# [483. Smallest Good Base](https://leetcode.com/problems/smallest-good-base/)


## 题目

For an integer n, we call k>=2 a **good base** of n, if all digits of n base k are 1.

Now given a string representing n, you should return the smallest good base of n in string format.

**Example 1:**

    Input: "13"
    Output: "3"
    Explanation: 13 base 3 is 111.

**Example 2:**

    Input: "4681"
    Output: "8"
    Explanation: 4681 base 8 is 11111.

**Example 3:**

    Input: "1000000000000000000"
    Output: "999999999999999999"
    Explanation: 1000000000000000000 base 999999999999999999 is 11.

**Note:**

1. The range of n is [3, 10^18].
2. The string representing n is always valid and will not have leading zeros.


## 题目大意


对于给定的整数 n, 如果n的k（k>=2）进制数的所有数位全为1，则称 k（k>=2）是 n 的一个好进制。

以字符串的形式给出 n, 以字符串的形式返回 n 的最小好进制。

提示：

- n 的取值范围是 [3, 10^18]。
- 输入总是有效且没有前导 0。



## 解题思路


- 给出一个数 n，要求找一个进制 k，使得数字 n 在 k 进制下每一位都是 1 。求最小的进制 k。
- 这一题等价于求最小的正整数 k，满足存在一个正整数 m 使得

<p align='center'>
<img src='https://img.halfrost.com/Leetcode/leetcode_483_1.png'>
</p>


- 这一题需要确定 k 和 m 两个数的值。m 和 k 是有关系的，确定了一个值，另外一个值也确定了。由

<p align='center'>
<img src='https://img.halfrost.com/Leetcode/leetcode_483_2.png'>
</p>


可得：

<p align='center'>
<img src='https://img.halfrost.com/Leetcode/leetcode_483_3.png'>
</p>


根据题意，可以知道 k ≥2，m ≥1 ，所以有:

<p align='center'>
<img src='https://img.halfrost.com/Leetcode/leetcode_483_4.png'>
</p>


所以 m 的取值范围确定了。那么外层循环从 1 到 log n 遍历。找到一个最小的 k ，能满足：

可以用二分搜索来逼近找到最小的 k。先找到 k 的取值范围。由 

<p align='center'>
<img src='https://img.halfrost.com/Leetcode/leetcode_483_5.png'>
</p>


可得，

<p align='center'>
<img src='https://img.halfrost.com/Leetcode/leetcode_483_6.png'>
</p>

所以 k 的取值范围是 [2, n*(1/m) ]。再利用二分搜索逼近找到最小的 k 即为答案。
