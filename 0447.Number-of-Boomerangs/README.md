# [447. Number of Boomerangs](https://leetcode.com/problems/number-of-boomerangs/)

## 题目

Given n points in the plane that are all pairwise distinct, a "boomerang" is a tuple of points (i, j, k) such that the distance between i and j equals the distance between i and k (the order of the tuple matters).

Find the number of boomerangs. You may assume that n will be at most 500 and coordinates of points are all in the range [-10000, 10000] (inclusive).



Example 1:

```c
Input:
[[0,0],[1,0],[2,0]]

Output:
2

Explanation:
The two boomerangs are [[1,0],[0,0],[2,0]] and [[1,0],[2,0],[0,0]]
```


## 题目大意

在一个 Point 的数组中求出 (i,j,k) 三元组，要求 j 和 i 的距离等于 k 和 i 的距离。这样的三元组有多少种？注意 (i,j,k) 和 (j,i,k) 是不同的解，即元素的顺序是有关系的。

## 解题思路

这道题考察的是哈希表的问题。

首先依次求出两两点之间的距离，然后把这些距离记录在 map 中，key 是距离，value 是这个距离出现了多少次。求距离一般都需要开根号，但是 key 如果为浮点数就会有一些误差，所以计算距离的时候最后一步不需要开根号，保留平方差即可。

最后求结果的时候，遍历 map，把里面距离大于 2 的 key 都拿出来，value 对应的是个数，在这些个数里面任取 2 个点就是解，所以利用排列组合，C n 2 就可以得到这个距离的结果，最后把这些排列组合的结果累积起来即可。
