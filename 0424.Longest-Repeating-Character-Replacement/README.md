# [424. Longest Repeating Character Replacement](https://leetcode.com/problems/longest-repeating-character-replacement/)

## 题目


Given a string that consists of only uppercase English letters, you can replace any letter in the string with another letter at most k times. Find the length of a longest substring containing all repeating letters you can get after performing the above operations.

Note:  

Both the string's length and k will not exceed 10^4.

Example 1:

```c
Input:
s = "ABAB", k = 2

Output:
4

Explanation:
Replace the two 'A's with two 'B's or vice versa.
```

Example 2:

```c
Input:
s = "AABABBA", k = 1

Output:
4

Explanation:
Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
```



## 题目大意


给一个字符串和变换次数 K，要求经过 K 次字符转换以后，输出相同字母能出现连续最长的长度。


## 解题思路

这道题笔者也提交了好几遍才通过。这一题是考察滑动窗口的题目，但是不能单纯的把左右窗口往右移动。因为有可能存在 ABBBBBA 的情况，这种情况需要从两边方向同时判断。正确的滑动窗口的做法应该是，边滑动的过程中边统计出现频次最多的字母，因为最后求得的最长长度的解，一定是在出现频次最多的字母上，再改变其他字母得到的最长连续长度。窗口滑动的过程中，用窗口的长度减去窗口中出现频次最大的长度，如果差值比 K 大，就代表需要缩小左窗口了直到差值等于 K。res 不断的取出窗口的长度的最大值就可以了。


