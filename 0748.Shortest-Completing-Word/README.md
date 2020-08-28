# [748. Shortest Completing Word](https://leetcode.com/problems/shortest-completing-word/)


## 题目

Find the minimum length word from a given dictionary `words`, which has all the letters from the string `licensePlate`. Such a word is said to complete the given string `licensePlate`

Here, for letters we ignore case. For example, `"P"` on the `licensePlate` still matches `"p"` on the word.

It is guaranteed an answer exists. If there are multiple answers, return the one that occurs first in the array.

The license plate might have the same letter occurring multiple times. For example, given a `licensePlate` of `"PP"`, the word `"pair"` does not complete the `licensePlate`, but the word `"supper"` does.

**Example 1:**

    Input: licensePlate = "1s3 PSt", words = ["step", "steps", "stripe", "stepple"]
    Output: "steps"
    Explanation: The smallest length word that contains the letters "S", "P", "S", and "T".
    Note that the answer is not "step", because the letter "s" must occur in the word twice.
    Also note that we ignored case for the purposes of comparing whether a letter exists in the word.

**Example 2:**

    Input: licensePlate = "1s3 456", words = ["looks", "pest", "stew", "show"]
    Output: "pest"
    Explanation: There are 3 smallest length words that contains the letters "s".
    We return the one that occurred first.

**Note:**

1. `licensePlate` will be a string with length in range `[1, 7]`.
2. `licensePlate` will contain digits, spaces, or letters (uppercase or lowercase).
3. `words` will have a length in the range `[10, 1000]`.
4. Every `words[i]` will consist of lowercase letters, and have length in range `[1, 15]`.


## 题目大意

如果单词列表（words）中的一个单词包含牌照（licensePlate）中所有的字母，那么我们称之为完整词。在所有完整词中，最短的单词我们称之为最短完整词。

单词在匹配牌照中的字母时不区分大小写，比如牌照中的 "P" 依然可以匹配单词中的 "p" 字母。我们保证一定存在一个最短完整词。当有多个单词都符合最短完整词的匹配条件时取单词列表中最靠前的一个。牌照中可能包含多个相同的字符，比如说：对于牌照 "PP"，单词 "pair" 无法匹配，但是 "supper" 可以匹配。

注意:

- 牌照（licensePlate）的长度在区域[1, 7]中。
- 牌照（licensePlate）将会包含数字、空格、或者字母（大写和小写）。
- 单词列表（words）长度在区间 [10, 1000] 中。
- 每一个单词 words[i] 都是小写，并且长度在区间 [1, 15] 中。



## 解题思路


- 给出一个数组，要求找出能包含 `licensePlate` 字符串中所有字符的最短长度的字符串。如果最短长度的字符串有多个，输出 word 下标小的那个。这一题也是简单题，不过有 2 个需要注意的点，第一点，`licensePlate` 中可能包含 `Unicode` 任意的字符，所以要先把字母的字符筛选出来，第二点是题目中保证了一定存在一个最短的单词能满足题意，并且忽略大小写。具体做法按照题意模拟即可。
