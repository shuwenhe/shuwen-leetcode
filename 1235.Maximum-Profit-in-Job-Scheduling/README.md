# [1235. Maximum Profit in Job Scheduling](https://leetcode.com/problems/maximum-profit-in-job-scheduling/)


## 题目

We have `n` jobs, where every job is scheduled to be done from `startTime[i]` to `endTime[i]`, obtaining a profit of `profit[i]`.

You're given the `startTime` , `endTime` and `profit` arrays, you need to output the maximum profit you can take such that there are no 2 jobs in the subset with overlapping time range.

If you choose a job that ends at time `X` you will be able to start another job that starts at time `X`.

**Example 1:**

![https://assets.leetcode.com/uploads/2019/10/10/sample1_1584.png](https://assets.leetcode.com/uploads/2019/10/10/sample1_1584.png)

    Input: startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
    Output: 120
    Explanation: The subset chosen is the first and fourth job. 
    Time range [1-3]+[3-6] , we get profit of 120 = 50 + 70.

**Example 2:**

![https://assets.leetcode.com/uploads/2019/10/10/sample22_1584.png](https://assets.leetcode.com/uploads/2019/10/10/sample22_1584.png)

    Input: startTime = [1,2,3,4,6], endTime = [3,5,10,6,9], profit = [20,20,100,70,60]
    Output: 150
    Explanation: The subset chosen is the first, fourth and fifth job. 
    Profit obtained 150 = 20 + 70 + 60.

**Example 3:**

![https://assets.leetcode.com/uploads/2019/10/10/sample3_1584.png](https://assets.leetcode.com/uploads/2019/10/10/sample3_1584.png)

    Input: startTime = [1,1,1], endTime = [2,3,4], profit = [5,6,4]
    Output: 6

**Constraints:**

- `1 <= startTime.length == endTime.length == profit.length <= 5 * 10^4`
- `1 <= startTime[i] < endTime[i] <= 10^9`
- `1 <= profit[i] <= 10^4`

## 题目大意


你打算利用空闲时间来做兼职工作赚些零花钱。这里有 n 份兼职工作，每份工作预计从 startTime[i] 开始到 endTime[i] 结束，报酬为 profit[i]。给你一份兼职工作表，包含开始时间 startTime，结束时间 endTime 和预计报酬 profit 三个数组，请你计算并返回可以获得的最大报酬。注意，时间上出现重叠的 2 份工作不能同时进行。如果你选择的工作在时间 X 结束，那么你可以立刻进行在时间 X 开始的下一份工作。


提示：

- 1 <= startTime.length == endTime.length == profit.length <= 5 * 10^4
- 1 <= startTime[i] < endTime[i] <= 10^9
- 1 <= profit[i] <= 10^4



## 解题思路

- 给出一组任务，任务有开始时间，结束时间，和任务收益。一个任务开始还没有结束，中间就不能再安排其他任务。问如何安排任务，能使得最后收益最大？
- 一般任务的题目，区间的题目，都会考虑是否能排序。这一题可以先按照任务的结束时间从小到大排序，如果结束时间相同，则按照收益从小到大排序。`dp[i]` 代表前 `i` 份工作能获得的最大收益。初始值，`dp[0] = job[1].profit` 。对于任意一个任务 `i` ，看能否找到满足 `jobs[j].enTime <= jobs[j].startTime && j < i` 条件的 `j`，即查找 `upper_bound` 。由于 `jobs` 被我们排序了，所以这里可以使用二分搜索来查找。如果能找到满足条件的任务 j，那么状态转移方程是：`dp[i] = max(dp[i-1], jobs[i].profit)`。如果能找到满足条件的任务 j，那么状态转移方程是：`dp[i] = max(dp[i-1], dp[low]+jobs[i].profit)`。最终求得的解在 `dp[len(startTime)-1]` 中。
