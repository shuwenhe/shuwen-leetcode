# [756. Pyramid Transition Matrix](https://leetcode.com/problems/pyramid-transition-matrix/)


## 题目

We are stacking blocks to form a pyramid. Each block has a color which is a one letter string.

We are allowed to place any color block `C` on top of two adjacent blocks of colors `A` and `B`, if and only if `ABC` is an allowed triple.

We start with a bottom row of `bottom`, represented as a single string. We also start with a list of allowed triples `allowed`. Each allowed triple is represented as a string of length 3.

Return true if we can build the pyramid all the way to the top, otherwise false.

**Example 1:**

    Input: bottom = "BCD", allowed = ["BCG", "CDE", "GEA", "FFF"]
    Output: true
    Explanation:
    We can stack the pyramid like this:
        A
       / \
      G   E
     / \ / \
    B   C   D
    
    We are allowed to place G on top of B and C because BCG is an allowed triple.  Similarly, we can place E on top of C and D, then A on top of G and E.

**Example 2:**

    Input: bottom = "AABA", allowed = ["AAA", "AAB", "ABA", "ABB", "BAC"]
    Output: false
    Explanation:
    We can't stack the pyramid to the top.
    Note that there could be allowed triples (A, B, C) and (A, B, D) with C != D.

**Note:**

1. `bottom` will be a string with length in range `[2, 8]`.
2. `allowed` will have length in range `[0, 200]`.
3. Letters in all strings will be chosen from the set `{'A', 'B', 'C', 'D', 'E', 'F', 'G'}`.


## 题目大意

现在，我们用一些方块来堆砌一个金字塔。 每个方块用仅包含一个字母的字符串表示，例如 “Z”。使用三元组表示金字塔的堆砌规则如下：

(A, B, C) 表示，“C” 为顶层方块，方块 “A”、“B” 分别作为方块 “C” 下一层的的左、右子块。当且仅当(A, B, C)是被允许的三元组，我们才可以将其堆砌上。

初始时，给定金字塔的基层 bottom，用一个字符串表示。一个允许的三元组列表 allowed，每个三元组用一个长度为 3 的字符串表示。如果可以由基层一直堆到塔尖返回 true，否则返回 false。



## 解题思路

- 这一题是一道 DFS 的题目。题目给出金字塔的底座字符串。然后还会给一个字符串数组，字符串数组里面代表的字符串的砖块。砖块是 3 个字符串组成的。前两个字符代表的是砖块的底边，后一个字符代表的是砖块的顶部。问给出的字符能拼成一个金字塔么？金字塔的特点是顶端就一个字符。  

- 这一题用 DFS 深搜每个砖块，从底层砖块开始逐渐往上层码。每递归一层，新一层底部的砖块都会变。当递归到了一层底部只有 2 个字符，顶部只有一个字符的时候，就到金字塔顶端了，就算是完成了。这一题为了挑选合适的砖块，需要把每个砖块底部的 2 个字符作为 key 放进 map 中，加速查找。题目中也给出了特殊情况，相同底部可能存在多种砖块，所以一个 key 可能对应多个 value 的情况，即可能存在多个顶部砖块的情况。这种情况在递归遍历中需要考虑。
