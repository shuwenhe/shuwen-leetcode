# [1275. Find Winner on a Tic Tac Toe Game](https://leetcode.com/problems/find-winner-on-a-tic-tac-toe-game/)


## 题目

Tic-tac-toe is played by two players *A* and *B* on a 3 x 3 grid.

Here are the rules of Tic-Tac-Toe:

- Players take turns placing characters into empty squares (" ").
- The first player *A* always places "X" characters, while the second player *B* always places "O" characters.
- "X" and "O" characters are always placed into empty squares, never on filled ones.
- The game ends when there are 3 of the same (non-empty) character filling any row, column, or diagonal.
- The game also ends if all squares are non-empty.
- No more moves can be played if the game is over.

Given an array `moves` where each element is another array of size 2 corresponding to the row and column of the grid where they mark their respective character in the order in which *A* and *B* play.

Return the winner of the game if it exists (*A* or *B*), in case the game ends in a draw return "Draw", if there are still movements to play return "Pending".

You can assume that `moves` is **valid** (It follows the rules of Tic-Tac-Toe), the grid is initially empty and *A* will play **first**.

**Example 1**:

```
Input: moves = [[0,0],[2,0],[1,1],[2,1],[2,2]]
Output: "A"
Explanation: "A" wins, he always plays first.
"X  "    "X  "    "X  "    "X  "    "X  "
"   " -> "   " -> " X " -> " X " -> " X "
"   "    "O  "    "O  "    "OO "    "OOX"

```

**Example 2**:

```
Input: moves = [[0,0],[1,1],[0,1],[0,2],[1,0],[2,0]]
Output: "B"
Explanation: "B" wins.
"X  "    "X  "    "XX "    "XXO"    "XXO"    "XXO"
"   " -> " O " -> " O " -> " O " -> "XO " -> "XO " 
"   "    "   "    "   "    "   "    "   "    "O  "

```

**Example 3**:

```
Input: moves = [[0,0],[1,1],[2,0],[1,0],[1,2],[2,1],[0,1],[0,2],[2,2]]
Output: "Draw"
Explanation: The game ends in a draw since there are no moves to make.
"XXO"
"OOX"
"XOX"

```

**Example 4**:

```
Input: moves = [[0,0],[1,1]]
Output: "Pending"
Explanation: The game has not finished yet.
"X  "
" O "
"   "

```

**Constraints:**

- `1 <= moves.length <= 9`
- `moves[i].length == 2`
- `0 <= moves[i][j] <= 2`
- There are no repeated elements on `moves`.
- `moves` follow the rules of tic tac toe.


## 题目大意

A 和 B 在一个 3 x 3 的网格上玩井字棋。井字棋游戏的规则如下：

- 玩家轮流将棋子放在空方格 (" ") 上。
- 第一个玩家 A 总是用 "X" 作为棋子，而第二个玩家 B 总是用 "O" 作为棋子。
- "X" 和 "O" 只能放在空方格中，而不能放在已经被占用的方格上。
- 只要有 3 个相同的（非空）棋子排成一条直线（行、列、对角线）时，游戏结束。
- 如果所有方块都放满棋子（不为空），游戏也会结束。
- 游戏结束后，棋子无法再进行任何移动。

给你一个数组 moves，其中每个元素是大小为 2 的另一个数组（元素分别对应网格的行和列），它按照 A 和 B 的行动顺序（先 A 后 B）记录了两人各自的棋子位置。如果游戏存在获胜者（A 或 B），就返回该游戏的获胜者；如果游戏以平局结束，则返回 "Draw"；如果仍会有行动（游戏未结束），则返回 "Pending"。你可以假设 moves 都 有效（遵循井字棋规则），网格最初是空的，A 将先行动。

提示：

- 1 <= moves.length <= 9
- moves[i].length == 2
- 0 <= moves[i][j] <= 2
- moves 里没有重复的元素。
- moves 遵循井字棋的规则。


## 解题思路

- 两人玩 3*3 井字棋，A 先走，B 再走。谁能获胜就输出谁，如果平局输出 “Draw”，如果游戏还未结束，输出 “Pending”。游戏规则：谁能先占满行、列或者对角线任意一条线，谁就赢。
- 简单题。题目给定 move 数组最多 3 步，而要赢得比赛，必须走满 3 步，所以可以先模拟，按照给的步数数组把 A 和 B 的步数都放在棋盘上。然后依次判断行、列，对角线的三种情况。如果都判完了，剩下的情况就是平局和死局的情况。

## 代码

```go
func tictactoe(moves [][]int) string {
	board := [3][3]byte{}
	for i := 0; i < len(moves); i++ {
		if i%2 == 0 {
			board[moves[i][0]][moves[i][1]] = 'X'
		} else {
			board[moves[i][0]][moves[i][1]] = 'O'
		}
	}
	for i := 0; i < 3; i++ {
		if board[i][0] == 'X' && board[i][1] == 'X' && board[i][2] == 'X' {
			return "A"
		}
		if board[i][0] == 'O' && board[i][1] == 'O' && board[i][2] == 'O' {
			return "B"
		}
		if board[0][i] == 'X' && board[1][i] == 'X' && board[2][i] == 'X' {
			return "A"
		}
		if board[0][i] == 'O' && board[1][i] == 'O' && board[2][i] == 'O' {
			return "B"
		}
	}
	if board[0][0] == 'X' && board[1][1] == 'X' && board[2][2] == 'X' {
		return "A"
	}
	if board[0][0] == 'O' && board[1][1] == 'O' && board[2][2] == 'O' {
		return "B"
	}
	if board[0][2] == 'X' && board[1][1] == 'X' && board[2][0] == 'X' {
		return "A"
	}
	if board[0][2] == 'O' && board[1][1] == 'O' && board[2][0] == 'O' {
		return "B"
	}
	if len(moves) < 9 {
		return "Pending"
	}
	return "Draw"
}
```