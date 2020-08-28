# [399. Evaluate Division](https://leetcode.com/problems/evaluate-division/)


## 题目

Equations are given in the format `A / B = k`, where `A` and `B` are variables represented as strings, and `k` is a real number (floating point number). Given some queries, return the answers. If the answer does not exist, return `-1.0`.

**Example:**

Given `a / b = 2.0, b / c = 3.0.`queries are: `a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ? .`return `[6.0, 0.5, -1.0, 1.0, -1.0 ].`

The input is: `vector<pair<string, string>> equations, vector<double>& values, vector<pair<string, string>> queries` , where `equations.size() == values.size()`, and the values are positive. This represents the equations. Return `vector<double>`.

According to the example above:

    equations = [ ["a", "b"], ["b", "c"] ],
    values = [2.0, 3.0],
    queries = [ ["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"] ].

The input is always valid. You may assume that evaluating the queries will result in no division by zero and there is no contradiction.


## 题目大意

给出方程式 A / B = k, 其中 A 和 B 均为代表字符串的变量， k 是一个浮点型数字。根据已知方程式求解问题，并返回计算结果。如果结果不存在，则返回 -1.0。

示例 :
给定 a / b = 2.0, b / c = 3.0
问题: a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ? 
返回 [6.0, 0.5, -1.0, 1.0, -1.0 ]

输入为: vector<pair<string, string>> equations, vector<double>& values, vector<pair<string, string>> queries(方程式，方程式结果，问题方程式)， 其中 equations.size() == values.size()，即方程式的长度与方程式结果长度相等（程式与结果一一对应），并且结果值均为正数。以上为方程式的描述。 返回vector<double>类型。

假设输入总是有效的。你可以假设除法运算中不会出现除数为0的情况，且不存在任何矛盾的结果。


## 解题思路


- 给出一些字母变量的倍数关系，问给出任意两个字母的倍数是多少。
- 这一题可以用 DFS 或者并查集来解题。先来看看 DFS 的做法。先建图。每个字母或者字母组合可以看做成一个节点，给出的 `equations` 关系可以看成两个节点之间的有向边。每条有向边都有权值。那么问题可以转换成是否存在一条从起点节点到终点节点的路径，如果存在，输出这条路径上所有有向边权值的累乘结果。如果不存在这条路径，就返回 -1 。如果给的起点和终点不在给出的节点集里面，也输出 -1 。
- 再来看看并查集的做法。先将每两个有倍数关系的节点做并查集 `union()` 操作。例如 A/B = 2，那么把 B 作为 `parent` 节点，`parents[A] = {B，2}`，`parents[B] = {B，1}`，B 指向自己是 1 。还有一个关系是 `B/C=3`，由于 B 已经在并查集中了，所以这个时候需要把这个关系反过来，处理成 `C/B = 1/3` ，即 `parents[C] = {B，1/3}`。这样把所有有关系的字母都 `union()` 起来。如何求任意两个字母的倍数关系呢？例如 `A/C = ?` 在并查集中查找，可以找到 `parents[C] == parents[A] == B`，那么就用 `parents[A]/parents[C] = 2/(1/3) = 6`。为什么可以这样做呢？因为 `A/B = 2`，`C/B = 1/3`，那么 `A/C = (A/B)/(C/B)` 即 `parents[A]/parents[C] = 2/(1/3) = 6`。
