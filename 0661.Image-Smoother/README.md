# [661. Image Smoother](https://leetcode.com/problems/image-smoother/)

## 题目

Given a 2D integer matrix M representing the gray scale of an image, you need to design a smoother to make the gray scale of each cell becomes the average gray scale (rounding down) of all the 8 surrounding cells and itself. If a cell has less than 8 surrounding cells, then use as many as you can.

**Example 1**:

```
Input:
[[1,1,1],
 [1,0,1],
 [1,1,1]]
Output:
[[0, 0, 0],
 [0, 0, 0],
 [0, 0, 0]]
Explanation:
For the point (0,0), (0,2), (2,0), (2,2): floor(3/4) = floor(0.75) = 0
For the point (0,1), (1,0), (1,2), (2,1): floor(5/6) = floor(0.83333333) = 0
For the point (1,1): floor(8/9) = floor(0.88888889) = 0

```

**Note**:

1. The value in the given matrix is in the range of [0, 255].
2. The length and width of the given matrix are in the range of [1, 150].


## 题目大意

包含整数的二维矩阵 M 表示一个图片的灰度。你需要设计一个平滑器来让每一个单元的灰度成为平均灰度 (向下舍入) ，平均灰度的计算是周围的8个单元和它本身的值求平均，如果周围的单元格不足八个，则尽可能多的利用它们。

注意:

- 给定矩阵中的整数范围为 [0, 255]。
- 矩阵的长和宽的范围均为 [1, 150]。


## 解题思路

- 将二维数组中的每个元素变为周围 9 个元素的平均值。
- 简单题，按照题意计算平均值即可。需要注意的是边界问题，四个角和边上的元素，这些点计算平均值的时候，计算平均值都不足 9 个元素。

## 代码

```go
func imageSmoother(M [][]int) [][]int {
	res := make([][]int, len(M))
	for i := range M {
		res[i] = make([]int, len(M[0]))
	}
	for y := 0; y < len(M); y++ {
		for x := 0; x < len(M[0]); x++ {
			res[y][x] = smooth(x, y, M)
		}
	}
	return res
}

func smooth(x, y int, M [][]int) int {
	count, sum := 1, M[y][x]
	// Check bottom
	if y+1 < len(M) {
		sum += M[y+1][x]
		count++
	}
	// Check Top
	if y-1 >= 0 {
		sum += M[y-1][x]
		count++
	}
	// Check left
	if x-1 >= 0 {
		sum += M[y][x-1]
		count++
	}
	// Check Right
	if x+1 < len(M[y]) {
		sum += M[y][x+1]
		count++
	}
	// Check Coners
	// Top Left
	if y-1 >= 0 && x-1 >= 0 {
		sum += M[y-1][x-1]
		count++
	}
	// Top Right
	if y-1 >= 0 && x+1 < len(M[0]) {
		sum += M[y-1][x+1]
		count++
	}
	// Bottom Left
	if y+1 < len(M) && x-1 >= 0 {
		sum += M[y+1][x-1]
		count++
	}
	//Bottom Right
	if y+1 < len(M) && x+1 < len(M[0]) {
		sum += M[y+1][x+1]
		count++
	}
	return sum / count
}
```