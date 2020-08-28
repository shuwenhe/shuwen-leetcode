# [357. Count Numbers with Unique Digits](https://leetcode.com/problems/count-numbers-with-unique-digits/)


## 题目

Given a **non-negative** integer n, count all numbers with unique digits, x, where 0 ≤ x < 10n.

**Example:**

    Input: 2
    Output: 91 
    Explanation: The answer should be the total numbers in the range of 0 ≤ x < 100, 
                 excluding 11,22,33,44,55,66,77,88,99


## 题目大意

给定一个非负整数 n，计算各位数字都不同的数字 x 的个数，其中 0 ≤ x < 10^n 。




## 解题思路

- 输出 n 位数中不出现重复数字的数字的个数
- 这道题摸清楚规律以后，可以直接写出最终所有答案，答案只有 11 个。
- 考虑不重复数字是如生成的。如果只是一位数，不存在重复的数字，结果是 10 。如果是二位数，第一位一定不能取 0，那么第一位有 1-9，9种取法，第二位为了和第一位不重复，只能有 0-9，10种取法中减去第一位取的数字，那么也是 9 种取法。以此类推，如果是三位数，第三位是 8 种取法；四位数，第四位是 7 种取法；五位数，第五位是 6 种取法；六位数，第六位是 5 种取法；七位数，第七位是 4 种取法；八位数，第八位是 3 种取法；九位数，第九位是 2 种取法；十位数，第十位是 1 种取法；十一位数，第十一位是 0 种取法；十二位数，第十二位是 0 种取法；那么第 11 位数以后，每个数都是重复数字的数字。知道这个规律以后，可以累积上面的结果，把结果直接存在数组里面，暴力打表即可。O(1) 的时间复杂度。

