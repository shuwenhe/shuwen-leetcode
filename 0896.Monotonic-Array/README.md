# [896. Monotonic Array](https://leetcode.com/problems/monotonic-array/)


## 题目

An array is *monotonic* if it is either monotone increasing or monotone decreasing.

An array `A` is monotone increasing if for all `i <= j`, `A[i] <= A[j]`. An array `A` is monotone decreasing if for all `i <= j`, `A[i] >= A[j]`.

Return `true` if and only if the given array `A` is monotonic.

**Example 1**:

```
Input: [1,2,2,3]
Output: true
```

**Example 2**:

```
Input: [6,5,4,4]
Output: true
```

**Example 3**:

```
Input: [1,3,2]
Output: false
```

**Example 4**:

```
Input: [1,2,4,5]
Output: true
```

**Example 5**:

```
Input: [1,1,1]
Output: true
```

**Note**:

1. `1 <= A.length <= 50000`
2. `-100000 <= A[i] <= 100000`

## 题目大意

如果数组是单调递增或单调递减的，那么它是单调的。如果对于所有 i <= j，A[i] <= A[j]，那么数组 A 是单调递增的。 如果对于所有 i <= j，A[i]> = A[j]，那么数组 A 是单调递减的。当给定的数组 A 是单调数组时返回 true，否则返回 false。


## 解题思路

- 判断给定的数组是不是单调(单调递增或者单调递减)的。
- 简单题，按照题意循环判断即可。

## 代码

```go
func isMonotonic(A []int) bool {
	if len(A) <= 1 {
		return true
	}
	if A[0] < A[1] {
		return inc(A[1:])
	}
	if A[0] > A[1] {
		return dec(A[1:])
	}
	return inc(A[1:]) || dec(A[1:])
}

func inc(A []int) bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			return false
		}
	}
	return true
}

func dec(A []int) bool {
	for i := 0; i < len(A)-1; i++ {
		if A[i] < A[i+1] {
			return false
		}
	}
	return true
}
```