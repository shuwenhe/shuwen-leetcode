# [137. Single Number II](https://leetcode.com/problems/single-number-ii/)


## 题目

Given a **non-empty** array of integers, every element appears *three* times except for one, which appears exactly once. Find that single one.

**Note:**

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

**Example 1:**

    Input: [2,2,3,2]
    Output: 3

**Example 2:**

    Input: [0,1,0,1,0,1,99]
    Output: 99


## 题目大意

给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现了三次。找出那个只出现了一次的元素。要求算法时间复杂度是线性的，并且不使用额外的辅助空间。




## 解题思路

- 这一题是第 136 题的加强版。这类题也可以扩展，在数组中每个元素都出现 5 次，找出只出现 1 次的数。
- 本题中要求找出只出现 1 次的数，出现 3 次的数都要被消除。第 136 题是消除出现 2 次的数。这一题也会相当相同的解法，出现 3 次的数也要被消除。定义状态，00、10、01，这 3 个状态。当一个数出现 3 次，那么它每个位置上的 1 出现的次数肯定是 3 的倍数，所以当 1 出现 3 次以后，就归零清除。如何能做到这点呢？仿造`三进制(00，10，01)` 就可以做到。
- 变量  ones 中记录遍历中每个位上出现 1 的个数。将它与 A[i] 进行异或，目的是：
    - 每位上两者都是 1 的，表示历史统计结果 ones 出现1次、A[i]中又出现1次，则是出现 2 次，需要进位到 twos 变量中。
    - 每位上两者分别为 0、1 的，加入到 ones 统计结果中。
    - 最后还要 & ^twos ，是为了能做到三进制，出现 3 次就清零。例如 ones = x，那么 twos = 0，当 twos = x，那么 ones = 0；
- 变量 twos 中记录遍历中每个位上出现 1 ，2次 的个数。与 A[i] 进行异或的目的和上述描述相同，不再赘述。

> 在 golang 中，&^ 表示 AND NOT 的意思。这里的 ^ 作为一元操作符，表示按位取反 (^0001 0100 = 1110 1011)，X &^ Y 的意思是将 X 中与 Y 相异的位保留，相同的位清零。

> 在 golang 中没有 Java 中的 ~ 位操作运算符，Java 中的 ~ 运算符代表按位取反。这个操作就想当于 golang 中的 ^ 运算符当做一元运算符使用的效果。

这一题还可以继续扩展，在数组中每个元素都出现 5 次，找出只出现 1 次的数。那该怎么做呢？思路还是一样的，模拟一个五进制，5 次就会消除。代码如下：

    // 解法一
    func singleNumberIII(nums []int) int {
    	na, nb, nc := 0, 0, 0
    	for i := 0; i < len(nums); i++ {
    		nb = nb ^ (nums[i] & na)
    		na = (na ^ nums[i]) & ^nc
    		nc = nc ^ (nums[i] & ^na & ^nb)
    	}
    	return na & ^nb & ^nc
    }
    
    // 解法二
    func singleNumberIIII(nums []int) int {
    	twos, threes, ones := 0xffffffff, 0xffffffff, 0
    	for i := 0; i < len(nums); i++ {
    		threes = threes ^ (nums[i] & twos)
    		twos = (twos ^ nums[i]) & ^ones
    		ones = ones ^ (nums[i] & ^twos & ^threes)
    	}
    	return ones
    }
