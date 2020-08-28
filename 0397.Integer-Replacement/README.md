# [397. Integer Replacement](https://leetcode.com/problems/integer-replacement/)


## 题目

Given a positive integer n and you can do operations as follow:

1. If  n is even, replace  n with `n/2`.
2. If  n is odd, you can replace n with either `n + 1` or `n - 1`.

What is the minimum number of replacements needed for n to become 1?

**Example 1:**

    Input:
    8
    
    Output:
    3
    
    Explanation:
    8 -> 4 -> 2 -> 1

**Example 2:**

    Input:
    7
    
    Output:
    4
    
    Explanation:
    7 -> 8 -> 4 -> 2 -> 1
    or
    7 -> 6 -> 3 -> 2 -> 1


## 题目大意

给定一个正整数 n，你可以做如下操作：

1. 如果 n 是偶数，则用 n / 2 替换 n。
2. 如果 n 是奇数，则可以用 n + 1 或 n - 1 替换 n。

问 n 变为 1 所需的最小替换次数是多少？


## 解题思路


- 题目给出一个整数 `n`，然后让我们通过变换将它为 1，如果 `n` 是偶数，可以直接变为 `n/2`，如果是奇数，可以先 `n+1` 或 `n-1`，问最终变为 1 的最少步骤。
- 当 n 为奇数的时候，什么时候需要加 1 ，什么时候需要减 1 ，通过观察规律可以发现，除了 3 和 7 以外，所有加 1 就变成 4 的倍数的奇数，都适合先加 1 运算，比如 15:

        15 -> 16 -> 8 -> 4 -> 2 -> 1
        15 -> 14 -> 7 -> 6 -> 3 -> 2 -> 1
        
        111011 -> 111010 -> 11101 -> 11100 -> 1110 -> 111 -> 1000 -> 100 -> 10 -> 1
        111011 -> 111100 -> 11110 -> 1111 -> 10000 -> 1000 -> 100 -> 10 -> 1

- 对于 7 来说，加 1 和减 1 的结果相同，可以不用管，对于 3 来说，减 1 的步骤更少，所以需要先去掉这种特殊情况。
- 最后如何判断某个数字加 1 后是 4 的倍数呢？这里有一个小技巧，由于之前判断了其是奇数了，那么最右边一位肯定是 1，如果其右边第二位也是 1 的话，那么进行加 1 运算，进位后右边肯定会出现两个 0，则一定是 4 的倍数。于是就可以判断出来了。剩下的情况就是偶数的情况，如果之前判定是偶数，那么直接除以 2 (右移一位)即可。

