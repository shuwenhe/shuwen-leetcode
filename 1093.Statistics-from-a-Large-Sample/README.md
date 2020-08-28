# [1093. Statistics from a Large Sample](https://leetcode.com/problems/statistics-from-a-large-sample/)


## 题目

We sampled integers between `0` and `255`, and stored the results in an array `count`: `count[k]` is the number of integers we sampled equal to `k`.

Return the minimum, maximum, mean, median, and mode of the sample respectively, as an array of **floating point numbers**. The mode is guaranteed to be unique.

*(Recall that the median of a sample is:*

- *The middle element, if the elements of the sample were sorted and the number of elements is odd;*
- *The average of the middle two elements, if the elements of the sample were sorted and the number of elements is even.)*

**Example 1:**

    Input: count = [0,1,3,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
    Output: [1.00000,3.00000,2.37500,2.50000,3.00000]

**Example 2:**

    Input: count = [0,4,3,2,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
    Output: [1.00000,4.00000,2.18182,2.00000,1.00000]

**Constraints:**

1. `count.length == 256`
2. `1 <= sum(count) <= 10^9`
3. The mode of the sample that count represents is unique.
4. Answers within `10^-5` of the true value will be accepted as correct.


## 题目大意

我们对 0 到 255 之间的整数进行采样，并将结果存储在数组 count 中：count[k] 就是整数 k 的采样个数。

我们以 浮点数 数组的形式，分别返回样本的最小值、最大值、平均值、中位数和众数。其中，众数是保证唯一的。我们先来回顾一下中位数的知识：

- 如果样本中的元素有序，并且元素数量为奇数时，中位数为最中间的那个元素；
- 如果样本中的元素有序，并且元素数量为偶数时，中位数为中间的两个元素的平均值。



## 解题思路


- 这个问题的关键需要理解题目的意思，什么是采样？`count[k]` 就是整数 `k` 的采样个数。
- 题目要求返回样本的最小值、最大值、平均值、中位数和众数。最大值和最小值就很好处理，只需要遍历 count 判断最小的非 0 的 index 就是最小值，最大的非 0 的 index 就是最大值。平均值也非常好处理，对于所有非 0 的 count，我们通过累加 count[k] * index 得到所有数的和，然后除上所有非 0 的 count 的和。![](https://latex.codecogs.com/svg.latex?\sum_{n=0}^{256}count[n](while\%20\%20count[n]!=0))

- 众数也非常容易，只需统计 count 值最大时的 index 即可。
- 中位数的处理相对麻烦一些，因为需要分非 0 的 count 之和是奇数还是偶数两种情况。先假设非 0 的 count 和为 cnt，那么如果 cnt 是奇数的话，只需要找到 cnt/2 的位置即可，通过不断累加 count 的值，直到累加和超过 ≥ cnt/2。如果 cnt 是偶数的话，需要找到 cnt/2 + 1 和 cnt/2 的位置，找法和奇数情况相同，不过需要找两次(可以放到一个循环中做两次判断)。