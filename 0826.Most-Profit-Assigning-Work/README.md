# [826. Most Profit Assigning Work](https://leetcode.com/problems/most-profit-assigning-work/)

## 题目

We have jobs: difficulty[i] is the difficulty of the ith job, and profit[i] is the profit of the ith job. 

Now we have some workers. worker[i] is the ability of the ith worker, which means that this worker can only complete a job with difficulty at most worker[i]. 

Every worker can be assigned at most one job, but one job can be completed multiple times.

For example, if 3 people attempt the same job that pays $1, then the total profit will be $3.  If a worker cannot complete any job, his profit is $0.

What is the most profit we can make?


Example 1:


```c
Input: difficulty = [2,4,6,8,10], profit = [10,20,30,40,50], worker = [4,5,6,7]
Output: 100 
Explanation: Workers are assigned jobs of difficulty [4,4,6,6] and they get profit of [20,20,30,30] seperately.
```

Note:

- 1 <= difficulty.length = profit.length <= 10000
- 1 <= worker.length <= 10000
- difficulty[i], profit[i], worker[i]  are in range [1, 10^5]


## 题目大意

这道题考察的是滑动窗口的问题，也是排序相关的问题。

给出一组任务，每个任务都有一定的难度，每个任务也都有完成以后对应的收益(完成难的任务不一定收益最高)。有一批工人，每个人能处理的任务难度不同。要求输出这批工人完成任务以后的最大收益。

## 解题思路

先将任务按照难度排序，工人也按照能处理任务难度的能力排序。用一个数组记录下，每个 i 下标，当前能达到的最大收益。计算这个收益只需要从下标为 1 开始，依次比较自己和前一个的收益即可(因为排过序，难度是依次递增的)。有了这个难度依次递增，并且记录了最大收益的数组以后，就可以计算最终结果了。遍历一遍工人数组，如果工人的能力大于任务的难度，就加上这个最大收益。遍历完工人数组，最终结果就是最大收益。


