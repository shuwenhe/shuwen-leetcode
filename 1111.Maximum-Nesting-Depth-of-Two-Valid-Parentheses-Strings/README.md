# [1111. Maximum Nesting Depth of Two Valid Parentheses Strings](https://leetcode.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/)


## 题目

A string is a *valid parentheses string* (denoted VPS) if and only if it consists of `"("` and `")"` characters only, and:

- It is the empty string, or
- It can be written as `AB` (`A` concatenated with `B`), where `A` and `B` are VPS's, or
- It can be written as `(A)`, where `A` is a VPS.

We can similarly define the *nesting depth* `depth(S)` of any VPS `S` as follows:

- `depth("") = 0`
- `depth(A + B) = max(depth(A), depth(B))`, where `A` and `B` are VPS's
- `depth("(" + A + ")") = 1 + depth(A)`, where `A` is a VPS.

For example, `""`, `"()()"`, and `"()(()())"` are VPS's (with nesting depths 0, 1, and 2), and `")("` and `"(()"` are not VPS's.

Given a VPS seq, split it into two disjoint subsequences `A` and `B`, such that `A` and `B` are VPS's (and `A.length + B.length = seq.length`).

Now choose **any** such `A` and `B` such that `max(depth(A), depth(B))` is the minimum possible value.

Return an `answer` array (of length `seq.length`) that encodes such a choice of `A` and `B`: `answer[i] = 0` if `seq[i]` is part of `A`, else `answer[i] = 1`. Note that even though multiple answers may exist, you may return any of them.

**Example 1:**

    Input: seq = "(()())"
    Output: [0,1,1,1,1,0]

**Example 2:**

    Input: seq = "()(())()"
    Output: [0,0,0,1,1,0,1,1]

**Constraints:**

- `1 <= seq.size <= 10000`


## 题目大意


有效括号字符串 仅由 "(" 和 ")" 构成，并符合下述几个条件之一：

- 空字符串
- 连接，可以记作 AB（A 与 B 连接），其中 A 和 B 都是有效括号字符串
- 嵌套，可以记作 (A)，其中 A 是有效括号字符串

类似地，我们可以定义任意有效括号字符串 s 的 嵌套深度 depth(S)：

- s 为空时，depth("") = 0
- s 为 A 与 B 连接时，depth(A + B) = max(depth(A), depth(B))，其中 A 和 B 都是有效括号字符串
- s 为嵌套情况，depth("(" + A + ")") = 1 + depth(A)，其中 A 是有效括号字符串


例如：""，"()()"，和 "()(()())" 都是有效括号字符串，嵌套深度分别为 0，1，2，而 ")(" 和 "(()" 都不是有效括号字符串。

 

给你一个有效括号字符串 seq，将其分成两个不相交的子序列 A 和 B，且 A 和 B 满足有效括号字符串的定义（注意：A.length + B.length = seq.length）。

现在，你需要从中选出 任意 一组有效括号字符串 A 和 B，使 max(depth(A), depth(B)) 的可能取值最小。

返回长度为 seq.length 答案数组 answer ，选择 A 还是 B 的编码规则是：如果 seq[i] 是 A 的一部分，那么 answer[i] = 0。否则，answer[i] = 1。即便有多个满足要求的答案存在，你也只需返回 一个。



## 解题思路

- 给出一个括号字符串。选出 A 部分和 B 部分，使得 `max(depth(A), depth(B))` 值最小。在最终的数组中输出 0 和 1，0 标识是 A 部分，1 标识是 B 部分。
- 这一题想要 `max(depth(A), depth(B))` 值最小，可以使用贪心思想。如果 A 部分和 B 部分都尽快括号匹配，不深层次嵌套，那么总的层次就会变小。只要让嵌套的括号中属于 A 的和属于 B 的间隔排列即可。例如：“`(((())))`”，上面的字符串的嵌套深度是 4，按照上述的贪心思想，则标记为 0101 1010。
- 这一题也可以用二分的思想来解答。把深度平分给 A 部分和 B 部分。
    - 第一次遍历，先计算最大深度
    - 第二次遍历，把深度小于等于最大深度一半的括号标记为 0(给 A 部分)，否则标记为 1(给 B 部分)
