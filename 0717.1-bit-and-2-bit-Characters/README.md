# [717. 1-bit and 2-bit Characters](https://leetcode.com/problems/1-bit-and-2-bit-characters/)


## 题目:

We have two special characters. The first character can be represented by one bit `0`. The second character can be represented by two bits (`10` or `11`).

Now given a string represented by several bits. Return whether the last character must be a one-bit character or not. The given string will always end with a zero.

**Example 1:**

    Input: 
    bits = [1, 0, 0]
    Output: True
    Explanation: 
    The only way to decode it is two-bit character and one-bit character. So the last character is one-bit character.

**Example 2:**

    Input: 
    bits = [1, 1, 1, 0]
    Output: False
    Explanation: 
    The only way to decode it is two-bit character and two-bit character. So the last character is NOT one-bit character.

**Note:**

- `1 <= len(bits) <= 1000`.
- `bits[i]` is always `0` or `1`.

## 题目大意

有两种特殊字符。第一种字符可以用一比特0来表示。第二种字符可以用两比特(10 或 11)来表示。

现给一个由若干比特组成的字符串。问最后一个字符是否必定为一个一比特字符。给定的字符串总是由0结束。

注意:

- 1 <= len(bits) <= 1000.
- bits[i] 总是0 或 1.


## 解题思路

- 给出一个数组，数组里面的元素只有 0 和 1，并且数组的最后一个元素一定是 0。有 2 种特殊的字符，第一类字符是 "0"，第二类字符是 "11" 和 "10"，请判断这个数组最后一个元素是否一定是属于第一类字符？
- 依题意， 0 的来源有 2 处，可以是第一类字符，也可以是第二类字符，1 的来源只有 1 处，一定出自第二类字符。最后一个 0 当前仅当为第一类字符的情况有 2 种，第一种情况，前面出现有 0，但是 0 和 1 配对形成了第二类字符。第二种情况，前面没有出现 0 。这两种情况的共同点是除去最后一个元素，数组中前面所有的1 都“结对子”。所以利用第二类字符的特征，"1X"，遍历整个数组，如果遇到 "1"，就跳 2 步，因为 1 后面出现什么数字( 0 或者 1 )并不需要关心。如果 `i` 能在 `len(bits) - 1` 的地方`(数组最后一个元素)`停下，那么对应的是情况一或者情况二，前面的 0 都和 1 匹配上了，最后一个 0 一定是第一类字符。如果 `i` 在 `len(bit)` 的位置`(超出数组下标)`停下，说明 `bits[len(bits) - 1] == 1`，这个时候最后一个 0 一定属于第二类字符。
