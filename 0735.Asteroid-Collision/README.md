# [735. Asteroid Collision](https://leetcode.com/problems/asteroid-collision/)

## 题目

We are given an array asteroids of integers representing asteroids in a row.

For each asteroid, the absolute value represents its size, and the sign represents its direction (positive meaning right, negative meaning left). Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions. If two asteroids meet, the smaller one will explode. If both are the same size, both will explode. Two asteroids moving in the same direction will never meet.

Example 1:

```c
Input: 
asteroids = [5, 10, -5]
Output: [5, 10]
Explanation: 
The 10 and -5 collide resulting in 10.  The 5 and 10 never collide.
```

Example 2:

```c
Input: 
asteroids = [8, -8]
Output: []
Explanation: 
The 8 and -8 collide exploding each other.
```

Example 3:

```c
Input: 
asteroids = [10, 2, -5]
Output: [10]
Explanation: 
The 2 and -5 collide resulting in -5.  The 10 and -5 collide resulting in 10.
```

Example 4:

```c
Input: 
asteroids = [-2, -1, 1, 2]
Output: [-2, -1, 1, 2]
Explanation: 
The -2 and -1 are moving left, while the 1 and 2 are moving right.
Asteroids moving the same direction never meet, so no asteroids will meet each other.
```

Note:

- The length of asteroids will be at most 10000.
- Each asteroid will be a non-zero integer in the range [-1000, 1000]..

## 题目大意

给定一个整数数组 asteroids，表示在同一行的行星。对于数组中的每一个元素，其绝对值表示行星的大小，正负表示行星的移动方向（正表示向右移动，负表示向左移动）。每一颗行星以相同的速度移动。找出碰撞后剩下的所有行星。碰撞规则：两个行星相互碰撞，较小的行星会爆炸。如果两颗行星大小相同，则两颗行星都会爆炸。两颗移动方向相同的行星，永远不会发生碰撞。

## 解题思路

这一题类似第 1047 题。这也是一个类似“对对碰”的游戏，不过这里的碰撞，大行星和小行星碰撞以后，大行星会胜出，小行星直接消失。按照题意的规则来，用栈模拟即可。考虑最终结果：

1. 所有向左飞的行星都向左，所有向右飞的行星都向右。
2. 向左飞的行星，如果飞行中没有向右飞行的行星，那么它将安全穿过。
3. 跟踪所有向右移动到右侧的行星，最右边的一个将是第一个面对向左飞行行星碰撞的。
4. 如果它幸存下来，继续前进，否则，任何之前的向右的行星都会被逐一被暴露出来碰撞。

所以先处理这种情况，一层循环把所有能碰撞的向右飞行的行星都碰撞完。碰撞完以后，如果栈顶行星向左飞，新来的行星向右飞，直接添加进来即可。否则栈顶行星向右飞，大小和向左飞的行星一样大小，两者都撞毁灭，弹出栈顶元素。



