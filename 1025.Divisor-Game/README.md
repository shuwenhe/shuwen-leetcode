# [1025. Divisor Game](https://leetcode.com/problems/divisor-game/)


## 题目

Alice and Bob take turns playing a game, with Alice starting first.

Initially, there is a number `N` on the chalkboard. On each player's turn, that player makes a *move* consisting of:

- Choosing any `x` with `0 < x < N` and `N % x == 0`.
- Replacing the number `N` on the chalkboard with `N - x`.

Also, if a player cannot make a move, they lose the game.

Return `True` if and only if Alice wins the game, assuming both players play optimally.

**Example 1:**

    Input: 2
    Output: true
    Explanation: Alice chooses 1, and Bob has no more moves.

**Example 2:**

    Input: 3
    Output: false
    Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more moves.

**Note:**

1. `1 <= N <= 1000`


## 题目大意


爱丽丝和鲍勃一起玩游戏，他们轮流行动。爱丽丝先手开局。最初，黑板上有一个数字 N 。在每个玩家的回合，玩家需要执行以下操作：

- 选出任一 x，满足 0 < x < N 且 N % x == 0 。
- 用 N - x 替换黑板上的数字 N 。

如果玩家无法执行这些操作，就会输掉游戏。只有在爱丽丝在游戏中取得胜利时才返回 True，否则返回 false。假设两个玩家都以最佳状态参与游戏。


## 解题思路


- 两人相互玩一个游戏，游戏初始有一个数 N，开始游戏的时候，任一方选择一个数 x，满足 `0 < x < N` 并且 `N % x == 0` 的条件，然后 `N-x` 为下一轮开始的数。此轮结束，该另外一个人继续选择数字，两人相互轮流选择。直到某一方再也没法选择数字的时候，输掉游戏。问如果你先手开始游戏，给出 N 的时候，能否直到这局你是否会必胜或者必输？
- 这一题当 `N = 1` 的时候，那一轮的人必输。因为没法找到一个数字能满足 `0 < x < N` 并且 `N % x == 0` 的条件了。必胜策略就是把对方逼至 `N = 1` 的情况。题目中假设了对手也是一个很有头脑的人。初始如果 `N 为偶数`，我就选择 x = 1，对手拿到的数字就是奇数。只要最终能让对手拿到奇数，他就会输。初始如果 `N 为奇数`，N = 1 的时候直接输了，N 为其他奇数的时候，我们也只能选择一个奇数 x，(因为 `N % x == 0` ，N 为奇数，x 一定不会是偶数，因为偶数就能被 2 整除了)，对手由于是一个很有头脑的人，当我们选完 N - x 是偶数的时候，他就选择 1，那么轮到我们拿到的数字又是奇数。对手只要一直保证我们拿到奇数，最终肯定会逼着我们拿到 1，最终他就会获得胜利。所以经过分析可得，初始数字如果是偶数，有必胜策略，如果初始数字是奇数，有必输的策略。
