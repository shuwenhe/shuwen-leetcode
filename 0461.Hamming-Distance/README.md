# [461. Hamming Distance](https://leetcode.com/problems/hamming-distance/)

## 题目

The [Hamming distance](https://en.wikipedia.org/wiki/Hamming_distance) between two integers is the number of positions at which the corresponding bits are different.

Given two integers `x` and `y`, calculate the Hamming distance.

**Note:**0 ≤ `x`, `y` < 231.

**Example:**

    Input: x = 1, y = 4
    
    Output: 2
    
    Explanation:
    1   (0 0 0 1)
    4   (0 1 0 0)
           ↑   ↑
    
    The above arrows point to positions where the corresponding bits are different.


## 题目大意

两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。给出两个整数 x 和 y，计算它们之间的汉明距离。

注意：
0 ≤ x, y < 231.



## 解题思路

- 求 2 个数的海明距离。海明距离的定义是两个数二进制位不同的总个数。这一题利用的位操作的是 X &= (X - 1) 不断的清除最低位的 1 。先将这两个数异或，异或以后清除低位的 1 就是最终答案。
