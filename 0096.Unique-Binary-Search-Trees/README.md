# [96. Unique Binary Search Trees](https://leetcode.com/problems/unique-binary-search-trees/)


## 题目

Given *n*, how many structurally unique **BST's** (binary search trees) that store values 1 ... *n*?

**Example:**

    Input: 3
    Output: 5
    Explanation:
    Given n = 3, there are a total of 5 unique BST's:
    
       1         3     3      2      1
        \       /     /      / \      \
         3     2     1      1   3      2
        /     /       \                 \
       2     1         2                 3

## 题目大意

给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？


## 解题思路

- 给出 n，要求利用 1-n 这些数字组成二叉排序树，有多少种不同的树的形态，输出这个个数。
- 这题的解题思路是 DP。`dp[n]` 代表 1-n 个数能组成多少个不同的二叉排序树，`F(i,n)` 代表以 `i` 为根节点，1-n 个数组成的二叉排序树的不同的个数。由于题意，我们可以得到这个等式：`dp[n] = F(1,n) + F(2,n) + F(3,n) + …… + F(n,n)` 。初始值 `dp[0] = 1`，`dp[1] = 1`。分析 `dp` 和 `F(i,n)` 的关系又可以得到下面这个等式 `F(i,n) = dp[i-1] * dp[n-i]` 。举例，[1,2,3,4,…, i ,…,n-1,n]，以 `i` 为 根节点，那么左半边 [1,2,3,……,i-1] 和 右半边 [i+1,i+2,……,n-1,n] 分别能组成二叉排序树的不同个数`相乘`，即为以 `i` 为根节点，1-n 个数组成的二叉排序树的不同的个数，也即 `F(i,n)`。

> 注意，由于二叉排序树本身的性质，右边的子树一定比左边的子树，值都要大。所以这里只需要根节点把树分成左右，不需要再关心左右两边数字的大小，只需要关心数字的个数。

- 所以状态转移方程是 `dp[i] = dp[0] * dp[n-1] + dp[1] * dp[n-2] + …… + dp[n-1] * dp[0]`，最终要求的结果是 `dp[n]` 。
