# [448. Find All Numbers Disappeared in an Array](https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/)


## 题目

Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements of [1, n] inclusive that do not appear in this array.

Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

**Example**:

```
Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]
```

## 题目大意

给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。找到所有在 [1, n] 范围之间没有出现在数组中的数字。你能在不使用额外空间且时间复杂度为 O(n) 的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。



## 解题思路

- 找出 [1,n] 范围内没有出现在数组中的数字。要求不使用额外空间，并且时间复杂度为 O(n)。
- 要求不能使用额外的空间，那么只能想办法在原有数组上进行修改，并且这个修改是可还原的。时间复杂度也只能允许我们一层循环。只要循环一次能标记出已经出现过的数字，这道题就可以按要求解答出来。这里笔者的标记方法是把 |nums[i]|-1 索引位置的元素标记为负数。即 nums[| nums[i] |- 1] * -1。这里需要注意的是，nums[i] 需要加绝对值，因为它可能被之前的数置为负数了，需要还原一下。最后再遍历一次数组，若当前数组元素 nums[i] 为负数，说明我们在数组中存在数字 i+1。把结果输出到最终数组里即可。

## 代码

```go
func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}
	}
	for i, v := range nums {
		if v > 0 {
			res = append(res, i+1)
		}
	}
	return res
}
```