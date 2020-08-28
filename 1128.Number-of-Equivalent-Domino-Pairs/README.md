# [1128. Number of Equivalent Domino Pairs](https://leetcode.com/problems/number-of-equivalent-domino-pairs/)


## 题目

Given a list of `dominoes`, `dominoes[i] = [a, b]` is *equivalent* to `dominoes[j] = [c, d]` if and only if either (`a==c` and `b==d`), or (`a==d` and `b==c`) - that is, one domino can be rotated to be equal to another domino.

Return the number of pairs `(i, j)` for which `0 <= i < j < dominoes.length`, and `dominoes[i]` is equivalent to `dominoes[j]`.

**Example 1:**

    Input: dominoes = [[1,2],[2,1],[3,4],[5,6]]
    Output: 1

**Constraints:**

- `1 <= dominoes.length <= 40000`
- `1 <= dominoes[i][j] <= 9`


## 题目大意

给你一个由一些多米诺骨牌组成的列表 dominoes。如果其中某一张多米诺骨牌可以通过旋转 0 度或 180 度得到另一张多米诺骨牌，我们就认为这两张牌是等价的。形式上，dominoes[i] = [a, b] 和 dominoes[j] = [c, d] 等价的前提是 a==c 且 b==d，或是 a==d 且 b==c。

在 0 <= i < j < dominoes.length 的前提下，找出满足 dominoes[i] 和 dominoes[j] 等价的骨牌对 (i, j) 的数量。

提示：

- 1 <= dominoes.length <= 40000
- 1 <= dominoes[i][j] <= 9



## 解题思路

- 给出一组多米诺骨牌，求出这组牌中相同牌的个数。牌相同的定义是：牌的 2 个数字相同(正序或者逆序相同都算相同)
- 简单题。由于牌是 2 个数，所以将牌的 2 个数 hash 成一个 2 位数，比较大小即可，正序和逆序都 hash 成 2 位数，然后在桶中比较是否已经存在，如果不存在，跳过，如果存在，计数。
