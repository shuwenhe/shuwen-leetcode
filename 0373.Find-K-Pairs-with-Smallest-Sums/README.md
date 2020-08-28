# [373. Find K Pairs with Smallest Sums](https://leetcode.com/problems/find-k-pairs-with-smallest-sums/)


## 题目

You are given two integer arrays **nums1** and **nums2** sorted in ascending order and an integer **k**.

Define a pair **(u,v)** which consists of one element from the first array and one element from the second array.

Find the k pairs **(u1,v1),(u2,v2) ...(uk,vk)** with the smallest sums.

**Example 1:**

    Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
    Output: [[1,2],[1,4],[1,6]] 
    Explanation: The first 3 pairs are returned from the sequence: 
                 [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]

**Example 2:**

    Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
    Output: [1,1],[1,1]
    Explanation: The first 2 pairs are returned from the sequence: 
                 [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]

**Example 3:**

    Input: nums1 = [1,2], nums2 = [3], k = 3
    Output: [1,3],[2,3]
    Explanation: All possible pairs are returned from the sequence: [1,3],[2,3]


## 题目大意


给定两个以升序排列的整形数组 nums1 和 nums2, 以及一个整数 k。

定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2。

找到和最小的 k 对数字 (u1,v1), (u2,v2) ... (uk,vk)。



## 解题思路


- 给出 2 个数组，和数字 k，要求找到 k 个数值对，数值对两个数的和最小。
- 这一题咋一看可以用二分搜索，两个数组两个组合有 `m * n` 个数值对。然后找到最小的和，最大的和，在这个范围内进行二分搜索，每分出一个 `mid`，再去找比 `mid` 小的数值对有多少个，如果个数小于 `k` 个，那么在右区间上继续二分，如果个数大于 `k` 个，那么在左区间上继续二分。到目前为止，这个思路看似可行。但是每次搜索的数值对是无序的。这会导致最终出现错误的结果。例如 `mid = 10` 的时候，小于 10 的和有 22 个，而 `k = 25` 。这说明 `mid` 偏小，`mid` 增大，`mid = 11` 的时候，小于 11 的和有 30 个，而 `k = 25` 。这时候应该从这 30 个和中取前 25 个。但是我们遍历数值对的时候，和并不是从小到大排序的。这时候还需要额外对这 30 个候选值进行排序。这样时间复杂度又增大了。
- 可以先用暴力解法解答。将所有的和都遍历出来，排序以后，取前 k 个。这个暴力方法可以 AC。
- 本题最优解应该是优先队列。维护一个最小堆。把数值对的和放在这个最小堆中，不断的 pop 出 k 个最小值到数组中，即为答案。
- 在已排序的矩阵中寻找最 K 小的元素这一系列的题目有：第 373 题，第 378 题，第 668 题，第 719 题，第 786 题。
