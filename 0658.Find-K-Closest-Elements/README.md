# [658. Find K Closest Elements](https://leetcode.com/problems/find-k-closest-elements/)


## 题目

Given a sorted array, two integers `k` and `x`, find the `k` closest elements to `x` in the array. The result should also be sorted in ascending order. If there is a tie, the smaller elements are always preferred.

**Example 1:**

    Input: [1,2,3,4,5], k=4, x=3
    Output: [1,2,3,4]

**Example 2:**

    Input: [1,2,3,4,5], k=4, x=-1
    Output: [1,2,3,4]

**Note:**

1. The value k is positive and will always be smaller than the length of the sorted array.
2. Length of the given array is positive and will not exceed 10^4
3. Absolute value of elements in the array and x will not exceed 10^4

---

**UPDATE (2017/9/19):**The arr parameter had been changed to an **array of integers** (instead of a list of integers). **Please reload the code definition to get the latest changes**.


## 题目大意


给定一个排序好的数组，两个整数 k 和 x，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。如果有两个数与 x 的差值一样，优先选择数值较小的那个数。


说明:

1. k 的值为正数，且总是小于给定排序数组的长度。
2. 数组不为空，且长度不超过 104
3. 数组里的每个元素与 x 的绝对值不超过 104
 

更新(2017/9/19):
这个参数 arr 已经被改变为一个整数数组（而不是整数列表）。 请重新加载代码定义以获取最新更改。




## 解题思路


- 给出一个数组，要求在数组中找到一个长度为 k 的区间，这个区间内每个元素距离 x 的距离都是整个数组里面最小的。
- 这一题可以用双指针解题，最优解法是二分搜索。由于区间长度固定是 K 个，所以左区间最大只能到 `len(arr) - K` (因为长度为 K 以后，正好右区间就到数组最右边了)，在 `[0,len(arr) - K]` 这个区间中进行二分搜索。如果发现 `a[mid]` 与 `x` 距离比 `a[mid + k]` 与 `x` 的距离要大，说明要找的区间一定在右侧，继续二分，直到最终 `low = high` 的时候退出。逼出的 `low` 值就是最终答案区间的左边界。
