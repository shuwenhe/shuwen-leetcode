# [327. Count of Range Sum](https://leetcode.com/problems/count-of-range-sum/)


## 题目

Given an integer array `nums`, return the number of range sums that lie in `[lower, upper]` inclusive.Range sum `S(i, j)` is defined as the sum of the elements in `nums` between indices `i` and `j` (`i` ≤ `j`), inclusive.

**Note:**A naive algorithm of O(n2) is trivial. You MUST do better than that.

**Example:**

    Input: nums = [-2,5,-1], lower = -2, upper = 2,
    Output: 3 
    Explanation: The three ranges are : [0,0], [2,2], [0,2] and their respective sums are: -2, -1, 2.


## 题目大意


给定一个整数数组 nums，返回区间和在 [lower, upper] 之间的个数，包含 lower 和 upper。区间和 S(i, j) 表示在 nums 中，位置从 i 到 j 的元素之和，包含 i 和 j (i ≤ j)。

说明:   
最直观的算法复杂度是 O(n^2) ，请在此基础上优化你的算法。


## 解题思路

- 给出一个数组，要求在这个数组中找出任意一段子区间的和，位于 [lower,upper] 之间。
- 这一题可以用暴力解法，2 层循环，遍历所有子区间，求和并判断是否位于 [lower,upper] 之间，时间复杂度 O(n^2)。
- 这一题当然还有更优的解法，用线段树或者树状数组，将时间复杂度降为 O(n log n)。题目中要求 `lower ≤ sum(i,j) ≤ upper`，`sum(i,j) = prefixSum(j) - prefixSum(i-1)`，那么 `lower + prefixSum(i-1) ≤ prefixSum(j) ≤ upper + prefixSum(i-1)`。所以利用前缀和将区间和转换成了前缀和在线段树中 `query` 的问题，只不过线段树中父节点中存的不是子节点的和，而应该是子节点出现的次数。第二个转换，由于前缀和会很大，所以需要离散化。例如 `prefixSum = [-3,-2,-1,0]`，用前缀和下标进行离散化，所以线段树中左右区间变成了 0-3 。

    ![](https://img.halfrost.com/Leetcode/leetcode_327_0.png)

    利用 `prefixSum` 下标离散化：

    ![](https://img.halfrost.com/Leetcode/leetcode_327_1.png)

- 还需要注意一些小细节，`prefixSum` 计算完以后需要去重，去重以后并排序，方便构造线段树的有效区间。如果不去重，线段树中可能出现非法区间(left > right)或者重叠区间。最后一步往线段树中倒序插入 `prefixSum` 的时候，用的是非去重的，插入 `prefixSum[j]` 代表 sum(i,j) 中的 j，例如往线段树中插入 `prefixSum[5]`，代表当前树中加入了 j = 5 的情况。query 操作实质是在做区间匹配，例如当前 i 循环到 i = 3，累计往线段树中插入了 `prefixSum[5]`，`prefixSum[4]`，`prefixSum[3]`，那么 query 操作实质是在判断：`lower ≤ sum(i=3,j=3) ≤ upper`，`lower ≤ sum(i=3,j=4) ≤ upper`，`lower ≤ sum(i=3,j=5) ≤ upper`，这 3 个等式是否成立，有几个成立就返回几个，即是最终要求得的结果的一部分。
- 举个例子，`nums = [-3,1,2,-2,2,-1]`，`prefixSum = [-3,-2,0,-2,0,-1]`，去重以后并排序得到 `sum = [-3,-2,-1,0]`。离散化构造线段树，这里出于演示的方便，下图中就不画出离散后的线段树了，用非离散的线段树展示：

    ![](https://img.halfrost.com/Leetcode/leetcode_327_2_.png)

    倒序插入 `len(prefixSum)-1 = prefixSum[5] = -1`：

    ![](https://img.halfrost.com/Leetcode/leetcode_327_3_.png)

    这时候查找区间变为了 `[-3 + prefixSum[5-1], -1 + prefixSum[5-1]] = [-3,-1]`，即判断 `-3 ≤ sum(5,5) ≤ -1`，满足等式的有几种情况，这里明显只有一种情况，即 `j = 5`，也满足等式，所以这一步 `res = 1`。

- 倒序插入 `len(prefixSum)-2 = prefixSum[4] = 0`：

    ![](https://img.halfrost.com/Leetcode/leetcode_327_4_.png)

    这时候查找区间变为了 `[-3 + prefixSum[4-1], -1 + prefixSum[4-1]] = [-5,-3]`，即判断 `-5 ≤ sum(4,  4,5) ≤ -3`，满足等式的有几种情况，这里有两种情况，即 `j = 4` 或者 `j = 5`，都不满足等式，所以这一步 `res = 0`。

- 倒序插入 `len(prefixSum)-3 = prefixSum[3] = -2`：

    ![](https://img.halfrost.com/Leetcode/leetcode_327_5_.png)

    这时候查找区间变为了 `[-3 + prefixSum[3-1], -1 + prefixSum[3-1]] = [-3,-1]`，即判断 `-3 ≤ sum(3,  3,4,5) ≤ -1`，满足等式的有几种情况，这里有三种情况，即 `j = 3` 、`j = 4` 或者 `j = 5`，满足等式的有 `j = 3` 和 `j = 5`，即 `-3 ≤ sum(3, 3) ≤ -1` 和 `-3 ≤ sum(3, 5) ≤ -1`。所以这一步 `res = 2`。

- 倒序插入 `len(prefixSum)-4 = prefixSum[2] = 0`：

    ![](https://img.halfrost.com/Leetcode/leetcode_327_6_.png)

    这时候查找区间变为了 `[-3 + prefixSum[2-1], -1 + prefixSum[2-1]] = [-5,-3]`，即判断 `-5 ≤ sum(2,  2,3,4,5) ≤ -3`，满足等式的有几种情况，这里有四种情况，即 `j = 2`、 `j = 3` 、`j = 4` 或者 `j = 5`，都不满足等式。所以这一步 `res = 0`。

- 倒序插入 `len(prefixSum)-5 = prefixSum[1] = -2`：

    ![](https://img.halfrost.com/Leetcode/leetcode_327_7_.png)

    这时候查找区间变为了 `[-3 + prefixSum[1-1], -1 + prefixSum[1-1]] = [-6,-4]`，即判断 `-6 ≤ sum(1,  1,2,3,4,5) ≤ -4`，满足等式的有几种情况，这里有五种情况，即 `j = 1`、 `j = 2`、 `j = 3` 、`j = 4` 或者 `j = 5`，都不满足等式。所以这一步 `res = 0`。

- 倒序插入 `len(prefixSum)-6 = prefixSum[0] = -3`：

    ![](https://img.halfrost.com/Leetcode/leetcode_327_8_.png)

    这时候查找区间变为了 `[-3 + prefixSum[0-1], -1 + prefixSum[0-1]] = [-3,-1]`，注意 `prefixSum[-1] = 0`，即判断 `-3 ≤ sum(0,  0,1,2,3,4,5) ≤ -1`，满足等式的有几种情况，这里有六种情况，即 `j = 0`、`j = 1`、`j = 2`、 `j = 3` 、`j = 4` 或者 `j = 5`，满足等式的有 `j = 0`、`j = 1`、 `j = 3` 和 `j = 5`，即 `-3 ≤ sum(0, 0) ≤ -1` 、 `-3 ≤ sum(0, 1) ≤ -1`、`-3 ≤ sum(0, 3) ≤ -1` 和 `-3 ≤ sum(0, 5) ≤ -1`。所以这一步 `res = 4`。最后的答案就是把每一步的结果都累加，`res = 1 + 0 + 2 + 0 + 0 + 4 = 7`。
