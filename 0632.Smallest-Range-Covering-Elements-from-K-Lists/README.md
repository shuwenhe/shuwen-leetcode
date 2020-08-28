# [632. Smallest Range Covering Elements from K Lists](https://leetcode.com/problems/smallest-range-covering-elements-from-k-lists/)


## 题目

You have `k` lists of sorted integers in ascending order. Find the **smallest** range that includes at least one number from each of the `k` lists.

We define the range [a,b] is smaller than range [c,d] if `b-a < d-c` or `a < c` if `b-a == d-c`.

**Example 1:**

    Input: [[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
    Output: [20,24]
    Explanation: 
    List 1: [4, 10, 15, 24,26], 24 is in range [20,24].
    List 2: [0, 9, 12, 20], 20 is in range [20,24].
    List 3: [5, 18, 22, 30], 22 is in range [20,24].

**Note:**

1. The given list may contain duplicates, so ascending order means >= here.
2. 1 <= `k` <= 3500
3. -10^5 <= `value of elements` <= 10^5.


## 题目大意

你有 k 个升序排列的整数数组。找到一个最小区间，使得 k 个列表中的每个列表至少有一个数包含在其中。

我们定义如果 b-a < d-c 或者在 b-a == d-c 时 a < c，则区间 [a,b] 比 [c,d] 小。

注意:

- 给定的列表可能包含重复元素，所以在这里升序表示 >= 。
- 1 <= k <= 3500
- -105 <= 元素的值 <= 105
- 对于使用Java的用户，请注意传入类型已修改为List<List<Integer>>。重置代码模板后可以看到这项改动。



## 解题思路


- 给出 K 个数组，要求在这 K 个数组中找到一个区间，至少能包含这 K 个数组中每个数组中的一个元素。
- 这一题是第 76 题的变种版。第 76 题是用滑动窗口来解答的，它要求在母字符串 S 中找到最小的子串能包含 T 串的所有字母。这一题类似的，可以把母字符串看成 K 个数组合并起来的大数组，那么 T 串是由 K 个数组中每个数组中抽一个元素出来组成的。求的区间相同，都是能包含 T 的最小区间。另外一个区别在于，第 76 题里面都是字符串，这一题都是数字，在最终拼接成 T 串的时候需要保证 K 个数组中每个都有一个元素，所以理所当然的想到需要维护每个元素所在数组编号。经过上述的转换，可以把这道题转换成第 76 题的解法了。
- 在具体解题过程中，用 map 来维护窗口内 K 个数组出现的频次。时间复杂度 O(n*log n)，空间复杂度是O(n)。