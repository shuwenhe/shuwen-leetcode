# [1184. Distance Between Bus Stops](https://leetcode.com/problems/distance-between-bus-stops/)

## 题目

A bus has `n` stops numbered from `0` to `n - 1` that form a circle. We know the distance between all pairs of neighboring stops where `distance[i]` is the distance between the stops number `i` and `(i + 1) % n`.

The bus goes along both directions i.e. clockwise and counterclockwise.

Return the shortest distance between the given `start` and `destination` stops.

**Example 1:**

![](https://assets.leetcode.com/uploads/2019/09/03/untitled-diagram-1.jpg)

    Input: distance = [1,2,3,4], start = 0, destination = 1
    Output: 1
    Explanation: Distance between 0 and 1 is 1 or 9, minimum is 1.

**Example 2:**

![](https://assets.leetcode.com/uploads/2019/09/03/untitled-diagram-1-1.jpg)

    Input: distance = [1,2,3,4], start = 0, destination = 2
    Output: 3
    Explanation: Distance between 0 and 2 is 3 or 7, minimum is 3.

**Example 3:**

![](https://assets.leetcode.com/uploads/2019/09/03/untitled-diagram-1-2.jpg)

    Input: distance = [1,2,3,4], start = 0, destination = 3
    Output: 4
    Explanation: Distance between 0 and 3 is 6 or 4, minimum is 4.

**Constraints:**

- `1 <= n <= 10^4`
- `distance.length == n`
- `0 <= start, destination < n`
- `0 <= distance[i] <= 10^4`


## 题目大意

环形公交路线上有 n 个站，按次序从 0 到 n - 1 进行编号。我们已知每一对相邻公交站之间的距离，distance[i] 表示编号为 i 的车站和编号为 (i + 1) % n 的车站之间的距离。环线上的公交车都可以按顺时针和逆时针的方向行驶。返回乘客从出发点 start 到目的地 destination 之间的最短距离。

提示：

- 1 <= n <= 10^4
- distance.length == n
- 0 <= start, destination < n
- 0 <= distance[i] <= 10^4


## 解题思路


- 给出一个数组，代表的是公交车站每站直接的距离。距离是按照数组下标的顺序给出的，公交车可以按照顺时针行驶，也可以按照逆时针行驶。问行驶的最短距离是多少。
- 按照题意，分别算出顺时针和逆时针的行驶距离，比较两者距离，取出小值就是结果。
