# [455. Assign Cookies](https://leetcode.com/problems/assign-cookies/)

## 题目

Assume you are an awesome parent and want to give your children some cookies. But, you should give each child at most one cookie. Each child i has a greed factor gi, which is the minimum size of a cookie that the child will be content with; and each cookie j has a size sj. If sj >= gi, we can assign the cookie j to the child i, and the child i will be content. Your goal is to maximize the number of your content children and output the maximum number.

**Note:**You may assume the greed factor is always positive. You cannot assign more than one cookie to one child.

**Example 1:**

    Input: [1,2,3], [1,1]
    
    Output: 1
    
    Explanation: You have 3 children and 2 cookies. The greed factors of 3 children are 1, 2, 3. 
    And even though you have 2 cookies, since their size is both 1, you could only make the child whose greed factor is 1 content.
    You need to output 1.

**Example 2:**

    Input: [1,2], [1,2,3]
    
    Output: 2
    
    Explanation: You have 2 children and 3 cookies. The greed factors of 2 children are 1, 2. 
    You have 3 cookies and their sizes are big enough to gratify all of the children, 
    You need to output 2.


## 题目大意

假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。对每个孩子 i ，都有一个胃口值 gi ，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j ，都有一个尺寸 sj 。如果 sj >= gi ，我们可以将这个饼干 j 分配给孩子 i ，这个孩子会得到满足。你的目标是尽可能满足越多数量的孩子，并输出这个最大数值。

注意：你可以假设胃口值为正。一个小朋友最多只能拥有一块饼干。



## 解题思路


- 假设你想给小朋友们饼干，每个小朋友最多能够给一块饼干。每个小朋友都有一个“贪心指数”，称为 `g[i]`，`g[i]` 表示的是这名小朋友需要的饼干大小的最小值。同时，每个饼干都有一个大小值 `s[i]`，如果 `s[j] ≥ g[i]`，我们将饼干 `j` 分给小朋友 `i` 后，小朋友会很开心。给定数组 `g[]` 和 `s[]`，问如何分配饼干，能让更多的小朋友开心。
- 这是一道典型的简单贪心题。贪心题一般都伴随着排序。将 `g[]` 和 `s[]` 分别排序。按照最难满足的小朋友开始给饼干，依次往下满足，最终能满足的小朋友数就是最终解。
