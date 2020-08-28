# [470. Implement Rand10() Using Rand7()](https://leetcode.com/problems/implement-rand10-using-rand7/)


## 题目

Given a function `rand7` which generates a uniform random integer in the range 1 to 7, write a function `rand10` which generates a uniform random integer in the range 1 to 10.

Do NOT use system's `Math.random()`.

**Example 1:**

    Input: 1
    Output: [7]

**Example 2:**

    Input: 2
    Output: [8,4]

**Example 3:**

    Input: 3
    Output: [8,1,10]

**Note:**

1. `rand7` is predefined.
2. Each testcase has one argument: `n`, the number of times that `rand10` is called.

**Follow up:**

1. What is the [expected value](https://en.wikipedia.org/wiki/Expected_value) for the number of calls to `rand7()` function?
2. Could you minimize the number of calls to `rand7()`?


## 题目大意

已有方法 rand7 可生成 1 到 7 范围内的均匀随机整数，试写一个方法 rand10 生成 1 到 10 范围内的均匀随机整数。不要使用系统的 Math.random() 方法。


提示:

- rand7 已定义。
- 传入参数: n 表示 rand10 的调用次数。
 

进阶:

- rand7()调用次数的 期望值 是多少 ?
- 你能否尽量少调用 rand7() ?



## 解题思路


- 给出 `rand7()` 要求实现 `rand10()`。
- `rand7()` 等概率地产生1，2，3，4，5，6，7。要想得到 `rand10()` 即等概率的生成 1-10 。解题思路是先构造一个 `randN()`，这个 N 必须是 10 的整数倍，然后 randN % 10 就可以得到 `rand10()` 了。所以可以从 `rand7()` 先构造出 `rand49()`，再把 `rand49()` 中大于等于 40 的都过滤掉，这样就得到了 `rand40()`，在对 10 取余即可。
- 具体构造步骤，`rand7() --> rand49() --> rand40() --> rand10()`：
    1. `rand7()` 等概率地产生 1,2,3,4,5,6,7.
    2. `rand7() - 1` 等概率地产生 [0,6].
    3. `(rand7() - 1) *7` 等概率地产生 0, 7, 14, 21, 28, 35, 42
    4. `(rand7() - 1) * 7 + (rand7() - 1)` 等概率地产生 [0, 48] 这 49 个数字
    5. 如果步骤 4 的结果大于等于 40，那么就重复步骤 4，直到产生的数小于 40
    6. 把步骤 5 的结果 mod 10 再加 1， 就会等概率的随机生成 [1, 10]
- 这道题可以推广到生成任意数的随机数问题。用 `randN()` 实现 `randM()`，`M>N`。步骤如下：
    1. 用 `randN()` 先实现 `randX()`，其中 X ≥ M，X 是 M 的整数倍。如这题中的 49 > 10；
    2. 再用 `randX()` 生成 `randM()`，如这题中的 49 —> 40 —> 10。
- 举个例子，用 `rand3()` 生成 `rand11()`，可以先生成 `rand27()`，然后变成以 22 为界限，因为 22 是 11 的倍数。生成 `rand27()` 的方式：`3 * 3 * (rand3() - 1) + 3 * (rand3() - 1) + (rand3() - 1)`，最后生成了 `rand11()`；用 `rand7()` 生成 `rand9()`，可以先生成 `rand49()`，然后变成以 45 为界限，因为 45 是 9 的倍数。生成 `rand49()` 的方式：`(rand7() - 1) * 7 + (rand7() - 1)`，最后生成了 `rand9()`；用 `rand6()` 生成 `rand13()`，可以先生成 `rand36()`，然后变成以 26 为界限，因为 26 是 13 的倍数。生成 `rand36()` 的方式：`(rand6() - 1) * 6 + (rand6() - 1)`，最后生成了 `rand13()`；
