# [1232. Check If It Is a Straight Line](https://leetcode.com/problems/check-if-it-is-a-straight-line/)


## 题目

You are given an array `coordinates`, `coordinates[i] = [x, y]`, where `[x, y]` represents the coordinate of a point. Check if these points make a straight line in the XY plane.

**Example 1**:

![](https://img.halfrost.com/Leetcode/leetcode_1232_1.png)

    Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
    Output: true

**Example 2**:


![](https://img.halfrost.com/Leetcode/leetcode_1232_2.png)

    Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
    Output: false

**Constraints**:

- `2 <= coordinates.length <= 1000`
- `coordinates[i].length == 2`
- `-10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4`
- `coordinates` contains no duplicate point.

## 题目大意


在一个 XY 坐标系中有一些点，我们用数组 coordinates 来分别记录它们的坐标，其中 coordinates[i] = [x, y] 表示横坐标为 x、纵坐标为 y 的点。

请你来判断，这些点是否在该坐标系中属于同一条直线上，是则返回 true，否则请返回 false。

提示：

- 2 <= coordinates.length <= 1000
- coordinates[i].length == 2
- -10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
- coordinates 中不含重复的点



## 解题思路

- 给出一组坐标点，要求判断这些点是否在同一直线上。
- 按照几何原理，依次计算这些点的斜率是否相等即可。斜率需要做除法，这里采用一个技巧是换成乘法。例如 `a/b = c/d` 换成乘法是 `a*d  = c*d`  。


## 代码

```go

package leetcode

func checkStraightLine(coordinates [][]int) bool {
	dx0 := coordinates[1][0] - coordinates[0][0]
	dy0 := coordinates[1][1] - coordinates[0][1]
	for i := 1; i < len(coordinates)-1; i++ {
		dx := coordinates[i+1][0] - coordinates[i][0]
		dy := coordinates[i+1][1] - coordinates[i][1]
		if dy*dx0 != dy0*dx { // check cross product
			return false
		}
	}
	return true
}


```