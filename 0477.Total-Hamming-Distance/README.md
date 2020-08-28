# [477. Total Hamming Distance](https://leetcode.com/problems/total-hamming-distance/)


## 题目

The [Hamming distance](https://en.wikipedia.org/wiki/Hamming_distance) between two integers is the number of positions at which the corresponding bits are different.

Now your job is to find the total Hamming distance between all pairs of the given numbers.

**Example:**

    Input: 4, 14, 2
    
    Output: 6
    
    Explanation: In binary representation, the 4 is 0100, 14 is 1110, and 2 is 0010 (just
    showing the four bits relevant in this case). So the answer will be:
    HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 + 2 + 2 = 6.

**Note:**

1. Elements of the given array are in the range of `0` to `10^9`
2. Length of the array will not exceed `10^4`.


## 题目大意

两个整数的[汉明距离](https://baike.baidu.com/item/%E6%B1%89%E6%98%8E%E8%B7%9D%E7%A6%BB/475174?fr=aladdin)指的是这两个数字的二进制数对应位不同的数量。计算一个数组中，任意两个数之间汉明距离的总和。


## 解题思路

- 计算一个数组内两两元素的海明距离总和。海明距离的定义是两个数二进制位不同的总个数。那么可以把数组中的每个元素 32 位的二进制位依次扫一遍，当扫到某一位上的时候，有 k 个元素在这个位上的值是 1，n - k 个元素在这个位上的值是 0，那么在这一位上所有两两元素的海明距离是 k*(n-k) ，当把 32 位全部都扫完以后，累加出来的海明距离就是所有两两元素的海明距离。

