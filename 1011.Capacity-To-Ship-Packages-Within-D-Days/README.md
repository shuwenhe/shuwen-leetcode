# [1011. Capacity To Ship Packages Within D Days](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/)


## 题目

A conveyor belt has packages that must be shipped from one port to another within `D` days.

The `i`-th package on the conveyor belt has a weight of `weights[i]`. Each day, we load the ship with packages on the conveyor belt (in the order given by `weights`). We may not load more weight than the maximum weight capacity of the ship.

Return the least weight capacity of the ship that will result in all the packages on the conveyor belt being shipped within `D` days.

**Example 1:**

    Input: weights = [1,2,3,4,5,6,7,8,9,10], D = 5
    Output: 15
    Explanation: 
    A ship capacity of 15 is the minimum to ship all the packages in 5 days like this:
    1st day: 1, 2, 3, 4, 5
    2nd day: 6, 7
    3rd day: 8
    4th day: 9
    5th day: 10
    
    Note that the cargo must be shipped in the order given, so using a ship of capacity 14 and splitting the packages into parts like (2, 3, 4, 5), (1, 6, 7), (8), (9), (10) is not allowed.

**Example 2:**

    Input: weights = [3,2,2,4,1,4], D = 3
    Output: 6
    Explanation: 
    A ship capacity of 6 is the minimum to ship all the packages in 3 days like this:
    1st day: 3, 2
    2nd day: 2, 4
    3rd day: 1, 4

**Example 3:**

    Input: weights = [1,2,3,1,1], D = 4
    Output: 3
    Explanation: 
    1st day: 1
    2nd day: 2
    3rd day: 3
    4th day: 1, 1

**Note:**

1. `1 <= D <= weights.length <= 50000`
2. `1 <= weights[i] <= 500`


## 题目大意

传送带上的包裹必须在 D 天内从一个港口运送到另一个港口。

传送带上的第 i 个包裹的重量为 weights[i]。每一天，我们都会按给出重量的顺序往传送带上装载包裹。我们装载的重量不会超过船的最大运载重量。

返回能在 D 天内将传送带上的所有包裹送达的船的最低运载能力。

提示：

- 1 <= D <= weights.length <= 50000
- 1 <= weights[i] <= 500


## 解题思路

- 给出一个数组和天数 D，要求正好在 D 天把数组中的货物都运完。求传输带上能承受的最小货物重量是多少。
- 这一题和第 410 题完全一样，只不过换了一个题面。代码完全不变。思路解析见第 410 题。
