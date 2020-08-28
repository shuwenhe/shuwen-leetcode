# [1380. Lucky Numbers in a Matrix](https://leetcode.com/problems/lucky-numbers-in-a-matrix/)


## 题目

Given a `m * n` matrix of **distinct** numbers, return all lucky numbers in the matrix in **any** order.

A lucky number is an element of the matrix such that it is the minimum element in its row and maximum in its column.

**Example 1**:

```
Input: matrix = [[3,7,8],[9,11,13],[15,16,17]]
Output: [15]
Explanation: 15 is the only lucky number since it is the minimum in its row and the maximum in its column
```

**Example 2**:

```
Input: matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
Output: [12]
Explanation: 12 is the only lucky number since it is the minimum in its row and the maximum in its column.
```

**Example 3**:

```
Input: matrix = [[7,8],[1,2]]
Output: [7]
```

**Constraints**:

- `m == mat.length`
- `n == mat[i].length`
- `1 <= n, m <= 50`
- `1 <= matrix[i][j] <= 10^5`.
- All elements in the matrix are distinct.

## 题目大意

给你一个 m * n 的矩阵，矩阵中的数字 各不相同 。请你按 任意 顺序返回矩阵中的所有幸运数。幸运数是指矩阵中满足同时下列两个条件的元素：

- 在同一行的所有元素中最小
- 在同一列的所有元素中最大



## 解题思路

- 找出矩阵中的幸运数。幸运数的定义：同时满足 2 个条件，在同一行的所有元素中最小并且在同一列的所有元素中最大。
- 简单题。按照题意遍历矩阵，找到同时满足 2 个条件的数输出即可。

## 代码

```go
func luckyNumbers(matrix [][]int) []int {
	t, r, res := make([]int, len(matrix[0])), make([]int, len(matrix[0])), []int{}
	for _, val := range matrix {
		m, k := val[0], 0
		for j := 0; j < len(matrix[0]); j++ {
			if val[j] < m {
				m = val[j]
				k = j
			}
			if t[j] < val[j] {
				t[j] = val[j]
			}
		}

		if t[k] == m {
			r[k] = m
		}
	}
	for k, v := range r {
		if v > 0 && v == t[k] {
			res = append(res, v)
		}
	}
	return res
}
```