# [207. Course Schedule](https://leetcode.com/problems/course-schedule/)

## 题目

There are a total of n courses you have to take, labeled from `0` to `n-1`.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: `[0,1]`

Given the total number of courses and a list of prerequisite **pairs**, is it possible for you to finish all courses?

**Example 1:**

    Input: 2, [[1,0]] 
    Output: true
    Explanation: There are a total of 2 courses to take. 
                 To take course 1 you should have finished course 0. So it is possible.

**Example 2:**

    Input: 2, [[1,0],[0,1]]
    Output: false
    Explanation: There are a total of 2 courses to take. 
                 To take course 1 you should have finished course 0, and to take course 0 you should
                 also have finished course 1. So it is impossible.

**Note:**

1. The input prerequisites is a graph represented by **a list of edges**, not adjacency matrices. Read more about [how a graph is represented](https://www.khanacademy.org/computing/computer-science/algorithms/graph-representation/a/representing-graphs).
2. You may assume that there are no duplicate edges in the input prerequisites.


## 题目大意

现在你总共有 n 门课需要选，记为 0 到 n-1。在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们: [0,1]。给定课程总量以及它们的先决条件，判断是否可能完成所有课程的学习？



## 解题思路

- 给出 n 个任务，每两个任务之间有相互依赖关系，比如 A 任务一定要在 B 任务之前完成才行。问是否可以完成所有任务。
- 这一题就是标准的 AOV 网的拓扑排序问题。拓扑排序问题的解决办法是主要是循环执行以下两步，直到不存在入度为0的顶点为止。
    - 1.  选择一个入度为0的顶点并输出之；
    - 2. 从网中删除此顶点及所有出边。

    循环结束后，若输出的顶点数小于网中的顶点数，则输出“有回路”信息，即无法完成所有任务；否则输出的顶点序列就是一种拓扑序列，即可以完成所有任务。

