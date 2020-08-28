# [1385. Find the Distance Value Between Two Arrays](https://leetcode.com/problems/find-the-distance-value-between-two-arrays/)


## 题目

Given two integer arrays `arr1` and `arr2`, and the integer `d`, *return the distance value between the two arrays*.

The distance value is defined as the number of elements `arr1[i]` such that there is not any element `arr2[j]` where `|arr1[i]-arr2[j]| <= d`.

**Example 1**:

```
Input: arr1 = [4,5,8], arr2 = [10,9,1,8], d = 2
Output: 2
Explanation: 
For arr1[0]=4 we have: 
|4-10|=6 > d=2 
|4-9|=5 > d=2 
|4-1|=3 > d=2 
|4-8|=4 > d=2 
For arr1[1]=5 we have: 
|5-10|=5 > d=2 
|5-9|=4 > d=2 
|5-1|=4 > d=2 
|5-8|=3 > d=2
For arr1[2]=8 we have:
|8-10|=2 <= d=2
|8-9|=1 <= d=2
|8-1|=7 > d=2
|8-8|=0 <= d=2
```

**Example 2**:

```
Input: arr1 = [1,4,2,3], arr2 = [-4,-3,6,10,20,30], d = 3
Output: 2
```

**Example 3**:

```
Input: arr1 = [2,1,100,3], arr2 = [-5,-2,10,-3,7], d = 6
Output: 1
```

**Constraints**:

- `1 <= arr1.length, arr2.length <= 500`
- `-10^3 <= arr1[i], arr2[j] <= 10^3`
- `0 <= d <= 100`


## 题目大意

给你两个整数数组 arr1 ， arr2 和一个整数 d ，请你返回两个数组之间的 距离值 。「距离值」 定义为符合此距离要求的元素数目：对于元素 arr1[i] ，不存在任何元素 arr2[j] 满足 |arr1[i]-arr2[j]| <= d 。

提示：

- 1 <= arr1.length, arr2.length <= 500
- -10^3 <= arr1[i], arr2[j] <= 10^3
- 0 <= d <= 100


## 解题思路

- 计算两个数组之间的距离，距离值的定义：满足对于元素 arr1[i] ，不存在任何元素 arr2[j] 满足 |arr1[i]-arr2[j]| <= d 这一条件的元素数目。
- 简单题，按照距离值的定义，双层循环计数即可。

## 代码

```go
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	for i := range arr1 {
		for j := range arr2 {
			if abs(arr1[i]-arr2[j]) <= d {
				break
			}
			if j == len(arr2)-1 {
				res++
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
```