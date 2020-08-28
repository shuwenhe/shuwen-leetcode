# [1003. Check If Word Is Valid After Substitutions](https://leetcode.com/problems/check-if-word-is-valid-after-substitutions/)

## 题目

We are given that the string "abc" is valid.

From any valid string V, we may split V into two pieces X and Y such that X + Y (X concatenated with Y) is equal to V.  (X or Y may be empty.)  Then, X + "abc" + Y is also valid.

If for example S = "abc", then examples of valid strings are: "abc", "aabcbc", "abcabc", "abcabcababcc".  Examples of invalid strings are: "abccba", "ab", "cababc", "bac".

Return true if and only if the given string S is valid.

 

Example 1:

```c
Input: "aabcbc"
Output: true
Explanation: 
We start with the valid string "abc".
Then we can insert another "abc" between "a" and "bc", resulting in "a" + "abc" + "bc" which is "aabcbc".
```

Example 2:

```c
Input: "abcabcababcc"
Output: true
Explanation: 
"abcabcabc" is valid after consecutive insertings of "abc".
Then we can insert "abc" before the last letter, resulting in "abcabcab" + "abc" + "c" which is "abcabcababcc".
```

Example 3:

```c
Input: "abccba"
Output: false
```

Example 4:

```c
Input: "cababc"
Output: false
```

Note:

1. 1 <= S.length <= 20000
2. S[i] is 'a', 'b', or 'c'

## 题目大意

假设 abc 是有效的字符串，对于任何 字符串 V，如果用 abc 把字符串 V 切成 2 半，X 和 Y，组成 X + abc + Y 的字符串，X + abc + Y 的这个字符串依旧是有效的。X 和 Y 可以是空字符串。

例如，"abc"( "" + "abc" + ""), "aabcbc"( "a" + "abc" + "bc"), "abcabc"( "" + "abc" + "abc"), "abcabcababcc"( "abc" + "abc" + "ababcc"，其中 "ababcc" 也是有效的，"ab" + "abc" + "c") 都是有效的字符串。

"abccba"( "" + "abc" + "cba"，"cba" 不是有效的字符串), "ab"("ab" 也不是有效字符串), "cababc"("c" + "abc" + "bc"，"c"，"bc" 都不是有效字符串), "bac" ("bac" 也不是有效字符串)这些都不是有效的字符串。

任意给一个字符串 S ，要求判断它是否有效，如果有效则输出 true。

## 解题思路

这一题可以类似括号匹配问题，因为 "abc" 这样的组合就代表是有效的，类似于括号匹配，遇到 "a" 就入栈，当遇到 "b" 字符的时候判断栈顶是不是 "a"，当遇到 "c" 字符的时候需要判断栈顶是不是 "a" 和 "b"。最后如果栈都清空了，就输出 true。