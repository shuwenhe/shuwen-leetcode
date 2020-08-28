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
	nums := []int{2, 7, 11, 15, 19} // 给定一个整数数组nums
	target := 30                    // 目标值target
	arr := twoSum(nums, target)     // 在该数组中找出和为目标值的那两个整数，并返回他们的数组下标
	fmt.Println("arr = ", arr)      // 打印下标
}

func twoSum(nums []int, target int) []int { // 哈希查找的时间复杂度为 O(1)
	map1 := make(map[int]int) // 哈希容器map降低时间复杂度
	for index, val := range nums {
		if wanted, ok := map1[val]; ok { // 判断val1是否在map中，第一次map中肯定没有，所以ok为false，将数组中的值val和位置pos对应的键值对放入map中
			return []int{wanted, index}
		}
		map1[target-val] = index // 通过hash表查找使val2 = target - val1 的值，如果找到则返回结果，找不到，则将当前值插入map
	}
	return nil
}
