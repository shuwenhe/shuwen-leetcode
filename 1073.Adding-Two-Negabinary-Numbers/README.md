# [1073. Adding Two Negabinary Numbers](https://leetcode.com/problems/adding-two-negabinary-numbers/)


## 题目

Given two numbers `arr1` and `arr2` in base **-2**, return the result of adding them together.

Each number is given in *array format*: as an array of 0s and 1s, from most significant bit to least significant bit. For example, `arr = [1,1,0,1]`represents the number `(-2)^3 + (-2)^2 + (-2)^0 = -3`. A number `arr` in *array format* is also guaranteed to have no leading zeros: either `arr == [0]` or `arr[0] == 1`.

Return the result of adding `arr1` and `arr2` in the same format: as an array of 0s and 1s with no leading zeros.

**Example 1:**

    Input: arr1 = [1,1,1,1,1], arr2 = [1,0,1]
    Output: [1,0,0,0,0]
    Explanation: arr1 represents 11, arr2 represents 5, the output represents 16.

**Note:**

1. `1 <= arr1.length <= 1000`
2. `1 <= arr2.length <= 1000`
3. `arr1` and `arr2` have no leading zeros
4. `arr1[i]` is `0` or `1`
5. `arr2[i]` is `0` or `1`


## 题目大意

给出基数为 -2 的两个数 arr1 和 arr2，返回两数相加的结果。数字以 数组形式 给出：数组由若干 0 和 1 组成，按最高有效位到最低有效位的顺序排列。例如，arr = [1,1,0,1] 表示数字 (-2)^3 + (-2)^2 + (-2)^0 = -3。数组形式 的数字也同样不含前导零：以 arr 为例，这意味着要么 arr == [0]，要么 arr[0] == 1。

返回相同表示形式的 arr1 和 arr2 相加的结果。两数的表示形式为：不含前导零、由若干 0 和 1 组成的数组。

提示：

- 1 <= arr1.length <= 1000
- 1 <= arr2.length <= 1000
- arr1 和 arr2 都不含前导零
- arr1[i] 为 0 或 1
- arr2[i] 为 0 或 1



## 解题思路

- 给出两个 -2 进制的数，要求计算出这两个数的和，最终表示形式还是 -2 进制。
- 这一题最先想到的思路是先把两个 -2 进制的数转成 10 进制以后做加法，然后把结果表示成 -2 进制。这个思路可行，但是在提交以后会发现数据溢出 int64 了。在第 257 / 267 组测试数据会出现 WA。测试数据见 test 文件。另外换成 big.Add 也不是很方便。所以考虑换一个思路。
- 这道题实际上就是求两个 -2 进制数的加法，为什么还要先转到 10 进制再换回 -2 进制呢？为何不直接进行 -2 进制的加法。所以开始尝试直接进行加法运算。加法是从低位到高位依次累加，遇到进位要从低往高位进位。所以从两个数组的末尾往前扫，模拟低位相加的过程。关键的是进位问题。进位分 3 种情况，依次来讨论：

1. 进位到高位 k ，高位 k 上的两个数数字分别是 0 和 0 。这种情况最终 k 位为 1 。
```c
        证明：由于进位是由 k - 1 位进过来的，所以 k - 1 位是 2 个 1 。现在 k 位是 2 个 0，
        所以加起来的和是 2 * (-2)^(k - 1)。
        当 k 为奇数的时候，2 * (-2)^(k - 1) = (-1)^(k - 1)* 2 * 2^(k - 1) = 2^k
        当 k 为偶数的时候，2 * (-2)^(k - 1) = (-1)^(k - 1)* 2 * 2^(k - 1) = -2^k
        综合起来就是 (-2)^k，所以最终 k 位上有一个 1
```
2. 进位到高位 k ，高位 k 上的两个数数字分别是 0 和 1 。这种情况最终 k 位为 0 。
```c
        证明：由于进位是由 k - 1 位进过来的，所以 k - 1 位是 2 个 1。现在 k 位是 1 个 0 和 1 个 1,
        所以加起来的和是 (-2)^k + 2 * (-2)^(k - 1)。
        当 k 为奇数的时候，(-2)^k + 2 * (-2)^(k - 1) = -2^k + 2^k = 0
        当 k 为偶数的时候，(-2)^k + 2 * (-2)^(k - 1) = 2^k - 2^k = 0
        综合起来就是 0，所以最终 k 位上有一个 0
```
3. 进位到高位 k ，高位 k 上的两个数数字分别是 1 和 1 。这种情况最终 k 位为 1 。
```c
        证明：由于进位是由 k - 1 位进过来的，所以 k - 1 位是 2 个 1 。现在 k 位是 2 个 1，
        所以加起来的和是 2 * (-2)^k + 2 * (-2)^(k - 1)。
        当 k 为奇数的时候，2 * (-2)^k + 2 * (-2)^(k - 1) = -2^(k + 1) + 2^k = 2^k*(1 - 2) = -2^k
        当 k 为偶数的时候，2 * (-2)^k + 2 * (-2)^(k - 1) = 2^(k + 1) - 2^k = 2^k*(2 - 1) = 2^k
        综合起来就是 (-2)^k，所以最终 k 位上有一个 1
```

- 所以综上所属，-2 进制的进位和 2 进制的进位原理是完全一致的，只不过 -2 进制的进位是 -1，而 2 进制的进位是 1 。由于进位可能在 -2 进制上出现前导 0 ，所以最终结果需要再去除前导 0 。

