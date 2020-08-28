# [724. Find Pivot Index](https://leetcode.com/problems/find-pivot-index/)


## 题目

Given an array of integers nums, write a method that returns the "pivot" index of this array.

We define the pivot index as the index where the sum of all the numbers to the left of the index is equal to the sum of all the numbers to the right of the index.

If no such index exists, we should return -1. If there are multiple pivot indexes, you should return the left-most pivot index.

 

**Example 1**:

    Input: nums = [1,7,3,6,5,6]
    Output: 3
    Explanation:
    The sum of the numbers to the left of index 3 (nums[3] = 6) is equal to the sum of numbers to the right of index 3.
    Also, 3 is the first index where this occurs.

**Example 2**:

    Input: nums = [1,2,3]
    Output: -1
    Explanation:
    There is no index that satisfies the conditions in the problem statement.

**Constraints**:

- The length of nums will be in the range [0, 10000].
- Each element nums[i] will be an integer in the range [-1000, 1000].



## 题目大意

给定一个整数类型的数组 nums，请编写一个能够返回数组 “中心索引” 的方法。我们是这样定义数组 中心索引 的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。



## 解题思路

- 在数组中，找到一个数，使得它左边的数之和等于它的右边的数之和，如果存在，则返回这个数的下标索引，否作返回 -1。
- 这里面存在一个等式，只需要满足这个等式即可满足条件：leftSum + num[i] = sum - leftSum => 2 * leftSum + num[i] = sum
- 题目提到如果存在多个索引，则返回最左边那个，因此从左开始求和，而不是从右边

## 代码

```go

package leetcode

// 2 * leftSum + num[i] = sum
// 时间: O(n)
// 空间: O(1)
func pivotIndex(nums []int) int {
	if len(nums) <= 0 {
		return -1
	}
	var sum, leftSum int
	for _, num := range nums {
		sum += num
	}
	for index, num := range nums {
		if leftSum*2+num == sum {
			return index
		}
		leftSum += num
	}
	return -1
}

```
