// hash table
// 给定一个整数数组nums和一个目标值target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。
// 示例:
// 给定 nums = [2, 7, 11, 15], target = 9
// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15} // 给定一个整数数组nums
	target := 9                 // 目标值target
	arr := twoSum(nums, target) // 在该数组中找出和为目标值的那两个整数，并返回他们的数组下标
	fmt.Println("arr = ", arr)  // 打印下标
}

func twoSum(nums []int, target int) []int { // 哈希查找的时间复杂度为 O(1)
	h := make(map[int]int)   // 哈希容器map降低时间复杂度
	for k, v := range nums { // 遍历数组nums
		if wanted, ok := h[v]; ok {
			return []int{wanted, k}
		} else {
			h[target-v] = k
		}
	}
	return nil
}
