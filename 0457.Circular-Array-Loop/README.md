# [457. Circular Array Loop](https://leetcode.com/problems/circular-array-loop/)


## 题目

You are given a **circular** array `nums` of positive and negative integers. If a number k at an index is positive, then move forward k steps. Conversely, if it's negative (-k), move backward k steps. Since the array is circular, you may assume that the last element's next element is the first element, and the first element's previous element is the last element.

Determine if there is a loop (or a cycle) in `nums`. A cycle must start and end at the same index and the cycle's length > 1. Furthermore, movements in a cycle must all follow a single direction. In other words, a cycle must not consist of both forward and backward movements.

**Example 1:**

    Input: [2,-1,1,2,2]
    Output: true
    Explanation: There is a cycle, from index 0 -> 2 -> 3 -> 0. The cycle's length is 3.

**Example 2:**

    Input: [-1,2]
    Output: false
    Explanation: The movement from index 1 -> 1 -> 1 ... is not a cycle, because the cycle's length is 1. By definition the cycle's length must be greater than 1.

**Example 3:**

    Input: [-2,1,-1,-2,-2]
    Output: false
    Explanation: The movement from index 1 -> 2 -> 1 -> ... is not a cycle, because movement from index 1 -> 2 is a forward movement, but movement from index 2 -> 1 is a backward movement. All movements in a cycle must follow a single direction.

**Note:**

1. -1000 ≤ nums[i] ≤ 1000
2. nums[i] ≠ 0
3. 1 ≤ nums.length ≤ 5000

**Follow up:**

Could you solve it in **O(n)** time complexity and **O(1)** extra space complexity?


## 题目大意

给定一个含有正整数和负整数的环形数组 nums。 如果某个索引中的数 k 为正数，则向前移动 k 个索引。相反，如果是负数 (-k)，则向后移动 k 个索引。因为数组是环形的，所以可以假设最后一个元素的下一个元素是第一个元素，而第一个元素的前一个元素是最后一个元素。

确定 nums 中是否存在循环（或周期）。循环必须在相同的索引处开始和结束并且循环长度 > 1。此外，一个循环中的所有运动都必须沿着同一方向进行。换句话说，一个循环中不能同时包括向前的运动和向后的运动。

提示：

- -1000 ≤ nums[i] ≤ 1000
- nums[i] ≠ 0
- 1 ≤ nums.length ≤ 5000

进阶：

- 你能写出时间时间复杂度为 O(n) 和额外空间复杂度为 O(1) 的算法吗？

## 解题思路


- 给出一个循环数组，数组的数字代表了前进和后退的步数，+ 代表往右(前进)，- 代表往左(后退)。问这个循环数组中是否存在一个循环，并且这个循环内不能只有一个元素，循环的方向都必须是同方向的。
- 遇到循环就可以优先考虑用快慢指针的方法判断循环，这一题对循环增加了一个条件，循环不能只是单元素的循环，所以在快慢指针中加入这个判断条件。还有一个判断条件是循环的方向必须是同向的，这个简单，用 `num[i] * num[j] > 0` 就可以判断出是同向的(如果是反向的，那么两者的乘积必然是负数)，如果没有找到循环，可以将当前已经走过的路径上的 num[] 值都置为 0，标记已经访问过了。下次循环遇到访问过的元素，`num[i] * num[j] > 0` 就会是 0，提前退出找循环的过程。
