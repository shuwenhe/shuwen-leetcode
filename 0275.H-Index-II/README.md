# [275. H-Index II](https://leetcode.com/problems/h-index-ii/)

## 题目

Given an array of citations **sorted in ascending order** (each citation is a non-negative integer) of a researcher, write a function to compute the researcher's h-index.

According to the [definition of h-index on Wikipedia](https://en.wikipedia.org/wiki/H-index): "A scientist has index h if h of his/her N papers have **at least** h citations each, and the other N − h papers have **no more than** h citations each."

**Example:**

    Input: citations = [0,1,3,5,6]
    Output: 3 
    Explanation: [0,1,3,5,6] means the researcher has 5 papers in total and each of them had 
                 received 0, 1, 3, 5, 6 citations respectively. 
                 Since the researcher has 3 papers with at least 3 citations each and the remaining 
                 two with no more than 3 citations each, her h-index is 3.

**Note:**

If there are several possible values for *h*, the maximum one is taken as the h-index.

**Follow up:**

- This is a follow up problem to [H-Index](https://leetcode.com/problems/h-index/description/), where `citations` is now guaranteed to be sorted in ascending order.
- Could you solve it in logarithmic time complexity?



## 题目大意


给定一位研究者论文被引用次数的数组（被引用次数是非负整数），数组已经按照升序排列。编写一个方法，计算出研究者的 h 指数。

h 指数的定义: “h 代表“高引用次数”（high citations），一名科研人员的 h 指数是指他（她）的 （N 篇论文中）至多有 h 篇论文分别被引用了至少 h 次。（其余的 N - h 篇论文每篇被引用次数不多于 h 次。）"

说明:

- 如果 h 有多有种可能的值 ，h 指数是其中最大的那个。

进阶：

- 这是 H 指数 的延伸题目，本题中的 citations 数组是保证有序的。
你可以优化你的算法到对数时间复杂度吗？


## 解题思路

- 给出一个数组，代表该作者论文被引用次数，要求这个作者的 h 指数。h 指数定义："高引用次数”（`high citations`），一名科研人员的 h 指数是指他（她）的 （N 篇论文中）至多有 h 篇论文分别被引用了至少 h 次。（其余的 N - h 篇论文每篇被引用次数不多于 h 次。）
- 这一题要找出 h 指数，即要找到一个边界，在这个边界上为最多的 h 指数。可以用二分搜索来解决这道题。当 `len(citations)-mid > citations[mid]` 时，说明 h 指数的边界一定在右边，因为最多 `len(citations)-mid` 篇数比引用数 `citations[mid]` 还要大。否则 h 指数的边界在左边界，缩小边界以后继续二分。找到边界以后，最终求的是 h 指数，用 `len(citations) - low` 即是结果。
