# [995. Minimum Number of K Consecutive Bit Flips](https://leetcode.com/problems/minimum-number-of-k-consecutive-bit-flips/)


## 题目

In an array `A` containing only 0s and 1s, a `K`-bit flip consists of choosing a (contiguous) subarray of length `K` and simultaneously changing every 0 in the subarray to 1, and every 1 in the subarray to 0.

Return the minimum number of `K`-bit flips required so that there is no 0 in the array. If it is not possible, return `-1`.

**Example 1:**

    Input: A = [0,1,0], K = 1
    Output: 2
    Explanation: Flip A[0], then flip A[2].

**Example 2:**

    Input: A = [1,1,0], K = 2
    Output: -1
    Explanation: No matter how we flip subarrays of size 2, we can't make the array become [1,1,1].

**Example 3:**

    Input: A = [0,0,0,1,0,1,1,0], K = 3
    Output: 3
    Explanation:
    Flip A[0],A[1],A[2]: A becomes [1,1,1,1,0,1,1,0]
    Flip A[4],A[5],A[6]: A becomes [1,1,1,1,1,0,0,0]
    Flip A[5],A[6],A[7]: A becomes [1,1,1,1,1,1,1,1]

**Note:**

1. `1 <= A.length <= 30000`
2. `1 <= K <= A.length`


## 题目大意

在仅包含 0 和 1 的数组 A 中，一次 K 位翻转包括选择一个长度为 K 的（连续）子数组，同时将子数组中的每个 0 更改为 1，而每个 1 更改为 0。返回所需的 K 位翻转的次数，以便数组没有值为 0 的元素。如果不可能，返回 -1。

提示：

1. 1 <= A.length <= 30000
2. 1 <= K <= A.length


## 解题思路


- 给出一个数组，数组里面的元素只有 0 和 1。给一个长度为 K 的窗口，在这个窗口内的所有元素都会 0-1 翻转。问最后需要翻转几次，使得整个数组都为 1 。如果不能翻转使得最后数组元素都为 1，则输出 -1。
- 拿到这一题首先想到的是贪心算法。例如第 765 题，这类题的描述都是这样的：在一个数组中或者环形数组中通过交换位置，或者翻转变换，达到最终结果，要求找到最少步数。贪心能保证是最小步数(证明略)。按照贪心的思想，这一题也这样做，从数组 0 下标开始往后扫，依次翻转每个 K 大小的窗口内元素。
- 由于窗口大小限制了，所以这题滑动窗口只需要一个边界坐标，用左边界就可以判断了。每一个 `A[i]` 是否需要翻转，是由于 `[ i-k+1，i ]`、`[ i-k+2，i+1 ]`、`[ i-k+3，i+2 ]`……`[ i-1，i+k ]` 这一系列的窗口翻转`累积影响`的。那如何之前这些窗口`累积`到 `A[i]` 上翻转的次数呢？可以动态的维护一个翻转次数，当 `i` 摆脱了上一次翻转窗口 `K` 的时候，翻转次数就 `-1` 。举个例子：

        A = [0 0 0 1 0 1 1 0] K = 3
        
        A = [2 0 0 1 0 1 1 0] i = 0 flippedTime = 1
        A = [2 0 0 1 0 1 1 0] i = 1 flippedTime = 1
        A = [2 0 0 1 0 1 1 0] i = 2 flippedTime = 1
        A = [2 0 0 1 0 1 1 0] i = 3 flippedTime = 0
        A = [2 0 0 1 2 1 1 0] i = 4 flippedTime = 1
        A = [2 0 0 1 2 2 1 0] i = 5 flippedTime = 2
        A = [2 0 0 1 2 2 1 0] i = 6 flippedTime = 2
        A = [2 0 0 1 2 2 1 0] i = 7 flippedTime = 1

    当判断 `A[i]` 是否需要翻转的时候，只需要留意每个宽度为 `K` 窗口的左边界。会影响到 A[i] 的窗口的左边界分别是 `i-k+1`、`i-k+2`、`i-k+3`、…… `i-1`，只需要分别看这些窗口有没有翻转就行。这里可以用特殊标记来记录这些窗口的左边界是否被翻转了。如果翻转过，则把窗口左边界的那个数字标记为 2 (为什么要标记为 2 呢？其实设置成什么都可以，只要不是 0 和 1 ，和原有的数字区分开就行)。当 `i≥k` 的时候，代表 `i` 已经脱离了 `i-k` 的这个窗口，因为能影响 `A[i]` 的窗口是从 `i-k+1` 开始的，如果 `A[i-k] == 2` 代表 `i-k` 窗口已经翻转过了，现在既然脱离了它的窗口影响，那么就要把累积的 `flippedTime - 1` 。这样就维护了累积 `flippedTime` 和滑动窗口中累积影响的关系。

- 接下来还需要处理的是 `flippedTime` 与当前 `A[i]` 是否翻转的问题。如果 `flippedTime` 是偶数次，原来的 0 还是 0，就需要再次翻转，如果 `flippedTime` 是奇数次，原来的 0 变成了 1 就不需要翻转了。总结成一条结论就是 `A[i]` 与 `flippedTime` 同奇偶性的时候就要翻转。当 `i + K` 比 `len(A)` 大的时候，代表剩下的这些元素肯定不能在一个窗口里面翻转，则输出 -1 。
