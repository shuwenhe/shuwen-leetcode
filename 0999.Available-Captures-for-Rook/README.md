# [999. Available Captures for Rook](https://leetcode.com/problems/available-captures-for-rook/)


## 题目

On an 8 x 8 chessboard, there is one white rook. There also may be empty squares, white bishops, and black pawns. These are given as characters 'R', '.', 'B', and 'p' respectively. Uppercase characters represent white pieces, and lowercase characters represent black pieces.

The rook moves as in the rules of Chess: it chooses one of four cardinal directions (north, east, west, and south), then moves in that direction until it chooses to stop, reaches the edge of the board, or captures an opposite colored pawn by moving to the same square it occupies. Also, rooks cannot move into the same square as other friendly bishops.

Return the number of pawns the rook can capture in one move.

**Example 1**:

![https://assets.leetcode.com/uploads/2019/02/20/1253_example_1_improved.PNG](https://assets.leetcode.com/uploads/2019/02/20/1253_example_1_improved.PNG)

```
Input: [[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".","R",".",".",".","p"],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."]]
Output: 3
Explanation: 
In this example the rook is able to capture all the pawns.
```

**Example 2**:

![https://assets.leetcode.com/uploads/2019/02/19/1253_example_2_improved.PNG](https://assets.leetcode.com/uploads/2019/02/19/1253_example_2_improved.PNG)

```
Input: [[".",".",".",".",".",".",".","."],[".","p","p","p","p","p",".","."],[".","p","p","B","p","p",".","."],[".","p","B","R","B","p",".","."],[".","p","p","B","p","p",".","."],[".","p","p","p","p","p",".","."],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."]]
Output: 0
Explanation: 
Bishops are blocking the rook to capture any pawn.
```

**Example 3**:

![https://assets.leetcode.com/uploads/2019/02/20/1253_example_3_improved.PNG](https://assets.leetcode.com/uploads/2019/02/20/1253_example_3_improved.PNG)

```
Input: [[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".","p",".",".",".","."],["p","p",".","R",".","p","B","."],[".",".",".",".",".",".",".","."],[".",".",".","B",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".",".",".",".",".","."]]
Output: 3
Explanation: 
The rook can capture the pawns at positions b5, d6 and f5.
```

**Note**:

1. `board.length == board[i].length == 8`
2. `board[i][j]` is either `'R'`, `'.'`, `'B'`, or `'p'`
3. There is exactly one cell with `board[i][j] == 'R'`

## 题目大意

在一个 8 x 8 的棋盘上，有一个白色的车（Rook），用字符 'R' 表示。棋盘上还可能存在空方块，白色的象（Bishop）以及黑色的卒（pawn），分别用字符 '.'，'B' 和 'p' 表示。不难看出，大写字符表示的是白棋，小写字符表示的是黑棋。车按国际象棋中的规则移动。东，西，南，北四个基本方向任选其一，然后一直向选定的方向移动，直到满足下列四个条件之一：

- 棋手选择主动停下来。
- 棋子因到达棋盘的边缘而停下。
- 棋子移动到某一方格来捕获位于该方格上敌方（黑色）的卒，停在该方格内。
- 车不能进入/越过已经放有其他友方棋子（白色的象）的方格，停在友方棋子前。

你现在可以控制车移动一次，请你统计有多少敌方的卒处于你的捕获范围内（即，可以被一步捕获的棋子数）。

## 解题思路

- 按照国际象棋的规则移动车，要求输出只移动一次，有多少个卒在车的捕获范围之内
- 简单题，按照国际象棋车的移动规则， 4 个方向分别枚举即可。

## 代码

```go
package leetcode

func numRookCaptures(board [][]byte) int {
	num := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'R' {
				num += caputure(board, i-1, j, -1, 0) // Up
				num += caputure(board, i+1, j, 1, 0)  // Down
				num += caputure(board, i, j-1, 0, -1) // Left
				num += caputure(board, i, j+1, 0, 1)  // Right
			}
		}
	}
	return num
}

func caputure(board [][]byte, x, y int, bx, by int) int {
	for x >= 0 && x < len(board) && y >= 0 && y < len(board[x]) && board[x][y] != 'B' {
		if board[x][y] == 'p' {
			return 1
		}
		x += bx
		y += by
	}
	return 0
}
```