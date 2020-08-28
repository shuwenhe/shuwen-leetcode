# [1122. Relative Sort Array](https://leetcode.com/problems/relative-sort-array/)


## 题目

Given two arrays `arr1` and `arr2`, the elements of `arr2` are distinct, and all elements in `arr2` are also in `arr1`.

Sort the elements of `arr1` such that the relative ordering of items in `arr1` are the same as in `arr2`. Elements that don't appear in `arr2` should be placed at the end of `arr1` in **ascending** order.

**Example 1:**

    Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
    Output: [2,2,2,1,4,3,3,9,6,7,19]

**Constraints:**

- `arr1.length, arr2.length <= 1000`
- `0 <= arr1[i], arr2[i] <= 1000`
- Each `arr2[i]` is distinct.
- Each `arr2[i]` is in `arr1`.


## 题目大意


给你两个数组，arr1 和 arr2，

- arr2 中的元素各不相同
- arr2 中的每个元素都出现在 arr1 中

对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。

提示：

- arr1.length, arr2.length <= 1000
- 0 <= arr1[i], arr2[i] <= 1000
- arr2 中的元素 arr2[i] 各不相同
- arr2 中的每个元素 arr2[i] 都出现在 arr1 中



## 解题思路

- 给出 2 个数组 A 和 B，A 中包含 B 中的所有元素。要求 A 按照 B 的元素顺序排序，B 中没有的元素再接着排在后面，从小到大排序。
- 这一题有多种解法。一种暴力的解法就是依照题意，先把 A 中的元素都统计频次放在 map 中，然后 按照 B 的顺序输出，接着再把剩下的元素排序接在后面。还有一种桶排序的思想，由于题目限定了 A 的大小是 1000，这个数量级很小，所以可以用 1001 个桶装所有的数，把数都放在桶里，这样默认就已经排好序了。接下来的做法和前面暴力解法差不多，按照频次输出。B 中以外的元素就按照桶的顺序依次输出即可。
