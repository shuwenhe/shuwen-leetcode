# [1234. Replace the Substring for Balanced String](https://leetcode.com/problems/replace-the-substring-for-balanced-string/)


## 题目

You are given a string containing only 4 kinds of characters `'Q',` `'W', 'E'` and `'R'`.

A string is said to be **balanced** **if each of its characters appears `n/4` times where `n` is the length of the string.

Return the minimum length of the substring that can be replaced with **any** other string of the same length to make the original string `s` **balanced**.

Return 0 if the string is already **balanced**.

**Example 1**:

    Input: s = "QWER"
    Output: 0
    Explanation: s is already balanced.

**Example 2**:

    Input: s = "QQWE"
    Output: 1
    Explanation: We need to replace a 'Q' to 'R', so that "RQWE" (or "QRWE") is balanced.

**Example 3**:

    Input: s = "QQQW"
    Output: 2
    Explanation: We can replace the first "QQ" to "ER".

**Example 4**:

    Input: s = "QQQQ"
    Output: 3
    Explanation: We can replace the last 3 'Q' to make s = "QWER".

**Constraints**:

- `1 <= s.length <= 10^5`
- `s.length` is a multiple of `4`
- `s` contains only `'Q'`, `'W'`, `'E'` and `'R'`.

## 题目大意


有一个只含有 'Q', 'W', 'E', 'R' 四种字符，且长度为 n 的字符串。假如在该字符串中，这四个字符都恰好出现 n/4 次，那么它就是一个「平衡字符串」。给你一个这样的字符串 s，请通过「替换一个子串」的方式，使原字符串 s 变成一个「平衡字符串」。你可以用和「待替换子串」长度相同的 任何 其他字符串来完成替换。请返回待替换子串的最小可能长度。如果原字符串自身就是一个平衡字符串，则返回 0。

提示：

- 1 <= s.length <= 10^5
- s.length 是 4 的倍数
- s 中只含有 'Q', 'W', 'E', 'R' 四种字符



## 解题思路

- 给出一个字符串，要求输出把这个字符串变成“平衡字符串”的最小替换字符串的长度(替换只能替换一串，不能单个字母替换)。“平衡字符串”的定义是：字符串中，`‘Q’`，`‘W’`，`‘E’`，`‘R’`，出现的次数当且仅当只有 `len(s)/4` 次。
- 这一题是滑动窗口的题目。先统计 4 个字母的频次并计算出 `k = len(s)/4` 。滑动窗口向右滑动一次，对应右窗口的那个字母频次减 1，直到滑到所有字母的频次都 `≤ k`  的地方停止。此时，窗口外的字母的频次都 `≤ k` 了。这是只要变换窗口内字符串即可。但是这个窗口内还可能包含本来频次就小于 `k`  的字母，如果能把它们剔除掉，窗口可以进一步的减少。所以继续移动左边界，试探移动完左边界以后，是否所有字母频次都 `≤ k`。在所有窗口移动过程中取出最小值，即为最终答案。
- 举个例子：`"WQWRQQQW"`。`w` 有 3 个，`Q` 有 4 个，`R` 有 1 个，`E` 有 0 个。最后平衡状态是每个字母 2 个，那么我们需要拿出 1 个 `W` 和 2 个 `Q` 替换掉。即要找到一个最短的字符串包含 1 个 `W` 和 2 个 `Q`。滑动窗口正好可以解决这个问题。向右滑到 `"WQWRQ"` 停止，这时窗口外的所有字母频次都  `≤ k` 了。这个窗口包含了多余的 1 个 `W`，和 1 个 `R`。`W` 可以踢除掉，那么要替换的字符串是 `"QWRQ"`。`R` 不能踢除了(因为要找包含 1 个 `W` 和 2 个 `Q` 的字符串) 。窗口不断的滑动，直到结束。这个例子中最小的字符串其实位于末尾，`"QQW"`。

## 代码

```go

package leetcode

func balancedString(s string) int {
	count, k := make([]int, 128), len(s)/4
	for _, v := range s {
		count[int(v)]++
	}
	left, right, res := 0, -1, len(s)
	for left < len(s) {
		if count['Q'] > k || count['W'] > k || count['E'] > k || count['R'] > k {
			if right+1 < len(s) {
				right++
				count[s[right]]--
			} else {
				break
			}
		} else {
			res = min(res, right-left+1)
			count[s[left]]++
			left++
		}
	}
	return res
}

```