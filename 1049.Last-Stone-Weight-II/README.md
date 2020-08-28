# [1049. Last Stone Weight II](https://leetcode.com/problems/last-stone-weight-ii/)

## 题目

We have a collection of rocks, each rock has a positive integer weight.

Each turn, we choose **any two rocks** and smash them together. Suppose the stones have weights `x` and `y` with `x <= y`. The result of this smash is:

- If `x == y`, both stones are totally destroyed;
- If `x != y`, the stone of weight `x` is totally destroyed, and the stone of weight `y`has new weight `y-x`.

At the end, there is at most 1 stone left. Return the **smallest possible** weight of this stone (the weight is 0 if there are no stones left.)

**Example 1:**

    Input: [2,7,4,1,8,1]
    Output: 1
    Explanation: 
    We can combine 2 and 4 to get 2 so the array converts to [2,7,1,8,1] then,
    we can combine 7 and 8 to get 1 so the array converts to [2,1,1,1] then,
    we can combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
    we can combine 1 and 1 to get 0 so the array converts to [1] then that's the optimal value.

**Note:**

1. `1 <= stones.length <= 30`
2. `1 <= stones[i] <= 100`



## 题目大意

有一堆石头，每块石头的重量都是正整数。每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：

如果 x == y，那么两块石头都会被完全粉碎；
如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
最后，最多只会剩下一块石头。返回此石头最小的可能重量。如果没有石头剩下，就返回 0。

提示：

1. 1 <= stones.length <= 30
2. 1 <= stones[i] <= 1000


## 解题思路


- 给出一个数组，数组里面的元素代表的是石头的重量。现在要求两个石头对碰，如果重量相同，两个石头都消失，如果一个重一个轻，剩下的石头是两者的差值。问经过这样的多次碰撞以后，能剩下的石头的重量最轻是多少？
- 由于两两石头要发生碰撞，所以可以将整个数组可以分为两部分，如果这两部分的石头重量总和相差不大，那么经过若干次碰撞以后，剩下的石头重量一定是最小的。现在就需要找到这样两堆总重量差不多的两堆石头。这个问题就可以转化为 01 背包问题。从数组中找到 `sum/2` 重量的石头集合，如果一半能尽量达到 `sum/2`，那么另外一半和 `sum/2` 的差是最小的，最好的情况就是两堆石头的重量都是 `sum/2`，那么两两石头对碰以后最后都能消失。01 背包的经典模板可以参考第 416 题。
