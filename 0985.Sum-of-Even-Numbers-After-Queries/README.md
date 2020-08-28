# [985. Sum of Even Numbers After Queries](https://leetcode.com/problems/sum-of-even-numbers-after-queries/)

## 题目

We have an array `A` of integers, and an array `queries` of queries.

For the `i`-th query `val = queries[i][0], index = queries[i][1]`, we add val to `A[index]`. Then, the answer to the `i`-th query is the sum of the even values of `A`.

*(Here, the given `index = queries[i][1]` is a 0-based index, and each query permanently modifies the array `A`.)*

Return the answer to all queries. Your `answer` array should have `answer[i]` as the answer to the `i`-th query.

**Example 1**:

```
Input: A = [1,2,3,4], queries = [[1,0],[-3,1],[-4,0],[2,3]]
Output: [8,6,2,4]
Explanation: 
At the beginning, the array is [1,2,3,4].
After adding 1 to A[0], the array is [2,2,3,4], and the sum of even values is 2 + 2 + 4 = 8.
After adding -3 to A[1], the array is [2,-1,3,4], and the sum of even values is 2 + 4 = 6.
After adding -4 to A[0], the array is [-2,-1,3,4], and the sum of even values is -2 + 4 = 2.
After adding 2 to A[3], the array is [-2,-1,3,6], and the sum of even values is -2 + 6 = 4.
```

**Note**:

1. `1 <= A.length <= 10000`
2. `-10000 <= A[i] <= 10000`
3. `1 <= queries.length <= 10000`
4. `-10000 <= queries[i][0] <= 10000`
5. `0 <= queries[i][1] < A.length`

## 题目大意

给出一个整数数组 A 和一个查询数组 queries。

对于第 i 次查询，有 val = queries[i][0], index = queries[i][1]，我们会把 val 加到 A[index] 上。然后，第 i 次查询的答案是 A 中偶数值的和。（此处给定的 index = queries[i][1] 是从 0 开始的索引，每次查询都会永久修改数组 A。）返回所有查询的答案。你的答案应当以数组 answer 给出，answer[i] 为第 i 次查询的答案。


## 解题思路

- 给出一个数组 A 和 query 数组。要求每次 query 操作都改变数组 A 中的元素值，并计算此次操作结束数组 A 中偶数值之和。
- 简单题，先计算 A 中所有偶数之和。再每次 query 操作的时候，动态维护这个偶数之和即可。

## 代码

```go
func sumEvenAfterQueries(A []int, queries [][]int) []int {
	cur, res := 0, []int{}
	for _, v := range A {
		if v%2 == 0 {
			cur += v
		}
	}
	for _, q := range queries {
		if A[q[1]]%2 == 0 {
			cur -= A[q[1]]
		}
		A[q[1]] += q[0]
		if A[q[1]]%2 == 0 {
			cur += A[q[1]]
		}
		res = append(res, cur)
	}
	return res
}
```