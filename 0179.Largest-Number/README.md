# [179. Largest Number](https://leetcode.com/problems/largest-number/)

## 题目

Given a list of non negative integers, arrange them such that they form the largest number.



Example 1:

```c
Input: [10,2]
Output: "210"
```


Example 2:

```c
Input: [3,30,34,5,9]
Output: "9534330"
```

Note: 

The result may be very large, so you need to return a string instead of an integer.



## 题目大意

给出一个数组，要求排列这些数组里的元素，使得最终排列出来的数字是最大的。


## 解题思路

这一题很容易想到把数字都转化为字符串，利用字符串比较，来排序，这样 9 开头的一定排在最前面。不过这样做有一个地方是错误的，比如："3" 和 "30" 比较，"30" 比 "3" 的字符序要大，这样排序以后就出错了。实际上就这道题而言， "3" 应该排在 "30" 前面。

在比较 2 个字符串大小的时候，不单纯的只用字符串顺序进行比较，还加入一个顺序。

```c
aStr := a + b
bStr := b + a
```

通过比较 aStr 和 bStr 的大小来得出是 a 大还是 b 大。

举个例子，还是 "3" 和 "30" 的例子，比较这 2 个字符串的大小。


```c
aStr := "3" + "30" = "330"
bStr := "30" + "3" = "303"
```

通过互相补齐位数之后再进行比较，就没有问题了。很显然这里 "3" 比 "30" 要大。



