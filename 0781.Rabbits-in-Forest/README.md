# [781. Rabbits in Forest](https://leetcode.com/problems/rabbits-in-forest/)


## 题目

In a forest, each rabbit has some color. Some subset of rabbits (possibly all of them) tell you how many other rabbits have the same color as them. Those `answers` are placed in an array.

Return the minimum number of rabbits that could be in the forest.

    Examples:
    Input: answers = [1, 1, 2]
    Output: 5
    Explanation:
    The two rabbits that answered "1" could both be the same color, say red.
    The rabbit than answered "2" can't be red or the answers would be inconsistent.
    Say the rabbit that answered "2" was blue.
    Then there should be 2 other blue rabbits in the forest that didn't answer into the array.
    The smallest possible number of rabbits in the forest is therefore 5: 3 that answered plus 2 that didn't.
    
    Input: answers = [10, 10, 10]
    Output: 11
    
    Input: answers = []
    Output: 0

**Note:**

1. `answers` will have length at most `1000`.
2. Each `answers[i]` will be an integer in the range `[0, 999]`.


## 题目大意

森林中，每个兔子都有颜色。其中一些兔子（可能是全部）告诉你还有多少其他的兔子和自己有相同的颜色。我们将这些回答放在 answers 数组里。返回森林中兔子的最少数量。

说明:

- answers 的长度最大为1000。
- answers[i] 是在 [0, 999] 范围内的整数。


## 解题思路


- 给出一个数组，数组里面代表的是每个兔子说自己同类还有多少个。要求输出总共有多少只兔子。数字中可能兔子汇报的人数小于总兔子数。
- 这一题关键在于如何划分不同种类的兔子，有可能相同种类的兔子的个数是一样的，比如 `[2,2,2,2,2,2]`，这其实是 3 个种类，总共 6 只兔子。用 map 去重相同种类的兔子，不断的减少，当有种类的兔子为 0 以后，还有该种类的兔子报数，需要当做另外一个种类的兔子来看待。
