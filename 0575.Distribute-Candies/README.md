# [575. Distribute Candies](https://leetcode.com/problems/distribute-candies/)


## 题目

Given an integer array with **even** length, where different numbers in this array represent different **kinds** of candies. Each number means one candy of the corresponding kind. You need to distribute these candies **equally** in number to brother and sister. Return the maximum number of **kinds** of candies the sister could gain.

**Example 1:**

    Input: candies = [1,1,2,2,3,3]
    Output: 3
    Explanation:
    There are three different kinds of candies (1, 2 and 3), and two candies for each kind.
    Optimal distribution: The sister has candies [1,2,3] and the brother has candies [1,2,3], too. 
    The sister has three different kinds of candies.

**Example 2:**

    Input: candies = [1,1,2,3]
    Output: 2
    Explanation: For example, the sister has candies [2,3] and the brother has candies [1,1]. 
    The sister has two different kinds of candies, the brother has only one kind of candies.

**Note:**

1. The length of the given array is in range [2, 10,000], and will be even.
2. The number in given array is in range [-100,000, 100,000].


## 题目大意

给定一个偶数长度的数组，其中不同的数字代表着不同种类的糖果，每一个数字代表一个糖果。你需要把这些糖果平均分给一个弟弟和一个妹妹。返回妹妹可以获得的最大糖果的种类数。


## 解题思路


- 给出一个糖果数组，里面每个元素代表糖果的种类，相同数字代表相同种类。把这些糖果分给兄弟姐妹，问姐妹最多可以分到多少种糖果。这一题比较简单，用 map 统计每个糖果的出现频次，如果总数比 `n/2` 小，那么就返回 `len(map)`，否则返回 `n/2` (即一半都分给姐妹)。
