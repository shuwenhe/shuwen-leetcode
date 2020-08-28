# [636. Exclusive Time of Functions](https://leetcode.com/problems/exclusive-time-of-functions/)


## 题目

On a single threaded CPU, we execute some functions. Each function has a unique id between `0` and `N-1`.

We store logs in timestamp order that describe when a function is entered or exited.

Each log is a string with this format: `"{function_id}:{"start" | "end"}:{timestamp}"`. For example, `"0:start:3"` means the function with id `0` started at the beginning of timestamp `3`. `"1:end:2"` means the function with id `1` ended at the end of timestamp `2`.

A function's *exclusive time* is the number of units of time spent in this function. Note that this does not include any recursive calls to child functions.

Return the exclusive time of each function, sorted by their function id.

Example 1:

![](https://assets.leetcode.com/uploads/2019/04/05/diag1b.png)

    Input:
    n = 2
    logs = ["0:start:0","1:start:2","1:end:5","0:end:6"]
    Output: [3, 4]
    Explanation:
    Function 0 starts at the beginning of time 0, then it executes 2 units of time and reaches the end of time 1.
    Now function 1 starts at the beginning of time 2, executes 4 units of time and ends at time 5.
    Function 0 is running again at the beginning of time 6, and also ends at the end of time 6, thus executing for 1 unit of time. 
    So function 0 spends 2 + 1 = 3 units of total time executing, and function 1 spends 4 units of total time executing.

**Note:**

1. `1 <= n <= 100`
2. Two functions won't start or end at the same time.
3. Functions will always log when they exit.



## 题目大意

给出一个非抢占单线程CPU的 n 个函数运行日志，找到函数的独占时间。每个函数都有一个唯一的 Id，从 0 到 n-1，函数可能会递归调用或者被其他函数调用。

日志是具有以下格式的字符串：function_id：start_or_end：timestamp。例如："0:start:0" 表示函数 0 从 0 时刻开始运行。"0:end:0" 表示函数 0 在 0 时刻结束。

函数的独占时间定义是在该方法中花费的时间，调用其他函数花费的时间不算该函数的独占时间。你需要根据函数的 Id 有序地返回每个函数的独占时间。


## 解题思路


- 利用栈记录每一个开始了但是未完成的任务，完成以后任务就 pop 一个。
- 注意题目中关于任务时长的定义，例如，start 7，end 7，这个任务执行了 1 秒而不是 0 秒
