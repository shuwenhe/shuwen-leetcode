# [1051. Height Checker](https://leetcode.com/problems/height-checker/)


## 题目

Students are asked to stand in non-decreasing order of heights for an annual photo.

Return the minimum number of students that must move in order for all students to be standing in non-decreasing order of height.

Notice that when a group of students is selected they can reorder in any possible way between themselves and the non selected students remain on their seats.

**Example 1**:

```
Input: heights = [1,1,4,2,1,3]
Output: 3
Explanation: 
Current array : [1,1,4,2,1,3]
Target array  : [1,1,1,2,3,4]
On index 2 (0-based) we have 4 vs 1 so we have to move this student.
On index 4 (0-based) we have 1 vs 3 so we have to move this student.
On index 5 (0-based) we have 3 vs 4 so we have to move this student.
```

**Example 2**:

```
Input: heights = [5,1,2,3,4]
Output: 5
```

**Example 3**:

```
Input: heights = [1,2,3,4,5]
Output: 0
```

**Constraints**:

- `1 <= heights.length <= 100`
- `1 <= heights[i] <= 100`

## 题目大意

学校在拍年度纪念照时，一般要求学生按照 非递减 的高度顺序排列。请你返回能让所有学生以 非递减 高度排列的最小必要移动人数。注意，当一组学生被选中时，他们之间可以以任何可能的方式重新排序，而未被选中的学生应该保持不动。


## 解题思路

- 给定一个高度数组，要求输出把这个数组按照非递减高度排列所需移动的最少次数。
- 简单题，最少次数意味着每次移动，一步到位，一步就移动到它所在的最终位置。那么用一个辅助排好序的数组，一一比对计数即可。

## 代码

```go
func heightChecker(heights []int) int {
	result, checker := 0, []int{}
	checker = append(checker, heights...)
	sort.Ints(checker)
	for i := 0; i < len(heights); i++ {
		if heights[i] != checker[i] {
			result++
		}
	}
	return result
}
```