# [1389. Create Target Array in the Given Order](https://leetcode.com/problems/create-target-array-in-the-given-order/)

## 题目

Given two arrays of integers `nums` and `index`. Your task is to create *target* array under the following rules:

- Initially *target* array is empty.
- From left to right read nums[i] and index[i], insert at index `index[i]` the value `nums[i]` in *target* array.
- Repeat the previous step until there are no elements to read in `nums` and `index.`

Return the *target* array.

It is guaranteed that the insertion operations will be valid.

**Example 1**:

```
Input: nums = [0,1,2,3,4], index = [0,1,2,2,1]
Output: [0,4,1,3,2]
Explanation:
nums       index     target
0            0        [0]
1            1        [0,1]
2            2        [0,1,2]
3            2        [0,1,3,2]
4            1        [0,4,1,3,2]
```

**Example 2**:

```
Input: nums = [1,2,3,4,0], index = [0,1,2,3,0]
Output: [0,1,2,3,4]
Explanation:
nums       index     target
1            0        [1]
2            1        [1,2]
3            2        [1,2,3]
4            3        [1,2,3,4]
0            0        [0,1,2,3,4]
```

**Example 3**:

```
Input: nums = [1], index = [0]
Output: [1]
```

**Constraints**:

- `1 <= nums.length, index.length <= 100`
- `nums.length == index.length`
- `0 <= nums[i] <= 100`
- `0 <= index[i] <= i`

## 题目大意

给你两个整数数组 nums 和 index。你需要按照以下规则创建目标数组：

- 目标数组 target 最初为空。
- 按从左到右的顺序依次读取 nums[i] 和 index[i]，在 target 数组中的下标 index[i] 处插入值 nums[i] 。
- 重复上一步，直到在 nums 和 index 中都没有要读取的元素。

请你返回目标数组。题目保证数字插入位置总是存在。




## 解题思路

- 给定 2 个数组，分别装的是待插入的元素和待插入的位置。最后输出操作完成的数组。
- 简单题，按照题意插入元素即可。

## 代码

```go
func createTargetArray(nums []int, index []int) []int {
	result := make([]int, len(nums))
	for i, pos := range index {
		copy(result[pos+1:i+1], result[pos:i])
		result[pos] = nums[i]
	}
	return result
}
```