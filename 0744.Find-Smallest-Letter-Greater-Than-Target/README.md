# [744. Find Smallest Letter Greater Than Target](https://leetcode.com/problems/find-smallest-letter-greater-than-target/)


## 题目

Given a list of sorted characters `letters` containing only lowercase letters, and given a target letter `target`, find the smallest element in the list that is larger than the given target.

Letters also wrap around. For example, if the target is `target = 'z'` and `letters = ['a', 'b']`, the answer is `'a'`.

**Examples:**

    Input:
    letters = ["c", "f", "j"]
    target = "a"
    Output: "c"
    
    Input:
    letters = ["c", "f", "j"]
    target = "c"
    Output: "f"
    
    Input:
    letters = ["c", "f", "j"]
    target = "d"
    Output: "f"
    
    Input:
    letters = ["c", "f", "j"]
    target = "g"
    Output: "j"
    
    Input:
    letters = ["c", "f", "j"]
    target = "j"
    Output: "c"
    
    Input:
    letters = ["c", "f", "j"]
    target = "k"
    Output: "c"

**Note:**

1. `letters` has a length in range `[2, 10000]`.
2. `letters` consists of lowercase letters, and contains at least 2 unique letters.
3. `target` is a lowercase letter.


## 题目大意

给定一个只包含小写字母的有序数组letters 和一个目标字母 target，寻找有序数组里面比目标字母大的最小字母。

数组里字母的顺序是循环的。举个例子，如果目标字母target = 'z' 并且有序数组为 letters = ['a', 'b']，则答案返回 'a'。

注:

1. letters长度范围在[2, 10000]区间内。
2. letters 仅由小写字母组成，最少包含两个不同的字母。
3. 目标字母target 是一个小写字母。



## 解题思路

- 给出一个字节数组，在这个字节数组中查找在 target 后面的第一个字母。数组是环形的。
- 这一题也是二分搜索的题目，先在数组里面查找 target，如果找到了，取这个字母的后一个字母。如果没有找到，就取 low 下标的那个字母。注意数组是环形的，所以最后结果需要对下标取余。
