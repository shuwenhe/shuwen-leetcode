# [969. Pancake Sorting](https://leetcode.com/problems/pancake-sorting/)

## 题目

Given an array A, we can perform a pancake flip: We choose some positive integer k <= A.length, then reverse the order of the first k elements of A.  We want to perform zero or more pancake flips (doing them one after another in succession) to sort the array A.

Return the k-values corresponding to a sequence of pancake flips that sort A.  Any valid answer that sorts the array within 10 * A.length flips will be judged as correct.

Example 1:

```c
Input: [3,2,4,1]
Output: [4,2,4,3]
Explanation: 
We perform 4 pancake flips, with k values 4, 2, 4, and 3.
Starting state: A = [3, 2, 4, 1]
After 1st flip (k=4): A = [1, 4, 2, 3]
After 2nd flip (k=2): A = [4, 1, 2, 3]
After 3rd flip (k=4): A = [3, 2, 1, 4]
After 4th flip (k=3): A = [1, 2, 3, 4], which is sorted. 
```

Example 2:

```c
Input: [1,2,3]
Output: []
Explanation: The input is already sorted, so there is no need to flip anything.
Note that other answers, such as [3, 3], would also be accepted.
```

Note:

- 1 <= A.length <= 100
- A[i] is a permutation of [1, 2, ..., A.length]

## 题目大意

给定一个数组，要求输出“煎饼排序”的步骤，使得最终数组是从小到大有序的。“煎饼排序”，每次排序都反转前 n 个数，n 小于数组的长度。

## 解题思路

这道题的思路是，每次找到当前数组中无序段中最大的值，（初始的时候，整个数组相当于都是无序段），将最大值的下标 i 进行“煎饼排序”，前 i 个元素都反转一遍。这样最大值就到了第一个位置了。然后紧接着再进行一次数组总长度 n 的“煎饼排序”，目的是使最大值到数组最后一位，这样它的位置就归位了。那么数组的无序段为 n-1 。然后用这个方法不断的循环，直至数组中每个元素都到了排序后最终的位置下标上了。最终数组就有序了。

这道题有一个特殊点在于，数组里面的元素都是自然整数，那么最终数组排序完成以后，数组的长度就是最大值。所以找最大值也不需要遍历一次数组了，直接取出长度就是最大值。
