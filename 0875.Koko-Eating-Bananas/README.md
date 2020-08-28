# [875. Koko Eating Bananas](https://leetcode.com/problems/koko-eating-bananas/)


## 题目

Koko loves to eat bananas. There are `N` piles of bananas, the `i`-th pile has `piles[i]` bananas. The guards have gone and will come back in `H` hours.

Koko can decide her bananas-per-hour eating speed of `K`. Each hour, she chooses some pile of bananas, and eats K bananas from that pile. If the pile has less than `K` bananas, she eats all of them instead, and won't eat any more bananas during this hour.

Koko likes to eat slowly, but still wants to finish eating all the bananas before the guards come back.

Return the minimum integer `K` such that she can eat all the bananas within `H` hours.

**Example 1:**

    Input: piles = [3,6,7,11], H = 8
    Output: 4

**Example 2:**

    Input: piles = [30,11,23,4,20], H = 5
    Output: 30

**Example 3:**

    Input: piles = [30,11,23,4,20], H = 6
    Output: 23

**Note:**

- `1 <= piles.length <= 10^4`
- `piles.length <= H <= 10^9`
- `1 <= piles[i] <= 10^9`


## 题目大意


珂珂喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。

珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。  

珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。

提示：

- 1 <= piles.length <= 10^4
- piles.length <= H <= 10^9
- 1 <= piles[i] <= 10^9



## 解题思路


- 给出一个数组，数组里面每个元素代表的是每个香蕉🍌串上香蕉的个数。koko 以 `k 个香蕉/小时`的速度吃这些香蕉。守卫会在 `H 小时`以后回来。问 k 至少为多少，能在守卫回来之前吃完所有的香蕉。当香蕉的个数小于 k 的时候，这个小时只能吃完这些香蕉，不能再吃其他串上的香蕉了。
- 这一题可以用二分搜索来解答。在 `[0 , max(piles)]` 的范围内搜索，二分的过程都是常规思路。判断是否左右边界如果划分的时候需要注意题目中给的限定条件。当香蕉个数小于 k 的时候，那个小时不能再吃其他香蕉了。
