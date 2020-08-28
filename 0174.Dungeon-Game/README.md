# [174. Dungeon Game](https://leetcode.com/problems/dungeon-game/)

## 题目

The demons had captured the princess (**P**) and imprisoned her in the bottom-right corner of a dungeon. The dungeon consists of M x N rooms laid out in a 2D grid. Our valiant knight (**K**) was initially positioned in the top-left room and must fight his way through the dungeon to rescue the princess.

The knight has an initial health point represented by a positive integer. If at any point his health point drops to 0 or below, he dies immediately.

Some of the rooms are guarded by demons, so the knight loses health (*negative* integers) upon entering these rooms; other rooms are either empty (*0's*) or contain magic orbs that increase the knight's health (*positive* integers).

In order to reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.

**Write a function to determine the knight's minimum initial health so that he is able to rescue the princess.**

For example, given the dungeon below, the initial health of the knight must be at least **7** if he follows the optimal path `RIGHT-> RIGHT -> DOWN -> DOWN`.


![](https://img.halfrost.com/Leetcode/leetcode_174_0.png)

**Note:**

- The knight's health has no upper bound.
- Any room can contain threats or power-ups, even the first room the knight enters and the bottom-right room where the princess is imprisoned.


## 题目大意

一些恶魔抓住了公主（P）并将她关在了地下城的右下角。地下城是由 M x N 个房间组成的二维网格。我们英勇的骑士（K）最初被安置在左上角的房间里，他必须穿过地下城并通过对抗恶魔来拯救公主。

骑士的初始健康点数为一个正整数。如果他的健康点数在某一时刻降至 0 或以下，他会立即死亡。

有些房间由恶魔守卫，因此骑士在进入这些房间时会失去健康点数（若房间里的值为负整数，则表示骑士将损失健康点数）；其他房间要么是空的（房间里的值为 0），要么包含增加骑士健康点数的魔法球（若房间里的值为正整数，则表示骑士将增加健康点数）。

为了尽快到达公主，骑士决定每次只向右或向下移动一步。编写一个函数来计算确保骑士能够拯救到公主所需的最低初始健康点数。

说明:

- 骑士的健康点数没有上限。
- 任何房间都可能对骑士的健康点数造成威胁，也可能增加骑士的健康点数，包括骑士进入的左上角房间以及公主被监禁的右下角房间。

## 解题思路

- 在二维地图上给出每个格子扣血数，负数代表扣血，正数代表补血。左上角第一个格子是起点，右下角最后一个格子是终点。问骑士初始最少多少血才能走完迷宫，顺利营救位于终点的公主。需要注意的是，起点和终点都会对血量进行影响。每到一个格子，骑士的血都不能少于 1，一旦少于 1 点血，骑士就会死去。
- 这一题首先想到的解题思路是动态规划。从终点逆推回起点。`dp[i][j]` 代表骑士进入坐标为 `(i,j)` 的格子之前最少的血量值。 那么 `dp[m-1][n-1]` 应该同时满足两个条件，`dp[m-1][n-1] + dungeon[m-1][n-1] ≥ 1` 并且 `dp[m-1][n-1] ≥ 1`，由于这两个不等式的方向是相同的，取交集以后，起决定作用的是数轴最右边的数，即 `max(1-dungeon[m-1][n-1] , 1)`。算出 `dp[m-1][n-1]` 以后，接着可以推出 `dp[m-1][i]` 这一行和 `dp[i][n-1]` 这一列的值。因为骑士只能往右走和往下走。往回推，即只能往上走和往左走。到这里，DP 的初始条件都准备好了。那么状态转移方程是什么呢？分析一般的情况，`dp[i][j]` 这个值应该是和 `dp[i+1][j]` 和 `dp[i][j+1]` 这两者有关系。即 `dp[i][j]` 经过自己本格子的扣血以后，要能至少满足下一行和右一列格子血量的最少要求。并且自己的血量也应该 `≥1`。即需要满足下面这两组不等式。   
	 ![](https://img.halfrost.com/Leetcode/leetcode_174_1.png)  
	 ![](https://img.halfrost.com/Leetcode/leetcode_174_2.png)    
    上面不等式中第一组不等式是满足下一行格子的最低血量要求，第二组不等式是满足右一列格子的最低血量要求。第一个式子化简即 `dp[i][j] = max(1, dp[i+1][j]-dungeon[i][j])`，第二个式子化简即 `dp[i][j] = max(1, dp[i][j+1]-dungeon[i][j])`。求得了这两种走法的最低血量值，从这两个值里面取最小，即是当前格子所需的最低血量，所以状态转移方程为 `dp[i][j] = min(max(1, dp[i][j+1]-dungeon[i][j]), max(1, dp[i+1][j]-dungeon[i][j]))`。DP 完成以后，`dp[0][0]` 中记录的就是骑士初始最低血量值。时间复杂度 O(m\*n)，空间复杂度 O(m\*n)。

- 这一题还可以用二分搜索来求解。骑士的血量取值范围一定是在 `[1，+∞)` 这个区间内。那么二分这个区间，每次二分的中间值，再用 dp 在地图中去判断是否能到达终点，如果能，就缩小搜索空间至 `[1，mid]`，否则搜索空间为 `[mid + 1，+∞)` 。时间复杂度 O(m\*n\* log math.MaxInt64)，空间复杂度 O(m\*n)。
