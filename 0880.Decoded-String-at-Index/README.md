# [880. Decoded String at Index](https://leetcode.com/problems/decoded-string-at-index/)

## 题目

An encoded string S is given.  To find and write the decoded string to a tape, the encoded string is read one character at a time and the following steps are taken:

If the character read is a letter, that letter is written onto the tape.
If the character read is a digit (say d), the entire current tape is repeatedly written d-1 more times in total.
Now for some encoded string S, and an index K, find and return the K-th letter (1 indexed) in the decoded string.

 

Example 1:

```c
Input: S = "leet2code3", K = 10
Output: "o"
Explanation: 
The decoded string is "leetleetcodeleetleetcodeleetleetcode".
The 10th letter in the string is "o".
```

Example 2:

```c
Input: S = "ha22", K = 5
Output: "h"
Explanation: 
The decoded string is "hahahaha".  The 5th letter is "h".
```

Example 3:

```c
Input: S = "a2345678999999999999999", K = 1
Output: "a"
Explanation: 
The decoded string is "a" repeated 8301530446056247680 times.  The 1st letter is "a".
```

Note:

1. 2 <= S.length <= 100
2. S will only contain lowercase letters and digits 2 through 9.
3. S starts with a letter.
4. 1 <= K <= 10^9
5. The decoded string is guaranteed to have less than 2^63 letters.

## 题目大意

给定一个编码字符串 S。为了找出解码字符串并将其写入磁带，从编码字符串中每次读取一个字符，并采取以下步骤：  

- 如果所读的字符是字母，则将该字母写在磁带上。
- 如果所读的字符是数字（例如 d），则整个当前磁带总共会被重复写 d-1 次。   
  
现在，对于给定的编码字符串 S 和索引 K，查找并返回解码字符串中的第 K 个字母。


## 解题思路

按照题意，扫描字符串扫到数字的时候，开始重复字符串，这里可以用递归。注意在重复字符串的时候到第 K 个字符的时候就可以返回了，不要等所有字符都扩展完成，这样会超时。d 有可能超大。
