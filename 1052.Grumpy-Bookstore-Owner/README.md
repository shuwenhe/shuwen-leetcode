# [1052. Grumpy Bookstore Owner](https://leetcode.com/problems/grumpy-bookstore-owner/)


## 题目

Today, the bookstore owner has a store open for `customers.length`minutes. Every minute, some number of customers (`customers[i]`) enter the store, and all those customers leave after the end of that minute.

On some minutes, the bookstore owner is grumpy. If the bookstore owner is grumpy on the i-th minute, `grumpy[i] = 1`, otherwise `grumpy[i] = 0`. When the bookstore owner is grumpy, the customers of that minute are not satisfied, otherwise they are satisfied.

The bookstore owner knows a secret technique to keep themselves not grumpy for `X` minutes straight, but can only use it once.

Return the maximum number of customers that can be satisfied throughout the day.

**Example 1:**

    Input: customers = [1,0,1,2,1,1,7,5], grumpy = [0,1,0,1,0,1,0,1], X = 3
    Output: 16
    Explanation: The bookstore owner keeps themselves not grumpy for the last 3 minutes. 
    The maximum number of customers that can be satisfied = 1 + 1 + 1 + 1 + 7 + 5 = 16.

**Note:**

- `1 <= X <= customers.length == grumpy.length <= 20000`
- `0 <= customers[i] <= 1000`
- `0 <= grumpy[i] <= 1`


## 题目大意

今天，书店老板有一家店打算试营业 customers.length 分钟。每分钟都有一些顾客（customers[i]）会进入书店，所有这些顾客都会在那一分钟结束后离开。在某些时候，书店老板会生气。 如果书店老板在第 i 分钟生气，那么 grumpy[i] = 1，否则 grumpy[i] = 0。 当书店老板生气时，那一分钟的顾客就会不满意，不生气则他们是满意的。书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 X 分钟不生气，但却只能使用一次。请你返回这一天营业下来，最多有多少客户能够感到满意的数量。

提示：

1. 1 <= X <= customers.length == grumpy.length <= 20000
2. 0 <= customers[i] <= 1000
3. 0 <= grumpy[i] <= 1



## 解题思路


- 给出一个顾客入店时间表和书店老板发脾气的时间表。两个数组的时间是一一对应的，即相同下标对应的相同的时间。书店老板可以控制自己在 X 分钟内不发火，但是只能控制一次。问有多少顾客能在书店老板不发火的时候在书店里看书。抽象一下，给出一个价值数组和一个装着 0 和 1 的数组，当价值数组的下标对应另外一个数组相同下标的值是 0 的时候，那么这个价值可以累加，当对应是 1 的时候，就不能加上这个价值。现在可以让装着 0 和 1 的数组中连续 X 个数都变成 0，问最终价值最大是多少？
- 这道题是典型的滑动窗口的题目。最暴力的解法是滑动窗口右边界，当与左边界的距离等于 X 的时候，计算此刻对应的数组的总价值。当整个宽度为 X 的窗口滑过整个数组以后，输出维护的最大值即可。这个方法耗时比较长。因为每次计算数组总价值的时候都要遍历整个数组。这里是可以优化的地方。
- 每次计算数组总价值的时候，其实目的是为了找到宽度为 X 的窗口对应里面为 1 的数累加和最大，因为如果把这个窗口里面的 1 都变成 0 以后，那么对最终价值的影响也最大。所以用一个变量 `customer0` 专门记录脾气数组中为 0 的对应的价值，累加起来。因为不管怎么改变，为 0 的永远为 0，唯一变化的是 1 变成 0 。用 `customer1` 专门记录脾气数组中为 1 的对应的价值。在窗口滑动过程中找到 `customer1` 的最大值。最终要求的最大值就是 `customer0 + maxCustomer1`。
