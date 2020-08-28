# [888. Fair Candy Swap](https://leetcode.com/problems/fair-candy-swap/)


## 题目

Alice and Bob have candy bars of different sizes: `A[i]` is the size of the `i`-th bar of candy that Alice has, and `B[j]` is the size of the `j`-th bar of candy that Bob has.

Since they are friends, they would like to exchange one candy bar each so that after the exchange, they both have the same total amount of candy. (*The total amount of candy a person has is the sum of the sizes of candy bars they have*.)

Return an integer array `ans` where `ans[0]` is the size of the candy bar that Alice must exchange, and `ans[1]` is the size of the candy bar that Bob must exchange.

If there are multiple answers, you may return any one of them. It is guaranteed an answer exists.

**Example 1**:

```
Input: A = [1,1], B = [2,2]
Output: [1,2]
```

**Example 2**:

```
Input: A = [1,2], B = [2,3]
Output: [1,2]
```

**Example 3**:

```
Input: A = [2], B = [1,3]
Output: [2,3]
```

**Example 4**:

```
Input: A = [1,2,5], B = [2,4]
Output: [5,4]
```

**Note**:

- `1 <= A.length <= 10000`
- `1 <= B.length <= 10000`
- `1 <= A[i] <= 100000`
- `1 <= B[i] <= 100000`
- It is guaranteed that Alice and Bob have different total amounts of candy.
- It is guaranteed there exists an answer.


## 题目大意

爱丽丝和鲍勃有不同大小的糖果棒：A[i] 是爱丽丝拥有的第 i 块糖的大小，B[j] 是鲍勃拥有的第 j 块糖的大小。因为他们是朋友，所以他们想交换一个糖果棒，这样交换后，他们都有相同的糖果总量。（一个人拥有的糖果总量是他们拥有的糖果棒大小的总和。）返回一个整数数组 ans，其中 ans[0] 是爱丽丝必须交换的糖果棒的大小，ans[1] 是 Bob 必须交换的糖果棒的大小。如果有多个答案，你可以返回其中任何一个。保证答案存在。

提示：

- 1 <= A.length <= 10000
- 1 <= B.length <= 10000
- 1 <= A[i] <= 100000
- 1 <= B[i] <= 100000
- 保证爱丽丝与鲍勃的糖果总量不同。
- 答案肯定存在。


## 解题思路

- 两人交换糖果，使得两人糖果相等。要求输出一个数组，里面分别包含两人必须交换的糖果大小。
- 首先这一题肯定了一定有解，其次只允许交换一次。有了这两个前提，使本题变成简单题。先计算出为了使得交换以后两个相同的糖果数，A 需要增加或者减少的糖果数 diff。然后遍历 B ，看 A 中是否存在一个元素，能使得 B 做了对应交换 diff 以后，两人糖果相等。(此题前提保证了一定能找到)。最后输出 A 中的这个元素和遍历到 B 的这个元素，即是两人要交换的糖果数。

## 代码

```go
func fairCandySwap(A []int, B []int) []int {
	hDiff, aMap := diff(A, B)/2, make(map[int]int, len(A))
	for _, a := range A {
		aMap[a] = a
	}
	for _, b := range B {
		if a, ok := aMap[hDiff+b]; ok {
			return []int{a, b}
		}
	}
	return nil
}

func diff(A []int, B []int) int {
	diff, maxLen := 0, max(len(A), len(B))
	for i := 0; i < maxLen; i++ {
		if i < len(A) {
			diff += A[i]
		}
		if i < len(B) {
			diff -= B[i]
		}
	}
	return diff
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```