# [1040. Moving Stones Until Consecutive II](https://leetcode.com/problems/moving-stones-until-consecutive-ii/)


## 题目

On an **infinite** number line, the position of the i-th stone is given by `stones[i]`. Call a stone an *endpoint stone* if it has the smallest or largest position.

Each turn, you pick up an endpoint stone and move it to an unoccupied position so that it is no longer an endpoint stone.

In particular, if the stones are at say, `stones = [1,2,5]`, you **cannot** move the endpoint stone at position 5, since moving it to any position (such as 0, or 3) will still keep that stone as an endpoint stone.

The game ends when you cannot make any more moves, ie. the stones are in consecutive positions.

When the game ends, what is the minimum and maximum number of moves that you could have made? Return the answer as an length 2 array: `answer = [minimum_moves, maximum_moves]`

**Example 1:**

    Input: [7,4,9]
    Output: [1,2]
    Explanation: 
    We can move 4 -> 8 for one move to finish the game.
    Or, we can move 9 -> 5, 4 -> 6 for two moves to finish the game.

**Example 2:**

    Input: [6,5,4,3,10]
    Output: [2,3]
    We can move 3 -> 8 then 10 -> 7 to finish the game.
    Or, we can move 3 -> 7, 4 -> 8, 5 -> 9 to finish the game.
    Notice we cannot move 10 -> 2 to finish the game, because that would be an illegal move.

**Example 3:**

    Input: [100,101,104,102,103]
    Output: [0,0]

**Note:**

1. `3 <= stones.length <= 10^4`
2. `1 <= stones[i] <= 10^9`
3. `stones[i]` have distinct values.


## 题目大意

在一个长度无限的数轴上，第 i 颗石子的位置为 stones[i]。如果一颗石子的位置最小/最大，那么该石子被称作端点石子。每个回合，你可以将一颗端点石子拿起并移动到一个未占用的位置，使得该石子不再是一颗端点石子。值得注意的是，如果石子像 stones = [1,2,5] 这样，你将无法移动位于位置 5 的端点石子，因为无论将它移动到任何位置（例如 0 或 3），该石子都仍然会是端点石子。当你无法进行任何移动时，即，这些石子的位置连续时，游戏结束。

要使游戏结束，你可以执行的最小和最大移动次数分别是多少？ 以长度为 2 的数组形式返回答案：answer = [minimum\_moves, maximum\_moves] 。

提示：

1. 3 <= stones.length <= 10^4
2. 1 <= stones[i] <= 10^9
3. stones[i] 的值各不相同。


## 解题思路


- 给出一个数组，数组里面代表的是石头的坐标。要求移动石头，最终使得这些石头的坐标是一个连续的自然数列。但是规定，当一个石头是端点的时候，是不能移动的，例如 [1,2,5]，5 是端点，不能把 5 移到 3 或者 0 的位置，因为移动之后，这个石头仍然是端点。最终输出将所有石头排成连续的自然数列所需的最小步数和最大步数。
- 这道题的关键就是如何保证端点石头不能再次移动到端点的限制。例如，[5,6,8,9,20]，20 是端点，但是 20 就可以移动到 7 的位置，最终形成 [5,6,7,8,9] 的连续序列。但是 [5,6,7,8,20]，这种情况 20 就不能移动到 9 了，只能让 8 移动到 9，20 再移动到 8 的位置，最终还是形成了 [5,6,7,8,9]，但是步数需要 2 步。经过上述分析，可以得到，端点石头只能往中间空挡的位置移动，如果中间没有空挡，那么需要借助一个石头先制造一个空挡，然后端点石头再插入到中间，这样最少是需要 2 步。
- 再来考虑极值的情况。先看最大步数，最大步数肯定慢慢移动，一次移动一格，并且移动的格数最多。这里有两个极端情况，把数组里面的数全部都移动到最左端点，把数组里面的数全部都移动到最右端点。每次只移动一格。例如，全部都移到最右端点：

        [3,4,5,6,10] // 初始状态，连续的情况
        [4,5,6,7,10] // 第一步，把 3 挪到右边第一个可以插入的位置，即 7
        [5,6,7,8,10] // 第二步，把 4 挪到右边第一个可以插入的位置，即 8
        [6,7,8,9,10] // 第三步，把 5 挪到右边第一个可以插入的位置，即 9
        
        
        [1,3,5,7,10] // 初始状态，不连续的情况
        [3,4,5,7,10] // 第一步，把 1 挪到右边第一个可以插入的位置，即 4
        [4,5,6,7,10] // 第二步，把 3 挪到右边第一个可以插入的位置，即 6
        [5,6,7,8,10] // 第三步，把 4 挪到右边第一个可以插入的位置，即 8
        [6,7,8,9,10] // 第四步，把 5 挪到右边第一个可以插入的位置，即 9

    挪动的过程类似翻滚，最左边的石头挪到右边第一个可以放下的地方。然后不断的往右翻滚。把数组中的数全部都移动到最左边也同理。对比这两种情况的最大值，即是移动的最大步数。

- 再看最小步数。这里就涉及到了滑动窗口了。由于最终是要形成连续的自然数列，所以滑动窗口的大小已经固定成 n 了，从数组的 0 下标可以往右滑动窗口，这个窗口中能包含的数字越多，代表窗口外的数字越少，那么把这些数字放进窗口内的步数也最小。于是可以求得最小步数。这里有一个比较坑的地方就是题目中的那个`“端点不能移动以后还是端点”`的限制。针对这种情况，需要额外的判断。如果当前窗口内有 n-1 个元素了，即只有一个端点在窗口外，并且窗口右边界的值减去左边界的值也等于 n-1，代表这个窗口内已经都是连续数字了。这种情况端点想融合到这个连续数列中，最少需要 2 步(上文已经分析过了)。
- 注意一些边界情况。如果窗口从左往右滑动，窗口右边界滑到最右边了，但是窗口右边界的数字减去左边界的数字还是小于窗口大小 n，代表已经滑到头了，可以直接 break 出去。为什么滑到头了呢？由于数组经过从小到大排序以后，数字越往右边越大，当前数字是小值，已经满足了 `stones[right]-stones[left] < n`，左边界继续往右移动只会使得 `stones[left]` 更大，就更加小于 n 了。而我们需要寻找的是 `stones[right]-stones[left] >= n` 的边界点，肯定再也找不到了。
